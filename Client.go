package main

import (
	"flag"
	"fmt"
	"net/http"
	"io"
	"os"
	"io/ioutil"
	"mime/multipart"
	"log"
	"bytes"
)
func print_help(){
	fmt.Println("Written by: Pch12_ -->> @https://github.com/Pch12")
	fmt.Println(" ")
	fmt.Println("Help commands:")
	fmt.Println(" ")
	fmt.Println(" -u / -d for mode")
	fmt.Println(" -H Remote ip/hostname/domain")
	fmt.Println(" -p Port")
	fmt.Println(" -f File to upload/download")
	fmt.Println(" ")
	os.Exit(2)
}
func incoming_flags() (string,string,string) {
        var host string
        var port string
        var file string
        var u bool
        var d bool
	url := ""
	mode := ""
        flag.BoolVar(&u,"u", false, "Set mode to u")
        flag.BoolVar(&d,"d", false, "Set mode to d")
        flag.StringVar(&host,"H", "", "Host/domain")
        flag.StringVar(&port,"p", "80", "Port")
        flag.StringVar(&file,"f", "", "file to u/d")
        flag.Parse()
	if !u && !d{
		fmt.Println("No arguments were set")
		print_help()}
        if os.Args[1] == "-u" || os.Args[1] == "-d"{
		if host != ""{
			url = host + ":" + port
			if port == "443" {
				url = "https://" + url
			}else {
				url = "http://" + url }
			if os.Args[1] == "-u" && file == "" || os.Args[1] == "-d" && file == "" {
				fmt.Println("You didn't specify a file for the mode")
				print_help()}
			if os.Args[1] == "-u" && file != "" {
				mode = "u"
				return url,file,mode
			}else if os.Args[1] == "-d" && file != ""{
				mode = "d"
				return url,file,mode }
		}else {
			fmt.Println("No host was set")
			print_help()}
        }else {
                fmt.Println("Bad syntax, no mode was set")
		print_help()}

	return "","",""
}
func DownloadFile(u_file string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
	return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(u_file)
	if err != nil {
	return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Println("Downloaded")
	return err
}

func newfileUploadRequest(url string, paramName,filename string) (*http.Request, error) {
	os.Chdir("/"+filename)
	extraParams := map[string]string{
		"title":        filename,
		"author":      " Pch12_",
		"description": " Retrieved from machine",
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	for key, val := range extraParams {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request,err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request,err
}


func main() {

	//Get url, and files from functions (Upload or Download)
	url,file,mode := incoming_flags()
	url = url +"/"

	//For u mode
	if mode == "u" { //check if a file was entered 
		if err := DownloadFile(file,url);err != nil {
		panic(err)
		}
	}


	//For d mode
	if mode == "d"{
		request, err := newfileUploadRequest(url,"fileUpload", file)
		if err != nil {
		log.Fatal(err)
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		} else {
			var bodyContent []byte
			fmt.Println("Response code:", resp.StatusCode)
			resp.Body.Read(bodyContent)
			resp.Body.Close()
			fmt.Println("---Uploaded----")
			}
		}
	}
