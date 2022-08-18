package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var URL = os.Getenv("GAS_VOCAB_URL")

type Data struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Example string `json:"example"`
}

func main() {
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		fmt.Printf(`Usage: vocab <word> [-m <meaning>] [-e <example>]
		
# alias
vocab <word> <meaning>
vocab <word> <meaning> <example>
`)
		os.Exit(1)
	}

	var (
		meaning string
		example string
	)

	flag.StringVar(&meaning, "m", "", "meaning of the word")
	flag.StringVar(&example, "e", "", "example of the word")
	flag.Parse()

	data := Data{
		Word:    os.Args[1],
		Meaning: meaning,
		Example: example,
	}

	// vocab <word> <meaning>
	if len(os.Args) >= 3 {
		data.Meaning = os.Args[2]
	}
	// vocab <word> <meaning> <example>
	if len(os.Args) >= 4 {
		data.Meaning = os.Args[3]
	}

	marshalData, _ := json.Marshal(data)

	res, err := http.Post(
		URL,
		"application/json",
		bytes.NewBuffer(marshalData),
	)
	if err != nil {
		fmt.Printf("[!] GAS_VOCAB_URL=%s\n", URL)
		fmt.Printf("[!] %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalf("[!] %s\n", err)
	// }
	// fmt.Fprintln(os.Stdout, string(body))
	fmt.Printf("registerd: %v\n", string(marshalData))
}
