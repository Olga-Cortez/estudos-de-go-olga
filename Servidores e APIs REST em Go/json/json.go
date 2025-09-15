package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	
)

type Usuarios struct{
	Usuarios []Usuario `json: "Usuarios"`
} 

type Usuario struct {
	Nome string `json:"Nome"`
	Idade int `json:"Idade"`
	Email string `json:"Email"`
	Social Social `json:"Social"`
}

type Social struct {
	Facebook string `json:"Facebook"`
	Twitter string `json:"Twitter"`
}

func main() {
	arquivoJson, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com sucesso")
	defer arquivoJson.Close()

	byteValue, _ := io.ReadAll(arquivoJson)

	var usuarios Usuarios

	json.Unmarshal(byteValue, &usuarios)

	for i:=0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Nome usuário: ", usuarios.Usuarios[i].Nome)
		fmt.Println("Idade usuário: ", usuarios.Usuarios[i].Idade)
		fmt.Println("email: ", usuarios.Usuarios[i].Email)
		
	}
}