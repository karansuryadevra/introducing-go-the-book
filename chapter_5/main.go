package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100

	fmt.Println(x)

	y := [5]float64{
		12,
		32,
		46,
		190,
		91,
	}

	for i, value := range y {
		fmt.Println(i, " ", value)
	}

	z := make([]float64, 5)
	fmt.Println(z)
	z = append(z, 1, 2, 3)
	fmt.Println(z)

	a := make([]float64, 8)
	copy(z, a)
	fmt.Println(a)

	newMap := make(map[string]int)
	newMap["key"] = 20
	fmt.Println(newMap)

	// check whether a an element exists, if it does, print it otherwise print "ok"

	if value, ok := newMap["karan"]; ok {
		fmt.Println(value, ok)
	} else {
		fmt.Println(ok)
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "Gas",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "Solid",
		},
	}

	if el, ok := elements["C"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
