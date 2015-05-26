/*
- Arrays
	* var x [5]int

- Slices
	* A slice is a segment of an array
		var x []float64		-> 		an array is the missing length between the brackets

	* For create a slice use make
		x := make([]float64, 5, 10)

		_ _ _ _ _ _ _ _ _ _
	   |_ _ _ _ _ _|
	         x

	* Another form
		arr := []float64{1,2,3,4,5}
		x := arr[0:5]

	* Slice functions
		slice1 := []int{1,2,3}
		slice2 := append(slice1, 4, 5)
		fmt.Println(slice1,slice2)

		slice1 -> [1,2,3]
		slice2 -> [1,2,3,4,5]

		--------------------------

		slice1 := []int{1,2,3}
		slice2 := make([]int, 2)
		copy(slice2, slice1)
		fmt.Println(slice1, slice2)

		slice1 -> [1,2,3]
		slice2 -> [1,2]

	* Maps: A map is an unordered collection of key-value pairs. Another names associative array, hash table, dictionary
		var x map[string]int

		x := make(map[string]int)
		x["key"] = 10
		fmt.Println(x["key"])    ->    10

		delete(x, 1)

		------------------------------------

		elements["Un"] == ""

		name, ok := elements["Un"]
		fmt.Println(name, ok)

		- is the same to:

		if name, ok := elements["Un"]; ok {
			fmt.Println(name, ok)
		}

		------------------------------------

		elements := map[string]string{
			"H": "Hydrogen",
			"He": "Helium",
		}

		elements := map[string]map[string]string{
			"H": map[string]string{
				"name":"Hydrogen",
				"state":"gas",
			},
			"He": map[string]string{
				"name":"Helium",
				"state":"gas",
			},
		}

		if el, ok := elements["Li"]; ok {
			fmt.Println(el["name"], el["state"])
		}

*/

package main

import "fmt"

func main() {
	/*var x [5]int
	x[4] = 100
	fmt.Println(x)*/

	/*var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83*/

	// 	x := [5]float64{ 98, 93, 77, 82, 83 }

	/*
		x := [5]float64{
			98,
			93,
			77,
			82,
			83, // The extra trailing, after 83. This is required by Go
		}
	*/

	/*var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)*/

	/*var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))*/

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println(elements["He"])
}
