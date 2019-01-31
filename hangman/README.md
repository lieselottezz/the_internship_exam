# the_internship_exam

A hangman game powered by Golang

### How to play
Run the game by using
```
go run hangman.go
```
Or build it with the source code
```
go build hangman.go
```
Then run the execuable file
```
./hangman    // for unix
hangman      // for windows
```

You can try one of three categories of words in the game
1. Famous places in Thailand
2. Japanese animation name
3. Province of Japan

### Add more category of words
You can add it more by follow this instruction
1. Create JSON file for words and hints in "words" folder, The format will be look like
```json
{
    "words": [
        {
            "word": "this is word 1",
            "hint": "this is hint 1"
        },
        {
            "word": "this is word 2", 
            "hint": "this is hint 2"
        },
        {
            "word": "this is word N", 
            "hint": "this is hint N"
        }
    ]
}
```
**Note:** In this JSON file, you can have any number of the word as you want but please be careful about JSON format when you are editing

2. Edit the categories.json in "words" folder, The format will be look like
```json
{
    "categories": [
        {
            "category_name": "Famous places in Thailand",
            "file_name": "th_places.json"
        },
        {
            "category_name": "Japanese animation name",
            "file_name": "anime_name.json"
        },
        {
            "category_name": "This is new category",
            "file_name": "new_category.json"
        }
    ]
}
```
### References
* Parsing JSON files With Golang
  * https://tutorialedge.net/golang/parsing-json-with-golang/
* Conversions to and from string representations of basic data types
  * https://golang.org/pkg/strconv/
* Simple functions to manipulate UTF-8 encoded strings
  * https://golang.org/pkg/strings/
* Random integer with random seed
  * https://flaviocopes.com/go-random/
* isAlpha function
  * https://gist.github.com/ammario/d61fb67d15077343e3dd17f2113e4c4b
* Go: Find an element in a slice
  * https://programming.guide/go/find-search-contains-slice.html
