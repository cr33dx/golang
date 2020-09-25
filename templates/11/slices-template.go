package main

import ("text/template"; "os"; "fmt")

var tmpl *template.Template

func init(){
	tmpl = template.Must(template.ParseGlob("../templates/*"))
}

func main(){
	fruits := []string{"apple", "banana", "pear", "peach"}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", fruits)
	if err != nil {
		fmt.Println(err)
	}
}
