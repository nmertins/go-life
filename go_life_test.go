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
}
