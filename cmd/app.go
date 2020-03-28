package main

import (
	"github.com/reasno/gotask/pkg/gotask"
	"log"
)

// App sample
type App struct{}

// Hi returns greeting message.
func (a *App) Hi(name interface{}, r *interface{}) error {
	*r = map[string]interface{}{
		"hello": name,
	}
	return nil
}

func main() {
	gotask.Register(new(App))
	if err := gotask.Run(); err != nil {
		log.Fatalln(err)
	}
}
