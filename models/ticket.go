package models

import (
	"context"
	"time"

 )

type Ticket struct {
	ID        uint `json:"id" gorm:"primarykey"`
	EventId   uint  `json:"eventId"`
	UserId    uint  `json:"userId" gorm:"foreignkey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event     *Event  `json:"event" gorm:"foreignkey:EventId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Entered   bool `json:"entered" default:"false"`
	CreateAt  time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TicketRepository interface {
	GetOne(ctx context.Context,userId uint ,TicketId uint) (*Ticket, error)
	GetMany(ctx context.Context,userId uint ) ([] *Ticket, error)
	CreateOne(ctx context.Context,userId uint , ticket *Ticket) (*Ticket, error)
	UpdateOne(ctx context.Context,userId uint ,TicketId uint, updateTicket map[string]any)   (*Ticket, error)


}


type ValidateTicket struct{
	TicketId uint `json:"ticketId"`
	OwnerId uint `json: "OwnerId"`
}