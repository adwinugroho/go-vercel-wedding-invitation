package service

import (
	"context"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newID, err := s.repo.Insert(ctx, data)
	if err != nil {
		log.Printf("Error while saving data wishes:%+v\n", err)
		return nil, err
	}

	return &newID, nil
}

func (s *wishesImp) List(offset, limit int) ([]model.Wishes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	publishedWishes, err := s.repo.List(ctx, offset, limit)
	if err != nil {
		log.Printf("Error while get list wishes:%+v\n", err)
		return nil, err
	} else if len(publishedWishes) == 0 {
		return nil, nil
	}

	return publishedWishes, nil
}
