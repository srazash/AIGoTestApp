package request

import "errors"

type Request struct {
	Model  string
	Prompt string
	Suffix string
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
