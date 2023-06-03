package main

import "fmt"

func main() {
    cards := []string{"A", "B", "C"}
    cards = append(cards, "D")

    for _, card := range cards {
        fmt.Println(card)
    }
}

func newCards() string {
    return "Ace of Spades"
}