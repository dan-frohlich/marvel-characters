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
	fileFlag      = flag.String("f", "character.json", "REQUIRED path to character json file")
	karmaLogFlag  = flag.Bool("k", false, "print karma log")
	charLogFlag   = flag.Bool("c", false, "print creation log")
	popLogFlag    = flag.Bool("p", false, "print popularity log")
	outFormatFlag = flag.String("o", "text", "output format: [text|pdf]")
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

		switch *outFormatFlag {
		case "text":
			fmt.Println(string(characters.AsciiCharacterSheet(c, *karmaLogFlag, *charLogFlag, *popLogFlag)))
		case "pdf":
			fallthrough
		default:
			fmt.Printf("output format [%s] not implemented.\n", *outFormatFlag)
			flag.Usage()
			os.Exit(1)
		}
		return
	}
	flag.Usage()
	os.Exit(1)
}
