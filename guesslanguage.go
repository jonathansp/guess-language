package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/jonathansp/guess-language/blocks"
	"github.com/jonathansp/guess-language/collections"
	"github.com/jonathansp/guess-language/languageinfo"
	"golang.org/x/text/unicode/norm"
)

const (
	minLength = 20
	maxGrams  = 300
)

type (
	//Guesser ...
	Guesser struct {
		trigrams map[string]map[string]int
	}
)

// GuessLanguage ...
func GuessLanguage() *Guesser {
	instance := Guesser{}
	instance.loadTrigrams()
	return &instance
}

// TODO: send to a help
func in(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (gl *Guesser) createModel(content string) []string {
	trigrams := make(map[string]int)
	content = strings.ToLower(content)
	for i := 0; i <= len(content)-2; i++ {
		end := i + 3
		if end > len(content) {
			end = len(content) - 1
		}
		_, ok := trigrams[content[i:end]]
		if !ok {
			trigrams[content[i:end]] = 0
		}
		trigrams[content[i:end]]++
	}
	var data []string
	for _, res := range collections.SortedKeys(trigrams) {
		data = append(data, res)
	}
	return data

}

func (gl *Guesser) distance(foundModels []string, knownModel map[string]int) float64 {
	var dist float64
	if len(foundModels) > maxGrams {
		foundModels = foundModels[:maxGrams]
	}

	for i, value := range foundModels {
		if !strings.Contains(value, "  ") {
			if _, ok := knownModel[value]; ok {
				dist += math.Abs(float64(i - knownModel[value]))
			} else {
				dist += maxGrams
			}
		}
	}
	return dist
}

func (gl *Guesser) check(sample string, languageSet []string) string {
	if len(sample) < minLength {
		return "UNKNOWN"
	}
	model := gl.createModel(sample)
	minDistance := math.Inf(1)
	lastLang := "UNKNOWN"
	for _, lang := range languageSet {
		lang := strings.ToLower(lang)
		if value, ok := gl.trigrams[lang]; ok {
			distance := gl.distance(model, value)
			if distance < minDistance {
				minDistance = distance
				lastLang = lang
			}
		}
	}
	return lastLang
}

func (gl *Guesser) identify(sample string, scripts []string) string {

	//todo: remove from here
	BasicLatin := strings.Split("en ceb ha so tlh id haw la sw eu nr nso zu xh ss st tn ts", " ")
	ExtendedLatin := strings.Split("cs af pl hr ro sk sl tr hu az et sq ca es fr de nl it da is nb sv fi lv pt ve lt tl cy", " ")
	Latin := append(BasicLatin, ExtendedLatin...)
	Cyrillic := strings.Split("ru uk kk uz mn sr mk bg ky", " ")
	Arabic := strings.Split("ar fa ps ur", " ")
	Devanagari := strings.Split("hi ne", " ")

	OtherLanguage := make(map[string]string)
	OtherLanguage["Armenian"] = "hy"
	OtherLanguage["Hebrew"] = "he"
	OtherLanguage["Bengali"] = "bn"
	OtherLanguage["Gurmukhi"] = "pa"
	OtherLanguage["Greek"] = "el"
	OtherLanguage["Gujarati"] = "gu"
	OtherLanguage["Oriya"] = "or"
	OtherLanguage["Tamil"] = "ta"
	OtherLanguage["Telugu"] = "te"
	OtherLanguage["Kannada"] = "kn"
	OtherLanguage["Malayalam"] = "ml"
	OtherLanguage["Sinhala"] = "si"
	OtherLanguage["Thai"] = "th"
	OtherLanguage["Lao"] = "lo"
	OtherLanguage["Tibetan"] = "bo"
	OtherLanguage["Burmese"] = "my"
	OtherLanguage["Georgian"] = "ka"
	OtherLanguage["Mongolian"] = "mn"
	OtherLanguage["Khmer"] = "km"

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
	if in("Cyrillic", scripts) {
		return gl.check(sample, Cyrillic)
	}
	if in("Arabic", scripts) || in("Arabic Presentation Forms-A", scripts) || in("Arabic Presentation Forms-B", scripts) {
		return gl.check(sample, Arabic)
	}
	if in("Devanagari", scripts) {
		return gl.check(sample, Devanagari)
	}
	if in("CJK Unified Ideographs", scripts) || in("Bopomofo", scripts) || in("Bopomofo Extended", scripts) || in("KangXi Radicals", scripts) {
		return "zh"
	}
	for block, language := range OtherLanguage {
		if in(block, scripts) {
			return language
		}
	}

	if in("Extended Latin", scripts) {
		return gl.check(sample, ExtendedLatin)
	}
	if in("Latin Extended Additional", scripts) {
		return "vi"
	}
	if in("Basic Latin", scripts) {
		return gl.check(sample, Latin)
	}
	return "UNKNOWN"

}

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
func (gl *Guesser) Parse(text string) languageinfo.Language {
	text = gl.normalize(text)
	runs := gl.findRuns(text)
	isocode := gl.identify(text, runs)
	lang := languageinfo.Languages[isocode]
	fmt.Println(lang)
	return lang
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
	x.Parse("Pedras no caminho? Eu guardo todas. Um dia vou construir um castelo.")
	x.Parse("Los siguientes tutoriales te darán las pautas principales para comenzar a utilizar el nuevo sistema.")
	x.Parse("თქვენ შეგიძლიათ ისარგებლოთ რეგისტრაციის განახლებული ვებ გვერდით. Ge დომენების რეგისტრაცია, დომენური ")
	x.Parse("Wir stellen für die Domainverwaltung ein automatisches elektronisches Registrierungssystem zur Verfügung und betreiben ein weltweites Netz von Nameservern, das sicherstellt, dass über 15 Millionen")
	x.Parse("目aНовинка! Благодаря сервису Google Мой бизнес вы можете бесплатно рассказать о себе клиентам с помощью")
	x.Parse("ᠮᠤᠩᠭᠤᠯᠤᠯᠤᠰ")
	x.Parse("The easy way to start building Golang command line application.")
	x.Parse("Uma  ①  frase em português")
	x.Parse("できる限りわかりやすい説明を目指しておりますが、Cookie、IP アドレス、ピクセル タグ、ブラウザなどの用語がご不明の場合は、先にこれらの主な用語についての説明をご覧ください。Google ではお客様のプライバシーを重視しております")
}
