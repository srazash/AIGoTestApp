package gollama

import (
	"aigotestapp/gollama/models"
	"aigotestapp/gollama/request"
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
		fmt.Printf("\t%s (%.02f GB)\n", mo.Name, models.ToGB(mo.Size))
	}

	for _, mn := range *m.ModelNames() {
		fmt.Printf("%s\n", mn)
	}

	r, err := request.Init("phi3.5", "How do I convert from C to F in Go?", "")
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	go r.Generate(ch, s)

	for c := range ch {
		fmt.Printf("%s", c)
	}
}
