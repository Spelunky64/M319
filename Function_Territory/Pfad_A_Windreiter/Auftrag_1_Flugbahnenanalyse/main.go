package main

import "fmt"

func berechneBasisAuftrieb(windgeschwindigkeit int, anflugwinkel int) int {
	ergebnis := windgeschwindigkeit + anflugwinkel
	return ergebnis
}

func berechneEnergieVerbrauch(gewicht int, strecke int) int {
	ergebnis := (gewicht * strecke) / 10
	return ergebnis
}

func checkFlugStatus(auftrieb int, verbrauch int) (int, string) {
	differenz := auftrieb - verbrauch
	status := ""
	if differenz > 0 {
		status = "Stabil"
	} else {
		status = "Absturzgefahr"
	}
	return differenz, status
}

func kreischen() {
	fmt.Println("KKKRREEEIIIISSSCCCHHHHH")
}

func main() {
	fmt.Println("--- Analyse der Variable-Falken startet ---")
	speed := 150
	winkel := 50
	aktuellerAuftrieb := berechneBasisAuftrieb(speed, winkel)
	fmt.Printf("Berechneter Auftrieb: %d Einheiten\n", aktuellerAuftrieb)
	gewicht := 5
	distanz := 10
	verbrauch := berechneEnergieVerbrauch(gewicht, distanz)
	fmt.Printf("Voraussichtlicher Verbrauch: %d Einheiten\n", verbrauch)
	kraftReserve, flugZustand := checkFlugStatus(aktuellerAuftrieb, verbrauch)
	fmt.Println("-------------------------------------------")
	fmt.Printf("Analyse-Ergebnis: %s (Reserve: %d)\n", flugZustand, kraftReserve)
	if flugZustand == "Stabil" && kraftReserve > 0 {
		fmt.Println("Mission abgeschlossen: Der Falke hält seine Bahn!")
		kreischen()
		kreischen()
		kreischen()
	} else {
		fmt.Println("Fehler: Die mathematische Kapselung ist noch instabil.")
	}
}
