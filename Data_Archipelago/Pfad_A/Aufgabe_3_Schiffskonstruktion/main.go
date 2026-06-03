package main

import "fmt"

type antrieb struct {
	schubkraft int
}

func (a antrieb) Beschleunigen() {
	fmt.Printf("Der Antrieb zündet mit %d Newton Schub!\n", a.schubkraft)
}

type digitalschiff struct {
	antrieb
	name   string
	modell string
}

func NewDigitalschiff(name string, schub int) digitalschiff {
	if schub < 10 {
		schub = 10
	}

	return digitalschiff{
		antrieb: antrieb{
			schubkraft: schub,
		},
		name:   name,
		modell: "Mark-I",
	}
}

func main() {
	fmt.Println("=== WERFT-PROTOKOLL: AKTIV ===")

	meinSchiff := NewDigitalschiff("WaveRunner", 50)

	meinSchiff.Beschleunigen()

	fmt.Printf("Das Schiff '%s' ist bereit für die Reise!\n", meinSchiff.name)

	fmt.Println("=== KONSTRUKTION ABGESCHLOSSEN ===")
}
