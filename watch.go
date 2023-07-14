package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func watchAndLoad(path string, v any, onLoad func()) {
	for {
		if err := loadJSONFromFile(path, v); err != nil {
			log.Printf("failed to load JSON from file %s: %v", path, err)
		}

		onLoad()

		if err := watchFile(path); err != nil {
			log.Printf("failed to watch file %s: %v", path, err)
		}
	}
}

func loadJSONFromFile(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}

func watchFile(path string) error {
	initialStat, err := os.Stat(path)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
