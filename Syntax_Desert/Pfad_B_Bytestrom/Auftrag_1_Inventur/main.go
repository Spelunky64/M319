package main

import "fmt"

func main() {
    fmt.Println("=== Mercators Inventar-Kontrolle ===")

    var teppiche int = 15
    var gewuerzGewicht float64 = 2.5
    var handelsRoute string = "Küste"

    neueTeppiche := -20
    neueGewuerze := 1.2

    fmt.Printf("Aktueller Bestand: %d Teppiche, %.2f kg Gewürze\n", teppiche, gewuerzGewicht)

    if teppiche+neueTeppiche < 0 {
        fmt.Println("STOPP! Der Bestand an Teppichen darf nicht negativ werden!")
    } else {
        teppiche += neueTeppiche
        fmt.Println("Inventar erfolgreich aktualisiert.")
    }

    if neueGewuerze > 0 {
        gewuerzGewicht += neueGewuerze
        fmt.Println("Gewürze erfolgreich verbucht.")
    } else {
        fmt.Println("Ungültiges Gewicht")
    }

    fmt.Printf("Endbestand auf Route %s: %d Teppiche, %.2f kg Gewürze\n", handelsRoute, teppiche, gewuerzGewicht)
}
