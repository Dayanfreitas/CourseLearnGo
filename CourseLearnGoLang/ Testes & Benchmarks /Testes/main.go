package teste

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}

	return total
}
