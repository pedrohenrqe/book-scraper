package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "[SCRAPER] ", log.LstdFlags)