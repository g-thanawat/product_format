package main

import (
	"fmt"
	"strings"

	"github.com/g-thanawat/product_format/pd_format"
	// "github.com/labstack/echo/v4"
)

// Example of struct
type InputOrder struct {
	No                int     `json:"no"`
	PlatformProductId string  `json:"platformProductId"`
	Qty               int     `json:"qty"`
	UnitPrice         float64 `json:"unitPrice"`
	TotalPrice        float64 `json:"totalPrice"`
}

type CleanedOrder struct {
	No         int     `json:"no"`
	ProductId  string  `json:"productId"`
	MaterialId string  `json:"materialId"`
	ModelId    string  `json:"modelId"`
	Qty        int     `json:"qty"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
}

func textureCount(s string) map[string]int {
	r := map[string]int{}
	word := strings.Fields(s)
	for _, v := range word {
		r[v] = r[v] + 1
	}
	return r
}

func main() {
	// e := echo.New()

	// e.GET("/", endpoint)
	fmt.Println("hello")

	// clean_and_split("FG0A-CLEAR-IPHONE16PROMAX"

	// clean_and_split("--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3")
	// fmt.Println("------------------------------------------------------------------------\n")
	d := pd_format.Clean_and_split("--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2")
	fmt.Printf("Final : %#v\n", d)

	// e.Start("localhost:3000")
}
