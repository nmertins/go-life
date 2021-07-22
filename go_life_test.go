package main

import (
	"reflect"
	"testing"
)

func createEmptyWorld(x, y int) [][]bool {
	ret := make([][]bool, y)
	for i := range ret {
		ret[i] = make([]bool, x)
	}
	return ret
}

func TestUpdate(t *testing.T) {

	t.Run("dimensions match", func(t *testing.T) {
		empty_world := createEmptyWorld(3, 3)
		updated_world := Update(empty_world)

		got := []int{len(updated_world), len(updated_world[0])}
		want := []int{3, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("empty world remains empty", func(t *testing.T) {
		empty_world := createEmptyWorld(3, 3)
		updated_world := Update(empty_world)

		if !reflect.DeepEqual(updated_world, empty_world) {
			t.Errorf("got %v want %v", updated_world, empty_world)
		}
	})

	t.Run("single living cell dies", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world[1][1] = true
		updated_world := Update(world)

		got := updated_world[1][1]
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("live cell with two or three live neighbors continues living", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world[0][0] = true
		world[0][1] = true
		world[0][2] = true

		got := Update(world)
		want := [][]bool{
			{false, true, false},
			{false, false, false},
			{false, false, false},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("live cell with more than three live neighbors dies", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world[0][0] = true
		world[0][1] = true
		world[0][2] = true
		world[1][1] = true
		world[2][1] = true

		// Starting state
		// [*][*][*]
		// [ ][*][ ]
		// [ ][*][ ]

		got := Update(world)
		want := [][]bool{
			{true, true, true},
			{false, false, false},
			{false, false, false},
		}

		// Expected state
		// [*][*][*]
		// [ ][ ][ ] <- dies due to too many neighbors
		// [ ][ ][ ] <- dies due to not enough neighbors

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
