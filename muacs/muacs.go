package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hkwi/h2c"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type muacsFlags struct {
	Port int16
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!\n", r.URL.Path[1:])
}

func processFlags(f *muacsFlags) {
	flag.ErrHelp = errors.New("") //hide error message from pflag with -h
	flag.Int16VarP(&f.Port, "port", "p", 8088, "Port to use")
	flag.Parse()
}

func initViper() {
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

func main() {
	var f muacsFlags
	processFlags(&f)
	initViper()
	fmt.Printf("Port from viper: %d\n", viper.GetInt("daemon.port"))
	listenAddress := "localhost:" + strconv.Itoa(int(f.Port))
	m := http.NewServeMux()
	m.HandleFunc("/", handler)
	s := &http.Server{
		Addr:    listenAddress,
		Handler: &h2c.Server{Handler: m, DisableDirect: false},
	}
	log.Printf("Launching Server at %s\n", listenAddress)
	log.Fatal(s.ListenAndServe())
}
