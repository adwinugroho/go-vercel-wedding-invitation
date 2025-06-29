package service

import (
	"context"
	"time"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/repository"
	"github.com/labstack/gommon/log"
)

type (
	GiftInterface interface {
		List(offset, limit int) ([]model.Gifts, error)
		New(data model.Gifts) (*int32, error)
		MarkedGiftAsBooked(data model.Gifts) error
	}
)

type giftImp struct {
	repo repository.GiftInterface
}

func NewServiceGift(repo repository.GiftInterface) GiftInterface {
	return &giftImp{
		repo: repo,
	}
}

func (s *giftImp) New(data model.Gifts) (*int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newID, err := s.repo.Insert(ctx, data)
	if err != nil {
		log.Printf("Error while saving data gifts:%+v\n", err)
		return nil, err
	}

	return &newID, nil
}

func (s *giftImp) List(offset, limit int) ([]model.Gifts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	data, err := s.repo.List(ctx, offset, limit)
	if err != nil {
		log.Printf("Error while get list gifts:%+v\n", err)
		return nil, err
	} else if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}

func (s *giftImp) MarkedGiftAsBooked(data model.Gifts) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	err := s.repo.UpdateIsBooked(ctx, data)
	if err != nil {
		log.Printf("Error while marked gift as booked:%+v\n", err)
		return err
	}

	return nil
}
