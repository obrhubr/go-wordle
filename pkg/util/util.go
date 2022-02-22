package util

import "fmt"

type Pattern struct {
	L1 int
	L2 int
	L3 int
	L4 int
	L5 int
}

type Result struct {
	Score float64
	Word  string
}

type ByScore []Result

func (s ByScore) Len() int {
	return len(s)
}
func (s ByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByScore) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func Generate_all_Patterns() []Pattern {
	var ps = []Pattern{}
	nums := []int{0, 1, 2}

	for _, l1 := range nums {
		for _, l2 := range nums {
			for _, l3 := range nums {
				for _, l4 := range nums {
					for _, l5 := range nums {
						ps = append(ps, Pattern{l1, l2, l3, l4, l5})
					}
				}
			}
		}
	}

	return ps
}

func Format_results(r []Result) {
	fmt.Println("-------------------------------")
	fmt.Println("|  word   |      score        |")
	fmt.Println("-------------------------------")
	for _, r := range r {
		fmt.Printf("| \x1b[31;5m%v\033[0m   | \x1b[32;5m%v\033[0m | \n", r.Word, r.Score)
	}
	fmt.Println("---------------------------")
}
