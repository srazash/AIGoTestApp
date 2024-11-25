package models

import (
	"aigotestapp/ollama/settings"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Details struct {
	Parent_Model       string
	Format             string
	Family             string
	Families           []string
	Parameter_Size     string
	Quantization_Level string
}

type Model struct {
	Name        string
	Model       string
	Modified_At time.Time
	Size        int
	Digest      string
	Details     Details
}

type Models struct {
	Models []Model
}

func Init(settings *settings.Settings) (*Models, error) {
	server := settings.ConnectionString()
	url := fmt.Sprintf("http://%s/api/tags", server)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	models := &Models{}
	err = json.Unmarshal(body, models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (m *Models) ModelNames() *[]string {
	modelnames := make([]string, len(m.Models))

	for i, mo := range m.Models {
		modelnames[i] = strings.Split(mo.Name, ":")[0]
	}

	return &modelnames
}
