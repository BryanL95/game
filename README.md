
# Game Combat

Narración del juego de combate


## API Reference

#### Home

```http
  GET /
```
Un pequeño mensaje de Bienvenidad

#### Create new Fight

```http
  POST /new-fight
```

| Type     | Description                       |
| :------- | :-------------------------------- |
| `json` | **Required**. Movimientos y golpes de los dos jugadores |

Example:

```json
{
   "player1":{
      "movimientos":[
         "D",
         "DSD",
         "S",
         "DSD",
         "SD"
      ],
      "golpes":[
         "K",
         "P",
         "",
         "K",
         "P"
      ]
   },
   "player2":{
      "movimientos":[
         "SA",
         "SA",
         "SA",
         "ASA",
         "SA"
      ],
      "golpes":[
         "K",
         "",
         "K",
         "P",
         "P"
      ]
   }
}
```


## Installation

Ubicarse en la raiz del proyecto. Ejecutar los siguientes comandos de Docker para generar la imagen y lanzar el contenedor

```bash
  docker build -t game-container:latest .
  docker run -d --name game-app -p 127.0.0.1:9001:9001 game-container
```
    
## Tech Stack

**Server:** 
- Golang (1.21.0)
- Fiber Framework V2 (https://docs.gofiber.io/)


## Demo

Se deja collection de postman para pruebas

