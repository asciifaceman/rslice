package rslice

func ProcessNewlines(slice []rune, wordwrap bool, width int, height int, truncate bool, truncateChar *rune) [][]rune {
	block := make([][]rune, height)
	for i := range block {
		block[i] = make([]rune, width)
	}

	row := 0
	for _, r := range slice {
		block[row] = append(block[row], r)
	}

	return block
}
