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

(alternatively a [pre-compiled binary for Linux x86_64](bin/gotemplate) is checked in - caveat emptor)

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

# Call out to a Golang API function from a template
$GOPATH/bin/gotemplate $GOPATH/src/github.com/jim-minter/gotemplate/examples/example3.tmpl
```

## Available functions

In addition to the [default functions](https://golang.org/pkg/text/template/#hdr-Functions) available using text/template, the following are also available:

- [Strings package](https://golang.org/pkg/strings/)
  - [contains](https://golang.org/pkg/strings/#Contains)
  - [containsAny](https://golang.org/pkg/strings/#ContainsAny)
  - [containsRune](https://golang.org/pkg/strings/#ContainsRune)
  - [count](https://golang.org/pkg/strings/#Count)
  - [equalFold](https://golang.org/pkg/strings/#EqualFold)
  - [fields](https://golang.org/pkg/strings/#Fields)
  - [hasPrefix](https://golang.org/pkg/strings/#HasPrefix)
  - [hasSuffix](https://golang.org/pkg/strings/#HasSuffix)
  - [index](https://golang.org/pkg/strings/#Index)
  - [indexAny](https://golang.org/pkg/strings/#IndexAny)
  - [indexByte](https://golang.org/pkg/strings/#IndexByte)
  - [indexRune](https://golang.org/pkg/strings/#IndexRune)
  - [join](https://golang.org/pkg/strings/#Join)
  - [lastIndex](https://golang.org/pkg/strings/#LastIndex)
  - [lastIndexAny](https://golang.org/pkg/strings/#LastIndexAny)
  - [lastIndexByte](https://golang.org/pkg/strings/#LastIndexByte)
  - [map](https://golang.org/pkg/strings/#Map)
  - [repeat](https://golang.org/pkg/strings/#Repeat)
  - [replace](https://golang.org/pkg/strings/#Replace)
  - [split](https://golang.org/pkg/strings/#Split)
  - [splitAfter](https://golang.org/pkg/strings/#SplitAfter)
  - [splitAfterN](https://golang.org/pkg/strings/#SplitAfterN)
  - [splitN](https://golang.org/pkg/strings/#SplitN)
  - [title](https://golang.org/pkg/strings/#Title)
  - [toLower](https://golang.org/pkg/strings/#ToLower)
  - [toTitle](https://golang.org/pkg/strings/#ToTitle)
  - [toUpper](https://golang.org/pkg/strings/#ToUpper)
  - [trim](https://golang.org/pkg/strings/#Trim)
  - [trimLeft](https://golang.org/pkg/strings/#TrimLeft)
  - [trimPrefix](https://golang.org/pkg/strings/#TrimPrefix)
  - [trimRight](https://golang.org/pkg/strings/#TrimRight)
  - [trimSpace](https://golang.org/pkg/strings/#TrimSpace)
  - [trimSuffix](https://golang.org/pkg/strings/#TrimSuffix)

Making additional functions available is quite straightforward: e-mail or send a pull request if you require others.
