package main

import (
	"context"
	"log"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connStr := "postgresql://test:123@localhost/test"
	ctx := context.Background()
	db := NewDatabase()
	err := db.Init(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to the database")

	var data []AccountItem

	err = pgxscan.Select(ctx, db.Pool, &data, "SELECT (list).*, full_count FROM get_account_list(10, 0)")
	if err == nil {
		for i, v := range data {
			log.Printf("%d) Name: %s; Login: %s; Password: %s; EMail: %s; Role: %s",
				i, v.Name, v.Login, v.Password, v.EMail, v.Role)
		}

		if len(data) > 0 {
			log.Println("----------------------------------------")
			log.Printf("Record count: %d", int(data[0].FullCount))
		} else {
			log.Println("there are no rows in the table")
		}
	} else {
		log.Println(err)
	}

	db.Done()
	log.Println("disconnected from the database")

}

type Account struct {
	Id       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Login    string    `db:"login"`
	Password string    `db:"password"`
	EMail    string    `db:"email"`
	Role     string    `db:"role"`
}

type AccountItem struct {
	Account
	FullCount int64 `db:"full_count"`
}

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Init(ctx context.Context, connStr string) (err error) {
	d.Done()
	d.Pool, err = pgxpool.New(ctx, connStr)
	return err
}

func (d *Database) Done() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}
