/*
A map is an unordered collection of key-value pairs (maps are also sometimes called associative arrays,
hash tables, or dictionaries
*/

package main 

import "fmt"

func main(){
	x := make(map[string]int) 
	/*
		This will create a map called "x" which has key-pairs of strings to int
		Also note that you HAVE to initialize a map before you can start using it
		You can initialize it by using "make"
	*/


	x["key"] = 10
	
	fmt.Println(x["key"])

	names := make(map[string]string) 
	/* the "make" function needs to be used if you want to create an uninitialized 
		slice or map
	*/
	names["P"] = "Phoebe"
	names["H"] = "Hoebe"
	names["O"] = "Oebe"
	names["E"] = "Ebe"
	names["B"] = "Be"
	names["E"] = "E"

	fmt.Println("The length of the map is : ", len(names))

	if result_of_lookup, was_lookup_successful := names["K"]; !was_lookup_successful{
			fmt.Println("result_of_lookup = ", result_of_lookup)
			fmt.Println("was_lookup_successful = " , was_lookup_successful)
	}

	/* Looking up a key in a map that doesn't exist can return 2 values :
		1. What the lookup result was
		2. Whether the lookup was successful

		The if block above looks for a non existing key in the map, and if the lookup was
		unsuccessful, it prints the actual result and the status of the lookup
	*/

	

	// The code below  would be the map equivalent of {"string:{"string1": "string", "string2" : "string"}}
	elements := map[string]map[string]string {
		"H": map[string]string {
			"name" : "Hydrogen",
			"state" : "gas",
		},
		"He": map[string]string {
			"name" : "Helium",
			"state" : "gas",
		},
		"C": map[string]string {
			"name" : "Carbon",
			"state" : "solid",
		},
	}	

	/*
		The code below is how you check for whether and element is present
		And if it is present then print it's name and state
		"el" basically contains "[name:Carbon state:solid]"
		So we use the "el" map and access the name and state of it
	*/
	if el, ok := elements["C"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
