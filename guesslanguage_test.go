package guesslanguage

import "testing"

var parseTests = []struct {
	outCode string
	outErr  error
	in      string
}{
	{"en", nil, "This is a test of the language checker"},
	{"fr", nil, "Verifions que le détecteur de langues marche"},
	{"pl", nil, "Sprawdźmy, czy odgadywacz języków pracuje"},
	{"ru", nil, "Давайте проверим, распознает ли программа русский язык?"},
	{"es", nil, "La respuesta de los acreedores a la oferta argentina para salir del default no ha sido muy positiv"},
	{"kk", nil, "Сайлау Сайлау нәтижесінде дауыстардың басым бөлігін ел премьер министрі Виктор Янукович пен оның қарсыласы, оппозиция жетекшісі Виктор Ющенко алды."},
	{"uz", nil, "милиция ва уч солиқ идораси ходимлари яраланган. Шаҳарда хавфсизлик чоралари кучайтирилган."},
	{"ky", nil, "көрбөгөндөй элдик толкундоо болуп, Кокон шаарынын көчөлөрүндө бир нече миң киши нааразылык билдирди."},
	{"tr", nil, "yakın tarihin en çekişmeli başkanlık seçiminde oy verme işlemi sürerken, katılımda rekor bekleniyor."},
	{"az", nil, "Daxil olan xəbərlərdə deyilir ki, 6 nəfər Bağdadın mərkəzində yerləşən Təhsil Nazirliyinin binası yaxınlığında baş vermiş partlayış zamanı həlak olub."},
	{"ar", nil, " ملايين الناخبين الأمريكيين يدلون بأصواتهم وسط إقبال قياسي على انتخابات هي الأشد تنافسا منذ عقود"},
	{"uk", nil, "Американське суспільство, поділене суперечностями, збирається взяти активну участь у голосуванні"},
	{"cs", nil, "Francouzský ministr financí zmírnil výhrady vůči nízkým firemním daním v nových členských státech EU"},
	{"hr", nil, "biće prilično izjednačena, sugerišu najnovije ankete. Oba kandidata tvrde da su sposobni da dobiju rat protiv terorizma"},
	{"bg", nil, " е готов да даде гаранции, че няма да прави ядрено оръжие, ако му се разреши мирна атомна програма"},
	{"mk", nil, "на јавното мислење покажуваат дека трката е толку тесна, што се очекува двајцата соперници да ја прекршат традицијата и да се појават и на самиот изборен ден."},
	{"ro", nil, "în acest sens aparţinînd Adunării Generale a organizaţiei, în ciuda faptului că mai multe dintre solicitările organizaţiei privind organizarea scrutinului nu au fost soluţionate"},
	{"sq", nil, "kaluan ditën e fundit të fushatës në shtetet kryesore për të siguruar sa më shumë votues."},
	{"el", nil, "αναμένεται να σπάσουν παράδοση δεκαετιών και να συνεχίσουν την εκστρατεία τους ακόμη και τη μέρα των εκλογών"},
	{"zh", nil, " 美国各州选民今天开始正式投票。据信，"},
	{"nl", nil, " Die kritiek was volgens hem bitter hard nodig, omdat Nederland binnen een paar jaar in een soort Belfast zou dreigen te veranderen"},
	{"da", nil, "På denne side bringer vi billeder fra de mange forskellige forberedelser til arrangementet, efterhånden som vi får dem "},
	{"sv", nil, "Vi säger att Frälsningen är en gåva till alla, fritt och för intet.  Men som vi nämnt så finns det två villkor som måste"},
	{"fi", nil, "on julkishallinnon verkkopalveluiden yhteinen osoite. Kansalaisten arkielämää helpottavaa tietoa on koottu eri aihealueisiin"},
	{"et", nil, "Ennetamaks reisil ebameeldivaid vahejuhtumeid vii end kurssi reisidokumentide ja viisade reeglitega ning muu praktilise informatsiooniga"},
	{"hu", nil, "Hiába jön létre az önkéntes magyar haderő, hiába nem lesz többé bevonulás, változatlanul fennmarad a hadkötelezettség intézménye"},
	{"hy", nil, "հարաբերական"},
	{"vi", nil, "Hai vấn đề khó chịu với màn hình thường gặp nhất khi bạn dùng laptop là vết trầy xước và điểm chết. Sau đây là vài cách xử lý chú"},
	{"ja", nil, "トヨタ自動車、フィリピンの植林活動で第三者認証取得　トヨタ自動車(株)（以下、トヨタ）は、2007年９月よりフィリピンのルソン島北部に位置するカガヤン州ペニャブラン"},
	{"mn", nil, "ᠮᠤᠩᠭᠤᠯᠤᠯᠤᠰ"},
	{"pt", nil, "Pedras no caminho? Eu guardo todas. Um dia vou construir um castelo."},
	{"no", nil, "Tjenestene er svært varierte, og derfor kan også ytterligere vilkår eller produktkrav (herunder alderskrav) gjelde for hver enkelt tjeneste."},

	{"pt", nil, "ⒻⓇⒶⓈⒺ Ⓔⓜ ⓟⓞⓇ⒯⒰GⓊⒺⓈ"},
	{"ru", nil, "შეგიძლიათ рассказать рассказать "},

	{"", ErrUnknownLanguage, "1234567890"},
	{"", ErrUnknownLanguage, "           "},
	{"", ErrUnknownLanguage, ""},
	{"", ErrStringTooShort, "a"},
}

func TestParseFunction(t *testing.T) {
	for _, tt := range parseTests {
		lang, err := Parse(tt.in)
		if lang.ISOcode != tt.outCode || err != tt.outErr {
			t.Errorf("Parse(%q): have: (%q, %v) expected: (%q, %v)", tt.in, lang.ISOcode, err, tt.outCode, tt.outErr)
		}
	}
}
