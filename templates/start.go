package main


import ("fmt"
	"text/template"
	"os"
	)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tmpl.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
	if err!=nil  {
		fmt.Println(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err!=nil  {
		fmt.Println(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	if err!=nil  {
		fmt.Println(err)
	}
}
