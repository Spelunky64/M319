package main

import "fmt"

type artefakt struct {
	name    string
	energie int
}

func (a artefakt) GetEnergie() int {
	return a.energie
}

func (a *artefakt) SetEnergie(neuerWert int) {
	if neuerWert >= 0 {
		a.energie = neuerWert
	} else {
		fmt.Println("Warnung: Instabile Energiequelle! Wert wurde nicht geändert.")
	}
}

func main() {
	meinSchatz := artefakt{
		name:    "Zepter der Ordnung",
		energie: 80,
	}

	fmt.Println("Artefakt gefunden:", meinSchatz.name)

	meinSchatz.SetEnergie(120)

	fmt.Println("Aktuelle Energie:", meinSchatz.GetEnergie())

	fmt.Println("Versuche, Energie auf einen gefährlichen Wert (-50) zu setzen...")
	meinSchatz.SetEnergie(-50)
	
	fmt.Println("Finale Energie:", meinSchatz.GetEnergie())
}
