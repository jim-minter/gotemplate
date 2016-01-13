// Copyright 2016 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

var jsonfile string
var yamlfile string
var missingkey string

// setFuncs imports certain functions from the Golang API to make them
// available to templates
func setFuncs(t *template.Template) {
	t.Funcs(template.FuncMap{
		"contains": strings.Contains,
		"containsAny": strings.ContainsAny,
		"containsRune": strings.ContainsRune,
		"count": strings.Count,
		"equalFold": strings.EqualFold,
		"fields": strings.Fields,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		"index": strings.Index,
		"indexAny": strings.IndexAny,
		"indexByte": strings.IndexByte,
		"indexRune": strings.IndexRune,
		"join": strings.Join,
		"lastIndex": strings.LastIndex,
		"lastIndexAny": strings.LastIndexAny,
		"lastIndexByte": strings.LastIndexByte,
		"repeat": strings.Repeat,
		"replace": strings.Replace,
		"split": strings.Split,
		"splitAfter": strings.SplitAfter,
		"splitAfterN": strings.SplitAfterN,
		"splitN": strings.SplitN,
		"title": strings.Title,
		"toLower": strings.ToLower,
		"toTitle": strings.ToTitle,
		"toUpper": strings.ToUpper,
		"trim": strings.Trim,
		"trimLeft": strings.TrimLeft,
		"trimPrefix": strings.TrimPrefix,
		"trimRight": strings.TrimRight,
		"trimSpace": strings.TrimSpace,
		"trimSuffix": strings.TrimSuffix,
	})
}

// readFile reads in a given file and unmarshals it according to the passed
// function pointer
func readFile(name string, unmarshal func([]byte, interface{}) error) (map[string]interface{}, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var v map[string]interface{}
	err = unmarshal(b, &v)
	return v, err
}

// environment returns a map of environment variables
func environment() map[string]string {
	rv := make(map[string]string)

	for _, env := range os.Environ() {
		a := strings.SplitN(env, "=", 2)
		rv[a[0]] = a[1]
	}

	return rv
}

// checkError bails with a friendly message if an error occurred
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// usage prints the application usage
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-j JSONFILE | -y YAMLFILE] [-missingkey default|invalid|zero|error] TEMPLATE...\n\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

// parseArgs parses and verifies command line arguments
func parseArgs() {
	flag.Usage = usage

	flag.StringVar(&jsonfile, "j", "", "read data from `JSONFILE`")
	flag.StringVar(&yamlfile, "y", "", "read data from `YAMLFILE`")
	flag.StringVar(&missingkey, "missingkey", "default", "`missingkey` option, see https://golang.org/pkg/text/template/#Template.Option")
	flag.Parse()

	missingkey = strings.ToLower(missingkey)

	if flag.NArg() == 0 ||
		(jsonfile != "" && yamlfile != "") ||
		(missingkey != "default" && missingkey != "invalid" && missingkey != "zero" && missingkey != "error") {
		flag.Usage()
	}
}

// main
func main() {
	// parse command line arguments
	parseArgs()

	_, name := path.Split(flag.Args()[0])
	t := template.New(name)
	setFuncs(t)

	// read in specified template files
	t, err := t.ParseFiles(flag.Args()...)
	checkError(err)

	// populate template data if available
	var data map[string]interface{}
	if jsonfile != "" {
		data, err = readFile(jsonfile, json.Unmarshal)
	} else if yamlfile != "" {
		data, err = readFile(yamlfile, yaml.Unmarshal)
	} else {
		data = make(map[string]interface{})
	}
	checkError(err)

	// always add the current environment under the "env" key
	data["env"] = environment()

	// execute template
	t.Option("missingkey=" + missingkey)
	err = t.Execute(os.Stdout, data)
	checkError(err)
}
