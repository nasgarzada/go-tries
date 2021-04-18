package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var guessCount = 0

func main() {
	var randomValue = generateRandom(200)
	for true {
		fmt.Print("Enter your guess : ")
		guessedNumber, _ := strconv.Atoi(strings.TrimSpace(readData()))
		if guess(randomValue, guessedNumber) || !continueGame(randomValue) {
			break
		}
	}
}

func readData() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func generateRandom(size int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(size) + 1
}

func continueGame(randomValue int) bool {
	if guessCount%5 != 0 {
		return true
	}
	fmt.Print("Do you wanna continue (y/n): ")
	choice := readData()
	if strings.ToLower(strings.TrimSpace(choice)) == "n" {
		log.Println("You gave up early ((. Number was:", randomValue)
		return false
	}
	return true
}

func guess(actualValue int, guessedValue int) bool {
	guessCount++
	if guessedValue > actualValue {
		log.Println("It's just a little bit HIGH")
	} else if guessedValue < actualValue {
		log.Println("It's just a little bit LOW")
	} else {
		log.Println("You find correctly!! Congrats!!")
		return true
	}
	return false
}
