package main

import ("fmt"
	"text/template"
	"strings"
	"os"
	)

var templ *template.Template

var fn = template.FuncMap{
	"uc": strings.ToUpper,
}

func init(){
	templ = template.Must(template.New("").Funcs(fn).ParseGlob("templates/*.gohtml"))
}

func main(){
	data := map[string]string{
		"tasty": "mango",
		"drink": "budweiser",
	}
	err := templ.ExecuteTemplate(os.Stdout,"fun-map.gohtml",data)
	if (err != nil){
		fmt.Println(err)
	}
}
