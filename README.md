# getfromurl

A simple CLI tool written in Go to check the status or fetch the content of a given URL.

## 🚀 Features

- `checkstatus`: Displays the HTTP status code and its meaning.
- `checkcontent`: Downloads the HTML content of a URL.
- `-file <filename>`: (Optional) Saves the content to a file instead of printing it.

## 📦 Installation

### 🔧 Build from source

```bash
git clone https://github.com/MayaSAKATA/getfromurl.git
cd getfromurl
go build -o getfromurl
```
### 📥 Or install using Go
```bash
go install github.com/MayaSAKATA/getfromurl@latest
```
Make sure your $GOPATH/bin is in your system's PATH.

## 🧪 Usage
### ✅ Check HTTP status

```bash
./getfromurl checkstatus https://example.com
```

### 📝 Get content from a URL

```bash
./getfromurl checkcontent https://example.com
```

### 💾 Save content to a file
```bash
./getfromurl checkcontent https://example.com -file output.html
```

## 📁 Project structure
```go
getfromurl/
├── cmd/              # Cobra commands
│   ├── checkstatus.go
│   ├── checkcontent.go
│   └── root.go
├── utils/            # Utility functions (e.g., HTTP helpers)
│   └── httputils.go
├── main.go           # Entry point (calls cmd.Execute)
├── go.mod            # Go module file
├── go.sum
└── README.md
```

## ✅ Requirements

Go 1.18 or later
