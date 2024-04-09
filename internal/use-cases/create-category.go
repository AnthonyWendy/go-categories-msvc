package use_cases

import (
	"log"

	"github.com/AnthonyWendy/go-categories-msvc/internal/entities"
)

type createCateryUseCase struct {

}

func NewCategoryUseCase() *createCateryUseCase {
	return &createCateryUseCase{}
}

func (u *createCateryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO persist entity to db
	log.Println(category)

	return nil
}