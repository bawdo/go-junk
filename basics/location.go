// Go Bootcamp: 2.7

package main

import "fmt"

// unamed result parameters
func location(city string) (string, string) {
	var region, continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "Melbourne":
		region, continent = "Victoria", "Australia"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

// named result parameters
func location2(city string) (region, country string) {
	switch city {
	case "Tokyo", "Chiba":
		region, country = "Kanto", "Japan"
	case "Osaka", "Kyoto":
		region, country = "Kansai", "Japan"
	default:
		region, country = "Unknown", "Unknown"
	}
	return
}

func main() {
	region, continent := location("Melbourne")
	fmt.Printf("Bawdo lives in %s, %s\n", region, continent)

	region, country := location2("Tokyo")
	fmt.Printf("Bawdo used to live in %s, %s\n", region, country)
}
