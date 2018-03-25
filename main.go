package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/gorilla/mux"
)

//func GetFunFact(w http.ResponseWriter, r *http.Request) {

//

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/funFact", GetFunFact).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8000", router))
	url := "https://www.reddit.com/r/todayilearned/top/.json?count=20"

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "rp4fx12")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Found error in getting url response")
		panic(err)
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Found error in reading content body")
			panic(err)
		}
		fmt.Printf("%s\n", string(contents))
	}

}
