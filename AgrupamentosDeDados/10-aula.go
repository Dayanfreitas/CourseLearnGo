package main

import "fmt"

func main() {
	qualquerCoisa := map[int]string{
		123: "é mto legal",
		98:  "men",
		11:  "a",
	}

	fmt.Println(qualquerCoisa)
	delete(qualquerCoisa, 123)
	for k, v := range qualquerCoisa {
		fmt.Println(k, v)
	}
}
