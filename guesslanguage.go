package guesslanguage

/*  Guess the natural language (idiom) of a text.

 	2015, Jonathan Simon Prates.
	Go library to identify the natural language of a text.

    Based on guesslanguage.py by Kent S Johnson (https://pypi.python.org/pypi/guess-language)
    Python version: 2008, Kent S Johnson
    C++ version: 2006 Jacob R Rideout <kde@jacobrideout.net>
    Perl version: 2004-6 Maciej Ceglowski

    This library is free software; you can redistribute it and/or
    modify it under the terms of the GNU GENERAL PUBLIC LICENSE.

    This library is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
	See LICENSE file for details.

*/

import (
	"errors"
	"math"
	"strings"
	"unicode"

	"github.com/jonathansp/guess-language/data"
	arrays "github.com/jonathansp/guess-language/utils"
	"golang.org/x/text/unicode/norm"
)

const (
	minLength = 20
	maxGrams  = 300
)

var (
	//ErrUnknownLanguage throwns when Parse wasn't able to find an idiom.
	ErrUnknownLanguage = errors.New("Unknown Language")
	//ErrStringTooShort throwns if the string has 3 chars or less.
	ErrStringTooShort = errors.New("String too short")
)

func createModel(content string) []string {
	trigrams := make(map[string]int)
	content = strings.ToLower(content)
	for i := 0; i <= len(content)-2; i++ {
		end := i + 3
		if end > len(content) {
			end = len(content) - 1
		}
		if _, ok := trigrams[content[i:end]]; !ok {
			trigrams[content[i:end]] = 0
		}
		trigrams[content[i:end]]++
	}
	var data []string
	for _, res := range arrays.SortedKeys(trigrams) {
		data = append(data, res)
	}
	return data
}

func distance(foundModel []string, knownModel map[string]int) float64 {
	var dist float64
	if len(foundModel) > maxGrams {
		foundModel = foundModel[:maxGrams]
	}
	for i, value := range foundModel {
		// must ignore samples with 2 spaces in a row.
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

func check(sample string, languageSet []string) (string, error) {
	model := createModel(sample)
	minDistance := math.Inf(1)
	var lastLang string
	for _, lang := range languageSet {
		lang := strings.ToLower(lang)
		if value, ok := data.Trigrams[lang]; ok {
			distance := distance(model, value)
			if distance < minDistance {
				minDistance = distance
				lastLang = lang
			}
		}
	}
	if lastLang == "" {
		return "", ErrUnknownLanguage
	}
	return lastLang, nil
}

func identify(sample string, scripts []string) (string, error) {
	if len(sample) < minLength {
		return "", ErrStringTooShort
	}

	switch {

	case arrays.HasOne([]string{"Hangul Syllables", "Hangul Jamo",
		"Hangul Compatibility Jamo", "Hangul"}, scripts):
		return "ko", nil

	case arrays.In("Greek and Coptic", scripts):
		return "el", nil

	case arrays.In("Katakana", scripts):
		return "ja", nil

	case arrays.In("Cyrillic", scripts):
		return check(sample, data.Alphabets["Cyrillic"])

	case arrays.HasOne([]string{"Arabic", "Arabic Presentation Forms-A",
		"Arabic Presentation Forms-B"}, scripts):
		return check(sample, data.Alphabets["Arabic"])

	case arrays.In("Devanagari", scripts):
		return check(sample, data.Alphabets["Devanagari"])

	case arrays.HasOne([]string{"CJK Unified Ideographs", "Bopomofo",
		"Bopomofo Extended", "KangXi Radicals"}, scripts):
		return "zh", nil
	}

	// We need to check LocalLanguage (specific idioms) before checking Latin.
	for block, language := range data.LocalLanguages {
		if arrays.In(block, scripts) {
			return language, nil
		}
	}

	switch {

	case arrays.In("Extended Latin", scripts):
		return check(sample, data.Alphabets["ExtendedLatin"])

	case arrays.In("Latin Extended Additional", scripts):
		return "vi", nil

	case arrays.In("Basic Latin", scripts):
		return check(sample, data.Alphabets["Latin"])
	}

	return "", ErrUnknownLanguage
}

func normalize(text string) string {
	buffer := []byte(text)
	buffer = norm.NFKC.Bytes(buffer)
	text = strings.Replace(string(buffer), "  ", " ", -1)
	return text
}

func findRuns(text string) []string {
	var runTypes = make(map[string]int)
	var relevantRuns []string
	var totalCount = 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			block := data.GetBlockFromChar(char)
			if _, ok := runTypes[block]; !ok {
				runTypes[block] = 0
			}
			runTypes[block] = +1
			totalCount = +1
		}
	}

	for key, value := range runTypes {
		percent := (value * 100) / totalCount
		if percent >= 40 ||
			(key == "Basic Latin" && percent >= 15) ||
			(key == "Latin Extended Additional" && percent >= 10) {
			relevantRuns = append(relevantRuns, key)
		}
	}
	return relevantRuns
}

// Parse a string in order to identify its natural language.
func Parse(text string) (data.Language, error) {
	text = normalize(text)
	runs := findRuns(text)
	if len(runs) == 0 {
		return data.Languages[""], ErrUnknownLanguage
	}
	isocode, err := identify(text, runs)
	return data.Languages[isocode], err
}
