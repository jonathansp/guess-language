package guesslanguage

import (
	"math"
	"strings"
	"unicode"

	"github.com/jonathansp/guess-language/data"
	"github.com/jonathansp/guess-language/utils"
	"golang.org/x/text/unicode/norm"
)

const (
	minLength = 20
	maxGrams  = 300
)

type (
	//Guesser ...
	Guesser struct {
	}
)

// GuessLanguage ...
func GuessLanguage() *Guesser {
	instance := Guesser{}
	return &instance
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
	for _, res := range utils.SortedKeys(trigrams) {
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
		if value, ok := data.Trigrams[lang]; ok {
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

	if len(sample) < 3 {
		return "UNKNOWN"
	}
	if utils.In("Hangul Syllables", scripts) || utils.In("Hangul Jamo", scripts) || utils.In("Hangul Compatibility Jamo", scripts) || utils.In("Hangul", scripts) {
		return "ko"
	}
	if utils.In("Greek and Coptic", scripts) {
		return "el"
	}
	if utils.In("Katakana", scripts) {
		return "ja"
	}
	if utils.In("Cyrillic", scripts) {
		return gl.check(sample, data.Cyrillic)
	}
	if utils.In("Arabic", scripts) || utils.In("Arabic Presentation Forms-A", scripts) || utils.In("Arabic Presentation Forms-B", scripts) {
		return gl.check(sample, data.Arabic)
	}
	if utils.In("Devanagari", scripts) {
		return gl.check(sample, data.Devanagari)
	}
	if utils.In("CJK Unified Ideographs", scripts) || utils.In("Bopomofo", scripts) || utils.In("Bopomofo Extended", scripts) || utils.In("KangXi Radicals", scripts) {
		return "zh"
	}
	for block, language := range data.OtherLanguage {
		if utils.In(block, scripts) {
			return language
		}
	}

	if utils.In("Extended Latin", scripts) {
		return gl.check(sample, data.ExtendedLatin)
	}
	if utils.In("Latin Extended Additional", scripts) {
		return "vi"
	}
	if utils.In("Basic Latin", scripts) {
		return gl.check(sample, data.Latin)
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
			block := data.GetBlock(char)
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
func (gl *Guesser) Parse(text string) data.Language {
	text = gl.normalize(text)
	runs := gl.findRuns(text)
	isocode := gl.identify(text, runs)
	lang := data.Languages[isocode]
	return lang
}
