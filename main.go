/*
 * Programski Jezici
 *
 * Ovaj API pruža pristup skupu podataka opisuje razne programske jezike te njihove detalje i funkcionalnosti.
 *
 */

package main

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	sw "github.com/mf51721/OR_lab/goapi"
	"github.com/mf51721/OR_lab/goapi/models"
	"github.com/mf51721/OR_lab/goapi/services"
)

func main() {
	rawLog, _ := zap.NewDevelopment()
	log := rawLog.Sugar()
	log.Info("Server started")

	router := sw.NewRouter()

	dsn := "host=localhost user=marko password=marko dbname=languages port=5301 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Error(err)
	}
	err = db.AutoMigrate(&models.Creator{}, &models.Language{})
	if err != nil {
		log.Error(err)
	}

	app := sw.NewServer(db, router)

	ls := services.NewLanguageService(db)

	// Configure the server
	app.SetLanguageService(ls)
	app.SetRoutes(app.GetRoutes())

	log.Fatal(app.Run(":8080"))
}
