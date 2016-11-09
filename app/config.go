package app

import (
	"os"
	"log"
	"strconv"
	"fmt"
)

func outputDir() string {
	dir := getenv("OUTPUT_DIR", false)
	return dir
}

var OutputDir string

type DbParams struct {
	User, Pass       string
	Name, Charset    string
	ProtocleAddr     string
	MaxOpen, MaxIdle int
}

var MyDbParams DbParams

func dbParams() DbParams {
	var p DbParams

	p.User = getenv("DB_USER", false)
	p.Pass = getenv("DB_PASS", false)

	host := getenv("DB_HOST", false)
	port := getenv("DB_PORT", false)

	p.ProtocleAddr = fmt.Sprintf("tcp(%s:%s)", host, port)

	p.Name = getenv("DB_NAME", false)
	maxOpen := getenv("DB_MAX_OPEN", true)
	if maxOpen == "" {
		maxOpen = "200"
	}

	maxIdle := getenv("DB_MAX_IDLE", true)
	if maxOpen == "" {
		maxOpen = "100"
	}

	p.MaxOpen, _ = strconv.Atoi(maxOpen)
	p.MaxIdle, _ = strconv.Atoi(maxIdle)

	p.Charset = "utf8mb4,utf8"

	return p
}

func getenv(key string, nullable bool) string {
	v := os.Getenv(key);
	if nullable == false && v == "" {
		nullErr(key)
	}
	return v
}

func nullErr(key string) {
	log.Fatal(`Environment "` + key + `" MUST be set.`)
}

func init() {
	MyDbParams = dbParams()
	OutputDir = outputDir()
}
