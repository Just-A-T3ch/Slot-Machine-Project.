package main

import (
	"fmt"
	"math/rand"
)

func getName() string {
	name := ""

	fmt.Println("Welcome to Jay's Casino Library...")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s, let's play!\n", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet, or 0 to quit (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than balance.")
		} else {
			break
		}
	}
	return bet
}

func generateSymbolArray(symbols map[string]uint) []string {
	symbolArray := []string{}
	for symbol, count := range symbols {
		for i := uint(0); i < count; i++ {
			symbolArray = append(symbolArray, symbol)
		}
	}
	return symbolArray

}

func getRandomNumber(min int, max int) int {
	//randomNumber := rand.Intn(max-min+1) + min
	randomNumber := rand.Intn(max-min+1) + min
	return randomNumber
}

func getSpin(reel []string, rows int, cols int) [][]string {
	result := [][]string{}

	for i := 0; i < rows; i++ {
		result = append(result, []string{})
	}

	for col := 0; col < cols; col++ {
		selected := map[int]bool{}
		for row := 0; row < rows; row++ {
			for true {
				//randomIndex := getRandomNumber(min: 0, len(reel)-1)
				randomIndex := getRandomNumber(0, len(reel)-1)
				_, exists := selected[randomIndex]
				if !exists {
					selected[randomIndex] = true
					result[row] = append(result[row], reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		for j, symbol := range row {
			fmt.Printf("%s", symbol)
			if j != len(row)-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Println()
	}
}

func checkWin(spin [][]string, multipliers map[string]uint) []uint {
	lines := []uint{}

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break

			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}

func main() {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}
	symbolArr := generateSymbolArray(symbols)
	balance := uint(1000)
	getName()

	for balance > 0 {
		bet := getBet(balance)
		if bet == 0 {
			break
		}

		balance -= bet
		spin := getSpin(symbolArr, 3, 3)
		//spin := getSpin(symbolArr, rows: 3, cols: 3)
		printSpin(spin)
		winningLines := checkWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("Congratulations! You won $%d with the line %d.\n", win, i+1)
			}
		}

	}

	fmt.Printf("You are left with, $%d.\n", balance)

}
