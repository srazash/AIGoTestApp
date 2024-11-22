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

	fmt.Printf("%d models(s):\n", len(m.Models))
	for _, mo := range m.Models {
		fmt.Printf("\t%s (%d)\n", mo.Name, mo.Size)
	}

	for _, mn := range *m.ModelNames() {
		fmt.Printf("%s\n", mn)
	}
}
