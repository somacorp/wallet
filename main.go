package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type server struct {
	conf Conf
	keys Keys
}

func main() {

	confPath := *flag.String("c", "config/conf.json", "conf file")
	keysPath := *flag.String("k", "config/keys.json", "keys file")

	flag.Parse()

	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("file '%s' does not exist", confPath)
	}
	if _, err := os.Stat(keysPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("file '%s' does not exist", keysPath)
	}

	s := server{}

	go watchAndLoad(confPath, &s.conf, func() {
		fmt.Printf("%s - %s\n\n", confPath, JSON(s.conf))
	})
	go watchAndLoad(keysPath, &s.keys.privById, func() {
		fmt.Printf("%s - %s\n\n", keysPath, JSON(s.keys.privById))
	})

	middleware := ChainMiddleware(s.Guard)

	http.HandleFunc("/gen", middleware.Then(Get(s.Gen)))
	http.HandleFunc("/ids", middleware.Then(Get(s.Ids)))
	http.HandleFunc("/keys", middleware.Then(Post(s.Keys)))
	http.HandleFunc("/sign", middleware.Then(Post(s.Sign)))
	http.HandleFunc("/ecdh", middleware.Then(Post(s.ECDH)))

	log.Fatal(http.ListenAndServe(":8800", nil))
}
