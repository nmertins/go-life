package main

// Iterate over each cell in the world applying the update rule.
func Update(world [][]bool) [][]bool {
	updated := make([][]bool, len(world))
	for i := range updated {
		updated[i] = make([]bool, len(world[i]))
	}

	rows := len(world)
	columns := len(world[0])
	for i, row := range world {
		for j := range row {
			top_edge := i == 0
			bottom_edge := i == rows-1
			left_edge := j == 0
			right_edge := j == columns-1

			var neighbors []bool
			if !top_edge {
				neighbors = append(neighbors, world[i-1][j])
			}

			if !bottom_edge {
				neighbors = append(neighbors, world[i+1][j])
			}

			if !left_edge {
				neighbors = append(neighbors, world[i][j-1])
			}

			if !right_edge {
				neighbors = append(neighbors, world[i][j+1])
			}

			if !(top_edge || left_edge) {
				neighbors = append(neighbors, world[i-1][j-1])
			}

			if !(top_edge || right_edge) {
				neighbors = append(neighbors, world[i-1][j+1])
			}

			if !(bottom_edge || left_edge) {
				neighbors = append(neighbors, world[i+1][j-1])
			}

			if !(bottom_edge || right_edge) {
				neighbors = append(neighbors, world[i+1][j+1])
			}

			live_neighbors := 0
			living := world[i][j]
			for _, neighbor := range neighbors {
				if neighbor {
					live_neighbors++
				}
			}

			if (live_neighbors == 2 || live_neighbors == 3) && living {
				updated[i][j] = true
			}
		}
	}

	return updated
}
