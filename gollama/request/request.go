package request

import (
	"aigotestapp/gollama/settings"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Request struct {
	Model  string
	Prompt string
	Suffix string
}

type Response struct {
	Model                string
	Created_At           time.Time
	Response             string
	Done                 bool
	Context              []int
	Total_Duration       int
	Load_Duration        int
	Prompt_Eval_Count    int
	Prompt_Eval_Duration int
	Eval_Count           int
	Eval_Duration        int
}

func Init(model string, prompt string, suffix string) (*Request, error) {
	if model == "" {
		return nil, errors.New("model is required")
	}
	return &Request{
		Model:  model,
		Prompt: prompt,
		Suffix: suffix,
	}, nil
}

func (r *Request) Generate(ch chan<- string, settings *settings.Settings) error {
	server := settings.GetConnectionString()
	url := fmt.Sprintf("http://%s/api/generate", server)

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("model", r.Model)
	request.Header.Add("prompt", r.Prompt)
	request.Header.Add("suffix", r.Suffix)
	request.Header.Add("format", "json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// TODO

	return nil
}
