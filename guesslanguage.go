package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/jonathansp/guess-language/blocks"
	"golang.org/x/text/unicode/norm"
)

type (
	//Guesser ...
	Guesser struct {
		trigrams map[string]map[string]int
	}
	// Language ...
	Language struct {
		id   int
		tag  int
		name string
		info string
	}
)

// GuessLanguage ...
func GuessLanguage() *Guesser {
	instance := Guesser{}
	instance.loadTrigrams()
	return &instance
}

func in(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (gl *Guesser) check(sample string, languageSet []string) string {
	return ""
}

func (gl *Guesser) identify(sample string, scripts []string) string {

	//todo: remove from here
	BasicLatin := strings.Split("en ceb ha so tlh id haw la sw eu nr nso zu xh ss st tn ts", " ")
	ExtendedLatin := strings.Split("cs af pl hr ro sk sl tr hu az et sq ca es fr de nl it da is nb sv fi lv pt ve lt tl cy", " ")
	Latin := append(BasicLatin, ExtendedLatin...)
	Cyrillic := strings.Split("ru uk kk uz mn sr mk bg ky", " ")
	Arabic := strings.Split("ar fa ps ur", " ")
	Devanagari := strings.Split("hi ne", " ")

	if len(sample) < 3 {
		return "UNKNOWN"
	}
	if in("Hangul Syllables", scripts) || in("Hangul Jamo", scripts) || in("Hangul Compatibility Jamo", scripts) || in("Hangul", scripts) {
		return "ko"
	}
	if in("Greek and Coptic", scripts) {
		return "el"
	}
	if in("Katakana", scripts) {
		return "ja"
	}
	if in("CJK Unified Ideographs", scripts) || in("Bopomofo", scripts) || in("Bopomofo Extended", scripts) || in("KangXi Radicals", scripts) {
		return "zh"
	}
	if in("Cyrillic", scripts) {
		return gl.check(sample, Cyrillic)
	}
	if in("Arabic", scripts) || in("Arabic Presentation Forms-A", scripts) || in("Arabic Presentation Forms-B", scripts) {
		return gl.check(sample, Arabic)
	}
	if in("Devanagari", scripts) {
		return gl.check(sample, Devanagari)
	}
	/*
	   # Try languages with unique scripts
	   for blockName, langName in SINGLETONS:
	       if blockName in scripts:
	           return langName
	*/

	if in("Latin Extended Additional", scripts) {
		return "vi"
	}
	/*
		if "Extended Latin" in scripts:
			latinLang = check(sample, EXTENDED_LATIN)
			if latinLang == "pt":
				return check(sample, PT)
			else:
				return latinLang

	*/
	if in("Basic Latin", scripts) {
		return gl.check(sample, Latin)
	}
	return "UNKNOWN"

}

/*







   return UNKNOWN
*/

func (gl *Guesser) normalize(text string) string {
	buffer := []byte(text)
	buffer = norm.NFKC.Bytes(buffer)
	text = strings.Replace(string(buffer), "  ", " ", -1)
	return text
}

func (gl *Guesser) findRuns(text string) []string {

	var runTypes = make(map[string]int)
	var relevantRuns []string
	var totalCount = 0

	for _, char := range text {
		if unicode.IsLetter(char) {
			block := blocks.GetBlock(char)
			_, ok := runTypes[block]
			if !ok {
				runTypes[block] = 0
			}
			runTypes[block] = +1
			totalCount = +1
		}
	}

	for key, value := range runTypes {
		percent := (value * 100) / totalCount
		if percent >= 40 || (key == "Basic Latin" && percent >= 15) || (key == "Latin Extended Additional" && percent >= 10) {
			relevantRuns = append(relevantRuns, key)
		}
	}
	return relevantRuns
}

// Parse ...
func (gl *Guesser) Parse(text string) Language {
	text = gl.normalize(text)
	runs := gl.findRuns(text)
	result := gl.identify(text, runs)

	fmt.Print(result)

	return Language{}
}

func (gl *Guesser) loadTrigrams() {
	r := regexp.MustCompile("(.{3})\\s+(.*)")

	gl.trigrams = make(map[string]map[string]int)
	files, err := ioutil.ReadDir("trigrams")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		iana := f.Name()
		file, err := os.Open(filepath.Join("trigrams", iana))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			_, ok := gl.trigrams[iana]
			if !ok {
				gl.trigrams[iana] = make(map[string]int)
			}
			parsed := r.FindStringSubmatch(line)
			key := parsed[1]
			value := parsed[2]
			val, _ := strconv.Atoi(value)
			gl.trigrams[iana][key] = val
		}
	}
}

func main() {
	x := GuessLanguage()
	//z := x.Parse("目aНовинка! Благодаря сервису Google Мой бизнес вы можете бесплатно рассказать о себе клиентам с помощью Поиска Google, Google+ и Карт Google.")
	z := x.Parse("漢")
	fmt.Print("%v", z)
	//fmt.Print(x.normalize("Uma  ①  frase em português"))
	//fmt.Print(blocks.Data)
}
