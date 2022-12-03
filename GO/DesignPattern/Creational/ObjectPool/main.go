package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
)

// __source__ = 'https://golangbyexample.com/golang-object-pool/'

type connection interface {
	create(interface{}) (interface{}, error)
}

type Connection struct {
	con string
	id  int
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
	idle          []*Connection
	active        []*Connection
	maxConnection uint
	lock          *sync.Mutex
}

func CreateConnection(host string, port uint) *Connection {
	con, err := createConnection(host, port)
	if err != nil {
		log.Fatal("Error", err.Error())
	}
	return con

}

func CreateMultiConnection(count uint, host string, port uint) []*Connection {
	connections := make([]*Connection, count)
	for index, _ := range connections {
		connections[index] = CreateConnection(host, port)
	}
	return connections
}

func New(maxCon uint, host string, port uint) (*Pool, error) {
	return &Pool{
		idle:          CreateMultiConnection(maxCon, host, port),
		active:        make([]*Connection, 0),
		maxConnection: maxCon,
		lock:          new(sync.Mutex),
	}, nil
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
	return conn
}

func (p *Pool) Release(id int) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.remove(id)
}

func (p *Pool) remove(id int) {
	for index, value := range p.active {
		value := *value
		if value.id == id {
			p.idle = append(p.idle, p.active[index])
			p.active = removeIndex(p.active, index)
			return
		}
	}
}

func removeIndex(arr []*Connection, pos int) []*Connection {
	newArr := make([]*Connection, len(arr)-1)
	k := 0
	for i := 0; i < (len(arr) - 1); {
		if i != pos {
			newArr[i] = arr[k]
			k++
		} else {
			k++
		}
		i++
	}

	return newArr
}

func main() {
	pool, _ := New(10, "Localhost", 8085)
	id1 := pool.Get().id
	id2 := pool.Get().id
	fmt.Println(id1)
	fmt.Println(id2)
	fmt.Println(len(pool.active), len(pool.idle))
	pool.Release(id1)
	pool.Release(id2)
	fmt.Println("after")
	fmt.Println(len(pool.active), len(pool.idle))

}
