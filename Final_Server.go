package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"flag"
	"net/http"
)

func main() {
	host,port:= input_flags()
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(host+":"+port, nil)
}
func print_help(){
	fmt.Println("Written by: PCH12_ -->> @https://github.com/Pch12")
        fmt.Println(" ")
        fmt.Println("Help commands:")
        fmt.Println(" ")
        fmt.Println(" -H Remote ip/hostname for server")
        fmt.Println(" -p Port")
        fmt.Println(" ")
}
func input_flags() (string,string){
	var host string
	var port string
	flag.StringVar(&host, "H", "127.0.0.1", "Address to listen on")
	flag.StringVar(&port, "p", "80", "Port to listen on")
	flag.Parse()
	return host,port
}

//Checking errors 
func check(e error) {
	if e!= nil {
		panic(e)
	}
}

//Handle requests function and store in /uploads directory
func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		// open
		f, _, err := req.FormFile("fileUpload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		filename := req.FormValue("title")
		// for your information
		//fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Print("Downloaded file: " + filename)
		s = string(bs)
		a := []byte(s)
		if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
			os.Mkdir("./uploads",0777)
			}

		if _, err := os.Stat("./uploads"); !os.IsNotExist(err) {
			err = ioutil.WriteFile("./uploads/"+ filename, a, 0644)
			check(err)
			}
		err = ioutil.WriteFile("./uploads/"+ filename, a, 0644)
		check(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="fileUpload">
	<input type="submit">
	</form>
	<br>`+s)
}
