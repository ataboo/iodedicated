package game

import "sync"

type Game struct{
	Hub EventHub
	Roster WsRoster
}

type EventHub struct {
	Incoming chan Event
}

type Event struct {
	origin *Player
	message []byte
}

var instance *Game
var once sync.Once

/**
 Threadsafe singleton
 */
func GetInstance() *Game {
	once.Do(func() {
		instance = &Game{EventHub{}, WsRoster{}}
	})
	return instance
}