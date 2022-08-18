package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
		fmt.Printf(`Usage: %s <word> [-m <meaning>] [-e <example>]`, filepath.Base(os.Args[0]))
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
