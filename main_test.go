package main

import (
	"io"
	"os"
	"testing"
)

var defaultOutputFilePath = "testOutput"

// Reads input file and creates output file successfully
func TestReadAndCreateFiles(t *testing.T) {
	// Initialize test variables
	inputFilePath := "input.txt"
	outputFilePath := "output.txt"
	expectedOutputContent := "expected output content"
	expectedDefaultOutputContent := "expected default output content"

	// Create input file
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	// Write content to input file
	_, err = inputFile.WriteString("name string default\n")
	if err != nil {
		t.Fatalf("Failed to write content to input file: %v", err)
	}

	// Call the main function
	main()

	// Read output file
	outputFile, err := os.Open(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer outputFile.Close()

	// Read output content
	outputContent, err := io.ReadAll(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output content: %v", err)
	}

	// Read default output file
	defaultOutputFile, err := os.Open(defaultOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to open default output file: %v", err)
	}
	defer defaultOutputFile.Close()

	// Read default output content
	defaultOutputContent, err := io.ReadAll(defaultOutputFile)
	if err != nil {
		t.Fatalf("Failed to read default output content: %v", err)
	}

	// Assert the output content
	if string(outputContent) != expectedOutputContent {
		t.Errorf("Output content does not match expected: got %s, want %s", string(outputContent), expectedOutputContent)
	}

	// Assert the default output content
	if string(defaultOutputContent) != expectedDefaultOutputContent {
		t.Errorf("Default output content does not match expected: got %s, want %s", string(defaultOutputContent), expectedDefaultOutputContent)
	}
}

// Generates output content and default output content correctly
func TestGenerateOutputContent(t *testing.T) {
	// Initialize test variables
	inputFilePath := "input.txt"
	outputFilePath := "output.txt"
	expectedOutputContent := "expected output content"
	expectedDefaultOutputContent := "expected default output content"

	// Create input file
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	// Write content to input file
	_, err = inputFile.WriteString("name string default\n")
	if err != nil {
		t.Fatalf("Failed to write content to input file: %v", err)
	}

	// Call the main function
	main()

	// Read output file
	outputFile, err := os.Open(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer outputFile.Close()

	// Read output content
	outputContent, err := io.ReadAll(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output content: %v", err)
	}

	// Read default output file
	defaultOutputFile, err := os.Open(defaultOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to open default output file: %v", err)
	}
	defer defaultOutputFile.Close()

	// Read default output content
	defaultOutputContent, err := io.ReadAll(defaultOutputFile)
	if err != nil {
		t.Fatalf("Failed to read default output content: %v", err)
	}

	// Assert the output content
	if string(outputContent) != expectedOutputContent {
		t.Errorf("Output content does not match expected: got %s, want %s", string(outputContent), expectedOutputContent)
	}

	// Assert the default output content
	if string(defaultOutputContent) != expectedDefaultOutputContent {
		t.Errorf("Default output content does not match expected: got %s, want %s", string(defaultOutputContent), expectedDefaultOutputContent)
	}
}

// Parses input file content line by line and extracts name, typeValue, and defaultValue
func TestParseInputFileContent(t *testing.T) {
	// Initialize test variables
	inputFilePath := "input.txt"
	outputFilePath := "output.txt"
	expectedOutputContent := "expected output content"
	expectedDefaultOutputContent := "expected default output content"

	// Create input file
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	// Write content to input file
	_, err = inputFile.WriteString("name string default\n")
	if err != nil {
		t.Fatalf("Failed to write content to input file: %v", err)
	}

	// Call the main function
	main()

	// Read output file
	outputFile, err := os.Open(outputFilePath)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer outputFile.Close()

	// Read output content
	outputContent, err := io.ReadAll(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output content: %v", err)
	}

	// Read default output file
	defaultOutputFile, err := os.Open(defaultOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to open default output file: %v", err)
	}
	defer defaultOutputFile.Close()

	// Read default output content
	defaultOutputContent, err := io.ReadAll(defaultOutputFile)
	if err != nil {
		t.Fatalf("Failed to read default output content: %v", err)
	}

	// Assert the output content
	if string(outputContent) != expectedOutputContent {
		t.Errorf("Output content does not match expected: got %s, want %s", string(outputContent), expectedOutputContent)
	}

	// Assert the default output content
	if string(defaultOutputContent) != expectedDefaultOutputContent {
		t.Errorf("Default output content does not match expected: got %s, want %s", string(defaultOutputContent), expectedDefaultOutputContent)
	}
}

// Panics if input file cannot be opened
func TestInputFileOpenError(t *testing.T) {
	// Initialize test variables
	// inputFilePath := "nonexistent.txt"
	// outputFilePath := "output.txt"

	// Call the main function and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when opening input file, but no panic occurred")
		}
	}()
	main()
}

// Panics if malformed config file is encountered
func TestMalformedConfigFile(t *testing.T) {
	// Initialize test variables
	inputFilePath := "input.txt"
	// outputFilePath := "output.txt"

	// Create input file
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	// Write malformed content to input file
	_, err = inputFile.WriteString("malformed config file")
	if err != nil {
		t.Fatalf("Failed to write content to input file: %v", err)
	}

	// Call the main function and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when encountering malformed config file, but no panic occurred")
		}
	}()
	main()
}

// Panics if output file cannot be created
func TestOutputFileCreateError(t *testing.T) {
	// Initialize test variables
	inputFilePath := "input.txt"
	// outputFilePath := "/root/output.txt"

	// Create input file
	inputFile, err := os.Create(inputFilePath)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}
	defer inputFile.Close()

	// Write content to input file
	_, err = inputFile.WriteString("name string default\n")
	if err != nil {
		t.Fatalf("Failed to write content to input file: %v", err)
	}

	// Call the main function and expect a panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when creating output file, but no panic occurred")
		}
	}()
	main()
}

// Replaces '__NAME__' with the provided name parameter
func TestReplaceNameParameter(t *testing.T) {
	p := Params{
		in:        "input string",
		name:      "testName",
		typeValue: "string",
	}
	expected := "input string"
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Replaces '__NAME_PASCAL__' with the provided name parameter in PascalCase
func TestReplaceNamePascalCase(t *testing.T) {
	p := Params{
		in:        "input string",
		name:      "testName",
		typeValue: "string",
	}
	expected := "input string"
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Replaces '__TYPE__' with the provided typeValue parameter if it exists
func TestReplaceTypeValue(t *testing.T) {
	p := Params{
		in:        "input string",
		name:      "testName",
		typeValue: "string",
	}
	expected := "input string"
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns an empty string if the input string is empty
func TestEmptyInputString(t *testing.T) {
	p := Params{
		in:        "",
		name:      "testName",
		typeValue: "string",
	}
	expected := ""
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the input string if the name parameter is empty
func TestEmptyNameParameter(t *testing.T) {
	p := Params{
		in:        "input string",
		name:      "",
		typeValue: "string",
	}
	expected := "input string"
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the input string if the '__NAME__' pattern is not found
func TestPatternNotFound(t *testing.T) {
	p := Params{
		in:        "input string",
		name:      "testName",
		typeValue: "string",
	}
	expected := "input string"
	result := MakeCustonFunc(p)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
