package main

import "fmt"

func main() {
	StudentAge := make(map[string]int)
	StudentAge["Aryya"] = 23
	StudentAge["Saurabh"] = 27
	StudentAge["Prerna"] = 21
	StudentAge["Akriti"] = 19
	StudentAge["Sahiti"] = 22
	StudentAge["Kirti"] = 22

	fmt.Println(StudentAge)
	fmt.Println(len(StudentAge))
	fmt.Println(StudentAge["Kirti"])

	superHero := map[string]map[string]string{
		"Superman": {
			"RealName": "Clark Kent",
			"City":     "Metrolpolis",
		},
		"Batman": {
			"RealName": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}

	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
