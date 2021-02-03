package main

const (
	Unknown = iota + 100
	Female
	Male
)

func main() {
	const LENGTH = 10
	const WIDTH = 5
	const a, b, c = 1, false, "OK"

	var area int
	area = LENGTH * WIDTH
	println("面积为：", area)
	println(a, b, c)
	println(Unknown, Female, Male)
}
