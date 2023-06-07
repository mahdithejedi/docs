package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// __source__ = 'https://golangbyexample.com/golang-object-pool/'

var RemoveInactiveDuration string = "30ms"

var RemovalDuration time.Duration

type connection interface {
	create(interface{}) (interface{}, error)
}

type Connection struct {
	con      string
	id       int
	lastUsed time.Time
}

func createConnection(host string, port uint) (*Connection, error) {
	return &Connection{
		con: fmt.Sprintf(
			"connection %s created on port %d", host, port,
		),
		id: rand.Int(),
	}, nil
}

type Pool struct {
	connectionInfo struct {
		string
		uint
	}
	idle          []*Connection
	active        []*Connection
	maxConnection int8
	lock          *sync.Mutex
	handler       chan int8
}

func CreateConnection(host string, port uint) *Connection {
	con, err := createConnection(host, port)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	return con

}

func CreateMultiConnection(count int8, host string, port uint) []*Connection {
	connections := make([]*Connection, count)
	for index, _ := range connections {
		connections[index] = CreateConnection(host, port)
	}
	return connections
}

func New(maxCon int8, host string, port uint, handler chan int8, closeSignal chan bool) (*Pool, error) {
	pool := &Pool{
		connectionInfo: struct {
			string
			uint
		}{string: host, uint: port},
		idle:          CreateMultiConnection(maxCon, host, port),
		active:        make([]*Connection, 0),
		maxConnection: maxCon,
		lock:          new(sync.Mutex),
		handler:       handler,
	}
	pool.manager(closeSignal)
	return pool, nil
}

func (p *Pool) Get() *Connection {
	if len(p.idle) == 0 {
		return nil
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	var conn *Connection
	conn, p.idle = p.idle[0], p.idle[1:]
	p.active = append(p.active, conn)
	conn.lastUsed = time.Now()
	return conn
}

func (p *Pool) Release(id int) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.remove(id)
}

func (p *Pool) remove(id int) bool {
	for index, value := range p.active {
		value := *value
		if value.id == id {
			p.idle = append(p.idle, p.active[index])
			p.active = removeIndex(p.active, index)
			return true
		}
	}
	return false
}

func (p *Pool) manager(closeSignal chan bool) {
	go func() {
		var data int8
		for {
			fmt.Println("waiting")
			data = <-p.handler
			fmt.Println("data is ", data)
			if data == 0 {
				closeSignal <- true
				break
			} else if data == -1 {
				p.autoShrink(-1)
			} else if data < -1 {
				p.autoShrink(data * -1)
			} else {
				p.extend(data)
			}
		}
	}()
}

func (p *Pool) autoShrink(count int8) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for index, idleConnection := range p.idle {
		if count == 0 {
			break
		}
		con := *idleConnection
		inactiveDuration := time.Now().Sub(con.lastUsed)
		fmt.Println("inactive duration", inactiveDuration, "removela duration(config)", RemovalDuration)
		if inactiveDuration > RemovalDuration {
			p.idle = removeIndex(p.idle, index)
			count--
		}
	}
	if count < 0 {
		return
	}
	for index, _ := range p.idle {
		if count == 0 {
			break
		}
		p.idle = removeIndex(p.idle, index)
	}
}

func (p *Pool) extend(count int8) {
	newConnections := CreateMultiConnection(count, p.connectionInfo.string, p.connectionInfo.uint)
	p.idle = append(p.idle, newConnections...)
}

func removeIndex(arr []*Connection, pos int) []*Connection {
	if pos+1 == len(arr) {
		//pop last element
		return arr[:pos]
	}
	if pos == 0 {
		if len(arr) == 1 {
			return []*Connection{}
		}
		return arr[pos+1:]
	}
	return append(arr[:pos], arr[pos+1:]...)
}

//func Pop(arr []interface{}, pos int) []interface{} {
//	return append(arr[:pos], arr[pos+1:]...)
//}

func test() {
	counterHandler := make(chan int8)
	closeSignal := make(chan bool, 1)
	pool, _ := New(10, "Localhost", 8085, counterHandler, closeSignal)
	id1 := pool.Get().id
	id2 := pool.Get().id
	id3 := pool.Get().id
	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(id3)
	fmt.Println(len(pool.active), len(pool.idle))
	pool.Release(id1)
	pool.Release(id2)
	fmt.Println("after")
	fmt.Println(len(pool.active), len(pool.idle))
	time.Sleep(time.Second * 1)
	counterHandler <- -8
	fmt.Println("After shrinking")
	fmt.Println(len(pool.active), len(pool.idle))
	fmt.Println("Going to release all")
	pool.Release(id3)
	fmt.Println(len(pool.active), len(pool.idle))
	counterHandler <- 0
	<-closeSignal
}

func init() {
	var err error
	RemovalDuration, err = time.ParseDuration(RemoveInactiveDuration)
	if err != nil {
		log.Fatal("error", err.Error())
		panic("error")
	}
}

func main() {
	test()
}
