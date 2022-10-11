package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service tickets.Service
}

func NewHandler(s tickets.Service) *Handler {
	return &Handler{
		service: s,
	}
}


func (h *Handler) GetAllTickets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		tickets, err := h.service.GetAllTickets()

		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		ctx.JSON(200, tickets)
	}
}

func (h *Handler) GetTicketsByDestination() gin.HandlerFunc {

	return func(c *gin.Context) {
		destination := c.Param("destination")

		tickets, err := h.service.GetTicketsByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (h *Handler) GetAveragePeopleByDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("destination")

		avg, err := h.service.GetAveragePeopleByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
