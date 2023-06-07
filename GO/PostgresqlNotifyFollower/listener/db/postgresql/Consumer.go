package postgresql

import (
	"Notifier/helpers"
	"Notifier/listener"
	"github.com/lib/pq"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Notify struct {
	Type      string  `json:"type"`
	Price     int     `json:"price"`
	Direction string  `json:"direction"`
	Volume    float64 `json:"volume"`
	Status    string  `json:"status"`
	MarketID  int     `json:"market_id"`
}

var wg sync.WaitGroup

//
//func receiver(channelName string, notifyChan chan *pq.Notification, stop chan bool) {
//	defer wg.Done()
//	for true {
//		select {
//		case data := <-notifyChan:
//			fmt.Println("data is", data)
//		case _ = <-stop:
//			break
//		}
//	}
//}

func Handler(config *listener.Config, db *pq.Listener) {
	print("config", config.TriggerName, db)
	helpers.UpdateLock()
	//runners := make([]func(channelName string, notifyChan chan *pq.Notification, stop chan bool), config.WorkerCount)
	//stoppers := make(chan bool, config.WorkerCount)
	//for i := range runners {
	//	wg.Add(1)
	//	go runners[i](config.TriggerName, notifyChan, stoppers)
	//}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	//wait for Interruption signal
	<-c
	//for _ = range runners {
	//	stoppers <- true
	//}
	print("stopped")
	//defer func() {
	//	wg.Wait()
	//	close(notifyChan)
	//	close(stoppers)
	//}()
}
