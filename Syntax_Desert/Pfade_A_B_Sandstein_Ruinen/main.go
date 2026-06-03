// definiert das Paket für dieses Programm
package main

// importiert das fmt Paket für formatierte Ausgabe
import "fmt"

// =================================================================
// ARCHITEKTUR-BEREICH (Fortgeschrittene Magie - Bitte nicht ändern)
// =================================================================

// Fahrzeug ist die Basisklasse (Struktur) für alle Fortbewegungsmittel
type Fahrzeug struct {
	Name string
}

// Barke erbt Eigenschaften vom Fahrzeug und erweitert diese
type Barke struct {
	Fahrzeug
	Tragkraft float64
	Passagiere int
}

// Methode zum Materialisieren der Barke
func (b Barke) beschwoere() {
	fmt.Printf("Die Barke '%s' beginnt zu schimmern...\n", b.Name)
	for i := 0; i < 3; i++ {
		fmt.Println("... [DATA-STREAM-STABILIZING] ...")
	}
}

// =================================================================
// IHRE AUFGABE: DIE LOGIK-WERKSTATT
// =================================================================

// Die Hauptfunktion führt die Beschwörung der Logik-Barke durch
func main() {
	fmt.Println("=== Die Beschwörung der Logik-Barke ===")

	// AUFGABE 1: Variablen-Initialisierung
	var tragkraft float64 = 500.5
	var anzahlPassagiere int = 5
	
	// float64 wird verwendet, da die Tragkraft dezimale Werte (z.B. 500.5 kg) darstellen kann
	
	// FIXME: Bedingung überprüft Stabilität
	istStabil := false
      
	if anzahlPassagiere > 6 || tragkraft < 400 {
		fmt.Println("[WARNUNG]: Die Logik-Parameter sind instabil!")
		istStabil = false
	} else {
		fmt.Println("[CHECK]: Auftriebs-Logik stabil.")
		istStabil = true
	}

	fmt.Println("---------------------------------------")

	// AUFGABE 2: Navigationstyp festlegen
	switch anzahlPassagiere {
	case 1, 2, 3:
		fmt.Println("Modus: Kleines Boot")
	case 4, 5, 6:
		fmt.Println("Modus: Gruppen-Barke")
	default:
		fmt.Println("Modus: Unbekannt")
	}

	/* AUFGABE 3: Das Switch-Statement prüft die Anzahl der Passagiere
	   und weist der Barke entsprechend einen Navigationsmodus zu.
	   Dies bestimmt, welche Route und Geschwindigkeit optimal ist. */

	// FINALER SCHRITT: Wenn alles stabil ist, wird die Barke erzeugt
	if istStabil {
		meineBarke := Barke{
			Fahrzeug:   Fahrzeug{Name: "Lumen Logica"},
			Tragkraft:  tragkraft,
			Passagiere: anzahlPassagiere,
		}

		meineBarke.beschwoere()
		fmt.Println("Erfolg: Die Barke ist bereit für das Bytestromdelta!")
	} else {
		fmt.Println("Fehler: Die Beschwörung ist fehlgeschlagen.")
	}
}
