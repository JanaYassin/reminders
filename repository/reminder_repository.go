package repository

import (
	_ "github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_ "net/http"
	"survivorcoders.com/reminders/entity"
	"time"
)

type ReminderRepository struct {
	DB *gorm.DB
}

func (r ReminderRepository) GetAll() []entity.Reminder {
	var reminders []entity.Reminder
	_ = r.DB.Find(&reminders)
	return reminders
}
func (r ReminderRepository) Get() []entity.Reminder {

	// retrieve a record by id from a database
	var reminders []entity.Reminder
	_ = r.DB.Find(&reminders, &entity.Reminder{
		Id:          1,
		Name:        "Call my mom1",
		RemindMeAt:  time.Now(),
		Description: "it's about my friend12",
	})
	// SELECT * FROM users WHERE id=1,Name="Call your mom".....;
	// handle error

	// map data into Entity
	return reminders
}

func (r *ReminderRepository) Create(db *gorm.DB) (*ReminderRepository, error) {

	var err error
	err = db.Debug().Create(&r).Error
	if err != nil {
		return &ReminderRepository{}, err
	}
	return r, nil
}

func (r ReminderRepository) PUT(reminder entity.Reminder, id int) bool {
	//get the old object based on the 'id' if exist modifies content
	var reminders = entity.Reminder{Id: id}
	result := r.DB.First(&reminders)
	if result.RowsAffected > 0 {
		reminders.Name = reminder.Name
		reminders.RemindMeAt = reminder.RemindMeAt
		reminders.Description = reminder.Description
		r.DB.Save(&reminders)
		return true
	}
	return false
}
func (r ReminderRepository) Delete(db *gorm.DB, id uint32) (int64, error) {

	//var reminders []entity.Reminder
	db = db.Debug().Model(&ReminderRepository{}).Where("id = ?", id).Take(&ReminderRepository{}).Delete(&ReminderRepository{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
