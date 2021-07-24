package main

type World struct {
	State [][]bool
}

// Iterate over each cell in the world applying the update rule.
func (w *World) Update() {
	updated := make([][]bool, len(w.State))
	for i := range updated {
		updated[i] = make([]bool, len(w.State[i]))
	}

	for i, row := range w.State {
		for j := range row {
			neighbors := w.getNeighbors(i, j)

			live_neighbors := 0
			living := w.State[i][j]
			for _, neighbor := range neighbors {
				if neighbor {
					live_neighbors++
				}
			}

			if (live_neighbors == 2 || live_neighbors == 3) && living {
				updated[i][j] = true
			}

			if (live_neighbors == 3) && !living {
				updated[i][j] = true
			}
		}
	}

	w.State = updated
}

func (w World) getNeighbors(i, j int) []bool {
	rows := len(w.State)
	columns := len(w.State[0])

	top_edge := i == 0
	bottom_edge := i == rows-1
	left_edge := j == 0
	right_edge := j == columns-1

	var neighbors []bool
	if !top_edge {
		neighbors = append(neighbors, w.State[i-1][j])
	}

	if !bottom_edge {
		neighbors = append(neighbors, w.State[i+1][j])
	}

	if !left_edge {
		neighbors = append(neighbors, w.State[i][j-1])
	}

	if !right_edge {
		neighbors = append(neighbors, w.State[i][j+1])
	}

	if !(top_edge || left_edge) {
		neighbors = append(neighbors, w.State[i-1][j-1])
	}

	if !(top_edge || right_edge) {
		neighbors = append(neighbors, w.State[i-1][j+1])
	}

	if !(bottom_edge || left_edge) {
		neighbors = append(neighbors, w.State[i+1][j-1])
	}

	if !(bottom_edge || right_edge) {
		neighbors = append(neighbors, w.State[i+1][j+1])
	}

	return neighbors
}
