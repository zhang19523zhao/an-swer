package main

import (
	"fmt"
	"html/template"
	"net/http"
	"web/module"
)

func main() {
	addr := ":9999"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("test")
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/ans.html"))
		tpl.ExecuteTemplate(w, "ans.html", module.List())
	})

	//添加
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id := r.PostFormValue("id")
			q := r.PostFormValue("qs")
			a := r.PostFormValue("a")
			b := r.PostFormValue("b")
			c := r.PostFormValue("c")
			d := r.PostFormValue("d")
			ans := r.PostFormValue("ans")
			module.ADD(id, q, ans, a, b, c, d)
			http.Redirect(w, r, "/", 302)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/add.html"))
		tpl.ExecuteTemplate(w, "add.html", nil)
	})

	//检查题目
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		Err := []*string{}
		if r.Method == http.MethodPost {
			r.ParseForm()
			//fmt.Println(r.Form["1"])
			Err = module.Check(w, r)
		}
		tpl := template.Must(template.New("tpl").ParseFiles("vieiws/check.html"))
		tpl.ExecuteTemplate(w, "check.html", Err)
	})

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
