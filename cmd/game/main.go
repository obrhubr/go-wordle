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
	var all = util.Get_words_from_file("./words/all.txt")
	var possible = util.Get_words_from_file("./words/all.txt")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("     Wordle Bot")
	fmt.Printf("--------------------- \n \n \n")

	// Print out the best starters (precomputed because it takes a long time)
	util.Format_results([]util.Result{
		{Word: "tares", Score: 8.464541687215732},
		{Word: "teras", Score: 8.384621063737569},
		{Word: "pelas", Score: 8.330901493791202},
		{Word: "peats", Score: 8.246791141310844},
		{Word: "tears", Score: 8.134442220668847},
		{Word: "bares", Score: 8.085543735774847},
		{Word: "dores", Score: 8.084691108022328},
		{Word: "pears", Score: 8.082520569242503},
		{Word: "pores", Score: 8.055957327018294},
		{Word: "peart", Score: 8.027039950859331},
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
		util.Format_results(game.Get_best_words(possible, all))
	}
}
