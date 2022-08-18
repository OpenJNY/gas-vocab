package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var URL = os.Getenv("GAS_VOCAB_URL")

type Data struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	Example string `json:"example"`
}

func flagArgOrDefault(index int, fallback string) string {
	if index < flag.NArg() {
		return flag.Arg(index)
	}
	return fallback
}

func main() {
	var (
		meaning string
		example string
	)

	flag.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "Usage: %s [options] <word> [<meaning> [<example>]]\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(o, "\n")
		fmt.Fprintf(o, "This is a client to send a POST request to the 'Vocab' GAS application.\n")
		fmt.Fprintf(o, "If you specify meaning/example as arguments, they prefer over ones specified as options.\n")
		fmt.Fprintf(o, "[options]\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&meaning, "m", "", "meaning of the word")
	flag.StringVar(&example, "e", "", "example of the word")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	data := Data{
		Word:    flag.Arg(0),
		Meaning: flagArgOrDefault(1, meaning),
		Example: flagArgOrDefault(2, example),
	}

	marshalData, _ := json.Marshal(data)
	fmt.Printf("[!] %v\n", string(marshalData))

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
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("[!] %s\n", err)
	}
	fmt.Fprintln(os.Stdout, string(resBody))
}
