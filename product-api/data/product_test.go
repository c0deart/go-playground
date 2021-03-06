package data

import (
	"fmt"
	"testing"
)

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

func TestArrayRemoval(t *testing.T) {
	productList := []string{"a", "b", "c"}
	fmt.Println(productList)

	removeIndex := 1
	//updatedArr := productList[:removeIndex]
	//fmt.Println(updatedArr)

	productList = append(productList[:removeIndex], productList[removeIndex+1])
	fmt.Println(productList)

}
