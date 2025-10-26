package models

import (
	"context"
	"time"

 )

type Ticket struct {
	ID        uint `json:"id" gorm:"primarykey"`
	EventId   uint  `json:"eventId"`
	Event     *Event  `json:"event" gorm:"foreignkey:EventId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered   bool `json:"entered" default:"false"`
	CreateAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetOne(ctx context.Context,TicketId uint) (*Ticket, error)
	GetMany(ctx context.Context) ([] *Ticket, error)
	CreateOne(ctx context.Context, ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context,TicketId uint, updateTicket map[string]any)   (*Ticket, error)


}


type ValidateTicket struct{
	TicketId uint `json:"ticketId"`
}