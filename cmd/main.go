package main

import (
	"github.com/alvaroeds/amaris-dragon-ball/cmd/bootstrap"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Run()
}
