package repository

import (
	"database/sql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go_project/internal/model"
	"gorm.io/gorm"
)

var (
	// ErrNotFound data is not exist
	ErrNotFound = gorm.ErrRecordNotFound
)

var _ Repository = (*repository)(nil)

type Repository interface {
	Find() (model.TTable, error)
}

// repository mysql struct
type repository struct {
	orm    *gorm.DB
	db     *sql.DB
	tracer trace.Tracer
}

// New a repository and return
func New(db *gorm.DB, sqlDb *sql.DB) Repository {
	return &repository{
		orm:    db,
		db:     sqlDb,
		tracer: otel.Tracer("repository"),
	}
}

// Close release mysql connection
func (d *repository) Close() {

}
