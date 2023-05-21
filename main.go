package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TechBowl-japan/go-stations/db"
	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/router"
	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	fmt.Println("welcome to go-server")
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	// DB接続確認
	ping := todoDB.Ping()
	if ping != nil {
		return err
	}

	// DBのテーブル情報を表示
	rows, err := todoDB.Query("PRAGMA table_info(todos)")
	if err != nil {
		return err
	}
	defer todoDB.Close()
	defer rows.Close()

	fmt.Println("schema")
	// 結果の取得
	for rows.Next() {
		var cid int
		var name string
		var typ string
		var notnull int
		var dflt_value interface{}
		var pk int
		if err := rows.Scan(&cid, &name, &typ, &notnull, &dflt_value, &pk); err != nil {
			return err
		}
		fmt.Println(cid, name, typ, notnull, dflt_value, pk)
	}
	fmt.Println("end")
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)
	svc := &service.TODOService{}
	HealthzHandler := &model.HealthzHandler{}

	mux.Handle("/healthz", HealthzHandler)
	mux.Handle("/todos", handler.NewTODOHandler(svc))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// TODO: サーバーをlistenする
	log.Println("\nServer is running on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		return err
	}

	return nil
}
