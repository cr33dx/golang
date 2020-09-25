package main

import ("fmt"
	"text/template"
	"os")

var templ *template.Template

func init(){
	templ = template.Must(template.ParseGlob("../templates/*"))
}

func main(){
	fruits := map[string]string{
		"mango" : "eat",
		"pineapple" : "juice",
		"banana" : "eat",
		"orange": "juice",
		"apple" : "eat",
	}
	//fmt.Println(fruits)
	err := templ.ExecuteTemplate(os.Stdout,"map.gohtml",fruits)
	if err != nil{
		fmt.Println(err)
	}
}
