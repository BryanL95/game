package controllers

import (
	"github.com/BryanL95/game/domain"
	"github.com/BryanL95/game/helpers"
	"github.com/BryanL95/game/usecase"
	"github.com/gofiber/fiber/v2"
)

const FIRST_POSITION = 0

type Movement struct {
	Movements []string `json:"movimientos"`
	Hits      []string `json:"golpes"`
}

type Game struct {
	PlayerOne Movement `json:"player1"`
	PlayerTwo Movement `json:"player2"`
}

func CreateFight(c *fiber.Ctx) error {
	game := new(Game)
	if err := c.BodyParser(game); err != nil {
		return err
	}

	playerOne := domain.CreatePlayer(game.PlayerOne.Movements, game.PlayerOne.Hits)
	playerTwo := domain.CreatePlayer(game.PlayerTwo.Movements, game.PlayerTwo.Hits)
	howStart := helpers.CreateCombo(*playerOne, *playerTwo)
	fight := usecase.NewFight(*playerOne, *playerTwo, howStart.HowStart())
	fight.Movement(FIRST_POSITION, false)

	c.JSON(fight.Result)
	return nil
}
