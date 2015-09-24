package main

import "fmt"
import "math/rand"
import "time"

var quotes = []string{
	"Chaos will reign",
	"One step closer to the end of the world",
	"The Sun shall be turned into darkness",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf(quotes[rand.Intn(len(quotes))] + "\n")
}
