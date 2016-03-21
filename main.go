package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jerrac/davidbeat/beater"
)

func main() {
	err := beat.Run("davidbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
