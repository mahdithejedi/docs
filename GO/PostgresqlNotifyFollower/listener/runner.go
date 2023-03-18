package listener

import (
	"Notifier/helpers"
	"fmt"
	"github.com/lib/pq"
)

func Run(config *Config, db *pq.Listener) {
	err := db.Listen(config.TriggerName)
	helpers.CheckErr("Error while listening to trigger", err)
	select {
	case dbInfo := <-db.Notify:
		fmt.Println(dbInfo.Channel, string(dbInfo.Extra))
	}

}
