package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["mr_bean"] = []string{"boink", "bing", "bong"}

	for k, v := range x {
		fmt.Println(k)
		for p, q := range v {	
			fmt.Println("\t", p, q)
		}
	}
	fmt.Println("____________________________________-")
	delete(x, "bond_james")

	for k, v := range x {
		fmt.Println(k)
		for p, q := range v {
			fmt.Println("\t", p, q)
		}
	}

}
