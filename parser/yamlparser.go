package parser

import (
	"errors"
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"lentille/fragments"
	"log"
	"strings"
)

const (
	COMMAND_SEP = "."
)

type YAMLParser struct {
}

// get list of fragment configurations
func (p *YAMLParser) getChildList(config *yaml.File) (childList yaml.List, err error) {
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

func (p *YAMLParser) buildConfDict(confSection *yaml.Map) (result fragments.ConfDict, err error) {
	result = make(fragments.ConfDict)
	for key, value := range *confSection {
		valueStr, ok := value.(yaml.Scalar)
		if !ok {
			return nil, errors.New("Expected scalar")
		}
		result[key] = valueStr.String()
	}
	return result, nil
}

func (p *YAMLParser) parseCommand(name string) (mainCommand string, subCommand string) {
	commandParts := strings.Split(name, COMMAND_SEP)
	mainCommand = commandParts[0]
	subCommand = strings.Join(commandParts[1:], COMMAND_SEP)
	return
}

func (p *YAMLParser) Parse(configFileName string) (result []fragments.Fragment, err error) {
	config, err := yaml.ReadFile(configFileName)
	if err != nil {
		log.Printf("Error reading YAML file %s: %s", configFileName, err)
		return nil, err
	}

	childList, err := p.getChildList(config)

	if err != nil {
		log.Printf("Could not get child list: %s", err)
		return nil, err
	}

	for _, item := range childList {
		item := item.(yaml.Map)

		confDict, err := p.buildConfDict(&item)
		if err != nil {
			return nil, err
		}

		name, ok := confDict["name"]
		if !ok {
			return nil, errors.New("Missing name")
		}

		isActive := confDict.GetOptionalParamBool("active", "true")
		if !isActive {
			log.Printf("Skipping inactive fragment '%s'", name)
			continue
		}

		confDict["mainCommand"], confDict["subCommand"] = p.parseCommand(name)

		var fragment fragments.Fragment

		if name == "literal" {
			fragment = fragments.NewLiteralFragment(confDict)
		}

		if fragment != nil {
			result = append(result, fragment)
		}
	}

	return
}
