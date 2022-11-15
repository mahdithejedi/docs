package helpers

import (
	"github.com/icrowley/fake"
	"time"
)

type coin struct {
	created_at time.Time
	updated_at time.Time
	unit       string
	name       string
}

func Coin() coin {
	return coin{
		created_at: time.Now(),
		updated_at: time.Now(),
		unit:       fake.Brand(),
		name:       fake.Words(),
	}
}
