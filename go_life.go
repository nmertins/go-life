package main

// Iterate over each cell in the world applying the update rule.
func Update(world [][]bool) [][]bool {
	updated := make([][]bool, len(world))
	for i := range updated {
		updated[i] = make([]bool, len(world[i]))
	}

	return updated
}
