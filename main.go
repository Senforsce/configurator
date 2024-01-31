package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

var inputFilePath string
var outputFilePath string

func init() {
	const (
		defaultInputFilePath  = "./.input"
		inputUsage            = "the input file values are for each line name type default"
		defaultOutputFilePath = "./output.go"
		outputUsage           = "the output go file path"
	)
	flag.StringVar(&inputFilePath, "input_file_path", defaultInputFilePath, inputUsage)
	flag.StringVar(&inputFilePath, "i", defaultInputFilePath, inputUsage+" (shorthand)")

	flag.StringVar(&outputFilePath, "output_file_path", defaultOutputFilePath, outputUsage)
	flag.StringVar(&outputFilePath, "o", defaultOutputFilePath, outputUsage+" (shorthand)")
}

var configTmplBegin = `
type Config struct {`
var configTmplEnd = `}`
var defaultConfigTmplBegin = `
func NewConfig() Config{
	return Config{`
var defaultConfigTmplEnd = `	}
}
`
var withFunc = `
func (config Config) With__NAME_PASCAL__(c __TYPE__) Config {
	config.__NAME__ = c
	return config
}
`
var withFuncTrue = `
func (config Config) With__NAME_PASCAL__() Config {
	config.__NAME__ = true
	return config
}
`
var withoutFunc = `
func (config Config) Without__NAME_PASCAL__() Config {
	config.__NAME__ = false
	return config
}
`

func main() {

	var outputContent = configTmplBegin
	var defaultOutputContent = defaultConfigTmplBegin
	flag.Parse()
	file, err := os.Open(inputFilePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	output, err := os.Create(outputFilePath)

	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		triple := strings.Split(line, " ")
		name := triple[0]
		typeValue := triple[1]
		defaultValue := triple[2]

		if name == "" || typeValue == "" {
			panic("malformed config file it should be : myConfigName string default for each line")
		}

		outputContent += "\n\t" + name + " " + typeValue
		defaultOutputContent += "\n\t\t" + name + ": " + defaultValue + ","

		if typeValue == "bool" {
			io.WriteString(output, MakeCustonFunc(Params{in: withFuncTrue, name: name, typeValue: ""}))
			io.WriteString(output, "\n")
			io.WriteString(output, MakeCustonFunc(Params{in: withoutFunc, name: name, typeValue: ""}))
			io.WriteString(output, "\n")
		} else {
			io.WriteString(output, MakeCustonFunc(Params{in: withFunc, name: name, typeValue: typeValue}))
			io.WriteString(output, "\n")
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	outputContent += "\n" + configTmplEnd
	defaultOutputContent += "\n" + defaultConfigTmplEnd

	io.WriteString(output, outputContent)
	io.WriteString(output, "\n")
	io.WriteString(output, defaultOutputContent)
	io.WriteString(output, "\n")

	fmt.Println("All Done !")

}

type Params struct {
	in, name, typeValue string
}

func MakeCustonFunc(p Params) string {
	re := regexp.MustCompile(`__NAME__`)
	withFuncWithName := re.ReplaceAllString(p.in, p.name)

	rePascal := regexp.MustCompile(`__NAME_PASCAL__`)
	withFuncWithNameFinal := rePascal.ReplaceAllString(withFuncWithName, strcase.ToCamel(p.name))

	reType := regexp.MustCompile(`__TYPE__`)

	if p.typeValue != "" {
		return reType.ReplaceAllString(withFuncWithNameFinal, p.typeValue)
	}

	return withFuncWithNameFinal

}
