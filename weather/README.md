# the_internship_exam
Parsing XML to JSON with Golang

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

### How to use
1. Install Golang
    - https://golang.org/doc/install
2. Run the source code in the command line by using
```
go run xml2json.go $name_of_xml.xml
```
  - For example
```
go run xml2json.go weather.xml
```
  - Or build from the source code
```
go build xml2json.go
```
  - Then run the executable file by using
```
./xml2json $name_of_xml.xml
```
