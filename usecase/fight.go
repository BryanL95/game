package usecase

import (
	"strings"

	"github.com/BryanL95/game/domain"
	"github.com/BryanL95/game/helpers"
)

type FightUsecase struct {
	Figth  domain.Fight
	Result []string
}

func NewFight(playerOne, playerTwo domain.Player, turn int) *FightUsecase {
	return &FightUsecase{
		Figth: domain.Fight{
			PlayerOne: playerOne,
			PlayerTwo: playerTwo,
			Turn:      turn,
		},
	}
}

func (fight *FightUsecase) Movement(position int, visited bool) {
	visit := false
	if fight.Figth.PlayerOne.Score > 0 && fight.Figth.PlayerTwo.Score > 0 {
		if fight.Figth.Turn == domain.MIN_TURN {
			fight.score(fight.Figth.PlayerOne.Movements, fight.Figth.PlayerOne.Hits, position)

			fight.Figth.Turn++
		} else {
			fight.score(fight.Figth.PlayerTwo.Movements, fight.Figth.PlayerTwo.Hits, position)
			fight.Figth.Turn--
			visit = true
		}

		if visit {
			position++
		}
		fight.Movement(position, visit)
	} else {
		currentPlayer := 2
		if fight.Figth.PlayerOne.Score > fight.Figth.PlayerTwo.Score {
			currentPlayer = 1
		}
		fight.Result = append(fight.Result, "Gana "+helpers.TranformPlayerToName(&currentPlayer))
	}
}

func (fight *FightUsecase) score(movements, hits []string, position int) {
	if position < len(movements) {
		movement := movements[position]
		hit := hits[position]
		fight.Result = append(fight.Result, fight.round(movement, hit))
	}
}

func (fight *FightUsecase) round(movement, hit string) string {
	speech := helpers.TranformPlayerToName(&fight.Figth.Turn) + " "
	movementCombo := helpers.GetString(movement, hit, &fight.Figth.Turn)

	switch movementCombo {
	case "DSDP":
		speech += domain.COMBO_TALADOKEN
		fight.Figth.PlayerTwo.Score -= domain.MAX_HIT
	case "SDK":
		speech += domain.COMBO_REMUYUKEN
		fight.Figth.PlayerTwo.Score -= domain.MIDDLE_HIT
	case "SAK":
		speech += domain.COMBO_REMUYUKEN
		fight.Figth.PlayerOne.Score -= domain.MAX_HIT
	case "ASAP":
		speech += domain.COMBO_TALADOKEN
		fight.Figth.PlayerOne.Score -= domain.MIDDLE_HIT
	default:
		for _, val := range movementCombo {
			index := string(val)
			speech += domain.Speech[index] + ", "
			if index == domain.PUNCH || index == domain.KICK {
				if fight.Figth.Turn == domain.MIN_TURN {
					fight.Figth.PlayerTwo.Score -= domain.LOW_HIT
				} else {
					fight.Figth.PlayerOne.Score -= domain.LOW_HIT
				}
			}
		}
	}
	return strings.TrimSpace(speech)
}
