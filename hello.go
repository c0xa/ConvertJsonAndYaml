package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func FromJsonToYaml(fileFrom string, FileTo string) {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		panic(err)
	}
	var u interface{}
	json.Unmarshal(file, &u)
	if err != nil {
		panic(err)
	}
	finishResult, err := yaml.Marshal(&u)

	filename, err := os.OpenFile(FileTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(FileTo, finishResult, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer filename.Close()
}

func FromYamlToJson(fileFrom string, FileTo string) {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		panic(err)
	}
	var u interface{}
	yaml.Unmarshal([]byte(file), &u)
	if err != nil {
		panic(err)
	}
	finishResult, err := json.Marshal(u)

	filename, err := os.OpenFile(FileTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print(finishResult)
	err = ioutil.WriteFile(FileTo, finishResult, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer filename.Close()
}

func main() {
	nameInput := flag.String("input", "file.yaml", "Path to config file")
	nameOutput := flag.String("output", "file.json", "Path to config file")
	from := flag.String("from", "yaml", "Path to config file")
	to := flag.String("to", "json", "Path to config file")
	flag.Parse()
	if *from == "yaml" && *to == "json" {
		FromYamlToJson(*nameInput, *nameOutput)
	}
	if *from == "json" && *to == "yaml" {
		FromJsonToYaml(*nameInput, *nameOutput)
	}
}
