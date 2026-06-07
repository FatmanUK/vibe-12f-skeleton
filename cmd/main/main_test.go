package main

import (
	"testing"
)

func TestInitDatabase(t *testing.T) {
	err := initDatabase()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
