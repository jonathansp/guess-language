package data

import (
	"sort"
	"strings"
)

var (

	//todo: FIX IT!
	BasicLatin    = strings.Split("en ceb ha so tlh id haw la sw eu nr nso zu xh ss st tn ts", " ")
	ExtendedLatin = strings.Split("cs af pl hr ro sk sl tr hu az et sq ca es fr de nl it da is nb sv fi lv pt ve lt tl cy", " ")
	Latin         = append(BasicLatin, ExtendedLatin...)
	Cyrillic      = strings.Split("ru uk kk uz mn sr mk bg ky", " ")
	Arabic        = strings.Split("ar fa ps ur", " ")
	Devanagari    = strings.Split("hi ne", " ")

	OtherLanguage = map[string]string{
		"Armenian":  "hy",
		"Hebrew":    "he",
		"Bengali":   "bn",
		"Gurmukhi":  "pa",
		"Greek":     "el",
		"Gujarati":  "gu",
		"Oriya":     "or",
		"Tamil":     "ta",
		"Telugu":    "te",
		"Kannada":   "kn",
		"Malayalam": "ml",
		"Sinhala":   "si",
		"Thai":      "th",
		"Lao":       "lo",
		"Tibetan":   "bo",
		"Burmese":   "my",
		"Georgian":  "ka",
		"Mongolian": "mn",
		"Khmer":     "km",
	}

	//=========
	blockNames = []string{"Basic Latin",
		"Basic Latin", "Latin-1 Supplement", "Latin-1 Supplement",
		"Latin Extended-A",
		"Latin Extended-A",
		"Latin Extended-B", "Latin Extended-B", "IPA Extensions", "IPA Extensions",
		"Spacing Modifier Letters",
		"Spacing Modifier Letters", "Combining Diacritical Marks",
		"Combining Diacritical Marks", "Greek and Coptic", "Greek and Coptic",
		"Cyrillic", "Cyrillic", "Cyrillic Supplement", "Cyrillic Supplement", "Armenian",
		"Armenian", "Hebrew", "Hebrew", "Arabic", "Arabic",
		"Syriac", "Syriac", "Arabic Supplement", "Arabic Supplement", "Thaana",
		"Thaana", "NKo", "NKo", "Samaritan", "Samaritan",
		"Mandaic", "Mandaic", "Arabic Extended-A", "Arabic Extended-A", "Devanagari",
		"Devanagari", "Bengali", "Bengali", "Gurmukhi", "Gurmukhi",
		"Gujarati", "Gujarati", "Oriya", "Oriya", "Tamil",
		"Tamil", "Telugu", "Telugu", "Kannada", "Kannada",
		"Malayalam", "Malayalam", "Sinhala", "Sinhala", "Thai",
		"Thai", "Lao", "Lao", "Tibetan", "Tibetan",
		"Myanmar", "Myanmar", "Georgian", "Georgian", "Hangul Jamo",
		"Hangul Jamo", "Ethiopic", "Ethiopic", "Ethiopic Supplement",
		"Ethiopic Supplement",
		"Cherokee", "Cherokee", "Unified Canadian Aboriginal Syllabics",
		"Unified Canadian Aboriginal Syllabics", "Ogham",
		"Ogham", "Runic", "Runic", "Tagalog", "Tagalog",
		"Hanunoo", "Hanunoo", "Buhid", "Buhid", "Tagbanwa",
		"Tagbanwa", "Khmer", "Khmer", "Mongolian", "Mongolian",
		"Unified Canadian Aboriginal Syllabics Extended",
		"Unified Canadian Aboriginal Syllabics Extended", "Limbu", "Limbu",
		"Tai Le", "Tai Le", "New Tai Lue", "New Tai Lue", "Khmer Symbols", "Khmer Symbols",
		"Buginese", "Buginese", "Tai Tham", "Tai Tham",
		"Combining Diacritical Marks Extended",
		"Combining Diacritical Marks Extended", "Balinese", "Balinese",
		"Sundanese", "Sundanese",
		"Batak", "Batak", "Lepcha", "Lepcha", "Ol Chiki",
		"Ol Chiki", "Sundanese Supplement", "Sundanese Supplement",
		"Vedic Extensions", "Vedic Extensions",
		"Phonetic Extensions", "Phonetic Extensions", "Phonetic Extensions Supplement",
		"Phonetic Extensions Supplement", "Combining Diacritical Marks Supplement",
		"Combining Diacritical Marks Supplement", "Latin Extended Additional",
		"Latin Extended Additional", "Greek Extended", "Greek Extended",
		"General Punctuation", "General Punctuation", "Superscripts and Subscripts",
		"Superscripts and Subscripts", "Currency Symbols",
		"Currency Symbols", "Combining Diacritical Marks for Symbols",
		"Combining Diacritical Marks for Symbols", "Letterlike Symbols", "Letterlike Symbols",
		"Number Forms", "Number Forms", "Arrows", "Arrows", "Mathematical Operators",
		"Mathematical Operators", "Miscellaneous Technical", "Miscellaneous Technical",
		"Control Pictures", "Control Pictures",
		"Optical Character Recognition", "Optical Character Recognition", "Enclosed Alphanumerics", "Enclosed Alphanumerics", "Box Drawing",
		"Box Drawing", "Block Elements", "Block Elements", "Geometric Shapes", "Geometric Shapes",
		"Miscellaneous Symbols", "Miscellaneous Symbols", "Dingbats", "Dingbats",
		"Miscellaneous Mathematical Symbols-A",
		"Miscellaneous Mathematical Symbols-A", "Supplemental Arrows-A",
		"Supplemental Arrows-A", "Braille Patterns", "Braille Patterns",
		"Supplemental Arrows-B", "Supplemental Arrows-B", "Miscellaneous Mathematical Symbols-B",
		"Miscellaneous Mathematical Symbols-B", "Supplemental Mathematical Operators",
		"Supplemental Mathematical Operators", "Miscellaneous Symbols and Arrows",
		"Miscellaneous Symbols and Arrows", "Glagolitic", "Glagolitic",
		"Latin Extended-C", "Latin Extended-C", "Coptic", "Coptic", "Georgian Supplement",
		"Georgian Supplement", "Tifinagh", "Tifinagh", "Ethiopic Extended", "Ethiopic Extended",
		"Cyrillic Extended-A", "Cyrillic Extended-A", "Supplemental Punctuation",
		"Supplemental Punctuation", "CJK Radicals Supplement",
		"CJK Radicals Supplement", "Kangxi Radicals", "Kangxi Radicals",
		"Ideographic Description Characters", "Ideographic Description Characters",
		"CJK Symbols and Punctuation", "CJK Symbols and Punctuation", "Hiragana",
		"Hiragana", "Katakana", "Katakana", "Bopomofo", "Bopomofo",
		"Hangul Compatibility Jamo", "Hangul Compatibility Jamo",
		"Kanbun", "Kanbun", "Bopomofo Extended", "Bopomofo Extended", "CJK Strokes",
		"CJK Strokes", "Katakana Phonetic Extensions", "Katakana Phonetic Extensions",
		"Enclosed CJK Letters and Months", "Enclosed CJK Letters and Months",
		"CJK Compatibility", "CJK Compatibility", "CJK Unified Ideographs Extension A",
		"CJK Unified Ideographs Extension A", "Yijing Hexagram Symbols",
		"Yijing Hexagram Symbols", "CJK Unified Ideographs",
		"CJK Unified Ideographs", "Yi Syllables", "Yi Syllables",
		"Yi Radicals", "Yi Radicals", "Lisu", "Lisu", "Vai",
		"Vai", "Cyrillic Extended-B", "Cyrillic Extended-B", "Bamum", "Bamum",
		"Modifier Tone Letters", "Modifier Tone Letters", "Latin Extended-D", "Latin Extended-D", "Syloti Nagri",
		"Syloti Nagri", "Common Indic Number Forms", "Common Indic Number Forms", "Phags-pa", "Phags-pa",
		"Saurashtra", "Saurashtra", "Devanagari Extended", "Devanagari Extended", "Kayah Li",
		"Kayah Li", "Rejang", "Rejang", "Hangul Jamo Extended-A", "Hangul Jamo Extended-A",
		"Javanese", "Javanese", "Myanmar Extended-B", "Myanmar Extended-B", "Cham",
		"Cham", "Myanmar Extended-A", "Myanmar Extended-A", "Tai Viet", "Tai Viet",
		"Meetei Mayek Extensions", "Meetei Mayek Extensions", "Ethiopic Extended-A",
		"Ethiopic Extended-A", "Latin Extended-E",
		"Latin Extended-E", "Cherokee Supplement", "Cherokee Supplement", "Meetei Mayek", "Meetei Mayek",
		"Hangul Syllables", "Hangul Syllables", "Hangul Jamo Extended-B",
		"Hangul Jamo Extended-B", "High Surrogates",
		"High Surrogates", "High Private Use Surrogates", "High Private Use Surrogates",
		"Low Surrogates", "Low Surrogates",
		"Private Use Area", "Private Use Area", "CJK Compatibility Ideographs",
		"CJK Compatibility Ideographs", "Alphabetic Presentation Forms",
		"Alphabetic Presentation Forms", "Arabic Presentation Forms-A",
		"Arabic Presentation Forms-A", "Variation Selectors", "Variation Selectors",
		"Vertical Forms", "Vertical Forms", "Combining Half Marks", "Combining Half Marks", "CJK Compatibility Forms",
		"CJK Compatibility Forms", "Small Form Variants", "Small Form Variants",
		"Arabic Presentation Forms-B", "Arabic Presentation Forms-B",
		"Halfwidth and Fullwidth Forms", "Halfwidth and Fullwidth Forms", "Specials", "Specials"}

	blockIDs = []int{0, 127, 128, 255, 256, 383, 384, 591, 592, 687, 688, 767, 768, 879, 880, 1023, 1024, 1279, 1280, 1327, 1328, 1423, 1424, 1535, 1536, 1791,
		1792, 1871, 1872, 1919, 1920, 1983, 1984, 2047, 2048, 2111, 2112, 2143, 2208, 2303, 2304, 2431, 2432, 2559, 2560, 2687, 2688, 2815, 2816, 2943, 2944,
		3071, 3072, 3199, 3200, 3327, 3328, 3455, 3456, 3583, 3584, 3711, 3712, 3839, 3840, 4095, 4096, 4255, 4256, 4351, 4352, 4607, 4608, 4991, 4992, 5023,
		5024, 5119, 5120, 5759, 5760, 5791, 5792, 5887, 5888, 5919, 5920, 5951, 5952, 5983, 5984, 6015, 6016, 6143, 6144, 6319, 6320, 6399, 6400, 6479, 6480,
		6527, 6528, 6623, 6624, 6655, 6656, 6687, 6688, 6831, 6832, 6911, 6912, 7039, 7040, 7103, 7104, 7167, 7168, 7247, 7248, 7295, 7360, 7375, 7376, 7423,
		7424, 7551, 7552, 7615, 7616, 7679, 7680, 7935, 7936, 8191, 8192, 8303, 8304, 8351, 8352, 8399, 8400, 8447, 8448, 8527, 8528, 8591, 8592, 8703, 8704,
		8959, 8960, 9215, 9216, 9279, 9280, 9311, 9312, 9471, 9472, 9599, 9600, 9631, 9632, 9727, 9728, 9983, 9984, 10175, 10176, 10223, 10224, 10239, 10240, 10495,
		10496, 10623, 10624, 10751, 10752, 11007, 11008, 11263, 11264, 11359, 11360, 11391, 11392, 11519, 11520, 11567, 11568, 11647, 11648, 11743, 11744, 11775, 11776, 11903, 11904,
		12031, 12032, 12255, 12272, 12287, 12288, 12351, 12352, 12447, 12448, 12543, 12544, 12591, 12592, 12687, 12688, 12703, 12704, 12735, 12736, 12783, 12784, 12799, 12800, 13055,
		13056, 13311, 13312, 19903, 19904, 19967, 19968, 40959, 40960, 42127, 42128, 42191, 42192, 42239, 42240, 42559, 42560, 42655, 42656, 42751, 42752, 42783, 42784, 43007, 43008,
		43055, 43056, 43071, 43072, 43135, 43136, 43231, 43232, 43263, 43264, 43311, 43312, 43359, 43360, 43391, 43392, 43487, 43488, 43519, 43520, 43615, 43616, 43647, 43648, 43743,
		43744, 43775, 43776, 43823, 43824, 43887, 43888, 43967, 43968, 44031, 44032, 55215, 55216, 55295, 55296, 56191, 56192, 56319, 56320, 57343, 57344, 63743, 63744, 64255, 64256,
		64335, 64336, 65023, 65024, 65039, 65040, 65055, 65056, 65071, 65072, 65103, 65104, 65135, 65136, 65279, 65280, 65519, 65520, 65535}
)

// GetBlock ...
func GetBlock(character rune) string {
	i := sort.Search(len(blockIDs), func(i int) bool {
		return blockIDs[i] >= int(character)
	})

	return blockNames[i]
}
