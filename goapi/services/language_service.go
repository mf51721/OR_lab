package services

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/mf51721/OR_lab/goapi/models"
)

type LanguageService interface {
	Add(language *models.Language) error
	Delete(id uint) error
	Get(id uint) (*models.Language, error)
	GetAll(params string) (*[]models.Language, error)
	Update(id uint, payload map[string]interface{}) error
}

type LanguageServiceImpl struct {
	db *gorm.DB
}

func NewLanguageService(db *gorm.DB) LanguageService {
	return &LanguageServiceImpl{db: db}
}

// Add - Add a new programming language to the collection
func (s *LanguageServiceImpl) Add(language *models.Language) error {
	return s.db.Create(&language).Error
}

// Delete - Deletes a language entry
func (s *LanguageServiceImpl) Delete(id uint) error {
	// Replace with s.db.Unscoped().Delete to hard delete an entry instead of a soft delete
	result := s.db.Delete(&models.Language{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return fmt.Errorf("no language entries were deleted")
	}
	return nil
}

// Get - Find language by Id
func (s *LanguageServiceImpl) Get(id uint) (*models.Language, error) {
	var l models.Language
	err := s.db.Preload(clause.Associations).First(&l, id).Error
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// GetAll -
// TODO: Add query parameters
func (s *LanguageServiceImpl) GetAll(params string) (*[]models.Language, error) {
	var res []models.Language
	err := s.db.Preload(clause.Associations).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Update - Updates a programming language entry using form data
func (s *LanguageServiceImpl) Update(id uint, payload map[string]interface{}) error {
	lang := models.Language{
		Model: gorm.Model{ID: id},
	}
	result := s.db.Model(&lang).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return fmt.Errorf("no language entries were updated")
	}
	return nil
}
