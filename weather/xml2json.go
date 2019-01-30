package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"github.com/clbanning/mxj"
	"encoding/json"
)

func main() {
	arg := os.Args

	// Check if the number of argument is matched to requirements
	if len(arg) != 2 {
		fmt.Printf("%s\n%s\n", "Failed to convert XML to JSON", "Invalid number of argument")
		return
	}
	
	result := parsingXMLToJSON(arg[1])
	fmt.Println(result)
}

// Get filename from CLI argument then parse map and JSON
func parsingXMLToJSON(filename string) string{
	file, err := os.Open(filename) // Open file with permission
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	xmlVal, err := ioutil.ReadAll(file) // Read xml to array of byte
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	mxj.PrependAttrWithHyphen(false) // Remove hyphen (default prefix of the key in child nodes)

	mapVal, err := mxj.NewMapXml(xmlVal) // Convert byte array to map
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	rootNodeName, err := mapVal.Root() // Get root node name to remove root node (current tag in XML)
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	data, err := json.MarshalIndent(mapVal[rootNodeName], "", "  ") // Get pretty JSON from map
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	err = ioutil.WriteFile(strings.TrimSuffix(filename, path.Ext(filename))+".json", data, 0644) // Write new json file
	if err != nil {
		return "Failed to convert XML to JSON\n" + err.Error()
	}

	return "Convert XML to JSON successfully"
}
