package api

import (
	"context"
	"testing"
)

func TestConnectToDB(t *testing.T) {

	con := ConnectToDB()
	if con == nil {
		t.Fatal("Gagal koneksi ke database, koneksi nil")
	}
	defer con.Close(context.Background())

	var result int
	err := con.QueryRow(context.Background(), "SELECT 1").Scan(&result)
	if err != nil {
		t.Fatalf("Gagal menjalankan query: %v", err)
	}

	if result != 1 {
		t.Fatalf("Data yang diterima tidak sesuai, expected 1, got: %v", result)
	}
}
