package main

import (
	"flag"
	"net/http"

	"github.com/codersgyan/students-api/internal/config"
)

func main() {
	println("🔥 MAIN STARTED")

	// ✅ DEFINE flag (THIS WAS MISSING)
	flag.String("config", "config/local.yaml", "path to config file")

	// ✅ parse flags
	flag.Parse()

	// ✅ load config
	cfg := config.MustLoad()

	println("✅ CONFIG LOADED:", cfg.Addr)

	server := http.Server{
		Addr: cfg.Addr,
	}

	println("🚀 STARTING SERVER ON", cfg.Addr)

	err := server.ListenAndServe()
	if err != nil {
		println("❌ SERVER ERROR:", err.Error())
	}
}
