package transaction

import "context"

type Transaction interface {
	Exec(ctx context.Context, query string, args ...interface{}) (interface{}, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) (interface{}, error)
	Query(ctx context.Context, query string, args ...interface{}) (interface{}, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type TransactionalDB interface {
	BeginTx(ctx context.Context) (Transaction, error)
}

func ExecuteInTransaction(ctx context.Context, db TransactionalDB, fn func(tx Transaction) error) error {
	tx, err := db.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback(ctx)
			panic(r)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	err = fn(tx)
	return err
}
