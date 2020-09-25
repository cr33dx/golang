package main

import("text/template"; "os"; "fmt")

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	err := tmpl.ExecuteTemplate(os.Stdout, "first.gohtml", 43)
	if err != nil{
		fmt.Println(err)
	}
}
