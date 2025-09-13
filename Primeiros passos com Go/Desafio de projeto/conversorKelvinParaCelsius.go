package main

import "fmt"

//c = k - 273
func main(){
	var tempKelvin float32
	fmt.Println("Conversor de Kelvin para Cesius")
	fmt.Print("Insira a temperatura em Kelvin: ")
	fmt.Scanln(&tempKelvin)
	resultado := tempKelvin - 273
	fmt.Printf("A temperatura equivalente é %g", resultado)
}