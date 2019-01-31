package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"math/rand"
	"time"
	"strconv"
	s "strings"
)

// Categories struct (metadata for each category)
type Categories struct {
	Categories []Category `json:"categories"`
}

// Category struct
type Category struct {
	CategoryName string `json:"category_name"`
	FileName     string `json:"file_name"`
}

// Words struct (a list of word)
type Words struct {
	Words []Word `json:"words"`
}

// Word struct
type Word struct {
	Word string `json:"word"`
	Hint string `json:"hint"`
}

func main() {
	categories, err := getIndexCategory()

	if err != nil {
		fmt.Println(err)
		return
	}

	words, err := selectCategory(categories)

	if err != nil {
		fmt.Println(err)
		return
	}

	randIndex := randomNumber(len(words.Words))
	word, hint := words.Words[randIndex].Word, words.Words[randIndex].Hint
	play(word, hint)
}

// Main process for running the game
func play(word string, hint string) {

	// Initialize basic parameters
	life, score := 10, 0
	var guessed []string
	var wrongGuessList []string
	displayWord := createDisplayWord(word)
	wordLowercase := s.ToLower(word)
	fmt.Printf("\nHint: \"%s\"\n\n", hint)
	
	// Initialize game process
	for life > 0 && containsInSlice(displayWord, "_") {

		if life == 10 {
			fmt.Printf("%v    Score %d, remaining wrong guess %d\n", displayWord, score, life)
		} else {
			fmt.Printf("%v    Score %d, remaining wrong guess %d, wrong guessed: %v\n", displayWord, score, life, wrongGuessList)
		}

		// Get the input character
		var guessChar string
		fmt.Printf("Please enter one character to guess\n> ")
		fmt.Scanln(&guessChar)
		guessChar = s.ToLower(guessChar)

		// Validate character input
		if !isAlpha(guessChar) || len(guessChar) != 1{
			fmt.Println("The input is invalid. Please enter a new one.")
			continue
		}

		if containsInSlice(guessed, guessChar) {
			fmt.Println("The input character has guessed already")
			continue
		}

		// Append a used character into the slice
		guessed = append(guessed, guessChar)
		
		// Compare between the input character and the answer
		if s.Contains(wordLowercase, guessChar) {
			numberOfChar := s.Count(wordLowercase, guessChar)
			fmt.Printf("\"%s\" is in the word\n%d point for the correct guess !!\n", guessChar, 15*numberOfChar)
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
			fmt.Printf("\"%s\" isn't in the word\n-10 point for the incorrect guess\n", guessChar)
			wrongGuessList = append(wrongGuessList, guessChar)
			life--
			score -= 10
		}
	}

	// Check the status of completing the game
	if life == 0 {
		fmt.Printf("Game over\nThe answer is \"%s\"\n", word)
	} else {
		fmt.Printf("Congratulation !!\nThe answer is \"%s\"\nYour score: %d\n", word, score)
	}
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

// Select category from the list
func selectCategory(categories Categories) (Words, error) {
	var index string
	var words Words
	var err error
	for true {
		fmt.Printf("Enter a number to select the category\n> ")
		fmt.Scan(&index)
		indexInt, _ := strconv.Atoi(index)
		if indexInt >= 1 && indexInt <= len(categories.Categories) {
			categoryFilename := categories.Categories[indexInt-1].FileName
			words, err = getWordlistFromFile(categoryFilename)
			break
		}
		fmt.Println("Invalid input, Please try again")
	}
	return words, err
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

// Ref: https://flaviocopes.com/go-random/
// Random number with random seed value
func randomNumber(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

// Create a slice of underscore to display the word within a length of the word
func createDisplayWord(word string) []string {
	var displayWord []string

	for i := 0; i < len(word); i++ {
		wordStr := string(word[i])
		if isAlpha(wordStr) {
			displayWord = append(displayWord, "_")
		} else {
			displayWord = append(displayWord, wordStr)
		}
	}

	return displayWord
}

// Ref: https://gist.github.com/ammario/d61fb67d15077343e3dd17f2113e4c4b
// Return true if string is alphabet
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

// Ref: https://programming.guide/go/find-search-contains-slice.html
// Return true if element is in the slice 
func containsInSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
