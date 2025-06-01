package service

import (
	"context"
	"time"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/adwinugroho/go-vercel-wedding-invitation/api/repository"
	"github.com/labstack/gommon/log"
)

type (
	RSVPInterface interface {
		List(offset, limit int, isAttending bool) ([]model.Reservation, error)
		New(data model.Reservation) (*string, error)
	}
)

type rsvpImp struct {
	repo repository.RSVPInterface
}

func NewServiceRSVP(repo repository.RSVPInterface) RSVPInterface {
	return &rsvpImp{
		repo: repo,
	}
}

func (s *rsvpImp) New(data model.Reservation) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	newID, err := s.repo.Insert(ctx, data)
	if err != nil {
		log.Printf("Error while saving data rsvp:%+v\n", err)
		return nil, err
	}

	return &newID, nil
}

func (s *rsvpImp) List(offset, limit int, isAttending bool) ([]model.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	data, err := s.repo.List(ctx, offset, limit, isAttending)
	if err != nil {
		log.Printf("Error while get list rsvp:%+v\n", err)
		return nil, err
	} else if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}
