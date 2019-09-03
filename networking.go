package main

import (
  "fmt"
  "net/http"
  "html/template"
  "strings"
  "bufio"
)

type NewsAggPage struct {
  Title string
  News string
}

type searchResults struct {
  PageData NewsAggPage
  Query string
  Result []string
}

func handleErr(err error, w http.ResponseWriter) {
  if err == nil {
    fmt.Println("200 OK")
  } else {
    fmt.Fprintf(w, "<h1><center>500 Internal Server Error =8-O <center></h1>")
    fmt.Println(err)
  }
}

func indexSearchHandler(w http.ResponseWriter, r *http.Request) {
  fSearch := r.FormValue("search") == "yes"
  fText := r.FormValue("text")
  var fResults []string

  res := searchResults {PageData : NewsAggPage{Title: "The Title", News: "The very new news"}, Query: "", Result: nil}
  t, _ := template.ParseFiles("templateSearch.html")

  if fSearch {
    ioScanner := bufio.NewScanner(strings.NewReader(fText))
    ioScanner.Split(bufio.ScanWords)
    for ioScanner.Scan() {
      fResults = append(fResults, ioScanner.Text())
    }

    res.Query = fText
    res.Result = fResults
  }

  handleErr ( t.Execute(w, res), w )
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/agg", http.StatusSeeOther)
}


func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/agg", indexSearchHandler)
  http.ListenAndServe(":8000", nil)
}
