package services

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/mf51721/OR_lab/goapi/models"
)

type CreatorService interface {
	Add(creator *models.Creator) error
	Delete(id uint) error
	Get(id uint) (*models.Creator, error)
	GetAll(params string) (*[]models.Creator, error)
	Update(id uint, payload map[string]interface{}) error
}

type CreatorServiceImpl struct {
	db *gorm.DB
}

func NewCreatorService(db *gorm.DB) CreatorService {
	return &CreatorServiceImpl{db: db}
}

// Add - Add a new programming language creator to the collection
func (s *CreatorServiceImpl) Add(creator *models.Creator) error {
	return s.db.Create(&creator).Error
}

// Delete - Deletes a creator entry
func (s *CreatorServiceImpl) Delete(id uint) error {
	// Replace with s.db.Unscoped().Delete to hard delete an entry instead of a soft delete
	result := s.db.Delete(&models.Creator{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return fmt.Errorf("no creator entries were deleted")
	}
	return nil
}

// Get - Find creator by Id
func (s *CreatorServiceImpl) Get(id uint) (*models.Creator, error) {
	var c models.Creator
	err := s.db.Preload(clause.Associations).First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// GetAll -
func (s *CreatorServiceImpl) GetAll(params string) (*[]models.Creator, error) {
	var res []models.Creator
	err := s.db.Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Update - Updates a programming language creator entry using form data
func (s *CreatorServiceImpl) Update(id uint, payload map[string]interface{}) error {
	lang := models.Creator{
		Model: gorm.Model{ID: id},
	}
	result := s.db.Model(&lang).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return fmt.Errorf("no creator entries were updated")
	}
	return nil
}
