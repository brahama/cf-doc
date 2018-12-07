/*
* Here we should parse/return all the Description, Usage; Parameters and Outputs!
*
 */

package doc

import (
	"bufio"
	"bytes"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// Doc represents the CF template Documentation
type Doc struct {
	Usage      string
	Parameters []Parameter
	Outputs    []Output
}

// Parameter represents CF Parameter block
type Parameter struct {
	Name          string
	Description   string
	Default       string
	Type          string
	AllowedValues string
}

// Output represents the CF Outputs block
type Output struct {
	Name        string
	Description string
	Export      string
}

// We need to parse the yaml from the template file
// **************************************************
type yamlParam struct {
	Type          string `yaml:"Type"`
	Description   string `yaml:"Description"`
	AllowedValues string `yaml:"AllowedValues"`
	Default       string `yaml:"Default"`
}
type yamlParameters struct {
	Parameters map[string]yamlParam `yaml:"Parameters"`
}

// Here we take Description general
type yamlDescription struct {
	Description string `yaml:"Description"`
}

// Here we take Outputs
type yamlOut struct {
	Description string `yaml:"Description"`
	Export      struct {
		Name string `yaml:"Name"`
	} `yaml:"Export"`
}
type yamlOutputs struct {
	Outputs map[string]yamlOut `yaml:"Outputs"`
}

// Lets parse first lines of comments. This is more of a convention
// than anything else.
func getUsage(cfTemplate *[]byte) {
	// The idea was to make the usage here, but dunno how to handle
	// byte type as string without bufio and as it only accepts byte
	// and not pointer, i will do it in same function to avoid copying
	// data. Perhaps later we found a way
}

// Create creates a Doc type with the contents of the Template
func Create(cfTemplate []byte) *Doc {
	cfParam := yamlParameters{}
	cfOut := yamlOutputs{}
	cfDescription := yamlDescription{}

	cfDoc := Doc{}

	// Lets Parse Yaml Parameters and Outputs
	yaml.Unmarshal(cfTemplate, &cfDescription)
	yaml.Unmarshal(cfTemplate, &cfParam)
	yaml.Unmarshal(cfTemplate, &cfOut)

	//The first thing that should appear is the CF Description
	cfDoc.Usage = cfDescription.Description + "\n"

	// Now lets get Usage
	// We transform bytes2Newreader to use bufio
	scanner := bufio.NewScanner(bytes.NewReader(cfTemplate))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			cfDoc.Usage += scanner.Text()[1:] + "\n"
		} else {
			break
		}
	}

	// Now we create and fill the Doc Parameters
	keys := make([]string, 0)
	for k := range cfParam.Parameters {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	i := 0
	for _, k := range keys {
		cfDoc.Parameters = append(cfDoc.Parameters, Parameter{})
		cfDoc.Parameters[i].Name = k
		cfDoc.Parameters[i].Type = cfParam.Parameters[k].Type
		cfDoc.Parameters[i].Description = cfParam.Parameters[k].Description
		cfDoc.Parameters[i].Default = cfParam.Parameters[k].Default
		cfDoc.Parameters[i].AllowedValues = cfParam.Parameters[k].AllowedValues
		i++
	}

	// Create Doc Outputs
	keys = make([]string, 0)
	for k := range cfOut.Outputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	i = 0
	for _, k := range keys {
		cfDoc.Outputs = append(cfDoc.Outputs, Output{})
		cfDoc.Outputs[i].Name = k
		cfDoc.Outputs[i].Description = cfOut.Outputs[k].Description
		cfDoc.Outputs[i].Export = cfOut.Outputs[k].Export.Name
		i++
	}

	return &cfDoc
}
