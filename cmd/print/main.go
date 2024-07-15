package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	characters "github.com/dan-frohlich/marvel-characters"
)

var (
	fileFlag = flag.String("f", "character.json", "path to character json file")
)

func main() {
	flag.Parse()
	if fileFlag != nil && *fileFlag != "" {
		c := characters.Character{}
		b, err := os.ReadFile(*fileFlag)
		if err != nil {
			log.Printf("could not read %s : %s", *fileFlag, err)
			flag.Usage()
			os.Exit(1)
		}
		err = json.Unmarshal(b, &c)
		if err != nil {
			log.Printf("could not read %s : %s", *fileFlag, err)
			flag.Usage()
			os.Exit(1)
		}

		fmt.Println(string(characters.AsciiCharacterSheet(c)))
		return
	}
	flag.Usage()
	os.Exit(1)
}
