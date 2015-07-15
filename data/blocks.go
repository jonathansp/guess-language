package data

import "sort"

var (
	// Alphabets database
	Alphabets map[string][]string

	// LocalLanguages database (Specific regions)
	LocalLanguages map[string]string

	// Blocks from http://www.unicode.org/Public/UNIDATA/Blocks.txt
	Blocks           map[int]string
	orderedBlockKeys []int
)

func init() {

	Alphabets = make(map[string][]string)
	Alphabets["BasicLatin"] = []string{"en", "ceb", "ha", "so", "tlh", "id", "haw", "la",
		"sw", "eu", "nr", "nso", "zu", "xh", "ss", "st", "tn", "ts"}

	Alphabets["ExtendedLatin"] = []string{"cs", "af", "pl", "hr", "ro", "sk", "sl", "tr",
		"hu", "az", "et", "sq", "ca", "es", "fr", "de", "nl", "it", "da", "is", "no", "sv",
		"fi", "lv", "pt", "ve", "lt", "tl", "cy"}

	Alphabets["Cyrillic"] = []string{"ru", "uk", "kk", "uz", "mn", "sr", "mk", "bg", "ky"}
	Alphabets["Arabic"] = []string{"ar", "fa", "ps", "ur"}
	Alphabets["Devanagari"] = []string{"hi", "ne"}

	// Latin is BasicLatin + ExtendedLatin
	Alphabets["Latin"] = append(Alphabets["BasicLatin"], Alphabets["ExtendedLatin"]...)

	LocalLanguages = map[string]string{
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

	Blocks = map[int]string{
		0x007F: "Basic Latin",
		0x017F: "Extended Latin",
		0x024F: "Latin Extended-B",
		0x02AF: "Extended Latin",
		0x02FF: "Spacing Modifier Letters",
		0x036F: "Combining Diacritical Marks",
		0x03FF: "Greek and Coptic",
		0x04FF: "Cyrillic",
		0x052F: "Cyrillic Supplement",
		0x058F: "Armenian",
		0x05FF: "Hebrew",
		0x06FF: "Arabic",
		0x074F: "Syriac",
		0x077F: "Arabic Supplement",
		0x07BF: "Thaana",
		0x07FF: "NKo",
		0x097F: "Devanagari",
		0x09FF: "Bengali",
		0x0A7F: "Gurmukhi",
		0x0AFF: "Gujarati",
		0x0B7F: "Oriya",
		0x0BFF: "Tamil",
		0x0C7F: "Telugu",
		0x0CFF: "Kannada",
		0x0D7F: "Malayalam",
		0x0DFF: "Sinhala",
		0x0E7F: "Thai",
		0x0EFF: "Lao",
		0x0FFF: "Tibetan",
		0x109F: "Myanmar",
		0x10FF: "Georgian",
		0x11FF: "Hangul Jamo",
		0x137F: "Ethiopic",
		0x139F: "Ethiopic Supplement",
		0x13FF: "Cherokee",
		0x167F: "Unified Canadian Aboriginal Syllabics",
		0x169F: "Ogham",
		0x16FF: "Runic",
		0x171F: "Tagalog",
		0x173F: "Hanunoo",
		0x175F: "Buhid",
		0x177F: "Tagbanwa",
		0x17FF: "Khmer",
		0x18AF: "Mongolian",
		0x194F: "Limbu",
		0x197F: "Tai Le",
		0x19DF: "New Tai Lue",
		0x19FF: "Khmer Symbols",
		0x1A1F: "Buginese",
		0x1B7F: "Balinese",
		0x1D7F: "Phonetic Extensions",
		0x1DBF: "Phonetic Extensions Supplement",
		0x1DFF: "Combining Diacritical Marks Supplement",
		0x1EFF: "Latin Extended Additional",
		0x1FFF: "Greek Extended",
		0x206F: "General Punctuation",
		0x209F: "Superscripts and Subscripts",
		0x20CF: "Currency Symbols",
		0x20FF: "Combining Diacritical Marks for Symbols",
		0x214F: "Letterlike Symbols",
		0x218F: "Number Forms",
		0x21FF: "Arrows",
		0x22FF: "Mathematical Operators",
		0x23FF: "Miscellaneous Technical",
		0x243F: "Control Pictures",
		0x245F: "Optical Character Recognition",
		0x24FF: "Enclosed Alphanumerics",
		0x257F: "Box Drawing",
		0x259F: "Block Elements",
		0x25FF: "Geometric Shapes",
		0x26FF: "Miscellaneous Symbols",
		0x27BF: "Dingbats",
		0x27EF: "Miscellaneous Mathematical Symbols-A",
		0x27FF: "Supplemental Arrows-A",
		0x28FF: "Braille Patterns",
		0x297F: "Supplemental Arrows-B",
		0x29FF: "Miscellaneous Mathematical Symbols-B",
		0x2AFF: "Supplemental Mathematical Operators",
		0x2BFF: "Miscellaneous Symbols and Arrows",
		0x2C5F: "Glagolitic",
		0x2C7F: "Latin Extended-C",
		0x2CFF: "Coptic",
		0x2D2F: "Georgian Supplement",
		0x2D7F: "Tifinagh",
		0x2DDF: "Ethiopic Extended",
		0x2E7F: "Supplemental Punctuation",
		0x2EFF: "CJK Radicals Supplement",
		0x2FDF: "Kangxi Radicals",
		0x2FFF: "Ideographic Description Characters",
		0x303F: "CJK Symbols and Punctuation",
		0x309F: "Katakana",
		0x30FF: "Katakana",
		0x312F: "Bopomofo",
		0x318F: "Hangul Compatibility Jamo",
		0x319F: "Kanbun",
		0x31BF: "Bopomofo Extended",
		0x31EF: "CJK Strokes",
		0x31FF: "Katakana",
		0x32FF: "Enclosed CJK Letters and Months",
		0x33FF: "CJK Compatibility",
		0x4DBF: "CJK Unified Ideographs Extension A",
		0x4DFF: "Yijing Hexagram Symbols",
		0x9FFF: "CJK Unified Ideographs",
		0xA48F: "Yi Syllables",
		0xA4CF: "Yi Radicals",
		0xA71F: "Modifier Tone Letters",
		0xA7FF: "Latin Extended-D",
		0xA82F: "Syloti Nagri",
		0xA87F: "Phags-pa",
		0xD7AF: "Hangul Syllables",
		0xDB7F: "High Surrogates",
		0xDBFF: "High Private Use Surrogates",
		0xDFFF: "Low Surrogates",
		0xF8FF: "Private Use Area",
		0xFAFF: "CJK Compatibility Ideographs",
		0xFB4F: "Alphabetic Presentation Forms",
		0xFDFF: "Arabic Presentation Forms-A",
		0xFE0F: "Variation Selectors",
		0xFE1F: "Vertical Forms",
		0xFE2F: "Combining Half Marks",
		0xFE4F: "CJK Compatibility Forms",
		0xFE6F: "Small Form Variants",
		0xFEFF: "Arabic Presentation Forms-B",
		0xFFEF: "Halfwidth and Fullwidth Forms",
		0xFFFF: "Specials",
	}

	for key := range Blocks {
		orderedBlockKeys = append(orderedBlockKeys, key)
	}
	sort.Ints(orderedBlockKeys)
}

// GetBlockFromChar returns the block name that this char belongs to.
func GetBlockFromChar(character rune) string {
	i := sort.Search(len(orderedBlockKeys), func(i int) bool {
		return orderedBlockKeys[i] >= int(character)
	})
	return Blocks[orderedBlockKeys[i]]
}
