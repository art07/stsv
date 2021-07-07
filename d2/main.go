package main

import (
	"log"
	"os"
)

func main() {
	// logger to Inject
	logger := log.New(os.Stdout, "INFO: ", 0)

	// rest of the code
	var dec Decoder
	dec = MyDecoder{
		logger: logger,
	}
	persons, places := dec.Decode(jsonStr)
	dec.Printlen(persons, places)
	dec.Sort(persons)
	dec.Sort(places)
	dec.Print(persons)
	dec.Print(places)
}
