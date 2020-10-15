package main                                                                      // inshorts backend API - SACHIT KALIA

import (                                                                          // imported packages
	"fmt"
	"log"
  "io/ioutil"
	"net/http"
  "encoding/json"
  "strings"
  "time"
)

type Article struct{                                                              // an article object
  Id string `json:"Id"`
  Title string `json:"Title"`
  Subtitle string `json: "desc"`
  Content string `json: "content"`
  Timestamp string `json: "time"`
}

var Articles []Article                                                            // Global array of articles

func allArticles(w http.ResponseWriter, r * http.Request){                        // 'GET' request showing all the articles
        switch r.Method {
        case "GET":
          fmt.Println("Here showing all the Articles as get request!")
          json.NewEncoder(w).Encode(Articles)
        case "POST":
          fmt.Println("Post request called!")
          reqBody, _ := ioutil.ReadAll(r.Body)
          var newArticle Article
          json.Unmarshal(reqBody, &newArticle)
          t := time.Now()                                                                //setting the current timestamp
          newArticle.Timestamp = t.String()
          Articles = append(Articles, newArticle)
          json.NewEncoder(w).Encode(newArticle)

        default:
          fmt.Println("Something went wrong!")
        }

}

func homePage(w http.ResponseWriter, r * http.Request){
  fmt.Fprintf(w, "Home")
}

func getArticleById(w http.ResponseWriter, r * http.Request){                     //getting individual articles by id
    key := r.URL.Path[len("/articles/"):]
    for _, article := range Articles{
         if article.Id == key{
           json.NewEncoder(w).Encode(article)
         }
    }
}

func Found(pattern string, article Article)(bool){
     text := article.Title
     text1 := article.Subtitle
     text2 := article.Content
     ans := (strings.Contains(text, pattern) || strings.Contains(text1, pattern) || strings.Contains(text2, pattern))

     return ans
}

func searchInArticles(w http.ResponseWriter, r * http.Request){
     key:= r.FormValue("q")

    var tempArticles []Article

     for _, article := range Articles{                                                       //searching for query in the articles
       isFound := Found(key, article)
       if isFound == true{
         tempArticles = append(tempArticles, article)
       }
     }

       json.NewEncoder(w).Encode(tempArticles)

}

func handleRequests(){
  http.HandleFunc("/", homePage)
  http.HandleFunc("/articles", allArticles)
  http.HandleFunc("/articles/", getArticleById)
  http.HandleFunc("/articles/search", searchInArticles)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Subtitle: "Article Description", Content: "Article Content", Timestamp: "2019-11-10 23:00:00 +0000 UTC"},       //default articles
        Article{Id: "2", Title: "Hello 2", Subtitle: "Article Description", Content: "Article Content", Timestamp: "2020-08-10 04:00:00 +0000 UTC"},
    }
    handleRequests()
}
