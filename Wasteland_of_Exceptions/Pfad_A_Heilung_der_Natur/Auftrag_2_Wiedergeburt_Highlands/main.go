package main

import (
	"fmt"
)

func main() {
	// --- 1. SYNTAXFEHLER ---
	// FIX: Missing closing parenthesis
	fmt.Println("Initialisiere Heilungs-Protokoll...")
    
	// --- 2. LAUFZEITFEHLER ---
	var heilungsKristall *int
	
	// FIX: Nil-Prüfung vor Dereferencierung
	if heilungsKristall != nil {
		fmt.Printf("Energie-Level des Kristalls: %d\n", *heilungsKristall)
	} else {
		fmt.Println("Energie-Level des Kristalls: [nil]")
	}

	// --- 3. LOGIKFEHLER ---
	energieFluss := 40
	istHeilungErfolgreich := false

	// FIX: < 50 zu >= 50 ändern
	if energieFluss >= 50 {
		istHeilungErfolgreich = true
	}

	if istHeilungErfolgreich {
		fmt.Println("✅ Die Aura breitet sich aus! Die Highlands erblühen.")
	} else {
		fmt.Println("❌ Die Energie reicht nicht aus. Die Highlands bleiben grau.")
	}

	// --- 4. SCHLEIFEN-FEHLER ---
	baeumeGepflanzt := 0
	fmt.Println("Beginne Aufforstung...")

	// FIX: > 5 zu < 5 ändern (damit Schleife startet)
	for baeumeGepflanzt < 5 {
		baeumeGepflanzt++
		fmt.Printf("Baum Nr. %d gepflanzt.\n", baeumeGepflanzt)
	}

	fmt.Println("Mission abgeschlossen: Der Weltencode ist rein!")
}
