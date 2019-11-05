package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/fabiossilva00/Go/Alura/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProduts := models.SearchAllProduts()
	temp.ExecuteTemplate(w, "Index", allProduts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func InsertProduts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Printf("Erro conversao preco", err)
		}
		quantConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Printf("Erro conversao quantidade", err)
		}

		models.CriarProduto(nome, descricao, precoConvertido, quantConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func DeleteProdut(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	idConvertido, err := strconv.Atoi(idProduto)
	if err != nil {
		panic(err.Error())
	}
	models.ExcluiProduto(idConvertido)
	http.Redirect(w, r, "/", 301)
}

func EditProdut(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func UpdateProdut(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro conversao Id", err)
		}

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro conversao Preco", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro conversao Quantidade", err)
		}

		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)

}
