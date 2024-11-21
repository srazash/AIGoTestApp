package settings

import "fmt"

type Settings struct {
	Server string
	Port   int
}

func Init(server string, port int) *Settings {
	return &Settings{
		Server: server,
		Port:   port,
	}
}

func (s *Settings) GetConnectionString() string {
	return fmt.Sprintf("%s:%d", s.Server, s.Port)
}
