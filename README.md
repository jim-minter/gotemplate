# gotemplate

Proof of concept portable/container-friendly command-line tool to wrap Golang's
text/template capabilities.

## Author

- Jim Minter (jminter at redhat.com)

## Installation

```bash
go get github.com/jim-minter/gotemplate
$GOPATH/bin/gotemplate
```

## Documentation

Golang's text/template capabilities are powerful and are fully documented at
https://golang.org/pkg/text/template/.

```
Usage: gotemplate [-j JSONFILE | -y YAMLFILE] [-missingkey default|invalid|zero|error] TEMPLATE...

  -j JSONFILE
    	read data from JSONFILE
  -missingkey missingkey
    	missingkey option, see https://golang.org/pkg/text/template/#Template.Option (default "default")
  -y YAMLFILE
    	read data from YAMLFILE
```

## Examples

```bash
# Populate a freeform output file based on environment variables
$GOPATH/bin/gotemplate $GOPATH/src/github.com/jim-minter/gotemplate/examples/example1.tmpl

# Same, but fail on any required but missing environment variables
$GOPATH/bin/gotemplate -missingkey error $GOPATH/src/github.com/jim-minter/gotemplate/examples/example1.tmpl
echo $?

# Populate an XML output file based on JSON input
$GOPATH/bin/gotemplate -j $GOPATH/src/github.com/jim-minter/gotemplate/examples/example2.json $GOPATH/src/github.com/jim-minter/gotemplate/examples/example2.tmpl

# Populate an XML output file based on YAML input
$GOPATH/bin/gotemplate -y $GOPATH/src/github.com/jim-minter/gotemplate/examples/example2.yaml $GOPATH/src/github.com/jim-minter/gotemplate/examples/example2.tmpl
```
