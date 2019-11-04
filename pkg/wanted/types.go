package wanted

import (
	"fmt"
	"strings"
)

type Vehicle struct {
	ID            string `json:"ID"`
	OVD           string `json:"OVD"`
	Brand         string `json:"BRAND"`
	Color         string `json:"COLOR"`
	VehicleNumber string `json:"VEHICLENUMBER"`
	BodyNumber    string `json:"BODYNUMBER"`
	ChassisNumber string `json:"CHASSISNUMBER"`
	EngineNumber  string `json:"ENGINENUMBER"`
	TheftData     string `json:"THEFT_DATA"`
	InsertDate    string `json:"INSERT_DATE"`
}

type Transport []*Vehicle

// Len is the number of elements in the collection.
func (t Transport) Len() int {
	return len(t)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (t Transport) Less(i, j int) bool {
	return t[i].ID <= t[j].ID
}

// Swap swaps the elements with indexes i and j.
func (t Transport) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Transport) Search(v *Vehicle) int {
	pivot := len(t) / 2
	return t.search(pivot, v)
}

func (t Transport) search(i int, v *Vehicle) int {
	if len(t) == 0 {
		return -1
	}

	if v.ID == t[i].ID {
		return i
	}

	if v.ID < t[i].ID {
		return Transport(t[:i]).search(len(t[:i])/2, v)
	} else {
		return Transport(t[i+1:]).search(len(t[i+1:])/2, v)
	}
}

func (v *Vehicle) PrettyString() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "ID: %s\n", v.ID)
	fmt.Fprintf(&builder, "OVD: %s\n", v.OVD)
	fmt.Fprintf(&builder, "Brand: %s\n", v.Brand)
	fmt.Fprintf(&builder, "Color: %s\n", v.Color)
	fmt.Fprintf(&builder, "VehicleNumber: %s\n", v.VehicleNumber)
	fmt.Fprintf(&builder, "BodyNumber: %s\n", v.BodyNumber)
	fmt.Fprintf(&builder, "ChassisNumber: %s\n", v.ChassisNumber)
	fmt.Fprintf(&builder, "EngineNumber: %s\n", v.EngineNumber)
	fmt.Fprintf(&builder, "TheftData: %s\n", v.TheftData)
	fmt.Fprintf(&builder, "InsertDate: %s\n", v.InsertDate)

	return builder.String()
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	  runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
  }

func parseKind(lexeme string) (string, string) {
	// reversed := reverse(lexeme) 

	stack := make([]rune, 0)
	// kind := make([]rune, 0)

	for i := len(lexeme) - 1; i >0; i-- {
		switch lexeme[i] {
		case ')':
			stack = append(stack, ')')
		case '(':
			stack = stack[:len(stack) - 1]
		}

		if len(stack) == 0 {
			return lexeme[i:], lexeme[:i]
		}
	}

	return "", ""
}


func (v *Vehicle) ParseBrand() (string, string, string) {
	kind, other := parseKind(v.Brand)
	
	kind = strings.TrimSpace(kind)

	if len(kind) > 2 {
		kind = strings.TrimPrefix(kind,"(")
		kind = strings.TrimSuffix(kind, ")")
	}

	fmt.Println(kind, other)

	brandModel := strings.TrimSpace(other)

	partsbyDash := strings.SplitN(brandModel, " - ", 2)
	if len(partsbyDash) == 2 {
		return strings.TrimSpace(partsbyDash[0]), strings.TrimSpace(partsbyDash[1]), kind
	}

	partsbySpace := strings.SplitN(brandModel, " ", 2)
	if len(partsbySpace) == 2 {
		return strings.TrimSpace(partsbySpace[0]), strings.TrimSpace(partsbySpace[1]), kind
	}

	// TODO: Select vehicle by number from database of operations.

	// TODO: Select vehicle by VIN from the database of registrations.

	return brandModel, "", kind
}
