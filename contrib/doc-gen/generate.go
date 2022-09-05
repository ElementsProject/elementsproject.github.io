// Copied from https://github.com/bitcoin-core/bitcoincore.org/blob/master/contrib/doc-gen/generate.go
// All changes are marked with `ELEMENTS:`

// The original file is licensed with the following terms:
//
// The MIT License (MIT)
//
// Copyright (c) 2014 Michael Rose
// Copyright (c) 2015 Sylvain Durand
// Copyright (c) 2016 Respective Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

// ELEMENTS: changed this to elements-cli
const BITCOIN_COMMAND = "elements-cli"

type Command struct {
	Name        string
	Description string
}

type Group struct {
	Index    int
	Name     string
	Commands []Command
}

type CommandData struct {
	Version     string
	Name        string
	Description string
	Group       string
	Permalink   string
}

func getVersion() string {
	allInfo := run("getnetworkinfo")
	var f interface{}
	err := json.Unmarshal([]byte(allInfo), &f)
	if err != nil {
		panic("Cannot read network info as JSON")
	}
	m := f.(map[string]interface{})

	numv := int(m["version"].(float64))
	// ELEMENTS: added the numv%100 extra version part here
	v := fmt.Sprintf("%d.%d.%d", numv/10000, (numv/100)%100, numv%100)
	return v
}

func main() {
	version := getVersion()

	first := run("help")
	split := strings.Split(first, "\n")

	groups := make([]Group, 0)
	commands := make([]Command, 0)
	lastGroupName := ""

	for _, line := range split {
		if len(line) > 0 {
			if strings.HasPrefix(line, "== ") {
				if len(commands) != 0 {
					g := Group{
						Name:     lastGroupName,
						Commands: commands,
						Index:    len(groups),
					}
					groups = append(groups, g)
					commands = make([]Command, 0)
				}
				lastGroupName = strings.ToLower(line[3 : len(line)-3])
			} else {
				name := strings.Split(line, " ")[0]
				desc := run("help", name)
				comm := Command{
					Name:        name,
					Description: desc,
				}
				commands = append(commands, comm)
			}
		}
	}

	g := Group{
		Name:     lastGroupName,
		Commands: commands,
		Index:    len(groups),
	}
	groups = append(groups, g)

	tmpl := template.Must(template.ParseFiles("command-template.html"))

	for _, group := range groups {
		groupname := group.Name
		dirname := fmt.Sprintf("../../_doc/en/%s/rpc/%s/", version, groupname)
		err := os.MkdirAll(dirname, 0777)
		if err != nil {
			log.Fatalf("Cannot make directory %s: %s", dirname, err.Error())
		}
		for _, command := range group.Commands {
			name := command.Name
			address := fmt.Sprintf("%s%s.html", dirname, name)
			permalink := fmt.Sprintf("/en/doc/%s/rpc/%s/%s/", version, groupname, name)
			err = tmpl.Execute(open(address), CommandData{
				Version:     version,
				Name:        name,
				Description: command.Description,
				Group:       groupname,
				Permalink:   permalink,
			})
			if err != nil {
				log.Fatalf("Cannot make command file %s: %s", name, err.Error())
			}
		}
	}

	address := fmt.Sprintf("../../_doc/en/%s/rpc/index.html", version)
	permalink := fmt.Sprintf("/en/doc/%s/rpc/", version)
	err := tmpl.Execute(open(address), CommandData{
		Version:     version,
		Name:        "rpcindex",
		Description: "",
		Group:       "index",
		Permalink:   permalink,
	})
	if err != nil {
		log.Fatalf("Cannot make index file: %s", err.Error())
	}

	address = fmt.Sprintf("../../_doc/en/%s/index.html", version)
	permalink = fmt.Sprintf("/en/doc/%s/", version)
	err = tmpl.Execute(open(address), CommandData{
		Version:     version,
		Name:        "index",
		Description: "",
		Group:       "index",
		Permalink:   permalink,
	})
	if err != nil {
		log.Fatalf("Cannot make index file: %s", err.Error())
	}
}

func open(path string) io.Writer {
	f, err := os.Create(path)
	// not closing, program will close sooner
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", path, err.Error())
	}
	return f
}

func run(args ...string) string {
	// ELEMENTS: added the -chain=elementsregtest argument here
	out, err := exec.Command(BITCOIN_COMMAND, append([]string{"-chain=elementsregtest"}, args...)...).CombinedOutput()
	if err != nil {
		log.Fatalf("Cannot run %s: %s, is bitcoind running?", BITCOIN_COMMAND, err.Error())
	}

	return string(out)
}
