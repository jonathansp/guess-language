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
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/jonathansp/guess-language/collections"
	"github.com/jonathansp/guess-language/data"
	"golang.org/x/text/unicode/norm"
)

const (
	minLength = 20
	maxGrams  = 300
)

var (
	//ErrUnknownLanguage returns when Parse() wasn't able to find an idiom.
	ErrUnknownLanguage = errors.New("Unknown Language")
	//ErrStringTooShort returns if string has 3 chars or less.
	ErrStringTooShort = errors.New("String too short")
)

func createModel(content string) []string {
	trigrams := make(map[string]int)
	content = strings.ToLower(content)
	size := utf8.RuneCountInString(content)
	for i := 0; i <= size-2; i++ {
		end := i + 3
		if end > size {
			end = size - 1
		}
		slice := string([]rune(content)[i:end])
		trigrams[slice]++
	}
	return collections.SortedKeys(trigrams)
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

func hasOne(items []string, sortedList sort.StringSlice) bool {

	for _, item := range items {
		i := sort.SearchStrings(sortedList, item)
		if i < len(sortedList) && sortedList[i] == item {
			return true
		}
	}
	return false
}

func identify(sample string, scripts sort.StringSlice) (string, error) {
	if len(sample) < minLength {
		return "", ErrStringTooShort
	}

	sort.Strings(scripts)

	switch {

	case hasOne([]string{"Hangul Syllables", "Hangul Jamo",
		"Hangul Compatibility Jamo", "Hangul"}, scripts):
		return "ko", nil

	case hasOne([]string{"Greek and Coptic"}, scripts):
		return "el", nil

	case hasOne([]string{"Katakana"}, scripts):
		return "ja", nil

	case hasOne([]string{"Cyrillic"}, scripts):
		return check(sample, data.Alphabets["Cyrillic"])

	case hasOne([]string{"Arabic", "Arabic Presentation Forms-A",
		"Arabic Presentation Forms-B"}, scripts):
		return check(sample, data.Alphabets["Arabic"])

	case hasOne([]string{"Devanagari"}, scripts):
		return check(sample, data.Alphabets["Devanagari"])

	case hasOne([]string{"CJK Unified Ideographs", "Bopomofo",
		"Bopomofo Extended", "KangXi Radicals"}, scripts):
		return "zh", nil
	}

	// We need to check LocalLanguage (specific idioms) before checking Latin.
	for block, language := range data.LocalLanguages {
		if hasOne([]string{block}, scripts) {
			return language, nil
		}
	}

	switch {

	case hasOne([]string{"Latin Extended Additional"}, scripts):
		return "vi", nil

	case hasOne([]string{"Extended Latin"}, scripts):
		return check(sample, data.Alphabets["ExtendedLatin"])

	case hasOne([]string{"Basic Latin"}, scripts):
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
	var relevantRuns sort.StringSlice
	var totalCount = 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			block := data.GetBlockFromChar(char)
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

// Parse function tries to identify natural language from a given string.
func Parse(text string) (data.Language, error) {
	text = normalize(text)
	runs := findRuns(text)
	if len(runs) == 0 {
		return data.Languages[""], ErrUnknownLanguage
	}
	isocode, err := identify(text, runs)
	return data.Languages[isocode], err
}
