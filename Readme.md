# Commander App

This project is a basic cross-platform application that can execute simple system commands (like network ping and getting system info) and return results as JSON over an HTTP API.

## Table of Contents

1. [Build Instructions](#build-instructions)
2. [API Documentation](#api-documentation)
3. [Installation Guide](#installation-guide)
4. [Testing](#testing)
5. [App in Action](#app-in-action)

---

## Build Instructions

To build the application locally, follow these steps:

1. **Install Go**: Ensure you have Go installed. You can download it from [here](https://golang.org/dl/).
   
2. **Clone the repository**:
   ```bash
   git clone https://github.com/monomial/commanderapp.git
   cd commanderapp
   ```

3. **Build the Go application**:
   ```bash
   go build -o commanderapp ./cmd/main.go
   ```

4. **Run the application**:
   ```bash
   ./commanderapp
   ```

By default, the application will start on port `8080`.

---

## API Documentation

### Base URL

- **http://localhost:8080**

### Endpoints

#### 1. **POST /execute**

Executes a command and returns the result in JSON format.

##### Request

- **Content-Type**: `application/json`

- **Body**:

  ```json
  {
      "type": "ping" | "sysinfo",
      "payload": "hostname or IP address for ping"
  }
  ```

##### Example: Request Body for `ping`:

```json
{
    "type": "ping",
    "payload": "google.com"
}
```

##### Example: Request Body for `sysinfo`:

```json
{
    "type": "sysinfo"
}
```

##### Response

- **200 OK**: Returns the result of the command.

- **Example**:

  ```json
  {
      "success": true,
      "data": {
          "hostname": "MyHost",
          "ip_address": "192.168.1.1"
      }
  }
  ```

- **Error Example**:

  ```json
  {
      "success": false,
      "error": "Invalid command type"
  }
  ```

---

## Installation Guide

### macOS Installation

To install the application on macOS, follow these steps:

1. **Run the make-pkg.sh file**:
   
   Make the make-pkg.sh file executable and run it.

   ```bash
   chmod +x make-pkg.sh
   ./make-pkg.sh
   ```

2. **Run the `.pkg` file**:

   Double-click the `commanderapp-installer.pkg` file to open the installer, and follow the on-screen instructions to install the application. The app will be installed in `/usr/local/bin`.

3. **Verify Installation**:

   Once installed, verify the installation by running the app:

   ```bash
   commanderapp
   ```

   The app should start, and you can interact with it using the API at `http://localhost:8080/execute`.

### Start on Boot (macOS)

The `.pkg` installer sets up a `LaunchAgent` that will automatically start the app on boot. If you need to manually control this, you can use the following command to manage the LaunchAgent:

```bash
launchctl load ~/Library/LaunchAgents/com.monomial.commanderapp.plist
```

To unload the agent (i.e., stop starting it on boot), use:

```bash
launchctl unload ~/Library/LaunchAgents/com.monomial.commanderapp.plist
```

---

## Testing

### Running Tests

The project includes a test suite to ensure that the application behaves as expected. To run the tests, use the following command:

```bash
go test ./...
```

### Test Coverage

To see the test coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Open the generated `coverage.html` file in your browser to view the test coverage report.

### Example Test Case

Here’s an example test case from the project for the `GetSystemInfo` command:

```go
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
```

This test ensures that the `GetSystemInfo` method returns valid hostname and IP address.

---

## App in Action

Here’s a short clip demonstrating the app in action:

[View the demo video](media/commanderapp.mp4)

The clip shows:
1. Sending a `ping` command and receiving the result.
2. Getting system information (hostname and IP address).
