package address

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

type Repository struct {
	pool *pgxpool.Pool
}

func (r *Repository) Insert(ctx context.Context, addr *Address) error {
	result, err := r.pool.Exec(ctx, "INSERT INTO addresses (id, street, locality, region, postal_code, country, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", addr.ID, addr.Street, addr.Locality, addr.Region, addr.PostalCode, addr.Country, addr.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
	return err
}

func (r *Repository) SelectByID(ctx context.Context, id string) (*Address, error) {
	addr := &Address{}
	err := r.pool.QueryRow(ctx, "SELECT id, street, locality, region, postal_code, country, created_at, updated_at, deleted_at FROM addresses WHERE id = $1", id).Scan(
		&addr.ID,
		&addr.Street,
		&addr.Locality,
		&addr.Region,
		&addr.PostalCode,
		&addr.Country,
		&addr.CreatedAt,
		&addr.UpdatedAt,
		&addr.DeletedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(addr)
	return addr, nil
}
