package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {

	query := `CREATE TABLE if not exists Account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	sqlStatement := `
	INSERT INTO Account (first_name, last_name, number, balance, created_at)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Query(sqlStatement, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	// id, err := resp.LastInsertId()
	if err != nil {
		return err
	}
	// log.Printf("Inserted Account into db with id: %v", id)
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	sqlStatement := `
	DELETE FROM Account 
	WHERE ID = $1
	`
	_, err := s.db.Query(sqlStatement, id)
	return err
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	sqlStatement := `
	SELECT * FROM Account
	WHERE id = $1
	`
	rows, err := s.db.Query(sqlStatement, id)
	
	if err != nil {
		return nil, err
	}

	rows.Next()
	account, err := scanToAccount(rows)

	if err != nil {
		return nil, fmt.Errorf("account %d not found", id)
	}
	return account, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	sqlStatement := `SELECT * FROM Account`
	rows, err := s.db.Query(sqlStatement)
	
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account, err := scanToAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}
	return accounts, nil
}

func scanToAccount(rows *sql.Rows) (*Account, error){

	account := new(Account)
	err := rows.Scan(
		&account.ID, 
		&account.FirstName, 
		&account.LastName, 
		&account.Number, 
		&account.Balance, 
		&account.CreatedAt, 
	)
	if err != nil {
		return nil, err
	}
	return account, err
}