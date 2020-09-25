package main

import ("fmt"
	"text/template"
	"os"
	)
var templ *template.Template

func init(){
	templ = template.Must(template.ParseGlob("../templates/*"))
}

type nameType struct {
	Name string
	Type string
}

func main(){
	a := nameType{
		Name: "pringles",
		Type: "chips",
	}
	b := nameType{
		Name: "cookie heaven",
		Type: "biscut",
	}
	c := nameType{
		Name: "mix fruit",
		Type: "juice",
	}
	products := []nameType{a, b, c}
	err := templ.ExecuteTemplate(os.Stdout, "slice_struct.gohtml", products)
	if err != nil{
		fmt.Println(err)
	}
}
