package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/model"
	"github.com/sulis96/quinzia-golang-instrumentations/pkg/db"
)

type repository struct {
	DB *sql.DB
}

type IRepository interface {
	InsertMember(ctx context.Context, data model.Member) error
	ReadMember(ctx context.Context) (data []model.Member, err error)
}

func NewRepository(database *db.Database) IRepository {
	return &repository{
		DB: database.Database,
	}
}

func (r *repository) InsertMember(ctx context.Context, data model.Member) error {
	sql, err := r.DB.Conn(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer sql.Close()

	id := uuid.New()
	query := `INSERT INTO member (
		id,
		name,
		country
	) VALUES ($1, $2, $3)`
	_, err = sql.ExecContext(ctx, query, id, data.Name, data.Country)
	return err
}

func (r *repository) ReadMember(ctx context.Context) (data []model.Member, err error) {
	sql, err := r.DB.Conn(ctx)
	if err != nil {
		log.Println(err)
	}
	defer sql.Close()

	query := `SELECT name, country FROM member`
	rows, err := sql.QueryContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}
	defer rows.Close()

	var row model.Member
	for rows.Next() {
		err = rows.Scan(
			&row.Name,
			&row.Country,
		)
		if err != nil {
			log.Println(err.Error())
			return data, err
		}
		data = append(data, row)
	}
	return data, err
}
