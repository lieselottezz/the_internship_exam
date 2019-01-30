package main

import (
    "fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"math/rand"
	"time"
	s "strings"
)

// Categories struct (metadata for each category)
type Categories struct {
	Categories		[]Category		`json:"categories"`
}

// Category struct
type Category struct {
	CategoryName	string			`json:"category_name"`
	FileName		string			`json:"file_name"`
}

// Words struct (a list of word)
type Words struct {
	Words 			[]Word 			`json:"words"`
}

// Word struct
type Word struct {
	Word			string			`json:"word"`
	Hint			string			`json:"hint"`
}

func main() {
	categories, err := getIndexCategory()

	if err != nil {
		fmt.Println(err)
		return
	}

	var index int
	fmt.Printf("> ")
	fmt.Scan(&index)
	categoryFilename := getCategoryFile(categories, index)
	words, err := getWordlistFromFile(categoryFilename)

	if err != nil {
		fmt.Println(err)
		return
	}

	randIndex := randomNumber(len(words.Words))
	word, hint := words.Words[randIndex].Word, words.Words[randIndex].Hint
	play(word, hint)
}

// Get metadata for each category
func getIndexCategory() (Categories, error) {
	var categories Categories
	categoriesFile, err := os.Open("./words/categories.json")
	defer categoriesFile.Close()
	byteValCategories, _ := ioutil.ReadAll(categoriesFile)
	json.Unmarshal(byteValCategories, &categories)
	fmt.Println("Select Category:")

	for i := 0; i < len(categories.Categories); i++ {
		fmt.Printf("%d: %s\n", i+1, categories.Categories[i].CategoryName)
	}

	json.Unmarshal(byteValCategories, &categories)
	return categories, err
}

// Get a category by index
func getCategoryFile(categories Categories, index int) string {
	fmt.Println("Select Category:")

	for i := 0; i < len(categories.Categories); i++ {
		fmt.Printf("%d: %s\n", i+1, categories.Categories[i].CategoryName)
	}

	// selectedCategoryName := categories.Categories[index-1].CategoryName
	// selectedCategoryFilename := categories.Categories[index-1].FileName
	// fmt.Printf("The selected category is %d : %s\n", index, selectedCategoryName)

	return categories.Categories[index-1].FileName
}

// Get word list from JSON file
func getWordlistFromFile(filename string) (Words, error) {
	var words Words
	wordsFile, err := os.Open("./words/"+filename)
	defer wordsFile.Close()
	byteValWords, _ := ioutil.ReadAll(wordsFile)
	json.Unmarshal(byteValWords, &words)
	return words, err
}

// Random number with random seed value
func randomNumber(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

// Main process for running the game
func play(word string, hint string) {

	// Init basic parameters
	life, score := 10, 0
	var guessed []string
	var wrongGuessList []string
	var displayWord []string

	for i := 0; i < len(word); i++ {
		wordStr := string(word[i])
		if isAlpha(wordStr) {
			displayWord = append(displayWord, "_")
		} else {
			displayWord = append(displayWord, wordStr)
		}
	}

	wordLowercase := s.ToLower(word)
	
	fmt.Printf("Hint: \"%s\"\n", hint)
	
	for life > 0 && containsInSlice(displayWord, "_") {

		if life == 10 {
			fmt.Printf("%v    Score %d, remaining wrong guess %d\n", displayWord, score, life)
		} else {
			fmt.Printf("%v    Score %d, remaining wrong guess %d, wrong guessed: %v\n", displayWord, score, life, wrongGuessList)
		}

		var guessChar string
		fmt.Printf("Please enter one character to guess\n> ")
		fmt.Scanln(&guessChar)
		guessChar = s.ToLower(guessChar)

		if !isAlpha(guessChar) || len(guessChar) != 1{
			fmt.Println("The input is invalid. Please enter a new one.")
			continue
		}

		if containsInSlice(guessed, guessChar) {
			fmt.Println("The input character has guessed already")
			continue
		}

		guessed = append(guessed, guessChar)
		
		if s.Contains(wordLowercase, guessChar) {
			numberOfChar := s.Count(wordLowercase, guessChar)
			fmt.Printf("%s is in the word\n%d point for corrected guess !!\n", guessChar, 15*numberOfChar)
			if numberOfChar == 1{
				index := s.Index(wordLowercase, guessChar)
				displayWord[index] = string(word[index])
			} else {
				for i, element := range wordLowercase {
					if string(element) == guessChar {
						displayWord[i] = string(word[i])
					}
				}
			}
			score += 15 * numberOfChar

		} else {
			fmt.Printf("%s isn't in the word\n", guessChar)
			wrongGuessList = append(wrongGuessList, guessChar)
			life--
			score -= 10
		}
	}
	if life == 0 {
		fmt.Printf("Game over\nThe answer is \"%s\"\n", word)
	} else {
		fmt.Printf("Congratulation !!\nThe answer is \"%s\"\nYour score: %d\n", word, score)
	}
}

// ref: https://gist.github.com/ammario/d61fb67d15077343e3dd17f2113e4c4b
func isAlpha(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'z' {
			return false
		} else if str[i] > 'Z' && str[i] < 'a' {
			return false
		}
	}
	return true
}

// ref: https://programming.guide/go/find-search-contains-slice.html
func containsInSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
