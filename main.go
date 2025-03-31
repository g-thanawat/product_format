package main

import (
	"fmt"
	"maps"

	"net/http"

	"github.com/g-thanawat/product_format/pd_format"
	// "github.com/labstack/echo"
	"github.com/labstack/echo/v4"
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

type Add_ons struct {
	No         int     `json:"no"`
	ProductId  string  `json:"productId"`
	Qty        int     `json:"qty"`
	UnitPrice  float64 `json:"unitPrice"`
	TotalPrice float64 `json:"totalPrice"`
}

func order(c echo.Context) error {

	od := []InputOrder{}
	err := c.Bind(&od)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Data not complete")
	}

	Texture_amount := map[string]int{}

	var No_pd int = 0
	var pd_amount int = 0
	var total_pd_amount int = 0

	var output []any

	for _, pd := range od {
		product := pd_format.Clean_and_split(pd.PlatformProductId)
		new_texture_amount := pd_format.TextureCount(product, Texture_amount)
		maps.Insert(Texture_amount, maps.All(new_texture_amount))

		//Cal unit Price
		for _, v := range product {
			pd_amount += v.Quantity
		}

		for _, data := range product {

			var Order = CleanedOrder{}
			No_pd += 1
			Order.No = No_pd
			Order.ProductId = data.FilmID + "-" + data.TextureID + "-" + data.Phone_modelID
			Order.MaterialId = data.FilmID + "-" + data.TextureID
			Order.ModelId = data.Phone_modelID
			Order.Qty = data.Quantity
			Order.UnitPrice = pd.TotalPrice / float64(pd_amount)
			Order.TotalPrice = Order.UnitPrice * float64(Order.Qty)

			output = append(output, Order)
		}

		//Total product amount
		total_pd_amount += pd_amount
		//clear product amount
		pd_amount = 0

	}

	fmt.Printf("Texture = %#v \n", Texture_amount)

	//Add ons

	//Add wipingcloth
	No_pd += 1
	wipingcloth := Add_ons{No_pd, "WIPING-CLOTH", total_pd_amount, 0.00, 0.00}
	output = append(output, wipingcloth)

	//Add the others
	for t := range maps.Keys(Texture_amount) {
		No_pd += 1
		Cleaner := Add_ons{No_pd, t + "-CLEANER", Texture_amount[t], 0.00, 0.00}
		output = append(output, Cleaner)
	}

	return c.JSON(http.StatusCreated, output)
}

func main() {
	e := echo.New()

	fmt.Println("hello")

	// clean_and_split("FG0A-CLEAR-IPHONE16PROMAX")

	// clean_and_split("--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3")
	// fmt.Println("------------------------------------------------------------------------\n")
	// d := pd_format.Clean_and_split("--FG0A-CLEAR-OPPOA3*2/FG0A-MATTE-OPPOA3*2")
	// d := pd_format.Clean_and_split("x2-3&FG0A-MATTE-IPHONE16PROMAX*3")
	// fmt.Printf("Final : %#v\n", d)

	// mm := map[string]int{}
	// a := pd_format.TextureCount(d, mm)
	// fmt.Printf("FinalTexture : %#v\n", a)
	// maps.Insert(mm, maps.All(a))
	// fmt.Printf("FinalTexture2 : %#v\n", mm)
	e.POST("/order", order)
	e.Start("localhost:3000")
}
