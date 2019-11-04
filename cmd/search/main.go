package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/opencars/wanted/pkg/wanted"
)

func main() {
	var path, value, searchBy string

	flag.StringVar(&path, "path", "", "Path to the data file")
	flag.StringVar(&value, "value", "", "")
	flag.StringVar(&searchBy, "search-by", "BODYNUMBER", "Name of the field to seach by")

	// for _, vehicle := range  {
	// 	kek := reflect.ValueOf(&vehicle).Elem()
	// 	res := kek.FieldByName(tag).String()

	// 	if strings.Contains(res, value) {
	// 		vehicles = append(vehicles, vehicle)
	// 	}
	// // }

	// var v wanted.Transport

	// kek := reflect.ValueOf(&v).Elem()
	// for i := 0; i < kek.NumField(); i++ {
	// 	kek.FieldByIndex(i).
	// }

	// flag.Parse()

	flag.Parse()

	if filepath.Ext(path) != ".json" {
		fmt.Fprintln(os.Stderr, "invalid file extension")
		os.Exit(1)
	}

	transport, err := wanted.ParseFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	res := transport.SearchBy(searchBy, value)

	// for i := len(res) - 1; i >= 0; i-- {
	// 	fmt.Println(res[i].PrettyString())
	// }

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].PrettyString())
	}

	fmt.Println("Total:", len(res))
}
