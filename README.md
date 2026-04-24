# Go Port Scanner

A concurrent port and URL scanner written in Go.

## Features

- TCP port scanning
- HTTP/HTTPS probing
- Worker pool concurrency
- JSON output support
- File output support

---

## Usage

### Scan ports on target

```bash go run main.go -t 192.168.1.1 -p 80,443,22 ```

### Scan URL

```bash go run main.go -u example.com -p 80,443 ```

### JSON output

```bash go run main.go -t 192.168.1.1 -json ```

### Pretty JSON

```bash go run main.go -t 192.168.1.1 -json -pretty ```

### Output to file

```bash go run main.go -t 192.168.1.1 -o result.txt ```

---

## Flags
```
-t string        Target IP or CIDR (e.g. 192.168.1.1 or 192.168.1.0/24)
-u string        URL to scan
-p string        Ports (e.g. 80,443 or 80-100) (default "80")
-w int           Number of workers (default 10)
-json            Output in JSON format
-pretty          Pretty print JSON (must be used with -json)
-o string        Output to file
-fallback        Try both HTTP and HTTPS (only with -u)
```
