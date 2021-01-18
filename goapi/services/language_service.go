package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

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
	SetCreators(langId uint, creators []models.Creator) error
	GetPic(id uint)
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

// GetPic - Get pic
func (s *LanguageServiceImpl) GetPic(id uint) {
	var l models.Language
	err := s.db.Preload(clause.Associations).First(&l, id).Error

	filename := "resources/" + fmt.Sprint(id) + ".png"
	fileinfo, err := os.Stat(filename)
	cntr := 0

	if os.IsNotExist(err) {
		cntr++
	} else {
		stat := fileinfo.Sys().(*syscall.Stat_t)
		ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
		if time.Now().Unix()-ctime.Unix() > 600 {
			cntr++
		}
	}

	if cntr > 0 {
		url := "https://en.wikipedia.org/api/rest_v1/page/summary/" + l.Wikipedia
		response, e := http.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		bodyStr := string(body)

		result := map[string]interface{}{}
		json.Unmarshal([]byte(bodyStr), &result)
		if result["originalimage"] == nil {
			return
		}
		originalimage := result["originalimage"].(map[string]interface{})
		var imgurl string

		for key, value := range originalimage {
			if key == "source" {
				imgurl = value.(string)
			}
		}

		imgresponse, e := http.Get(imgurl)
		if e != nil {
			log.Fatal(e)
		}
		defer imgresponse.Body.Close()

		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = io.Copy(file, imgresponse.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Success!")

	}
	//log.Println(fileinfo)

	return
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

func (s *LanguageServiceImpl) SetCreators(langId uint, creators []models.Creator) error {
	lang := models.Language{
		Model: gorm.Model{ID: langId},
	}
	for _, c := range creators {
		if c.Name == "" {
			return fmt.Errorf("creator's name must be set")
		}
	}
	return s.db.Model(&lang).Association("Creators").Replace(creators)
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
