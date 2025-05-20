package main

import (
    "fmt"
    "strings"

    "github.com/smartystreets/goconvey/gocv"

    "github.com/alexcrichton/go-simplejson/jsonparser"
)

type SomeData struct {
    Name   string
    Age    int
    Gender bool
}

func main() {
    typeList := []string{"Alice", "Bob"}
    jsonString, _ := jsonparser.NewParser(strings.NewReader(`{"name": "Alice", "age": 30, "gender": true}`)).ParseAndReturnJSON()
    jsonData, err := jsonparser.Unmarshal([]byte(jsonString), &SomeData{})
    check(err)
}
