package game

import (
	"github.com/ataboo/iodedicated/wsserver"
)

type Player struct{
	id int
	user *wsserver.User
	username string
}

func (player Player) SetUser(user wsserver.User) {
	if player.user != nil {
		user.Close()
	}

	player.user = &user
}


