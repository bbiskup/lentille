package parser

import (
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"reflect"
)

type PromptDictList []yaml.Map

func GetChildList(config *yaml.File) (childList yaml.List, err error) {
	child, err := yaml.Child(config.Root, "fragments")
	if err != nil {
		fmt.Print("Child error:", err)
		return nil, err
	}
	fmt.Println("child", child, reflect.TypeOf(child))

	childList, childCastOK := child.(yaml.List)
	if !childCastOK {
		fmt.Println("Child cast error")
		return nil, errors.New("Cast failed")
	}
	return childList, nil
}

func Parse(configFileName string) (result PromptDictList, err error) {
	config, err := yaml.ReadFile(configFileName)
	if err != nil {
		log.Printf("Error reading YAML file %s: %s", configFileName, err)
		return nil, err
	}

	childList, err := GetChildList(config)

	if err != nil {
		log.Printf("Could not get child list: %s", err)
	}

	for _, item := range childList {
		item := item.(yaml.Map)
		fmt.Println("item", item, reflect.TypeOf(item), item["name"])
		//result[item["name"]] = item

		dummy := yaml.File{Root: item}
		num, err := dummy.GetInt("num")

		if err != nil {
			log.Printf("Could not get num : %s", err)
		} else {
			fmt.Println("##", num)
		}
	}

	return
}
