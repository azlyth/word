package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bradfitz/iter"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "word"
	app.Version = Version
	app.Usage = "generates random words"
	app.Author = "Peter Valdez"
	app.Email = "peter@nycmesh.net"

	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "number, n",
			Usage: "how many words to generate",
			Value: 1,
		},
	}
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
	// Get the length of the word list and create a random generator with the time as the seed
	totalWords := len(Words)
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Print the requested number of randomly selected words
	var chosen []string
	for range iter.N(c.Int("number")) {
		chosen = append(chosen, Words[gen.Intn(totalWords)])
	}
	fmt.Println(strings.Join(chosen, " "))
}
