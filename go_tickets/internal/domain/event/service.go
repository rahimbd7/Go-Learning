package event

import "go_tickets/internal/domain/event/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateEvent(req dto.CreateRequest) (*dto.Response, error) {
	event := Event{
		TITLE:           req.TITLE,
		DESCRIPTION:     req.DESCRIPTION,
		LOCATION:        req.LOCATION,
		STARTSAT:        req.STARTSAT,
		TOTALTICKETS:    req.TOTALTICKETS,
		AVAILABLETICKET: req.TOTALTICKETS,
		PRICE:           req.PRICE,
	}

	if err := s.repo.Create(&event); err != nil {
		return nil, err
	}

	return event.ToResponse(), nil

}

func (s *service) GetAllEvent() ([]dto.Response, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	responses := make([]dto.Response, len(events))

	for i, e := range events {
		responses[i] = *e.ToResponse()
	}
	return responses, nil
}

func (s *service) GetById(eventId uint) (*dto.Response, error) {
	event, err := s.repo.GetById(eventId)
	if err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

func (s *service) Update(eventId uint, req dto.UpdateRequest) (*dto.Response, error) {
	event, err := s.repo.GetById(eventId) //find existing event first
	if err != nil {
		return nil, err
	}
	if req.TITLE != "" {
		event.TITLE = req.TITLE
	}
	if req.DESCRIPTION != "" {
		event.DESCRIPTION = req.DESCRIPTION
	}

	if req.LOCATION != "" {
		event.LOCATION = req.LOCATION
	}

	if !req.STARTSAT.IsZero() {
		event.STARTSAT = req.STARTSAT
	}

	if req.PRICE != 0 {
		event.PRICE = req.PRICE
	}

	if err := s.repo.Update(event); err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

func (s *service) HardDelete(eventId uint) error {
	err := s.repo.HardDelete(eventId)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) SoftDelete(eventId uint) (*dto.DeleteResponse, error) {
	deletedEvent, err := s.repo.SoftDelete(eventId)
	if err != nil {
		return nil, err
	}
	return deletedEvent.ToDeletedResponse(), nil
}
