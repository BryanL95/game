package helpers

import (
	"strings"

	"github.com/BryanL95/game/domain"
)

const (
	COMBO_PLAYER_1_1 = 13
	COMBO_PLAYER_1_2 = 12
	COMBO_PLAYER_2_1 = 23
	COMBO_PLAYER_2_2 = 22
)

var name = map[int]string{1: "Tonyn", 2: "Arnaldor"}
var comboType = map[string]int{
	"DSDP": COMBO_PLAYER_1_1,
	"SDK":  COMBO_PLAYER_1_2,
	"SAK":  COMBO_PLAYER_2_1,
	"ASAP": COMBO_PLAYER_2_2,
}

type CountPlayerCombo struct {
	PlayerOne []int
	PlayerTwo []int
}

func CreateCombo(playerOne, playerTwo domain.Player) *CountPlayerCombo {
	return &CountPlayerCombo{
		PlayerOne: []int{size(playerOne.Movements), size(playerOne.Hits)},
		PlayerTwo: []int{size(playerTwo.Movements), size(playerTwo.Hits)},
	}
}

func size(size []string) int {
	counter := 0
	for i := 0; i < len(size); i++ {
		if size[i] != "" {
			counter++
		}
	}

	return counter
}

func clearDeniendMovement(movement string, player *int) string {
	if *player == domain.MIN_TURN && strings.Index(movement, domain.LETTER_DENIED_PLAYER_1) >= 0 {
		movement = strings.Trim(movement, domain.LETTER_DENIED_PLAYER_1)
	} else if *player == domain.MAX_TURN && strings.Index(movement, domain.LETTER_DENIED_PLAYER_2) >= 0 {
		movement = strings.Trim(movement, domain.LETTER_DENIED_PLAYER_2)
	}
	return movement
}

func (s *CountPlayerCombo) HowStart() int {
	counter := 1

	if s.PlayerOne[0]+s.PlayerOne[1] > s.PlayerTwo[0]+s.PlayerTwo[1] ||
		s.PlayerOne[0] > s.PlayerTwo[0] ||
		s.PlayerOne[1] > s.PlayerTwo[1] {
		counter = 2
	}

	return counter
}

func GetString(stringOne, stringTwo string, playerTurn *int) (int, string) {
	stringOne = clearDeniendMovement(stringOne, playerTurn)
	reponse := stringOne + stringTwo
	typeCombo := 0
	if len(stringOne) > 2 {
		typeCombo = comboType[stringOne[len(stringOne)-3:]+stringTwo]
	} else if len(stringOne) == 2 {
		typeCombo = comboType[stringOne[len(stringOne)-2:]+stringTwo]
	}
	return typeCombo, reponse
}

func TranformPlayerToName(player *int) string {
	return name[*player]
}
