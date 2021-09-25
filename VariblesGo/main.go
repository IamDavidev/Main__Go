package main

import (
	"fmt"
	"strconv"
)

//para poder exportar varibles o funciones  tienen que empezar con Mayuscula
// => ejmplo
var Export string = "exportvariable exportable"
func FuncionExportando (){
	fmt.Println(Export)



}
// ..
var num int = 10 
var text string 
var status bool = false || true 

func main() {
	/*
	valueUno, valueDos, valueTres := 1 , "valor dos tomado en go",true
	fmt.Println(valueUno)
	fmt.Println(valueDos)
	fmt.Println(valueTres)
	*/
	Variables()
}

func Variables () {
 	var valor int64 = 1369
	var text = strconv.Itoa(int(valor))
	fmt.Println(text)
} 
