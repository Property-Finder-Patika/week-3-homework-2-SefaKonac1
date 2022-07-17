package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	/*prints statistics belong to game*/
	statistics := numberGuessingGame()

	fmt.Println("Guessed Numbers History >")
	fmt.Printf("%+v\n", statistics)
	fmt.Printf("In total maden guess number is %d", len(statistics))

}

/*generates numbers as string ranging from between 1000 - 9999*/
func generateSecretNumber() string {

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(9999-1000) + 1000)

}

func numberGuessingGame() []string {

	var guessedNumbers []string

	secretNumber := generateSecretNumber()

	for {

		var guess string
		fmt.Println("Enter Your Guess: ")
		fmt.Scanln(&guess)

		guessedNumbers = append(guessedNumbers, guess)

		if guess == secretNumber {
			fmt.Println("You found it ! : %d", guess)
			return guessedNumbers
		}

		fmt.Println("Try again due following hint !")
		pos, nev := takeClue(secretNumber, guess)

		fmt.Println("+%d, %d", pos, nev)

	}

	return guessedNumbers
}

func takeClue(secretNumber string, guess string) (int, int) {

	pos, nev := 0, 0

	for i := 0; i < len(secretNumber); i++ {

		if secretNumber[i] == guess[i] {
			pos += 1
		} else {

			if strings.Index(secretNumber, guess[i]) != -1 {
				nev -= 1
			}
		}

	}
	return pos, nev
}
