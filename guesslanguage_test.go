package guesslanguage

import "testing"

const (
	// Phrases from internet.
	pt  = "Pedras no caminho? Eu guardo todas. Um dia vou construir um castelo."
	es  = "Los siguientes tutoriales te darán las pautas principales para comenzar a utilizar el nuevo sistema."
	ka  = "თქვენ შეგიძლიათ ისარგებლოთ რეგისტრაციის განახლებული ვებ გვერდით. Ge დომენების რეგისტრაცია, დომენური "
	de  = "Wir stellen für die Domainverwaltung ein automatisches elektronisches Registrierungssystem zur Verfügung und betreiben ein weltweites Netz von Nameservern, das sicherstellt, dass über 15 Millionen"
	ru  = "овинка! Благодаря сервису Google Мой бизнес вы можете бесплатно рассказать о себе клиентам с помощью"
	en  = "The easy way to start building Golang command line application."
	mn  = "ᠮᠤᠩᠭᠤᠯᠤᠯᠤᠰ"
	ja  = "できる限りわかりやすい説明を目指しておりますが、Cookie、IP アドレス、ピクセル タグ、ブラウザなどの用語がご不明の場合は、先にこれらの主な用語についての説明をご覧ください。Google ではお客様のプライバシーを重視しており"
	big = `It was a few minutes of eleven o’clock at night. One of the many editions of the great New York Herald had just gone to press. But in the big, half-lit room where editors, copy readers, reporters and telegraph operators were busy on the later editions to follow, there was no let-up in the work of making a world-known newspaper.
There was the noise of many persons working swiftly; the staccato of typewriters, the drone of telegraph sounders and now and then the sharp inquiry of some bent-over copy reader as he struggled to turn reportorial inexperience into a finished story. But there was no confusion and none of the wild rush and clatter that fiction uses in describing newspaper offices; copy boys were not dashing in all directions and the floor was not knee deep with newspapers and print paper.
Calmest of all was the night city editor. With a mind full of the work already done and in progress, he was as alert mentally as if he had just reached his desk. Five hours yet remained in which New York had to be watched; five hours, in any one minute of which the biggest news on hand might fade into nothing in the face of the one big story that every editor waits for night after night. And the night city editor, knowing this, dropped his half-lit pipe when his desk telephone buzzed.`
)

func TestCommonLanguages(t *testing.T) {

	if lang, _ := Parse(pt); lang.ISOcode != "pt" {
		t.Fatalf("Expected pt, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(es); lang.ISOcode != "es" {
		t.Fatalf("Expected es, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(ka); lang.ISOcode != "ka" {
		t.Fatalf("Expected ka, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(de); lang.ISOcode != "de" {
		t.Fatalf("Expected de, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(ru); lang.ISOcode != "ru" {
		t.Fatalf("Expected ru, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(en); lang.ISOcode != "en" {
		t.Fatalf("Expected en, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(mn); lang.ISOcode != "mn" {
		t.Fatalf("Expected mn, got %s.", lang.ISOcode)
	}
	if lang, _ := Parse(ja); lang.ISOcode != "ja" {
		t.Fatalf("Expected ja, got %s.", lang.ISOcode)
	}
}

func TestUnnormalizedChars(t *testing.T) {
	other := "ⒻⓇⒶⓈⒺ Ⓔⓜ ⓟⓞⓇ⒯⒰GⓊⒺⓈ"
	if lang, _ := Parse(other); lang.ISOcode != "pt" {
		t.Fatalf("Expected pt, got %s.", lang.ISOcode)
	}
}

func TestMixedChars(t *testing.T) {
	other := "ます შეგიძლიათ рассказать" //ru is bigger
	if lang, _ := Parse(other); lang.ISOcode != "ru" {
		t.Fatalf("Expected ru, got %s.", lang.ISOcode)
	}
}

func TestBigText(t *testing.T) {
	if lang, _ := Parse(big + big + big); lang.ISOcode != "en" {
		t.Fatalf("Expected en, got %s.", lang.ISOcode)
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
