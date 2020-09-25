package main

import ("fmt"
	"text/template"
	"os"
	)

var templ *template.Template

type productName struct{
	Name string "sex"
	Type string
}

type review struct{
	Name string
	Comment string
}


func init(){
	templ = template.Must(template.ParseGlob("../templates/*"))
}

func main(){
	a := productName{
		Name: "prigles",
		Type: "biscut",
	}
	b := productName{
                Name: "xrigles",
                Type: "no",
        }
	c := productName{
                Name: "mingles",
                Type: "biscut",
        }
	ar := review{
                Name: "prigles",
                Comment: "tasty chips",
        }
	br := review{
                Name: "xrigles",
                Comment: "noo tasty chips",
        }
	cr := review{
                Name: "mingles",
                Comment: "yummmmm",
        }

	reviewProduct := struct{
		Product []productName
		ProductReview []review
	}{[]productName{a,b,c}, []review{ar,br,cr}}
	err := templ.ExecuteTemplate(os.Stdout, "struct_slice_struct.gohtml", reviewProduct)
	if err != nil{
		fmt.Println(err)
	}
}
