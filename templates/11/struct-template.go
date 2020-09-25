package main

import ("fmt"
	"text/template"
	"os"
	)

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("../templates/*"))
}

type fruitsSeason struct {
	Name string
	Season string
}
func main(){
	fruit := fruitsSeason{
		Name: "mango",
		Season: "summer",
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "struct.gohtml", fruit)
	if err != nil{
		fmt.Println(err)
	}
}

