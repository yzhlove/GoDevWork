package main

import (
	"log"
	"os"
	"text/template"
)

const templateText = `
    # GENERAL VALUE
    NAME: {{.Name}}
    # IF STRING
    {{if ne .Name "Bob"}}No, I'm Not Bob{{end}}

    # IF NUMERIC
    {{- if le .Age 30}}
    I am a senior one
    {{else}}
    I am a little one
    {{end}}

    # IF BOOLEAN
    {{- if .Boy}}
    It's a Boy 
    {{else}}
    It's a Girl
    {{end}}

    # RANGE
    {{- range $index, $friend := .Friends}}
    Friend {{$index}}: {{$friend}}
    {{- end}}

    # EXISTENCE
    {{- with .Gift -}}
    I have a gift: {{.}}
    {{else}}
    I have not a gift.
    {{end}}

`

func main() {
	type Recipient struct {
		Name    string
		Age     int
		Boy     bool
		Friends []string
		Gift    string
	}

	recipient := Recipient{
		Name:    "Jack",
		Age:     30,
		Friends: []string{"Bob", "Json"},
		Boy:     true,
	}

	t := template.Must(template.New("anyname").Parse(templateText))
	err := t.Execute(os.Stdout, recipient)
	if err != nil {
		log.Println("Executing template:", err)
	}
}
