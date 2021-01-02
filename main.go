/*
 * Programski Jezici
 *
 * Ovaj API pruža pristup skupu podataka opisuje razne programske jezike te njihove detalje i funkcionalnosti.
 *
 */

package main

import (
	"github.com/mf51721/OR_lab/goapi/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	sw "github.com/mf51721/OR_lab/goapi"
)

func main() {
	rawLog, _ := zap.NewDevelopment()
	log := rawLog.Sugar()
	log.Info("Server started")

	router := sw.NewRouter()

	dsn := "host=localhost user=marko password=marko dbname=languages port=5301 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}
	err = db.AutoMigrate(&models.Creator{}, &models.Language{})
	if err != nil {
		log.Error(err)
	}

	app := sw.NewServer(db, router)

	log.Fatal(app.Run(":8080"))
}
