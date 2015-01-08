package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"text/template"
)

type Options struct {
	Output      string
	MethodName  string
	ApiEndpoint string
	OutputType  string
}

var methodTemplate string = `
package thegamesdb

import (
        "encoding/xml"
        "net/http"
)

func {{.MethodName}}(parameters map[string]string) (*{{.OutputType}}, error) {
        var output {{.OutputType}}
        endpoint := "{{.ApiEndpoint}}?" + ConvertMapIntoGetParams(parameters)
        resp, err := http.Get(apiEndpoint + endpoint)
        if err != nil {
                return nil, err
        }

        defer resp.Body.Close()
        decoder := xml.NewDecoder(resp.Body)
        decoder.Decode(&output)
        return &output, nil
}
`

func main() {
	options := parseFlags()
	tmpl, err := template.New("method").Parse(methodTemplate)
	if err != nil {
		panic(err)
	}
	createDirectoryIfNotExist(options.Output)
	f, err := os.Create(options.Output)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	defer func() {
		writer.Flush()
		f.Close()
	}()

	err = tmpl.Execute(writer, options)
	if err != nil {
		panic(err)
	}
}

func createDirectoryIfNotExist(file string) {
	directory := filepath.Dir(file)
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		os.MkdirAll(directory, 0777)
	}
}

func parseFlags() Options {
	output := flag.String("output", "", "Output file")
	methodName := flag.String("method", "", "Method name")
	apiEndpoint := flag.String("endpoint", "", "Endpoint of API")
	outputType := flag.String("type", "", "Output Type")
	flag.Parse()
	return Options{
		Output:      *output,
		MethodName:  *methodName,
		ApiEndpoint: *apiEndpoint,
		OutputType:  *outputType,
	}
}
