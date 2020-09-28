package main

import("fmt"
	"text/template"
	"os"
	"time"
)

var templ *template.Template

var fn = template.FuncMap{
	"double": twoTimes,
	"fmtTime": fmtTime,
}

func twoTimes(x float64) float64{
	return 2 * x
}

func fmtTime(t time.Time) string{
	return t.Format("02-01-2006")
}

func init(){
	//Here ParseGlob won't work becase as soon as we do it will parse all files and map-fun.gohtml has a func call which won't be possible as func is not defined here
	templ = template.Must(template.New("").Funcs(fn).ParseFiles("templates/time_pipeline.gohtml"))
}

func main(){
	data := struct{
		Nums float64
		Time time.Time
	}{34,time.Now()}
	err := templ.ExecuteTemplate(os.Stdout,"time_pipeline.gohtml", data)
	if (err != nil){
		fmt.Println(err)
	}
}
