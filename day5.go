package main

func findSeat(seat string) (int, int) {
	rowBounds := [2]int{0, 127}
	columnBounds := [2]int{0, 7}

	rowDirections := string(seat[:7])
	columnDirections := string(seat[7:])

	for _, dir := range rowDirections {
		spread := (rowBounds[1] - rowBounds[0]) / 2
		if dir == 'F' {
			rowBounds[1] -= spread
		} else if dir == 'B' {
			rowBounds[0] += spread
		}
	}

	for _, dir := range columnDirections {
		spread := (columnBounds[1] - columnBounds[0]) / 2
		if dir == 'R' {
			columnBounds[0] += spread
		} else if dir == 'L' {
			rowBounds[1] -= spread
		}
	}

	return rowBounds[0], columnBounds[0]
}
