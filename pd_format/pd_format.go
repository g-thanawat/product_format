package pd_format

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

type Product_Code_Format struct {
	FilmID        string
	textureID     string
	phone_modelID string
	quantity      int
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
				p_format.quantity = 1
			} else {
				p_format.quantity = val
			}

		} else {
			p_format.quantity = 1
		}

		// Split each ID
		data := strings.Split(q_data[0], "-")
		// fmt.Printf(" val product2 %#v\n", data)

		// Fill each data
		p_format.FilmID = data[0]
		p_format.textureID = data[1]
		p_format.phone_modelID = data[2]

		result = append(result, p_format)

	}

	// fmt.Printf("SPLIT : %#v\n", result)

	return result
}

func main() {
	fmt.Println("hello")
}
