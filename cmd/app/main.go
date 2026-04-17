package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"golang-trainning-backend/pkg/config"
	"golang-trainning-backend/pkg/infrastructure/datastore"
	"golang-trainning-backend/pkg/infrastructure/router"
	"golang-trainning-backend/pkg/infrastructure/storage"
	"golang-trainning-backend/pkg/infrastructure/validator"
	"golang-trainning-backend/pkg/registry"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	if config.C.SkipDB {
		log.Println("SkipDB=true: starting without DB connection")
	} else {
		db := datastore.NewDB()
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalln(err)
		}
		defer sqlDB.Close()

		s3Client, err := storage.NewS3Client()
		if err != nil {
			log.Fatalln(err)
		}
		if err := storage.CreateBuckets(s3Client); err != nil {
			log.Fatalln(err)
		}
		storageRepo := storage.NewStorageRepository(s3Client)

		r := registry.NewRegistry(db, storageRepo)
		e = router.NewRouter(e, r.NewAppController())
	}

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
