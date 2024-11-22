package request

import (
	"aigotestapp/gollama/settings"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func (r *Request) Generate(ch chan string, settings *settings.Settings) error {
	server := settings.GetConnectionString()
	url := fmt.Sprintf("http://%s/api/generate", server)

	payloadString, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	payload := strings.NewReader(string(payloadString))

	request, err := http.NewRequest(http.MethodPost, url, payload)
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

	if response.StatusCode != http.StatusOK {
		log.Println(response)
		panic(response.Status)
	}

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		data := Response{}
		err := json.Unmarshal([]byte(scanner.Text()), &data)
		if err != nil {
			return err
		}
		ch <- data.Response

		if data.Done {
			close(ch)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
