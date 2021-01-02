package services

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/mf51721/OR_lab/goapi/models"
)

type LanguageService interface {
	Add(language *models.Language) error
	Delete(language models.Language) error
	Get(id uint) (*models.Language, error)
	GetAll(params string) (*[]models.Language, error)
	Update(language models.Language) error
}

type LanguageServiceImpl struct {
	db *gorm.DB
}

func NewLanguageService(db *gorm.DB) LanguageService {
	return &LanguageServiceImpl{db: db}
}

// AddLanguage - Add a new programming language to the collection
func (s *LanguageServiceImpl) Add(language *models.Language) error {
	s.db.Create(&language)
	return nil
}

// DeleteLang - Deletes a language entry
func (s *LanguageServiceImpl) Delete(language models.Language) error {
	return nil
}

// GetLang - Find language by Id
func (s *LanguageServiceImpl) Get(id uint) (*models.Language, error) {
	return nil, nil
}

// LanguagesGet -
func (s *LanguageServiceImpl) GetAll(params string) (*[]models.Language, error) {
	var res []models.Language
	s.db.Preload(clause.Associations).Find(&res)
	return &res, nil
}

// UpdateLangWithForm - Updates a programming language entry using form data
func (s *LanguageServiceImpl) Update(language models.Language) error {
	return nil
}
