package main 

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", func( w  http.ResponseWriter , r *http.Request){
		fmt.Println(w ,"Hello, you've requested: %s\n", r.URL.Path);
	})

	fmt.Println("Server Runnig in 80 ");

	http.ListenAndServe(":80", nil );

}