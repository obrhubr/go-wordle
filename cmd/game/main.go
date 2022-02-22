package main

import (
	"bufio"
	"fmt"
	"obrhubr/gordle/pkg/game"
	"obrhubr/gordle/pkg/util"
	"os"
	"strings"
)

func main() {
	var all = util.Get_words_from_file("./words/answers.txt")
	var possible = util.Get_words_from_file("./words/all.txt")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("     Wordle Bot")
	fmt.Printf("--------------------- \n \n \n")

	// Print out the best starters (precomputed because it takes a long time)
	util.Format_results([]util.Result{
		{Score: 5.377566338603813, Word: "trace"},
		{Score: 5.303611974124567, Word: "salet"},
		{Score: 5.294365346728447, Word: "parse"},
		{Score: 5.287056928899425, Word: "crate"},
		{Score: 5.230307126327275, Word: "carte"},
		{Score: 5.223432537531815, Word: "peart"},
		{Score: 5.221421829329374, Word: "slate"},
		{Score: 5.2029841856220935, Word: "reast"},
		{Score: 5.186305073551834, Word: "caret"},
		{Score: 5.172858791190428, Word: "carle"},
	})

	for {
		fmt.Printf("Enter the word you have chosen \n")
		fmt.Print("-> ")

		// Get chosen word
		word, _ := reader.ReadString('\n')
		word = strings.Replace(word, "\n", "", -1)

		fmt.Printf("Enter the pattern (ints seperated by spaces; 0 is gray, 1 is yellow and 2 is green; ex: 1 0 0 0 2) \n")
		fmt.Print("-> ")

		// Get pattern
		p, _ := reader.ReadString('\n')
		p = strings.Replace(p, "\n", "", -1)
		pat := util.Pattern{L1: int(p[0] - '0'), L2: int(p[2] - '0'), L3: int(p[4] - '0'), L4: int(p[6] - '0'), L5: int(p[8] - '0')}

		possible = game.Filter_words(possible, pat, word)
		if len(possible) == 0 {
			fmt.Println("Something went wrong, try again.")
			continue
		}
		if len(possible) == 1 {
			fmt.Println("The secret word is:", possible[0])
			break
		}

		fmt.Printf("There are %v words left. \n", len(possible))

		// Get best word
		util.Format_results(game.Get_best_words(all, possible))
	}
}
