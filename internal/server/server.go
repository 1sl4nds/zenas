package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/1sl4nds/zenas/internal/server/address"
	"github.com/1sl4nds/zenas/internal/server/health"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewServer() *Server {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgress@localhost:5433/postgres")
	if err != nil {
		log.Fatal((err))
	}
	addressRep := address.NewRepository(pool)
	addressSvc := address.NewService(addressRep)
	healthSvc := &health.Service{}
	return &Server{
		addressSvc: addressSvc,
		healthSvc:  healthSvc,
	}
}

type Server struct {
	addressSvc *address.Service
	healthSvc  *health.Service
}

func (s *Server) ListenAndServe() {
	log.Print("Listening... (port: 8000)")
	mux := http.NewServeMux()
	mux.HandleFunc("/address", s.getAddress)
	mux.HandleFunc("/health", s.getHealth)
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) getAddress(w http.ResponseWriter, r *http.Request) {
	log.Print("getting address...")
	ctx := context.Background()
	addr, err := s.addressSvc.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(addr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (s *Server) getHealth(w http.ResponseWriter, r *http.Request) {
	log.Print("getting health...")
	health := s.healthSvc.Get()
	data, err := json.Marshal(health)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
