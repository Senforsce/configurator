# configurator

Helps make config files using the builder pattern in go
Just create a file a write a triple per configuration option to generate go code for each file that needs configuration.
(this is a work in progress and the functionalities are limited)

# how to build

```
    go build
    go install
```

# Commands

The CLI is simple, provide an input file path flag `-i` and an output file path `-o`

```
    configurator -i ./input.example -o ./generated.go
```

This generate an incomplete go file, just add your package name

# Example

.input.example

```
tls bool true
myTurtleName string "Hobby"
```

generated go code (in generated.go)

```go

func (config Config) WithTls() Config {
	config.tls = true
	return config
}


func (config Config) WithoutTls() Config {
	config.tls = false
	return config
}


func (config Config) WithMyTurtleName(c string) Config {
	config.myTurtleName = c
	return config
}


type Config struct {
	tls bool
	myTurtleName string
}

func NewConfig() Config{
	return Config{
		tls: true,
		myTurtleName: "Hobby"
	}
}

```

# Roadmap

    - [] Add an option to remove the input
    - [] make the option to remove the input the default and add an option to keep the input
    - [] create an injestible template file to be edited by users
    - [] inject generated code directly inside files that have a special tag

# Credits

    - Author: Abdoul Sy <dreescan@gmail.com>
    - Inspiration & Thanks: anthdm <https://github.com/anthdm>
