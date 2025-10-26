package repositories

import (
	"context"
	"errors"

 
	"github.com/Kartik30R/Tiket.git/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

 func (e *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {

if err:= e.db.Create(event).Error; err!=nil{
	return nil,err
}

return event,nil
}

 func (e *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	res := e.db.Delete(&models.Event{}, eventId)
	if res.RowsAffected == 0 {
		return errors.New("event not found")
	}
	return res.Error}


 func (e *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	res := e.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}


 func (e *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {

    var event models.Event
	if err:= e.db.Find(&event, eventId).Error ; err!=nil{
		
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	
	}

	return &event, nil

}

 func (e *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateEvent map[string]any) (*models.Event, error) {
var event models.Event
	if err := e.db.First(&event, eventId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}

	if err := e.db.Model(&event).Updates(updateEvent).Error; err != nil {
		return nil, err
	}

	return &event, nil}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}

}
