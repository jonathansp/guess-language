package guesslanguage

import "testing"

const (
	// Phrases from internet.
	pt = "Pedras no caminho? Eu guardo todas. Um dia vou construir um castelo."
	es = "Los siguientes tutoriales te darán las pautas principales para comenzar a utilizar el nuevo sistema."
	ka = "თქვენ შეგიძლიათ ისარგებლოთ რეგისტრაციის განახლებული ვებ გვერდით. Ge დომენების რეგისტრაცია, დომენური "
	de = "Wir stellen für die Domainverwaltung ein automatisches elektronisches Registrierungssystem zur Verfügung und betreiben ein weltweites Netz von Nameservern, das sicherstellt, dass über 15 Millionen"
	ru = "овинка! Благодаря сервису Google Мой бизнес вы можете бесплатно рассказать о себе клиентам с помощью"
	en = "The easy way to start building Golang command line application."
	mn = "ᠮᠤᠩᠭᠤᠯᠤᠯᠤᠰ"
	ja = "できる限りわかりやすい説明を目指しておりますが、Cookie、IP アドレス、ピクセル タグ、ブラウザなどの用語がご不明の場合は、先にこれらの主な用語についての説明をご覧ください。Google ではお客様のプライバシーを重視しており"
)

func TestCommonLanguages(t *testing.T) {
	g := GuessLanguage()
	if lang := g.Parse(pt); lang.ISOcode != "pt" {
		t.Fatalf("Expected pt, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(es); lang.ISOcode != "es" {
		t.Fatalf("Expected es, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(ka); lang.ISOcode != "ka" {
		t.Fatalf("Expected ka, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(de); lang.ISOcode != "de" {
		t.Fatalf("Expected de, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(ru); lang.ISOcode != "ru" {
		t.Fatalf("Expected ru, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(en); lang.ISOcode != "en" {
		t.Fatalf("Expected en, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(mn); lang.ISOcode != "mn" {
		t.Fatalf("Expected mn, got %s.", lang.ISOcode)
	}
	if lang := g.Parse(ja); lang.ISOcode != "ja" {
		t.Fatalf("Expected ja, got %s.", lang.ISOcode)
	}
}

func TestNonNormalizedChars(t *testing.T) {
	g := GuessLanguage()
	other := "ⒻⓇⒶⓈⒺ Ⓔⓜ ⓟⓞⓇ⒯⒰GⓊⒺⓈ"
	if lang := g.Parse(other); lang.ISOcode != "pt" {
		t.Fatalf("Expected pt, got %s.", lang.ISOcode)
	}
}

func TestMixedChars(t *testing.T) {
	g := GuessLanguage()
	other := "ます შეგიძლიათ рассказать" //ru is bigger
	if lang := g.Parse(other); lang.ISOcode != "ru" {
		t.Fatalf("Expected ru, got %s.", lang.ISOcode)
	}
}

func TestShortAndEmptyText(t *testing.T) {
	g := GuessLanguage()
	short := "sht" // impossible to identify
	if lang := g.Parse(short); lang.ISOcode != "" {
		t.Fatalf("Expected empty, got %s.", lang.ISOcode)
	}
	var empty string
	if lang := g.Parse(empty); lang.ISOcode != "" {
		t.Fatalf("Expected empty, got %s.", lang.ISOcode)
	}
}
