package repositories

import (
	"context"
	"github.com/Kartik30R/Tiket.git/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

// CreateOne implements models.TicketRepository.
func (t *TicketRepository) CreateOne(ctx context.Context, userId uint, ticket *models.Ticket) (*models.Ticket, error) {
	ticket.UserId=userId
	
	if err := t.db.Model(&models.Ticket{}).Create(ticket).Error; err != nil {
		return nil, err
	}

	return t.GetOne(ctx, userId, ticket.ID)
}

// GetMany implements models.TicketRepository.
func (t *TicketRepository) GetMany(ctx context.Context, userId uint) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	if err := t.db.Model(&models.Ticket{}).Where("user_id= ?", userId).
		Preload("Event").Order("updated_at desc").Find(&tickets).Error; err != nil {
		return nil, err
	}

	return tickets, nil

}

// GetOne implements models.TicketRepository.
func (t *TicketRepository) GetOne(ctx context.Context, userId uint, TicketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	err := t.db.
		Model(&models.Ticket{}).
		Where("id = ?", TicketId).
		Where("user_id= ?", userId).
		Preload("Event").
		First(ticket).Error

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

// UpdateOne implements models.TicketRepository.
func (t *TicketRepository) UpdateOne(ctx context.Context, userId uint, TicketId uint, updateTicket map[string]any) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	updateRes := t.db.Model(ticket).Where("id = ?", TicketId).Updates(updateTicket)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return t.GetOne(ctx, userId, TicketId)
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
