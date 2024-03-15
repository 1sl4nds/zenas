package address_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/1sl4nds/zenas/internal/server/address"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Test_Insert(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgress@localhost:5433/postgres")
	if err != nil {
		log.Fatal((err))
	}
	repo := address.NewRepository(pool)
	now := time.Now()
	addr := &address.Address{
		ID:         uuid.NewString(),
		Street:     "abcd-1234",
		Locality:   "abcd-1234",
		Region:     "abcd-1234",
		PostalCode: "abcd-1234",
		Country:    "abcd-1234",
		CreatedAt:  &now,
	}
	repo.Insert(ctx, addr)
}

func Test_SelectByID(t *testing.T) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgress@localhost:5433/postgres")
	if err != nil {
		log.Fatal((err))
	}
	id := "9db0a034-dcfd-428d-b121-ee1034bd4568"
	repo := address.NewRepository(pool)
	addr, err := repo.SelectByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(addr)
}
