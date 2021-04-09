package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:        "cacao",
		Description: "Nice cacao",
		Price:       1.00,
		SKU:         "abs-def-adg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
