package handler

import (
	"context"
	"darius/pkg/proto/suggest"
)

func (h *handler) SuggestOptions(ctx context.Context, req *suggest.SuggestOptionsRequest) (*suggest.SuggestOptionsResponse, error) {
	return &suggest.SuggestOptionsResponse{
		CriteriaList: &suggest.CriteriaEleResponse{
			Criteria:   "criteria1",
			OptionList: []string{"option1", "option2", "option3"},
		},
	}, nil
}

func (h *handler) SuggestCriteria(ctx context.Context, req *suggest.SuggestCriteriaRequest) (*suggest.SuggestCriteriaResponse, error) {
	return &suggest.SuggestCriteriaResponse{
		CriteriaList: []*suggest.CriteriaEleResponse{
			{
				Criteria:   "criteria1",
				OptionList: []string{"option1", "option2", "option3"},
			},
		},
	}, nil
}

func (h *handler) SuggestQuestions(ctx context.Context, req *suggest.SuggestQuestionsRequest) (*suggest.SuggestQuestionsResponse, error) {
	return &suggest.SuggestQuestionsResponse{
		QuestionList: []*suggest.Question{
			{
				QuestionContent: "question1",
				OptionList: []*suggest.AnswerOption{
					{
						OptionContent: "option1",
						IsCorrect:     true,
					},
				},
			},
		},
	}, nil
}
