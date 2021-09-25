package main

import "fmt"

//para poder exportar varibles o funciones  tienen que empezar con Mayuscula
// => ejmplo
var Export string = "exportvariable exportable"
func FuncionExportando (){
	fmt.Println(Export)



}
// ..
var num int = 10 
var text string = "hola soy de tipo texto"
var status bool = false || true 

func main() {
	valueUno, valueDos, valueTres := 1 , "valor dos tomado en go",true
	fmt.Println(valueUno)
	fmt.Println(valueDos)
	fmt.Println(valueTres)
}

