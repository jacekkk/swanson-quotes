package models

import (
	"database/sql"
	"fmt"
)

type Quote struct {
	ID    int    `json:"id" binding:"required"`
	Quote string `json:"quote" binding:"required"`
}

func GetAllQuotes() ([]*Quote, error) {
	rows, err := db.Query("SELECT * FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quotes := make([]*Quote, 0)
	for rows.Next() {
		quote := new(Quote)
		err := rows.Scan(&quote.ID, &quote.Quote)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func GetQuote(id string) (*Quote, error) {
	row := db.QueryRow("SELECT * FROM quotes WHERE quote_id=?", id)
	quote := new(Quote)
	err := row.Scan(&quote.ID, &quote.Quote)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
		return nil, err
	}

	return quote, nil
}
