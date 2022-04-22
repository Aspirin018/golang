package main

import (
	"flag"
	"fmt"
	"github.com/gogf/gf/os/glog"
	"log"
	"net/http"
	"os"
)

func main() {
	_ = flag.Set("v", "4")
	glog.Level(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Enter RootHandler...")
	for k, vArr := range r.Header {
		for _, v := range vArr {
			w.Header().Add(k, v)
		}
	}
	versionEnv := os.Getenv("VERSION")
	w.Header().Set("VERSION", versionEnv)
	fmt.Println("Client address:", r.RemoteAddr)
}
