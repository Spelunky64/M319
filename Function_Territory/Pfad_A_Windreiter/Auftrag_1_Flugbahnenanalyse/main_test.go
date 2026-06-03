package main

import "testing"

func TestBerechneBasisAuftrieb(t *testing.T) {
	result := berechneBasisAuftrieb(50, 30)
	if result != 80 {
		t.Errorf("Expected 80, got %d", result)
	}
}
