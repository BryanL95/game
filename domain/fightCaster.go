package domain

const (
	MAX_HIT         = 3
	MIDDLE_HIT      = 2
	LOW_HIT         = 1
	PUNCH           = "P"
	KICK            = "K"
	COMBO_TALADOKEN = "conecta un Taladoken 😲"
	COMBO_REMUYUKEN = "conecta un Remuyuken 😲"
)

var Speech = map[string]string{
	"A": "Se mueve a la izquierda",
	"S": "Da un paso hacia atras",
	"D": "Da un paso a la derecha",
	"W": "Se adelanta",
	"P": "Conecta un puño 🥊 en el ojo un poco más y lo saca",
	"K": "Lanza un patada y casi le arranca la cabeza 🤯",
}
