package main

type Register struct {
	sequence string
	open     int
	close    int
}

func generateParenthesis(n int) []string {
	var combinations []string

	recursiveGenerator(n, Register{"", 0, 0}, &combinations)

	return combinations
}

func recursiveGenerator(n int, sequence Register, combination *[]string) {

	if len(sequence.sequence) == n*2 {
		if sequence.open == sequence.close {
			*combination = append(*combination, sequence.sequence)
		}
		return
	}

	if sequence.close < sequence.open {
		recursiveGenerator(n, Register{sequence.sequence + ")", sequence.open, sequence.close + 1}, combination)
	}
	if sequence.open < n {
		recursiveGenerator(n, Register{sequence.sequence + "(", sequence.open + 1, sequence.close}, combination)
	}
}
