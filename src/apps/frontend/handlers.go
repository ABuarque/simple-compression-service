package main

import (
	"context"
	amassa "frontend/proto"
	"html/template"
	"log"
	"net/http"
)

// Home constrols the state of main html file
type Home struct {
	Sucess bool
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	h := Home{Sucess: false}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		showMessage("Eita, deu pau ai visse, tenta ai de novo...", w)
		return
	}
	t.Execute(w, h)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	email := r.FormValue("email")
	option := r.FormValue("options")
	file, header, err := r.FormFile("file")
	encoded, err := inputFileResourcesTobase64(file, header)
	if err != nil {
		log.Println(err.Error())
		showMessage("Deu pau ai visse, tenta de novo!", w)
		return
	}
	if option == "decompress" {
		err := isValidToDecompress(header.Filename)
		if err != nil {
			log.Println(err.Error())
			showMessage("Só sei descomprimir arquivo .huff :(", w)
			return
		}
	}
	p := amassa.Request{
		Email:   email,
		Name:    header.Filename,
		Command: option,
		Bytes:   encoded,
	}
	_, e := client.RequestAction(context.Background(), &p)
	if e != nil {
		log.Println("error on gRPC call: ", e)
		showMessage("Eita, deu pau ai visse, tenta ai de novo...", w)
		return
	}
	log.Println("rpc call made to be processed")
	h := Home{Sucess: true}
	tmpl.Execute(w, h)
}
