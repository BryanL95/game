package helpers

import (
	"testing"

	"github.com/BryanL95/game/domain"
)

func createPlayerForTest() (*domain.Player, *domain.Player) {
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
	return domain.CreatePlayer(movements, hits), domain.CreatePlayer(movements, hits)
}

func TestCreateCombinations(t *testing.T) {
	player1, player2 := createPlayerForTest()
	getCountCombo := CreateCombo(*player1, *player2)

	expectedMov := 5
	expectedHit := 4
	if getCountCombo.PlayerOne[0] != expectedMov || getCountCombo.PlayerOne[1] != expectedHit {
		t.Error("Quantity movements or hits doesn't match for player 1")
	}

	if getCountCombo.PlayerTwo[0] != expectedMov || getCountCombo.PlayerTwo[1] != expectedHit {
		t.Error("Quantity movements doesn't match for player 2")
	}
}

func TestHowStart(t *testing.T) {
	player1, player2 := createPlayerForTest()
	createCombo := CreateCombo(*player1, *player2)
	howStart := createCombo.HowStart()
	expected := 1

	if howStart != expected {
		t.Error("I don't who should start!!! 😵‍💫😵‍💫😵‍💫")
	}
}

func TestGetString(t *testing.T) {
	testCases := []struct {
		movement string
		hit      string
		want     string
		player   int
	}{
		{
			movement: "DSD",
			hit:      "P",
			want:     "DSDP",
			player:   1,
		},
		{
			movement: "SA",
			hit:      "K",
			want:     "SAK",
			player:   2,
		},
		{
			movement: "ASD",
			hit:      "P",
			want:     "SDP",
			player:   1,
		},
		{
			movement: "SD",
			hit:      "",
			want:     "S",
			player:   2,
		},
	}

	for _, testCase := range testCases {
		response := GetString(testCase.movement, testCase.hit, &testCase.player)
		if response != testCase.want {
			t.Errorf("Unexpected movement %s when wait for %s for player %d", testCase.movement, testCase.want, testCase.player)
		}
	}
}
