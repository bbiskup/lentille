package parser

import (
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"lentille/fragments"
	"log"
)

// get list of fragment configurations
func GetChildList(config *yaml.File) (childList yaml.List, err error) {
	child, err := yaml.Child(config.Root, "fragments")
	if err != nil {
		fmt.Print("Child error:", err)
		return nil, err
	}

	childList, childCastOK := child.(yaml.List)
	if !childCastOK {
		fmt.Println("Child cast error")
		return nil, errors.New("Cast failed")
	}
	return childList, nil
}

func Parse(configFileName string) (result []fragments.Fragment, err error) {
	config, err := yaml.ReadFile(configFileName)
	if err != nil {
		log.Printf("Error reading YAML file %s: %s", configFileName, err)
		return nil, err
	}

	childList, err := GetChildList(config)

	if err != nil {
		log.Printf("Could not get child list: %s", err)
		return nil, err
	}

	for _, item := range childList {
		item := item.(yaml.Map)

		// name, err := item["name"]
		// fmt.Println("item", item, reflect.TypeOf(item), item["name"])

		fragmentConf := yaml.File{Root: item}
		name, name_err := fragmentConf.Get("name")
		if name_err != nil {
			log.Fatalf("Missing mandatory parameter 'name'", name_err)
		}

		var fragment fragments.Fragment

		if name == "literal" {
			fragment = fragments.NewLiteralFragment(&fragmentConf)
		}

		if fragment != nil {
			result = append(result, fragment)
		}
		/*if err != nil {
			log.Printf("Could not get num : %s", err)
		} else {
			fmt.Println("##", num)
		}*/
	}

	return
}
