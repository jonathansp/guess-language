package guesslanguage

import "testing"

var sentences = map[string]string{
	"en": "This is a test of the language checker",
	"fr": "Verifions que le détecteur de langues marche",
	"pl": "Sprawdźmy, czy odgadywacz języków pracuje",
	"ru": "авай проверить  узнает ли наш угадатель русски язык",
	"es": "La respuesta de los acreedores a la oferta argentina para salir del default no ha sido muy positiv",
	"kk": "Сайлау Сайлау нәтижесінде дауыстардың басым бөлігін ел премьер министрі Виктор Янукович пен оның қарсыласы, оппозиция жетекшісі Виктор Ющенко алды.",
	"uz": "милиция ва уч солиқ идораси ходимлари яраланган. Шаҳарда хавфсизлик чоралари кучайтирилган.",
	"ky": "көрбөгөндөй элдик толкундоо болуп, Кокон шаарынын көчөлөрүндө бир нече миң киши нааразылык билдирди.",
	"tr": "yakın tarihin en çekişmeli başkanlık seçiminde oy verme işlemi sürerken, katılımda rekor bekleniyor.",
	"az": "Daxil olan xəbərlərdə deyilir ki, 6 nəfər Bağdadın mərkəzində yerləşən Təhsil Nazirliyinin binası yaxınlığında baş vermiş partlayış zamanı həlak olub.",
	"ar": " ملايين الناخبين الأمريكيين يدلون بأصواتهم وسط إقبال قياسي على انتخابات هي الأشد تنافسا منذ عقود",
	"uk": "Американське суспільство, поділене суперечностями, збирається взяти активну участь у голосуванні",
	"cs": "Francouzský ministr financí zmírnil výhrady vůči nízkým firemním daním v nových členských státech EU",
	"hr": "biće prilično izjednačena, sugerišu najnovije ankete. Oba kandidata tvrde da su sposobni da dobiju rat protiv terorizma",
	"bg": " е готов да даде гаранции, че няма да прави ядрено оръжие, ако му се разреши мирна атомна програма",
	"mk": "на јавното мислење покажуваат дека трката е толку тесна, што се очекува двајцата соперници да ја прекршат традицијата и да се појават и на самиот изборен ден.",
	"ro": "în acest sens aparţinînd Adunării Generale a organizaţiei, în ciuda faptului că mai multe dintre solicitările organizaţiei privind organizarea scrutinului nu au fost soluţionate",
	"sq": "kaluan ditën e fundit të fushatës në shtetet kryesore për të siguruar sa më shumë votues.",
	"el": "αναμένεται να σπάσουν παράδοση δεκαετιών και να συνεχίσουν την εκστρατεία τους ακόμη και τη μέρα των εκλογών",
	"zh": " 美国各州选民今天开始正式投票。据信，",
	"nl": " Die kritiek was volgens hem bitter hard nodig, omdat Nederland binnen een paar jaar in een soort Belfast zou dreigen te veranderen",
	"da": "På denne side bringer vi billeder fra de mange forskellige forberedelser til arrangementet, efterhånden som vi får dem ",
	"sv": "Vi säger att Frälsningen är en gåva till alla, fritt och för intet.  Men som vi nämnt så finns det två villkor som måste",
	"fi": "on julkishallinnon verkkopalveluiden yhteinen osoite. Kansalaisten arkielämää helpottavaa tietoa on koottu eri aihealueisiin",
	"et": "Ennetamaks reisil ebameeldivaid vahejuhtumeid vii end kurssi reisidokumentide ja viisade reeglitega ning muu praktilise informatsiooniga",
	"hu": "Hiába jön létre az önkéntes magyar haderő, hiába nem lesz többé bevonulás, változatlanul fennmarad a hadkötelezettség intézménye",
	"hy": "հարաբերական",
	"vi": "Hai vấn đề khó chịu với màn hình thường gặp nhất khi bạn dùng laptop là vết trầy xước và điểm chết. Sau đây là vài cách xử lý chú",
	"ja": "トヨタ自動車、フィリピンの植林活動で第三者認証取得　トヨタ自動車(株)（以下、トヨタ）は、2007年９月よりフィリピンのルソン島北部に位置するカガヤン州ペニャブラン",
	"mn": "ᠮᠤᠩᠭᠤᠯᠤᠯᠤᠰ",
	"pt": "Pedras no caminho? Eu guardo todas. Um dia vou construir um castelo.",
	"no": "Tjenestene er svært varierte, og derfor kan også ytterligere vilkår eller produktkrav (herunder alderskrav) gjelde for hver enkelt tjeneste.",
}

func TestOneLanguage(t *testing.T) {
	code := "pt"
	sentence := sentences[code]
	if lang, err := Parse(sentence); lang.ISOcode != code {
		t.Fatalf("Expected %s, got %s. Sentence: %s\n%s", code, lang.ISOcode, sentence, err)
	}

}

func TestCommonLanguages(t *testing.T) {
	for code, sentence := range sentences {
		if lang, _ := Parse(sentence); lang.ISOcode != code {
			t.Fatalf("Expected %s, got %s. Sentence: %s", code, lang.ISOcode, sentence)
		}
	}
}

func TestUnnormalizedChars(t *testing.T) {
	other := "ⒻⓇⒶⓈⒺ Ⓔⓜ ⓟⓞⓇ⒯⒰GⓊⒺⓈ"
	if lang, _ := Parse(other); lang.ISOcode != "pt" {
		t.Fatalf("Expected pt, got %s.", lang.ISOcode)
	}
}

func TestMixedChars(t *testing.T) {
	other := "შეგიძლიათ рассказать рассказать "
	if lang, _ := Parse(other); lang.ISOcode != "ru" {
		t.Fatalf("Expected ru, got %s.", lang.ISOcode)
	}
}

func TestNumbersAndSpaces(t *testing.T) {
	numbers := "1234567890"
	if _, err := Parse(numbers); err == nil {
		t.Fatalf("Expected err, got nil.")
	}
	spaces := "           "
	if _, err := Parse(spaces); err == nil {
		t.Fatalf("Expected err, got nil.")
	}

}

func TestShortAndEmptyText(t *testing.T) {
	// impossible to identify
	short := "a"
	if _, err := Parse(short); err == nil {
		t.Fatalf("Expected err, got nil.")
	}
	var empty string
	if _, err := Parse(empty); err == nil {
		t.Fatalf("Expected err, got nil.")
	}
}
