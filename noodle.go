package main

import (
	"flag"
	"fmt"
	"os"
)
var (
	h  bool
	v  bool
	p  string
	f  string
	na string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.StringVar(&p, "p", "./noodle", "create router file dir")
	flag.StringVar(&f, "f", "noodle", "create router file")
	flag.StringVar(&na, "name", "Noodle", "func name")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `cedar version: 0.1
Usage: noodle [-name funcname] [-p filepath] [-f filename]

Options:
`)
	flag.PrintDefaults()
}
func createDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			er := os.Mkdir(dir, os.ModeDir)
			if er != nil {
				panic("create dir err")
			}
		}
	}
}
func createFile(path string, str string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, 755)
	if err != nil {
		panic("open file err")
	}
	_, err = f.WriteString(str)
	if err != nil {
		panic("write file err")
	}
	_ = f.Close()
}
func main() {
	flag.Parse()
	var yes bool
	_, err := os.Stat(p + "/" + f + ".go")
	if err != nil {
		if os.IsNotExist(err) {
			yes = true
		}
	}
	if len(os.Args) == 1 {
		h = true
	}
	if h {
		flag.Usage()
	}
	if len(f) != 0 {
		createDir(p)
	}
	if len(f) != 0 {
		if !yes {
			createFile(p+"/"+f+".go", "\nfunc "+na+"(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {\nreturn func(w http.ResponseWriter, r *http.Request) {\n\n\n\nnext(w,r)\n}}\n")
		} else {
			createFile(p+"/"+f+".go", "package noodle\nimport(\n\t\"net/http\"\n)\n\n\nfunc "+na+"(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {\nreturn func(w http.ResponseWriter, r *http.Request) {\n\n\n\nnext(w,r)\n}}\n")
		}
	} else {
		if !yes {
			createFile(p+"/"+f+".go", "\nfunc "+na+"(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {\nreturn func(w http.ResponseWriter, r *http.Request) {\n\n\n\nnext(w,r)\n})}\n")
		} else {
			createFile(p+"/"+f+".go", "package noodle\nimport(\n\t\"net/http\"\n)\n\n\nfunc "+na+"(next func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {\nreturn func(w http.ResponseWriter, r *http.Request) {\n\n\n\nnext(w,r)\n}}\n")
		}
	}
}