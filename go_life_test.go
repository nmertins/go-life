package main

import (
	"reflect"
	"testing"
)

func createEmptyWorld(x, y int) World {
	state := make([][]bool, y)
	for i := range state {
		state[i] = make([]bool, x)
	}

	ret := World{state}

	return ret
}

func TestUpdate(t *testing.T) {

	t.Run("dimensions match", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world.Update()

		got := []int{len(world.State), len(world.State[0])}
		want := []int{3, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("empty world remains empty", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world_copy := world
		world.Update()

		if !reflect.DeepEqual(world, world_copy) {
			t.Errorf("got %v want %v", world, world_copy)
		}
	})

	t.Run("single living cell dies", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world.State[1][1] = true
		world.Update()

		got := world.State[1][1]
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("live cell with two or three live neighbors continues living", func(t *testing.T) {
		world := World{
			[][]bool{
				{true, true, true},
			},
		}

		world.Update()
		got := world.State
		want := [][]bool{
			{false, true, false},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		world.State = [][]bool{
			{true, true},
			{true, true},
		}

		want = world.State

		world.Update()
		got = world.State

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("live cell with more than three live neighbors dies", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world.State[0][0] = true
		world.State[0][1] = true
		world.State[0][2] = true
		world.State[1][1] = true
		world.State[2][1] = true

		// Starting state
		// [*][*][*]
		// [ ][*][ ]
		// [ ][*][ ]

		world.Update()
		got := world.State
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

	t.Run("dead cell with exactly three live neighbors becomes live", func(t *testing.T) {
		world := createEmptyWorld(3, 3)
		world.State[0][0] = true
		world.State[0][1] = true
		world.State[1][0] = true

		// Starting state
		// [*][*][ ]
		// [*][ ][ ]
		// [ ][ ][ ]

		world.Update()
		got := world.State
		want := [][]bool{
			{true, true, false},
			{true, true, false},
			{false, false, false},
		}

		// Expected state
		// [*][*][ ]
		// [*][*][ ] <- dies due to too many neighbors
		// [ ][ ][ ] <- dies due to not enough neighbors

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
