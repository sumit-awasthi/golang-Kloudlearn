//server for rest api
package main
import(
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct{
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}
type Articles []Article 

func allArticles(w http.ResponseWriter,r,*http.Resquest){
    articles:=Articles{
        Article{Title:"Test Title",Desc:"test Description",Content:"Hello World"},
    }
    fmt.Println("Endpoint Hit: All Articles Endpoints")
    json.NewEncoder(w).Encode(articles)
}
 func testPostArticles(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"Test post Endpoint worked") 
 }



func homepage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w,"Homepage Endpoint Hit")

}
func handleRequests(){
    
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/"homepage)
    myRouter.HandleFunc("/articles",allArticles).Methods("GET")
    myRouter.HandleFunc("/articles",testPostArticles).Methods("POST")
    log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main(){
handleRequests()
}