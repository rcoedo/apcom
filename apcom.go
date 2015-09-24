package main

import "fmt"
import "math/rand"
import "time"

var quotes = []string{
	"Chaos will reign",
	"One step closer to the end of the world",
	"The Sun shall be turned into darkness",
	"The whole land will be a burning waste of salt and sulfur; nothing planted, nothing sprouting, no vegetation growing on it.",
	"There the fire will consume you, the sword will cut you down; they will devour you like a swarm of locusts.",
	"All these are the beginning of sorrows",
	"And many false prophets will appear and will deceive many people.",
	"Not even death can save you",
	"Beware! Beyond lies mortal danger for the likes of you!",
	"This holy place has been desecrated",
	"Turn back! I can tell that you need more experience to fight safely in the next wilderness.",

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf(quotes[rand.Intn(len(quotes))] + "\n")
}
