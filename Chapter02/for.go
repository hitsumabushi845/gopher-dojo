package main

func main() {
	for i := 0; i <= 10; i = i + 1 {
		println(i)
	}

	for i, v := range []int{1, 2, 3} {
		println(i, v)
	}

	var i int
	for {
		if i%2 == 1 {
			break
		}
		i++
	}
}
