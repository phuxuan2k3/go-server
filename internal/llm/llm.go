package llm

import (
	"context"
	"fmt"
)

type LLM interface {
	GenerateCriteria(ctx context.Context, req *LLMRequest) (*LLMResponse, error)
	GenerateQuestion(ctx context.Context, req *LLMRequest) (*LLMResponse, error)
}

type llmInstance struct {
	config *Config
}

func NewLLM(config *Config) *llmInstance {
	if config == nil {
		config = &Config{
			Host: "localhost",
			Port: "8080", //xtodo: set port of llm server here
		}
	}
	return &llmInstance{
		config: config,
	}
}

func (l *llmInstance) GenerateCriteria(ctx context.Context, req *LLMRequest) (*LLMResponse, error) {

	//xtodo : implement the logic to generate LLM

	return &LLMResponse{
		Response: fmt.Sprintf(`
		[
  {
    "criteria": "Question Type",
    "optionList": [
      "Short Answer",
      "True or False",
      "Fill in the Blank",
      "Matching",
      "Word Problems"
    ]
  },
  {
    "criteria": "Topics Covered",
    "optionList": [
      "Geometry",
      "Trigonometry",
      "Statistics and Probability",
      "Number Theory",
      "Functions and Graphs"
    ]
  }
]
		`),
	}, nil
}
func (l *llmInstance) GenerateQuestion(ctx context.Context, req *LLMRequest) (*LLMResponse, error) {

	//xtodo : implement the logic to generate LLM

	return &LLMResponse{
		Response: fmt.Sprintf(`
		[
    {
        "questionContent": "What is the solution to the equation 2x + 3 = 7?",
        "optionList": [
            {
                "optionContent": "x = 2",
                "isCorrect": true
            },
            {
                "optionContent": "x = 3",
                "isCorrect": false
            },
            {
                "optionContent": "x = 4",
                "isCorrect": false
            },
            {
                "optionContent": "x = 5",
                "isCorrect": false
            }
        ]
    },
    {
        "questionContent": "Which of the following is a quadratic equation?",
        "optionList": [
            {
                "optionContent": "x + 2 = 5",
                "isCorrect": false
            },
            {
                "optionContent": "x² + 3x + 2 = 0",
                "isCorrect": true
            },
            {
                "optionContent": "2x + 3 = 7",
                "isCorrect": false
            },
            {
                "optionContent": "3x³ + 2x² + x = 0",
                "isCorrect": false
            }
        ]
    },
    {
        "questionContent": "What is the slope of the line y = 3x + 2?",
        "optionList": [
            {
                "optionContent": "2",
                "isCorrect": false
            },
            {
                "optionContent": "3",
                "isCorrect": true
            },
            {
                "optionContent": "1",
                "isCorrect": false
            },
            {
                "optionContent": "0",
                "isCorrect": false
            }
        ]
    }
]
		`),
	}, nil
}
