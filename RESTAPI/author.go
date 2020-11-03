package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article ...
type Article struct {
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

//array of  Articles ...
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage hope you like it !")//print in host server
	fmt.Println("Endpoint Hit: homePage here")//print in console
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

//array of article stores the content of structure
func main() {
	Articles = []Article{
		Article{Title: "Why Not me?",
			Author: "Anubhav agarwal",
			Link:   "https://www.flipkart.com/why-not-me/p/itma05b57b0b275e?pid=9789390351473&lid=LSTBOK9789390351473XB5UWP&marketplace=FLIPKART&srno=s_1_4&otracker=search&otracker1=search&fm=SEARCH&iid=b149037b-e850-4b7b-a6a6-d516cf67dcb6.9789390351473.SEARCH&ppt=sp&ppn=sp&ssid=ckkaml8ly80000001604402496290&qH=7d8949bcbf85067f"},
		Article{Title: "ATTITUDE IS EVERYTHING",
			Author: "JEFF KELLER",
			Link:   "https://www.flipkart.com/attitude-everything-change-your-life/p/itm4433e7e899444?pid=9789351772071&marketplace=FLIPKART"},
		Article{Title: "the Theory of Everything",
			Author: "Stephen Hawking",
			Link:   "https://www.flipkart.com/the-theory-of-everything/p/itmfc4ubkshsj8px?pid=9788179925911&lid=LSTBOK9788179925911C0KCWP&marketplace=FLIPKART&fm=productRecommendation%2Fsimilar&iid=R%3As%3Bp%3A9789351772071%3Bpt%3App%3Buid%3A2ddd3067-1dc7-11eb-bf3f-5fb95fa3edc1%3B.9788179925911&ppt=pp&ppn=pp&ssid=z0od2pfykg0000001604402612049&otracker=pp_reco_Similar%2BProducts_29_The%2BTheory%2BOf%2BEverything_9788179925911_productRecommendation%2Fsimilar_28&otracker1=pp_reco_NA_productRecommendation%2Fsimilar_Similar%2BProducts_DESKTOP_HORIZONTAL_productCard_cc_29_NA_view-all&cid=9788179925911"},
	}
	handleRequests()
}
