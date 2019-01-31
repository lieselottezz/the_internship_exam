# the_internship_exam
Parsing XML to JSON with Golang

### How to use
1. Install library
```
go get github.com/clbanning/mxj
```
2. The XML file must be in same directory with xml2json.go
3. Run the source code in the command line by using
```
go run xml2json.go $name_of_xml.xml
```
&nbsp; &nbsp; &nbsp; For example
```
go run xml2json.go weather.xml
```
&nbsp; &nbsp; &nbsp; Or build from the source code
```
go build xml2json.go
```
&nbsp; &nbsp; &nbsp; Then run the executable file by using
```
./xml2json $name_of_xml.xml  // for unix
xml2json $name_of_xml.xml    // for windows
```
4. The JSON file will appear in the same directory with xml2json.go

### References
- Trim suffix from filename
  - https://stackoverflow.com/a/21538822
- Get arguements from CLI
  - https://gobyexample.com/command-line-arguments
- Library for parsing XML to Map
  - https://github.com/clbanning/mxj
  - https://godoc.org/github.com/clbanning/mxj
- Library for parsing Map to JSON
  - https://golang.org/pkg/encoding/json/#MarshalIndent
