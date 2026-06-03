package main

import "fmt"

type Artefakt struct {
	Name      string
	Wert      int
	EnergieTyp string
	IstAktiv  bool
}

func (a Artefakt) Info() {
	fmt.Printf("Das Artefakt %s leuchtet mit %s-Energie und hat einen Wert von %d Goldstücken.\n", a.Name, a.EnergieTyp, a.Wert)
}

func main() {
	kristall := Artefakt{
		Name:       "Navigations-Kristall",
		Wert:       150,
		EnergieTyp: "Luminanz",
		IstAktiv:   true,
	}

	kristall.Info()
}
