package languageinfo

type (
	LanguageInfo struct {
		basicLatin     []string
		extendedLatin  []string
		latin          []string
		cyrillic       []string
		arabic         []string
		devanagari     []string
		languageMinLen int
	}
)

func New() *LanguageInfo {
	
	lang := LanguageInfo{}
	lang.languageMinLen = 20
	lang.basicLatin = []string{"en", "ceb", "ha", "so", "tlh", "id", "haw",
		"la", "sw", "eu", "nr", "nso", "zu", "xh", "ss", "st", "tn", "ts"}

	lang.extendedLatin = []string{"cs", "af", "pl", "hr", "ro", "sk", "sl",
		"tr", "hu", "az", "et", "sq", "ca", "es", "fr", "de", "nl", "it", "da",
		"is", "nb", "sv", "fi", "lv", "pt", "ve", "lt", "tl", "cy"}

	lang.latin = append(lang.basicLatin, lang.extendedLatin...)
	lang.cyrillic = []string{"ru", "uk", "kk", "uz", "mn", "sr", "mk", "bg", "ky"}
	lang.arabic = []string{"ar", "fa", "ps", "ur"}
	lang.devanagari = []string{"hi", "ne"}
	return &lang
}
