# Go SOAP File Service

This project is a web service implemented in Go that provides a SOAP interface for saving and accumulating files. It includes both a server and a demo client to interact with the service.

## Project Structure

```
go-soap-file-service
├── cmd
│   ├── server
│   │   └── main.go        # Entry point for the web service
│   └── client
│       └── main.go        # Entry point for the demo client
├── internal
│   ├── service
│   │   └── file_service.go # Business logic for file operations
│   └── soap
│       └── wsdl.go        # WSDL definitions for the SOAP service
├── go.mod                  # Module definition
├── go.sum                  # Module dependency checksums
└── README.md               # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/ArtemDaemon/GoFileSOAP.git
   cd GoFileSOAP
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the SOAP server:**
   ```
   go run cmd/server/main.go
   ```

4. **Run the demo client:**
   ```
   go run cmd/client/main.go
   ```

## Usage Examples

### Saving a File

To save a file, the client will send a SOAP request to the server with the file data. The server will process the request and store the file.

### Accumulating Files

The client can also request to accumulate files, which will retrieve all saved files from the server and return them in a single response.

## Additional Information

- Ensure that you have Go installed on your machine.
- The service uses SOAP for communication, so make sure to handle the requests and responses accordingly.
- For more details on the WSDL and available operations, refer to the `internal/soap/wsdl.go` file.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.