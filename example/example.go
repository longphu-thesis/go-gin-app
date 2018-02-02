package example

import (
	"github.com/longphu-thesis/go-gin-app/example/redis"
)

func Run() {
	f := &Foo{
		FirstName: "FirstName",
		LastName:  "LastName",
		Age:       50,
	}

	f.reflect()

	Sum(1,1)

	redis.Connect()
}
