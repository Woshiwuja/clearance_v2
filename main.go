package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Woshiwuja/clearance_v2/appconfig"
	//"github.com/Woshiwuja/clearance_v2/sql"
	"github.com/Woshiwuja/clearance_v2/static"
	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5"
)

func main() {
	// Config load
	cfg, err := appconfig.LoadFromPath(context.Background(), "pkl/AppConfig.pkl")
	if err != nil {
		fmt.Println("cant load config file check pkl/config.pkl")
		panic(err)
	}

	// DB Connection
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.DBhost, cfg.DBport, cfg.DBname)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	}
	defer conn.Close(ctx)

	// load various templ components
	indexComp := static.Index()

	// paths
	http.Handle("GET /", templ.Handler(indexComp))

	// Start webserver
	// mux := http.NewServeMux()
	servErr := http.ListenAndServe(cfg.Port, nil)
	if servErr != nil {
		panic(servErr)
	}
}
