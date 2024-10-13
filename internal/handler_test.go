package internal

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandleCommandSysInfo(t *testing.T) {
    cmdr := NewCommander()
    handler := HandleRequests(cmdr)

    requestBody := CommandRequest{
        Type: "sysinfo",
    }
    body, _ := json.Marshal(requestBody)

    req, err := http.NewRequest("POST", "/execute", bytes.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
    }

    var response CommandResponse
    responseData, _ := ioutil.ReadAll(rr.Body)
    json.Unmarshal(responseData, &response)

    if !response.Success {
        t.Errorf("Expected success true, got false with error: %s", response.Error)
    }

    dataMap, ok := response.Data.(map[string]interface{})
    if !ok {
        t.Fatal("Expected data to be a map")
    }

    if dataMap["hostname"] == "" {
        t.Error("Expected hostname to be non-empty")
    }

    if dataMap["ip_address"] == "" {
        t.Error("Expected IP address to be non-empty")
    }
}

func TestHandleCommandPing(t *testing.T) {
    cmdr := NewCommander()
    handler := HandleRequests(cmdr)

    requestBody := CommandRequest{
        Type:    "ping",
        Payload: "example.com",
    }
    body, _ := json.Marshal(requestBody)

    req, err := http.NewRequest("POST", "/execute", bytes.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
    }

    var response CommandResponse
    responseData, _ := ioutil.ReadAll(rr.Body)
    json.Unmarshal(responseData, &response)

    if !response.Success {
        t.Errorf("Expected success true, got false with error: %s", response.Error)
    }

    dataMap, ok := response.Data.(map[string]interface{})
    if !ok {
        t.Fatal("Expected data to be a map")
    }

    if !dataMap["successful"].(bool) {
        t.Error("Expected ping to be successful")
    }

    if dataMap["time"] == "" {
        t.Error("Expected time to be non-empty")
    }
}

func TestHandleCommandInvalidType(t *testing.T) {
    cmdr := NewCommander()
    handler := HandleRequests(cmdr)

    requestBody := CommandRequest{
        Type: "invalid_command",
    }
    body, _ := json.Marshal(requestBody)

    req, err := http.NewRequest("POST", "/execute", bytes.NewReader(body))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Expected status code %v, got %v", http.StatusOK, status)
    }

    var response CommandResponse
    responseData, _ := ioutil.ReadAll(rr.Body)
    json.Unmarshal(responseData, &response)

    if response.Success {
        t.Error("Expected success false, got true")
    }

    if response.Error == "" {
        t.Error("Expected an error message")
    }
}

