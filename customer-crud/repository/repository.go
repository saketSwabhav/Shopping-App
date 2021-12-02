package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Repository represents generic interface for interacting with DB
type Repository interface {
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error
	Find(uow *UnitOfWork, out interface{}, queryProcessor []QueryProcessor) error
	GetAll(uow *UnitOfWork, out interface{}, preloadAssociations ...string) error
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}, where ...map[string]interface{}) error
	Save(uow *UnitOfWork, out interface{}) error
}

// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{DB: db.New(), committed: false, readOnly: true}
	}
	return &UnitOfWork{DB: db.New().Begin(), committed: false, readOnly: false}
}

// Complete marks end of unit of work
func (uow *UnitOfWork) Complete() {
	if !uow.committed && !uow.readOnly {
		uow.DB.Rollback()
	}
}

// Commit the transaction
func (uow *UnitOfWork) Commit() {
	if !uow.readOnly {
		uow.DB.Commit()
	}
	uow.committed = true
}

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
	return &GormRepository{}
}

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.First(out, "id = ?", id).Error
}

func Where(condition string, value ...interface{}) QueryProcessor {

	// log.Println("Args ->", value)
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Model(out).Where(condition, value...)
		return db, nil
	}
}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

//find
func (g *GormRepository) Find(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {

	db := uow.DB
	var err error

	for _, queryProcessor := range queryProcessors {
		db, err = queryProcessor(db, out)
		if err != nil {
			return err
		}
	}

	if err = db.Debug().Find(out).Error; err != nil {
		return err
	}

	return nil
}

//Preload
func Preload(preloadAssociation []string) QueryProcessor {

	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		for _, association := range preloadAssociation {
			db = db.Debug().Preload(association)
		}

		return db, nil
	}
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, preloadAssociations ...string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Find(out).Error
}

// GetAllForTenant returns all objects of specifeid tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Where("tenantID = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Create(entity).Error
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Model(entity).Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}, where ...map[string]interface{}) error {
	fmt.Println(where)
	// value:=where[10]
	db := uow.DB

	for _, value := range where {
		fmt.Println(value)

		for key, val := range value {
			db = db.Debug().Where(key, val)
		}
	}

	return db.Delete(entity).Error
}

//Save specified Entity
func (repository *GormRepository) Save(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Save(entity).Error
}
