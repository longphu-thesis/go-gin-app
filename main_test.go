package main

import (
	"testing"
	"os"
)

func TestMain(t *testing.T) {
	main()
}

func TestMainFailConnectMySql(t *testing.T) {
	os.Setenv("DB_PASS", "DB_PASS")
	main()
}