package game

import (
	"log"
	"strconv"
)

type WsRoster struct {
	players map[int]Player
	userNames map[string]int
	topId int
}

func (roster WsRoster) FindOrNew(userName string) *Player {
	if player, ok := roster.FindByUsername(userName); ok {
		return player
	}

	return roster.createPlayer(userName)
}

func (roster WsRoster) FindByUsername(userName string) (*Player, bool) {
	if id, ok := roster.userNames[userName]; ok {
		return roster.FindOrFail(id), true
	}

	return nil, false
}

func (roster WsRoster) FindOrFail(id int) *Player {
	if player, ok := roster.players[id]; ok {
		return &player
	}

	log.Fatal("Failed to find player with id "+strconv.Itoa(id))
}

func (roster WsRoster) Load() {
	//TODO: loading

	roster.topId = 0
}

func (roster WsRoster) nextId() int {
	roster.topId++
	return roster.topId
}

func (roster WsRoster) createPlayer(userName string) *Player {
	player := Player {
		id: roster.nextId(),
		username: userName,
	}

	roster.players[player.id] = player
	roster.userNames[player.username] = player.id

	return &player
}
