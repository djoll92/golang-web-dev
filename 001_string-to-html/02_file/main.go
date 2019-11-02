package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Djordje Marjanovic"

	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	newFile, err := os.Create("001_string-to-html/02_file/index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer newFile.Close()

	io.Copy(newFile, strings.NewReader(str))
}
