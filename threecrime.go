package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Criminal is, well, a criminal that may or may not be a perpetrator
type Criminal struct {
	name string
	perp bool
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	// ask how many players
	nPlayers := toInt(getInput("Enter number of players (1-3):"))
	if nPlayers < 1 {
		nPlayers = 1
	}
	if nPlayers > 3 {
		nPlayers = 3
	}
	nPlayersPlaying := nPlayers
	stillPlaying := make([]bool, nPlayers)
	for i := range stillPlaying {
		stillPlaying[i] = true
	}

	criminalNames := [7]string{"Capone", "Siegel", "Lansky",
		"Dillinger", "Giancana", "Anastasia", "Cohen"}
	// criminalNames := [7]string{"Al Capone", "Bugsy Siegel", "Meyer Lansky",
	//	"John Dillinger", "Sam Giancana", "Albert Anastasia", "Mickey Cohen"}
	criminals := make([]Criminal, 7, 7)
	for i, crimName := range criminalNames {
		criminals[i].name = crimName
	}
	// show the criminals to the player
	fmt.Printf("\nThese are the suspects:\n")
	for i := 0; i < len(criminals)-1; i++ {
		fmt.Printf("%s, ", criminals[i].name)
	}
	fmt.Printf("and %s.\n\n", criminals[len(criminals)-1].name)
	// choose the three perpetrators
	setPerp := func(c *Criminal) { c.perp = true }
	ChooseN(criminals, 3, 7, setPerp) // for 3 random unique criminals, set them to perpetrators

	// main game loop
	finished := false
	for !finished {
		// print three criminal names, no more than 2 of which are criminals
		nPerp := ChooseN(criminals, 3, 2, func(c *Criminal) { fmt.Printf("%s ", c.name) })
		// tell players how many of those are criminals
		fmt.Printf("\n%d of these are perpetrators.\n\n", nPerp)
		// let players guess or pass
		for p := 0; p < nPlayers; p++ {
			if !stillPlaying[p] { // if the player is out, skip ahead
				continue
			}
			fmt.Printf("Player %d, type 'pass' or 'guess':", p+1)
			pInput := getInput("")
			if pInput == "guess" || pInput == "g" {
				fmt.Println("Enter your guess as space-separated names")
				var g1, g2, g3 string
				fmt.Scanf("%s %s %s", &g1, &g2, &g3)
				// check for duplicates
				if g1 == g2 || g1 == g3 || g2 == g3 {
					fmt.Println("You guessed the same name twice. Nice try.")
					p-- // try again
				} else {
					// check the guess for accuracy (accurate = perp(1) and perp(2) and perp(3))
					accurate := selectByName(g1, criminals).perp &&
						selectByName(g2, criminals).perp &&
						selectByName(g3, criminals).perp
					if accurate {
						fmt.Printf("Player %d wins! The perpetrators were %s, %s, and %s!\n", p+1, g1, g2, g3)
						finished = true
					} else {
						// this player is out
						fmt.Printf("Sorry, player %d, you're out!\n", p+1)
						nPlayersPlaying--
						stillPlaying[p] = false
					}
				}
			} else if pInput == "pass" || pInput == "p" {
				// don't do anything, really
			} else {
				fmt.Printf("Invalid input %s\n", pInput)
				p-- // let this player try again
			}
		}
		// exit if everyone has already guessed wrong
		if nPlayersPlaying <= 0 {
			finished = true
			fmt.Println("All players are out, game over!")
		}
	}

}

// ChooseN performs the supplied function to n unique, randomly chosen criminals
// It returns the number of perpetrators among the criminals it chose
// Also limits the number of perpetrators chosen to maxPerp
func ChooseN(items []Criminal, n int, maxPerp int, set func(*Criminal)) int {
	if len(items) > 32 {
		log.Fatal("ChooseN called with too large a criminal list (> 32)! What did you do??\n")
	}
	var visited uint32
	nPerps := 0
	for n > 0 {
		i := uint32(rand.Intn(len(items)))
		if (visited & (1 << (i + 1))) == 0 { // if not yet visited
			if !items[i].perp || nPerps < maxPerp {
				n--
				set(&items[i])
				if items[i].perp {
					nPerps++
				}
			}
		}
		visited = visited | (1 << (i + 1))
	}
	return nPerps
}

// find the criminal with the supplied name, or return a non-existent criminal if not found
// note that this means a typo will mean you're out
func selectByName(name string, c []Criminal) *Criminal {
	for i := range c {
		if c[i].name == name {
			return &c[i]
		}
	}
	log.Printf("%s not found!\n", name)
	return &Criminal{name: "", perp: false}
}

// display a prompt, then get input from stdin
func getInput(prompt string) string {
	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	return strings.TrimRight(nStr, "\r\n")
}

// just a helper to convert a string to an int
func toInt(val string) int {
	n64, err := strconv.ParseInt(val, 10, 32)
	n := int(n64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}
