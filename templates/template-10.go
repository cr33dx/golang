package main

import("text/template"; "os"; "fmt")

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseFiles("vartemplate.gohtml"))
}

func main(){
	err := tmpl.Execute(os.Stdout,`temprature is 43 C`)
	if err != nil{
		fmt.Println(err)
	}
}
