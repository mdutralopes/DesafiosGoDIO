package main

import "fmt"

const ebulicaoK float64 = 373.0
const fusaoK float64 = 273.0

func main() {

	tempC := ebulicaoK - fusaoK

	//fmt.Println("A temperatura de ebulição da água em Farenheit = ", tempF)
	//fmt.Println("A temperatura de ebulição da água em Celsius = ", tempC)

	fmt.Printf("A temperatura de ebulição da água em K = %g e em C = %g", ebulicaoK, tempC)
}
