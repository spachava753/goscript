package main

func a(v int) func(int) int {
	i := 100
	i = i + v
	return func(j int) int {
		return i + j
	}
}

func main() {
	b := a(30)
	c := b(2)
}
//132