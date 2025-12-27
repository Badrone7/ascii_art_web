# ASCII Art Web

## Description

**ASCII Art Web** is a Go web application that takes user input text and converts it into **ASCII art**.  
Users can choose between different ASCII art styles:

- **Standard**
- **Shadow**
- **Thinkertoy**

The generated ASCII art is displayed directly in the browser.

---

## Features

- Web interface for text input
- Multiple ASCII art styles
- Server-side ASCII art generation
- Simple and clean project structure
- Built using Go’s standard libraries

---

## Project Structure

```
.
├── functions/
│   ├── ArtHandler.go
│   ├── ArtMaker.go
│   └── HostLaunch.go
│
├── static/
│   ├── favicon.ico
│   ├── 500.html
│   ├── 400.html
│   ├── 404.html
│   ├── index.html
│   └── styles/
│       └── style.css
│
├── resources/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── go.mod
├── main.go
└── README.md
```

---

## How It Works

1. The web server is started from `main.go`.
2. `HostLaunch.go` initializes and runs the HTTP server.
3. `ArtHandler.go` handles HTTP requests and form data.
4. `ArtMaker.go` converts input text into ASCII art using banner files.
5. Banner files are loaded from the `resources/` directory.
6. The result is rendered using `index.html`.

---

## Usage

### Run the application

```bash
go run main.go
```

### Open in your browser

```
http://localhost:8080
```

---

## ASCII Art Styles

- **Standard** – basic ASCII art font
- **Shadow** – adds depth and shading
- **Thinkertoy** – decorative and complex style

---

## Technologies Used

- Go (Golang)
- net/http
- html/template

---

## Purpose

This project was created to practice:
- Web servers in Go
- HTTP request handling
- Template rendering
- File parsing and text processing

---

## Authors

- bguitoni
- msarar

