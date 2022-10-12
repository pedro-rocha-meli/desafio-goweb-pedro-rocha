package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"
)

type Repository interface {
	GetAllTickets()([]domain.Ticket, error)
	GetTicketsByDestination(destination string) ([]domain.Ticket, error)
	GetAveragePeopleByDestination(destination string) (int, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllTickets() ([]domain.Ticket, error) {
	
	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketsByDestination(destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAveragePeopleByDestination(destination string) (int, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return len(ticketsDest), fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return len(ticketsDest), nil
}