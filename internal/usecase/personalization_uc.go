package usecase

import (
	"errors"
	"github.com/google/uuid"
	"projectIntern/internal/repository"
	"projectIntern/model"
)

type PersonalizationItf interface {
	Analyze(id uuid.UUID, req model.PersonalizationReq) (uint, error)
}

type Personalization struct {
	user repository.UserRepoItf
}

type SkinProblem struct {
	ProblemId uint     `json:"problem_id"`
	Answers   []string `json:"answers"`
}

func NewPersonalization(user repository.UserRepoItf) PersonalizationItf {
	return &Personalization{user: user}
}

var knowledgeBase = []SkinProblem{
	{1, []string{"d", "a", "b", "a", "a", "b"}},
	{2, []string{"a", "a", "a", "b", "a", "a"}},
	{3, []string{"c", "b", "b", "b", "b", "b"}},
	{4, []string{"c", "a", "b", "c", "b", "d"}},
}

func (p Personalization) Analyze(id uuid.UUID, req model.PersonalizationReq) (uint, error) {
	var foundProblem []SkinProblem

	user, err := p.user.GetById(id)
	if err != nil {
		return 0, err
	}

	if user.ProblemID != 0 {
		user.ProblemID = 0
	}

	for _, problem := range knowledgeBase {
		if isEqual(problem.Answers, req.Answers) {
			foundProblem = append(foundProblem, problem)
		}
	}

	if len(foundProblem) == 0 {
		return 0, errors.New("skin problem not found")

	}

	user.ProblemID = foundProblem[0].ProblemId

	err = p.user.Update(user, model.UserUpdate{ProblemId: user.ProblemID})
	if err != nil {
		return 0, err
	}

	return foundProblem[0].ProblemId, nil
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}
