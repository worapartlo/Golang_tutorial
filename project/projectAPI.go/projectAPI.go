package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Cargo struct {
	CargoID     int    `json: "cargo_id"`
	CargoName   string `json: "cargo_name"`
	Amount      int    `json: "amount"`
	CargoDetail string `json: "cargo_detail"`
}

var Db *sql.DB
var CargoList []Cargo

const basePath = "/api"
const cargoPath = "cargo"

func getCargo(cargo_id int) (*Cargo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := Db.QueryRowContext(ctx, `SELECT
	cargo_id,
	cargo_name,
	amount,
	cargo_detail
	FROM cargo
	WHERE cargo_id = ?`, cargo_id)

	cargo := &Cargo{}
	err := row.Scan(
		&cargo.CargoID,
		&cargo.CargoName,
		&cargo.Amount,
		&cargo.CargoDetail)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return cargo, nil
}

func getCargoList() ([]Cargo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := Db.QueryContext(ctx, `SELECT
	cargo_id,
	cargo_name,
	amount,
	cargo_detail
	FROM cargo`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer result.Close()
	cargoes := make([]Cargo, 0)
	for result.Next() {
		var cargo Cargo
		result.Scan(&cargo.CargoID,
			&cargo.CargoName,
			&cargo.Amount,
			&cargo.CargoDetail)

		cargoes = append(cargoes, cargo)
	}
	return cargoes, nil
}

func insertCargo(cargo Cargo) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := Db.ExecContext(ctx, `INSERT INTO cargo
	(cargo_id,
	cargo_name,
	amount,
	cargo_detail
	) VALUES (?, ?, ?, ?)`,
		cargo.CargoID,
		cargo.CargoName,
		cargo.Amount,
		cargo.CargoDetail)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertCargoID, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertCargoID), nil
}

func handleCargoes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cargoList, err := getCargoList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(cargoList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var cargo Cargo
		err := json.NewDecoder(r.Body).Decode(&cargo)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Cargo_ID, err := insertCargo(cargo)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"cargo_id":%d}`, Cargo_ID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleCargo(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", cargoPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cargo_ID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		cargo, err := getCargo(cargo_ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if cargo == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(cargo)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization, X-Custom-Header, x-requested-with")
		handler.ServeHTTP(w, r)
	})
}

func SetupRoutes(apiBasePath string) {
	cargoesHandler := http.HandlerFunc(handleCargoes)
	cargoHandler := http.HandlerFunc(handleCargo)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, cargoPath), corsMiddleware(cargoesHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, cargoPath), corsMiddleware(cargoHandler))
}

func SetupDB() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/course")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Db)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}

func main() {
	SetupDB()
	SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
