package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array()
	slice()
	channel()
	maps()
}

func array() {
	fmt.Println("数组--------------------------")
	var numbers0 = [5]float32{10.0, 9.0, 8, 7, 6}
	numbers1 := [...]float32{20.0, 9.0, 8, 7, 6}
	numbers2 := [...]float32{1: 2, 9: 4}
	matrix := [...][3]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(numbers0, numbers1, numbers2)
	fmt.Println(numbers0[2])
	fmt.Println(matrix)
	fmt.Println(matrix)
}

func slice() {
	fmt.Println("切片--------------------------")
	intSlice := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		intSlice[i] = rand.Intn(20)
	}
	fmt.Println(intSlice)
	fmt.Println(intSlice[7:])
	fmt.Println(intSlice[:3])
	intSlice = append(intSlice, 100)
	fmt.Println(intSlice)
	fmt.Printf("大小：%d,容量：%d\n", len(intSlice), cap(intSlice))

	intSlice = []int{}
	oldSlice := []int{1, 2, 3}
	fmt.Printf("大小：%d,容量：%d\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 100)
	intSlice = append(intSlice, oldSlice...)
	fmt.Printf("大小：%d,容量：%d,values->%v\n", len(intSlice), cap(intSlice), intSlice)

	intSlice = []int{1, 2, 3, 4, 5}
	length := len(intSlice)
	splitSlice := append(intSlice[:length/2], intSlice[length/2+1:]...)
	fmt.Println(splitSlice)
}

func channel() {
	fmt.Println("通道--------------------------")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func maps() {
	fmt.Println("map映射--------------------------")
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Japan"] = "Tokyo"
	for k, e := range countryCapitalMap {
		fmt.Println("country:", k, "capital:", e)
	}
	China := "China"
	capital, ok := countryCapitalMap[China]
	if ok {
		fmt.Printf("%s -> %s\n", China, capital)
	} else {
		fmt.Println("China's capital not found")
	}
	delete(countryCapitalMap, "Japan")
	countryCapitalMap["China"] = "Beijing"
	fmt.Println(countryCapitalMap)
}
