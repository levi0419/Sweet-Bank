package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

//store provides all function to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB

}

//NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// exexTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// begin a new transaction
	tx, err :=  store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

    // Create a new Queries instance associated with the transaction
	q := New(tx)

	// Execute the provided function (fn) with the Queries instance
	err = fn(q)
	if err != nil {
        // If there is an error, attempt to rollback the transaction
		if rbbError := tx.Rollback(); rbbError != nil {
		    // If rollback fails, return an error with details
			return fmt.Errorf("tx error: %v rb error: %v",  err, rbbError)
		}
		// return the original error if the rollback its successful
		return err
	}

    // If the function executed successfully, commit the transaction
	return tx.Commit();

}

// fields that contains the params
type TransferTxParams struct {
	FromAccountID   int64	`json:"from_account_id"`
	ToAccountID 	int64	`json:"to_account_id"`
	Amount		 	int64	`json:"amount"`
}

// the result of the transfer transaction
type TransferTxResult struct {
	Transfer 		Transfer 	`json:"transfer"`
	FromAccount 	Account 	`json:"from_account"`
	ToAccount 		Account 	`json:"to_account"`
	FromEntry 		Entry   	`json:"from_entry"`
	ToEntry   		Entry 		`json:"to_entry"`
}


// TransferTx performs a money transfer from one account to the other
// It creates a transfer record add account entries and update account balance within a single db transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams)(TransferTxResult, error){
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID:  arg.FromAccountID,
			ToAccountID: 	arg.ToAccountID,
			Amount: 		arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID:      arg.FromAccountID,
			Amount:			-arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID:      arg.ToAccountID,
			Amount:			arg.Amount,
		})
		if err != nil {
			return err
		}

		account1, err := q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount, 
		})
		if err != nil {
			return err
		}

		return nil
	})



	return result, err
}
