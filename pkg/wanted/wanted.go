package wanted

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

func ParseFile(name string) (Transport, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return Parse(file)
}

func Parse(reader io.Reader) (Transport, error) {
	vehicles := make([]*Vehicle, 0, 80000)

	dec := json.NewDecoder(reader)

	if _, err := dec.Token(); err != nil {
		return nil, err
	}

	for dec.More() {
		tmp := &Vehicle{}
		err := dec.Decode(&tmp)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// Replace with binary search.
		vehicles = append(vehicles, tmp)
	}

	transport := Transport(vehicles)
	sort.Sort(transport)

	return transport, nil
}

func (transport Transport) SearchByFirst(tag, value string) *Vehicle {
	for _, vehicle := range transport {
		kek := reflect.ValueOf(&vehicle).Elem()
		res := kek.FieldByName(tag).String()

		if strings.Contains(res, value) {
			result := vehicle
			return result
		}
	}

	return nil
}
