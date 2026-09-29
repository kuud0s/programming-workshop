package main

import "fmt"

/*
Задача 1.1: объявить переменные всех типов, 3 способами var с типом
var без тпиа
:=
Вывести значения через %T
*/
func main() {
	fmt.Println("\nРезультат Задачи 1.1: ")

	var i int = -1
	var u uint = 1
	var f32 float32 = 3.14
	var b byte = 'A'
	var r rune = 'B'
	var bl bool = true
	var str string = "Hello"

	fmt.Printf("int: %T\n, %v\n", i, i)
	fmt.Printf("uint: %T\n, %v\n", u, u)
	fmt.Printf("float32: %T\n, %v\n", f32, f32)
	fmt.Printf("byte: %T\n, %v\n", b, b)
	fmt.Printf("rune: %T\n, %v\n", r, r)
	fmt.Printf("bool: %T\n, %v\n", bl, bl)
	fmt.Printf("string: %T\n, %v\n", str, str)

	var i2 int = -1
	var u2 uint = 1
	var f32_2 float32 = 3.14
	var b2 byte = 'A'
	var r2 rune = 'B'
	var bl2 bool = true
	var str2 string = "Hello"

	fmt.Printf("int2: %T\n, %v\n", i2, i2)
	fmt.Printf("uint2: %T\n, %v\n", u2, u2)
	fmt.Printf("float32_2: %T\n, %v\n", f32_2, f32_2)
	fmt.Printf("byte2: %T\n, %v\n", b2, b2)
	fmt.Printf("rune2: %T\n, %v\n", r2, r2)
	fmt.Printf("bool2: %T\n, %v\n", bl2, bl2)
	fmt.Printf("string2: %T\n, %v\n", str2, str2)

	i3 := -1
	u3 := uint(1)
	f32_3 := float32(3.14)
	b3 := byte('A')
	r3 := 'B'
	bl3 := true
	str3 := "Hello"

	fmt.Printf("int3: %T\n, %v\n", i3, i3)
	fmt.Printf("uint3: %T\n, %v\n", u3, u3)
	fmt.Printf("float32_3: %T\n, %v\n", f32_3, f32_3)
	fmt.Printf("byte3: %T\n, %v\n", b3, b3)
	fmt.Printf("rune3: %T\n, %v\n", r3, r3)
	fmt.Printf("bool3: %T\n, %v\n", bl3, bl3)
	fmt.Printf("string3: %T\n, %v\n", str3, str3)

	task1_2()
}

/*
Задача 1.2: поменять местами значение 2х переменных
- через третью переменную
- без третьей переменной
*/
func task1_2() {
	fmt.Println("\nРезультат Задачи 1.2: ")

	a := 5
	b := 10
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

	x := 5
	y := 10
	x, y = y, x
	fmt.Println(x, y)
}
