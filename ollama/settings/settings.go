package settings

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Settings struct {
	Server string `toml:"server"`
	Port   int    `toml:"port"`
}

func Init(server string, port int) *Settings {
	return &Settings{
		Server: server,
		Port:   port,
	}
}

func Load() (string, int) {
	p := "ollama.toml"

	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var s Settings
	err = toml.Unmarshal(c, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s.Server, s.Port
}

func LoadAndInit() *Settings {
	h, p := Load()
	return Init(h, p)
}

func (s *Settings) ConnectionString() string {
	return fmt.Sprintf("%s:%d", s.Server, s.Port)
}
