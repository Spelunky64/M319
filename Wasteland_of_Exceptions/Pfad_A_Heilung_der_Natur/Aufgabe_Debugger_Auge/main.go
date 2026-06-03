package main

import (
	"fmt"
)

func main() {
	fmt.Println("Beschwörung des Debugger-Auges gestartet...")

	var DebuggerAuge *string 
	
	if DebuggerAuge != nil {
		fmt.Println("Aktiviere Kraft von:", *DebuggerAuge)
	} else {
		fmt.Println("Aktiviere Kraft von: [nil]")
	}

	energieQuelle := 100
	aktiveGatter := 0
	
	if aktiveGatter != 0 {
		lastProGatter := energieQuelle / aktiveGatter
		fmt.Println("Lastverteilung:", lastProGatter)
	} else {
		fmt.Println("Fehler: Keine aktiven Gatter!")
	}

	fmt.Println("Aktiviere Schutzsiegel...")
	siegelZaehler := 1
	
	for siegelZaehler <= 3 {
		fmt.Printf("Siegel Nr. %d fixiert.\n", siegelZaehler)
		siegelZaehler++
	}

	fmt.Println("Die Kathedrale ist gereinigt!")
}
