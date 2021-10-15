package main

import (
	"fmt"
	"log"
)

func main(){
	u := Usuario{
		nickname: "juanmiloz",
		email: "juanmiloz@hotmail.com",
		firstname: "Juan",
		lastname: "Zorrilla",
		password: "juanca",
		country: "Colombia",
	}

	err := newUsuario(u)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("creado exitosamente")
}
