package domain

const (
	MIN_TURN               = 1
	MAX_TURN               = 2
	LETTER_DENIED_PLAYER_1 = "A"
	LETTER_DENIED_PLAYER_2 = "D"
)

type Player struct {
	Movements []string
	Hits      []string
	Score     int
}

func CreatePlayer(dataMovements []string, dataHits []string) *Player {
	return &Player{
		Movements: dataMovements,
		Hits:      dataHits,
		Score:     5,
	}
}
