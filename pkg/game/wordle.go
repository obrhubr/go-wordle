package game

import (
	"fmt"
	"math"
	"obrhubr/gordle/pkg/util"
	"sort"
	"strings"
)

func get_case(w string, cw string, i int) int {
	if w[i] == cw[i] {
		return 2
	}
	if strings.Contains(cw, string(w[i])) {
		return 1
	}
	return 0
}

func get_pattern_word(w string, cw string) util.Pattern {
	// TODO: fix edge cases with repeated letters
	return util.Pattern{L1: get_case(w, cw, 0), L2: get_case(w, cw, 1), L3: get_case(w, cw, 2), L4: get_case(w, cw, 3), L5: get_case(w, cw, 4)}
}

func is_possible_word(pat util.Pattern, w string, cw string) bool {
	return pat == get_pattern_word(w, cw)
}

func get_matches(ws []string, w string, p util.Pattern) int {
	sum := 0
	for _, e := range ws {
		if is_possible_word(p, w, e) {
			sum += 1
		}
	}
	return sum
}

func get_information_for_word(ps []util.Pattern, w string, total_words int, possible_words []string) float64 {
	var sum = 0.0

	for _, p := range ps {
		px := float64(get_matches(possible_words, w, p)) / float64(total_words)
		if px != 0 {
			sum += math.Log2(1 / px)
		}
	}

	return sum / float64(len(ps))
}

func Filter_words(possible_words []string, p util.Pattern, w string) []string {
	var words = []string{}

	for _, e := range possible_words {
		if is_possible_word(p, w, e) {
			words = append(words, e)
		}
	}

	return words
}

func Get_best_words(possible_words []string, all_words []string) []util.Result {
	var results = []util.Result{}

	for i, w := range all_words {
		fmt.Printf("Getting score for %v (num: %v)", w, i)
		results = append(results, util.Result{Score: get_information_for_word(util.Generate_all_Patterns(), w, len(possible_words), possible_words), Word: w})
		fmt.Printf("\r\033[K")
	}

	sort.Sort(util.ByScore(results))
	if len(results) < 10 {
		return results
	} else {
		return results[0:10]
	}
}
