package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/oswaldom-code/skaypet-api/pkg/config"
	"github.com/oswaldom-code/skaypet-api/pkg/log"
	"github.com/oswaldom-code/skaypet-api/src/domain/models"
)

// store handles the database context
type store struct {
	db *gorm.DB
}

type Repository interface {
	// Pet
	GetPets() ([]models.Pet, error)
}

var dbStorage *store

// New returns a new instance of a Store
func New(dsn config.DBConfig) Repository {
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=America/Bogota",
		dsn.Host, dsn.User, dsn.Password, dsn.Database, dsn.Port, dsn.SSLMode)
	db, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		log.ErrorWithFields("error connecting to db ", log.Fields{
			"database": dsn.Database,
			"username": dsn.User,
			"err":      err,
		})
		os.Exit(1)
	}
	return &store{db: db.Set("gorm:auto_preload", true)}
}

func GetPersistence() Repository {
	if dbStorage == nil {
		fmt.Println("Es NIL")
		return New(config.GetDBConfig())
	}
	fmt.Println("NO Es NIL")

	return dbStorage
}

func (s *store) OrderByLimitOffset(orderBy, direction *string, limit, offset *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderBy != nil {
			if direction != nil {
				db = db.Order(fmt.Sprintf("%s %s", *orderBy, *direction))
			} else {
				db = db.Order(fmt.Sprintf("%s asc", *orderBy))
			}
		}

		if limit != nil {
			db = db.Limit(int(*limit))
		}

		if offset != nil {
			db = db.Offset(int(*offset))
		}

		return db
	}
}
