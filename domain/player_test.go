package domain

import (
	"reflect"
	"testing"
)

func TestCreatePlayer(t *testing.T) {
	movements := []string{
		"D",
		"DSD",
		"S",
		"DSD",
		"SD",
	}

	hits := []string{
		"K",
		"P",
		"",
		"K",
		"P",
	}
	player := CreatePlayer(movements, hits)

	if !reflect.DeepEqual(player.Movements, movements) || !reflect.DeepEqual(player.Hits, hits) {
		t.Error("Error creating player")
	}
}
