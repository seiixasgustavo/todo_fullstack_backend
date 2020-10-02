package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text        string `gorm:"text" json:"text"`
	Due         string `gorm:"due" json:"due"`
	IsCompleted bool   `gorm:"is_completed" json:"isCompleted"`
	UserID      uint   `gorm:"user_id" json:"userId"`
}

func (t *Todo) Create(db *gorm.DB) error {
	if err := db.Create(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *Todo) Delete(db *gorm.DB, id uint) error {
	if err := db.Where("id = ?", id).Delete(Todo{}).Error; err != nil {
		return err
	}
	return nil
}

func (t *Todo) GetByPk(db *gorm.DB, id uint) (*Todo, error) {
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *Todo) GetByUser(db *gorm.DB, userID uint) (*[]Todo, error) {
	var todos []Todo
	if err := db.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return &todos, nil
}

func (t *Todo) Update(db *gorm.DB, id uint) error {
	if err := db.Where("id = ?", id).Save(&t).Error; err != nil {
		return err
	}
	return nil
}
