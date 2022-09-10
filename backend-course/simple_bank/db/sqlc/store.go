package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Make new *Queries object
	q := New(tx)

	// Pass the *Queries object here to do a callback
	err = fn(q)

	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("Tx Error %v, rb Err: %v\n", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// Contains the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"fromAccount"`
	ToAccount   Account  `json:"toAccount"`
	FromEntry   Entry    `json:"fromEntry"`
	ToEntry     Entry    `json:"toEntry"`
}

// TransferTx perform a money transfer from one account to another
// It creates a transfer record, add account entries, and update account balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	// Make Closure
	var result TransferTxResult

	err := store.execTx(ctx, func(queries *Queries) error { // Use the callback function
		var err error

		// Create the FromTransfer transaction
		result.Transfer, err = queries.CreateTransfer
		return nil
	})

	// Return the result and possible error from the transaction
	return result, err
}
