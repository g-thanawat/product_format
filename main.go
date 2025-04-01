package main

import (
	"fmt"
	"maps"
	"net/http"

	"github.com/g-thanawat/product_format/pd_format"
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

	var No_pd int = 0
	var total_pd_amount int = 0
	var output []any
	Texture_amount := map[string]int{}

	//Each Order
	for _, pd := range od {
		//Clean and Split data
		product := pd_format.Clean_and_split(pd.PlatformProductId)

		//TextureCount
		new_texture_amount := pd_format.TextureCount(product, Texture_amount)
		maps.Insert(Texture_amount, maps.All(new_texture_amount))

		//Cal unit Price
		product_amount := pd_format.ProductAmount(product)
		unit_price := pd.TotalPrice / float64(product_amount)

		//Build Output
		for _, data := range product {

			No_pd += 1

			//Build CleanedOrder
			Order := pd_format.Cleaned_Order(No_pd, data, pd.Qty, unit_price)

			//Add CleanedOrder to output
			output = append(output, Order)
		}

		//Total product amount
		total_pd_amount += product_amount
		//Clear product amount
		product_amount = 0

	}

	fmt.Printf("Texture = %#v \n", Texture_amount)

	//ADD ONS

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

	e.POST("/order", order)

	e.Start("localhost:3000")
}
