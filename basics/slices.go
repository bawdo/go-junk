// Go Bootcamp: 4.2
package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)
	// [2 3 5 7 11 13]

	fmt.Println(p[1:4])
	// [3 5 7]

	// mising low index implies 0
	fmt.Println(p[:3])
	// [2 3 5]

	// missing high index imples len(s)
	fmt.Println(p[4:])
	// [11 13]

	cities := make([]string, 3)
	cities[0] = "Melbourne"
	cities[1] = "Tokyo"
	cities[2] = "Perth"
	fmt.Printf("%q\n", cities)

	machi := []string{}
	// machi[0] = "Armadale" Boom cannot do this
	machi = append(machi, "Armadale")
	fmt.Println(machi)
	machi = append(machi, "Meguro", "Chiba")
	fmt.Println(machi)

	fmt.Println(len(cities))
	cities = append(cities, machi...)
	fmt.Println("%q", cities)
	fmt.Println(len(cities))

	countries := make([]string, 42)
	fmt.Println(len(countries))

	// 4.2.5 Nil slices
	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
	}

	// 4.3 Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

	// 4.3.1 Break and Continue
	pow3 := make([]int, 10)
	for i := range pow3 {
		pow3[i] = 1 << uint(i)
		if pow3[i] >= 16 {
			break
		}
	}
	fmt.Println(pow3)
	// [1 2 4 8 16 0 0 0 0 0]

	pow4 := make([]int, 10)
	for i := range pow4 {
		if i%2 == 0 {
			continue
		}
		pow4[i] = 1 << uint(i)
	}
	fmt.Println(pow4)
	// [0 2 0 8 0 32 0 128 0 512]

	// 4.3.2 Range and maps
	city_map := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range city_map {
		fmt.Printf("%s has %d inhabitants\n", key, value)
	}

}
