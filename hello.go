package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v2"
)

func FromJsonToYaml(fileFrom string, FileTo string) error {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		return err
	}

	var u interface{}
	if err := jsoniter.Unmarshal(file, &u); err != nil {
		return err
	} //< Лучше не игнорировать err возвращаемый после парсинга. Вдруг файл файл не корректный
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

func FromYamlToJson(fileFrom string, FileTo string) {
	file, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		panic(err)
	}
	var u interface{}
	err = yaml.Unmarshal(file, &u)
	if err != nil {
		panic(err)
	}
	finishResult, err := jsoniter.Marshal(&u)
	if err != nil {
		panic(err)
	}

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
		if err := FromJsonToYaml(*nameInput, *nameOutput); err != nil {
			panic(err)
		}
	}
}
