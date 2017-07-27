package main

import (
  "fmt"
  "log"
  "github.com/PuerkitoBio/goquery" // jQuery for Go
)

func getInfo() {
  doc, err := goquery.NewDocument("https://es.wikipedia.org/wiki/The_City_of_New_York_vs._Homer_Simpson")
  if err != nil {
    log.Fatal(err)
  }

  title := doc.Find("h1").Text()
  fmt.Printf("Location: %s\n", title)
}

func main() {
  getInfo()
}
