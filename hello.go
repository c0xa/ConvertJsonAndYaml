package main

import (
	"flag"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func fromJtoY(fileFrom string, FileTo string) error {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		return err
	}

	var u interface{}
	if err := jsoniter.Unmarshal(file, &u); err != nil {
		return err
	}
	finishResult, _ := yaml.Marshal(&u)

	filename, err := os.OpenFile(FileTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	_ = filename.Close()

	if err := ioutil.WriteFile(FileTo, finishResult, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func fromYtoJ(fileFrom string, FileTo string) error {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		return err
	}
	var u interface{}
	err = yaml.Unmarshal(file, &u)
	if err != nil {
		return err
	}
	finishResult, err := jsoniter.Marshal(&u)
	if err != nil {
		return err
	}

	filename, err := os.OpenFile(FileTo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer filename.Close()
	if err := ioutil.WriteFile(FileTo, finishResult, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func main() {
	nameInput := flag.String("input", "file.yaml", "Path to config file")
	nameOutput := flag.String("output", "file.json", "Path to config file")
	from := flag.String("from", "yaml", "Path to config file")
	to := flag.String("to", "json", "Path to config file")
	flag.Parse()
	if *from == "yaml" && *to == "json" {
		if err := fromYtoJ(*nameInput, *nameOutput); err != nil {
			panic(err)
		}
	}
	if *from == "json" && *to == "yaml" {
		if err := fromJtoY(*nameInput, *nameOutput); err != nil {
			panic(err)
		}
	}
}
