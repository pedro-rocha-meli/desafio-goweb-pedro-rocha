package tickets

import (
	"desafio-go-web/internal/domain"
)


type Service interface{
	GetAllTickets() ([]domain.Ticket, error)
	GetTicketsByDestination(destination string) ([]domain.Ticket, error)
	GetAveragePeopleByDestination(destination string) (int, error)
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllTickets() ([]domain.Ticket, error){
	return s.repository.GetAllTickets()
}

func (s *service) GetTicketsByDestination(destination string) ([]domain.Ticket, error){
	return s.repository.GetTicketsByDestination(destination)
}

func (s *service) GetAveragePeopleByDestination(destination string) (int, error){
	return s.repository.GetAveragePeopleByDestination(destination)
}