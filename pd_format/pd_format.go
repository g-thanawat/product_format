package pd_format

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

var Texture_amount map[string]int

type Product_Code_Format struct {
	FilmID        string
	TextureID     string
	Phone_modelID string
	Quantity      int
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
				fmt.Println("Not number\n")
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

func main() {
	fmt.Println("hello")
}
