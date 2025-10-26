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
func (t *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	if err:= t.db.Model(&models.Ticket{}).Create(ticket).Error; err!=nil{
		return nil, err
	}

	return ticket,nil
}

// GetMany implements models.TicketRepository.
func (t *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	if err:= t.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets).Error; err!=nil{
		return nil, err
	}

	return tickets,nil

}

// GetOne implements models.TicketRepository.
func (t *TicketRepository) GetOne(ctx context.Context, TicketId uint) (*models.Ticket, error) {
 	ticket := &models.Ticket{}

	err := t.db.
		Model(&models.Ticket{}).
		Where("id = ?", TicketId).
		Preload("Event").
		First(ticket).Error

	if err != nil {
		return nil, err
	}

	return ticket, nil
}


// UpdateOne implements models.TicketRepository.
func (t *TicketRepository) UpdateOne(ctx context.Context, TicketId uint, updateTicket map[string]any) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	updateRes := t.db.Model(ticket).Where("id = ?", TicketId).Updates(updateTicket)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	return t.GetOne(ctx, TicketId)
}




func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
