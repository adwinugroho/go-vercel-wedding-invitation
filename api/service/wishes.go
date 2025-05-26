package service

import (
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/repository"
	"github.com/labstack/gommon/log"
)

type (
	WishesInterface interface {
		List(offset int, limit int) ([]model.Wishes, error)
		New(data model.Wishes) (*string, error)
	}
)

type wishesImp struct {
	repo repository.WishesInterface
}

func NewServiceWishes(repo repository.WishesInterface) WishesInterface {
	return &wishesImp{
		repo: repo,
	}
}

func (s *wishesImp) New(data model.Wishes) (*string, error) {
	// ctx := context.Background()
	newID, err := s.repo.InsertWithSupabaseClient(data)
	if err != nil {
		log.Printf("Error while saving data:%+v\n", err)
		return nil, err
	}

	return &newID, nil
}

func (s *wishesImp) List(offset, limit int) ([]model.Wishes, error) {
	// ctx := context.Background()
	publishedWishes, err := s.repo.ListWithSupabaseClient(offset, limit)
	if err != nil {
		log.Printf("Error while get list:%+v\n", err)
		return nil, err
	} else if len(publishedWishes) == 0 {
		return nil, nil
	}
	return publishedWishes, nil
}
