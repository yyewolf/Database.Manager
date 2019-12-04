//go:generate goversioninfo
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/websocket"
)

var box packr.Box
var Connections []*websocket.Conn
var Websocket_Receive_Functions map[string]func(request Receive_Request)
var database *sql.DB

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func OpenBrowser(url string) {
	var err error
	time.Sleep(2 * time.Second)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Println(err)
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func DBSetup() {
	//Opens DB + Init default params if not already done.
	var err error
	database, err = sql.Open("sqlite3", "./db.db")
	if err != nil {
		return
	}
	//Request for the default DB columns
	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS DB
	(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    nom TEXT NOT NULL,
	    prenom TEXT NOT NULL,
	    telephone TEXT NOT NULL,
			instrument TEXT NOT NULL,
			niveauinstrument TEXT NOT NULL,
			solfege BOOLEAN NOT NULL,
			niveausolfege TEXT NOT NULL,
	    email TEXT NOT NULL,
	    adresse TEXT NOT NULL
	);`)
	if err != nil {
		return
	}
	statement.Exec()
}

func main() {
	box = packr.NewBox("./www")
	Websocket_Receive_Functions = make(map[string]func(request Receive_Request))
	//Default setup functions
	DBSetup()
	Websocket_Receive_AllFunctions()
	OpenBrowser("http://127.0.0.1/")
	//Setup Over
	http.HandleFunc("/ws", Websocket_Connection)
	http.HandleFunc("/", handler)
	http.Handle("/b/", http.StripPrefix("/b/", http.FileServer(box)))
	log.Fatal(http.ListenAndServe(":80", nil))
	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
