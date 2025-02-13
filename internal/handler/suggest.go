package handler

import (
	"context"
	llm "darius/internal/llm"
	"darius/pkg/proto/suggest"
	"encoding/json"
	"fmt"
	"log"
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
	generalInfo := req.GetGeneralInfo()
	if generalInfo == nil {
		log.Println("generalInfo is nil")
		return nil, nil
	}

	criteriaList := req.GetCriteriaList()
	if criteriaList == nil {
		log.Println("criteriaList is nil")
		return nil, nil
	}

	prompt := fmt.Sprintf(`
You are an expert in designing tests and assessments. Your task is to analyze the provided input, which includes general information about the test and a list of criteria with the user's chosen options. Based on this input, suggest additional criteria and options that will help the user provide more detailed information for generating test questions. Follow these steps:
1. Input Provided by the User:
   - General Information:
     %v
   - Criteria List:
     %v
2. Your Task:
   - Review the general information about the test to understand its context, purpose, and constraints.
   - Analyze the list of criteria and the user's chosen options for each criterion.
   - Suggest additional criteria and options that will help the user provide more detailed information for generating test questions. Ensure the suggestions are relevant to the test's context and align with the user's chosen options.
   - Provide the output in the specified JSON format.

3. Output Format:
   [
     {
       criteria: '[Suggested Criterion 1]',
       optionList: [
         "[Suggested Option 1 for Criterion 1]",
         "[Suggested Option 2 for Criterion 1]",
       ]
     },
     {
       criteria: '[Suggested Criterion 2]',
       optionList: [
         "[Suggested Option 1 for Criterion 2]",
         "[Suggested Option 2 for Criterion 2]",
       ]
     },
     ...
   ]
Now, based on the user's input, generate the output in the specified format.
`, generalInfo, criteriaList)

	llmResponse, err := h.llmService.GenerateCriteria(ctx, &llm.LLMRequest{
		Prompt: prompt,
	})
	if err != nil {
		return nil, err
	}

	var criteriaListResp []*suggest.CriteriaEleResponse
	if err := json.Unmarshal([]byte(llmResponse.Response), &criteriaListResp); err != nil {
		return nil, err
	}

	log.Println("criteriaListResp", criteriaListResp)

	return &suggest.SuggestCriteriaResponse{
		// CriteriaList: []*suggest.CriteriaEleResponse{
		// 	{
		// 		Criteria:   "criteria1",
		// 		OptionList: []string{"option1", "option2", "option3"},
		// 	},
		// },
		CriteriaList: criteriaListResp,
	}, nil
}

func (h *handler) SuggestQuestions(ctx context.Context, req *suggest.SuggestQuestionsRequest) (*suggest.SuggestQuestionsResponse, error) {
	generalInfo := req.GetGeneralInfo()
	if generalInfo == nil {
		log.Println("generalInfo is nil")
		return nil, nil
	}

	criteriaList := req.GetCriteriaList()
	if criteriaList == nil {
		log.Println("criteriaList is nil")
		return nil, nil
	}

	prompt := fmt.Sprintf(`
You are an expert in creating test questions and answers. Your task is to generate a set of questions and answers based on the provided test information and guidelines. Follow these steps:
1. Input Provided by the User:
   - General Information:
     %v
   - Criteria List (guidelines):
     %v
2. Your Task:
   - Review the general information about the test to understand its context, purpose, and constraints.
   - Analyze the list of criteria and the user's chosen options for each criterion.
   - Generate a set of questions and answers that align with the test's context, difficulty level, and chosen criteria.
   - Ensure the questions are clear, precise, and meaningful.
   - Provide the output in the specified JSON format.
3. Output Format:
   [
     {
       questionContent: "[Question 1]",
       optionList: [
         {
           optionContent: "[Option 1]",
           isCorrect: [true/false]
         },
         {
           optionContent: "[Option 2]",
           isCorrect: [true/false]
         },
         ...
       ]
     },
     {
       questionContent: "[Question 2]",
       optionList: [
         {
           optionContent: "[Option 1]",
           isCorrect: [true/false]
         },
         {
           optionContent: "[Option 2]",
           isCorrect: [true/false]
         },
         ...
       ]
     },
     ...
   ]
Now, based on the user's input, generate the output in the specified format.

	`, generalInfo, criteriaList)
	llmResponse, err := h.llmService.GenerateQuestion(ctx, &llm.LLMRequest{
		Prompt: prompt,
	})
	if err != nil {
		return nil, err
	}

	var questionListResp []*suggest.Question
	if err := json.Unmarshal([]byte(llmResponse.Response), &questionListResp); err != nil {
		return nil, err
	}

	return &suggest.SuggestQuestionsResponse{
		QuestionList: questionListResp,
	}, nil
}
