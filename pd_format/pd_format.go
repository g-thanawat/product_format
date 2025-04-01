package pd_format

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

type Product_Code_Format struct {
	FilmID        string
	TextureID     string
	Phone_modelID string
	Quantity      int
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

func Clean_and_split(s string) []Product_Code_Format {

	var p_format Product_Code_Format
	var result []Product_Code_Format
	// Split product
	product := strings.Split(s, "/")
	// fmt.Printf("Result: %#v\n", product)

	for idxx, val := range product {
		//clean wrong prefix data
		//Assume that all filmtypeID has start with "FG"
		j := strings.LastIndex(val, "FG")
		if j != -1 {
			product[idxx] = product[idxx][j:]
		}
		// fmt.Printf("Result2: %#v\n", product)

	}

	for _, v := range product {

		//Split quantity
		q_data := strings.Split(v, "*")
		// fmt.Printf(" val product %#v\n", q_data)

		if len(q_data) == 2 {
			val, err := strconv.Atoi(q_data[1])

			if err != nil {
				fmt.Printf("Not_number\n")
				p_format.Quantity = 1
			} else {
				p_format.Quantity = val
			}

		} else {
			p_format.Quantity = 1
		}

		// Split each ID
		data := strings.Split(q_data[0], "-")
		// fmt.Printf(" val product2 %#v\n", data)

		// Fill each data
		p_format.FilmID = data[0]
		p_format.TextureID = data[1]
		p_format.Phone_modelID = data[2]

		result = append(result, p_format)

	}

	fmt.Printf("SPLIT : %#v\n", result)

	return result
}

func TextureCount(p []Product_Code_Format, m map[string]int) map[string]int {
	// r := map[string]int{}
	old := maps.Clone(m)
	// old := map[string]int{}
	for _, v := range p {
		old[v.TextureID] = old[v.TextureID] + v.Quantity
	}
	return old
}

func ProductAmount(p []Product_Code_Format) int {

	var product_amount int

	for _, v := range p {
		product_amount += v.Quantity
	}

	return product_amount
}

func Cleaned_Order(No int, data Product_Code_Format, q int, u float64) CleanedOrder {

	var Order = CleanedOrder{}

	Order.No = No
	Order.ProductId = data.FilmID + "-" + data.TextureID + "-" + data.Phone_modelID
	Order.MaterialId = data.FilmID + "-" + data.TextureID
	Order.ModelId = data.Phone_modelID
	Order.Qty = data.Quantity * q
	Order.UnitPrice = u
	Order.TotalPrice = Order.UnitPrice * float64(Order.Qty)

	return Order

}

func main() {
	fmt.Println("hello")
}
