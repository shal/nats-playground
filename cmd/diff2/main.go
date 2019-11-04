package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/opencars/wanted/pkg/wanted"
)

func main() {
	var config string

	flag.StringVar(&config, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	// oldPath := flag.Arg(0)
	newPath := flag.Arg(1)

	if filepath.Ext(newPath) != ".json" {
		fmt.Fprintln(os.Stderr, "invalid file extension")
		os.Exit(1)
	}

	// transport, err := wanted.ParseFile()
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// TODO: SQL select all current cars.

	f, err := os.Open(newPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dec := json.NewDecoder(f)

	// Read the array open bracket.
	if _, err := dec.Token(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Amount of vehicle in %s: %d\n", oldPath, len(transport))

	newTransport := make([]wanted.Vehicle, 0, 100)
	for dec.More() {
		var tmp wanted.Vehicle
		err := dec.Decode(&tmp)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if res := transport.Search(&tmp); res == -1 {
			newTransport = append(newTransport, tmp)
		}
	}

	// TODO: Find transport that was removed.

	// New stolen cars.
	fmt.Println("Total:", len(newTransport))
}
