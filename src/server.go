package main

import (
	// "encoding/json"
	// "fmt"
	// "io/ioutil"

	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	tarantool "github.com/tarantool/go-tarantool"
)

type ServerController struct {
	Client   *tarantool.Connection
	Mutex    *sync.Mutex
	HeadHash []byte
	BodyHash []byte
}

func (s *ServerController) handleRoot(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

	case http.MethodGet:
		s.serveMain(w, r)
	default:

	}
}

func (s *ServerController) serveMain(w http.ResponseWriter, r *http.Request) {
	head, _ := ioutil.ReadFile("./head.html")
	body, _ := ioutil.ReadFile("./body.html")
	bodyS := string(body)
	headS := string(head)
	s.HeadHash = fnv.New32a().Sum(head)
	s.BodyHash = fnv.New32a().Sum(body)
	strings.Replace(bodyS, "RANDOM_VALUE", strconv.FormatInt(rand.Int63n(9223372036854775807), 10), 1)
	fmt.Fprintf(w, "<html>%s%s</html>", headS, bodyS)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// log.Fatal(err)
		}
		// log.Println(r.Header)
		// printed, _ := httputil.DumpRequest(r, true)
		// log.Println(printed)
		// log.Println(r.Header)
		log.Println(body)
		// time.Sleep(10 * time.Second)
		// log.Println(body)
		defer r.Body.Close()

		fmt.Fprintf(w, "hewwo uwu")

	}
}

func main() {

	var err error
	genServer := &ServerController{}

	genServer.Client, err = tarantool.Connect("tarantoolinst:3301",
		tarantool.Opts{
			Timeout: 500 * time.Millisecond,
			// Reconnect:     1 * time.Second,
			// MaxReconnects: 3,
			// User:          "guest",
		})
	if err != nil {
		log.Fatalf("Failed to connect to tarantool: %s", err.Error())
	}

	mux := http.NewServeMux()
	// mux.HandleFunc("/", http.FileServer(http.Dir("./main.html")))
	mux.HandleFunc("/", genServer.handleRoot)
	mux.HandleFunc("/login", login)

	http.ListenAndServe(":8080", mux)
	//title := r.URL.Path[len("/view/"):]
}
