package internal

import (
    "testing"
)

func TestGetSystemInfo(t *testing.T) {
    cmdr := NewCommander()
    info, err := cmdr.GetSystemInfo()

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if info.Hostname == "" {
        t.Error("Expected hostname to be non-empty")
    }

    if info.IPAddress == "" {
        t.Error("Expected IP address to be non-empty")
    }
}

func TestPing(t *testing.T) {
    cmdr := NewCommander()
    result, err := cmdr.Ping("example.com")

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    if !result.Successful {
        t.Error("Expected ping to be successful")
    }

    if result.Time == 0 {
        t.Error("Expected time to be non-zero")
    }
}
