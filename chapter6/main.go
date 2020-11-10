package main

import "fmt"

func main() {
	/*x := make(map[string]string)
	x["H"] = "Hidrogen"
	x["He"] = "Helium"
	x["Li"] = "Lithium"
	x["Be"] = "Beryllium"
	x["B"] = "Boron"
	x["C"] = "Carbon"
	x["N"] = "Nitrogen"
	x["O"] = "Oxygen"
	x["F"] = "Flourine"
	x["Ne"] = "Neon"*/

	/*x := map[string]string{
		"He": "Helium",
		"H":  "Hidrogen",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Flourine",
		"Ne": "Neon",
		"Li": "Lithium",
	}
	fmt.Println(x["F"])*/
	x := map[string]map[string]string{
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"H": map[string]string{
			"name":  "Hidrogen",
			"state": "gas",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Flourine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Println(x["Ne"])
}
