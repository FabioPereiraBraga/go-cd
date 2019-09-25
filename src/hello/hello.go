package main

import (
	"fmt"
	"net/http"
	"math"
)

func helloWord() string {
 return "Code.education Rocks!";
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, helloWord() )
	});

	fmt.Println(http.ListenAndServe(":80", nil))

}