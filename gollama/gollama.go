package gollama

import (
	"aigotestapp/gollama/models"
	"aigotestapp/gollama/settings"
	"fmt"
)

func Run() {
	s := settings.Init("localhost", 11434)
	m, err := models.Init(s)
	if err != nil {
		panic(err)
	}

	for _, mo := range m.Models {
		fmt.Println(mo)
	}
}
