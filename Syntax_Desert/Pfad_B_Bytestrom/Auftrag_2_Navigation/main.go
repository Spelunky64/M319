package main

import "fmt"

func main() {
    bitFrequenz := 42
    windRichtung := "Süd"
    stromFarbe := "Blau"

    navigiereBytestrom(bitFrequenz, windRichtung, stromFarbe)
}

func navigiereBytestrom(frequenz int, wind string, farbe string) {
    fmt.Println("--- Analyse der Strom-Fatamorgana ---")

    if frequenz%2 != 0 {
        fmt.Println("Ergebnis: Der linke Strom ist eine Falle (Frequenzfehler)!")
    } else {
        if wind == "Nord" {
            fmt.Println("Ergebnis: Der mittlere Strom ist eine Sackgasse (Nordwind)!")
        } else {
            switch farbe {
            case "Blau":
                fmt.Println("Ergebnis: Pfad zu den Ruinen gefunden!")
            case "Rot":
                fmt.Println("Ergebnis: Vorsicht: Firewall-Alarm!")
            case "Gold":
                fmt.Println("Ergebnis: Weg zur Schatzkammer, aber nicht zu den Ruinen.")
            default:
                fmt.Println("Ergebnis: Der Strom ist instabil.")
            }
        }
    }
}
