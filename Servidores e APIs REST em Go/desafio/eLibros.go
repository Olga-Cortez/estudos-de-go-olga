package main 

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"time"
)

type Usuario struct {
	id int
	nome string
	email string
	senha string
	cpf string
	foto_de_perfil string
	genero string
	data_de_nascimento time.Time
}

type Livro struct {
	id int
	titulo string
	subtitulo string
	ISBN string
	sinopse string
	capa string
	autor string
	editora string
	data_de_publicacao time.Time
	preco float32
	desconto int
	data_adicao time.Time
	quantidade_vendida int
	estoque int
}