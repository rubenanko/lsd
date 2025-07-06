package main

import "fmt"
import "os"
import "log"
import "io"
import "encoding/json"
import "path/filepath"
import "strings"


var COLORS = map[string]string {
	"grey": "0",
	"black":"30",
	"red": "31",
	"green":"32",
	"yellow":"33",
	"blue":"34",
	"purple":"35",
	"cyan":"36",
	"white":"37",
}

var STYLES = map[string]string {
	"regular":"0",
	"bold": "1",
	"background" : "7",
}

func main() {
	content,err := os.ReadDir("./")
    if err != nil {log.Fatal(err)}

	path, err := os.Executable()
	if err != nil {log.Fatal(err)}

	splitedPath := strings.Split(path,"/")
	finalPath := strings.Join(splitedPath[0:len(splitedPath)-1],"/")

	rawFile,err := os.Open(finalPath + "/config.json")
	if err != nil {log.Fatal(err) }

	file, err := io.ReadAll(rawFile)
	if err != nil {log.Fatal(err) }

	var parsedJson map[string]map[string]string

	json.Unmarshal(file,&parsedJson)

    for _, e := range content {
		if parsedJson[filepath.Ext(e.Name())] != nil {
			tmp := parsedJson[filepath.Ext(e.Name())]
			fmt.Println("\033[" + STYLES[tmp["style"]] + ";"+COLORS[tmp["color"]] + "m" + e.Name() + "\033[0;0m")
		} else {
			fileInfo,err := os.Stat(e.Name())
			if err != nil {log.Fatal(err) }
			
			if fileInfo.Mode().IsDir() {
				tmp := parsedJson["/"]
				fmt.Println("\033[" + STYLES[tmp["style"]] + ";"+COLORS[tmp["color"]] + "m" + e.Name() + "\033[0;0m")
			} else {
				fmt.Println(e.Name())
			}
		}
    }
}