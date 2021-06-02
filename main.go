package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JoelGauci/apigeeids/feodotracker"
	"github.com/JoelGauci/apigeeids/ruler"
)

// URL of feodo tracker
const feodotracker_url string = "https://feodotracker.abuse.ch/downloads/feodotracker.rules"
const feodotracker_desc string = "FeodoTracker rules"
const feodotracker_file string = "feodotracker.rules"

func main() {

	// create a file for writing rules into it
	// include date and time in the file
	file, err := os.Create(feodotracker_file)
	defer file.Close()

	if err != nil {
		fmt.Println("Cannot create file feodotracker.rules")
		os.Exit(1)
	}

	w := bufio.NewWriter(file)
	defer w.Flush()

	// Request to get ip ranges of Google API endpoints
	var r ruler.Ruler
	r, err = feodotracker.New(feodotracker_url, w, feodotracker_desc)
	if err != nil {
		fmt.Printf("Error while creating FeodoTracker instance (url=%v)\n", feodotracker_url)
		os.Exit(1)
	}
	// get ip addresses/rules/... from a source of trust
	err = r.GetRules()
	if err != nil {
		fmt.Printf("Error when getting rules (err=%v)", err)
		os.Exit(1)
	}
	// mediate ip addresses/rules into snort compliant rules
	r.SetRules()
	// write rules into a target file (.rules extension)
	r.WriteRules()
}
