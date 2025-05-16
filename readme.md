## GetDiskSpace

GetDiskSpace is a cross-platform CLI tool written in Go that reports disk usage statistics for a specified path. It shows total, used, free space, and usage percentage in a clear table format. The tool supports Windows, Linux (including WSL2), and macOS without any third-party dependencies.

---

## Features

- Cross-platform support with platform-specific disk usage detection  
- Simple and clean table output in the terminal  
- Built with Go standard library and Cobra CLI framework  

---

## Installation

### Prerequisites

- Go 1.20+ installed (for building from source)  
- Make

---

## Usage

### Running from source

1. Clone or download the repository.

2. Build the executable:

```bash 
# For Linux/WSL
go build -o getdiskspace
``` 

```pwsh
# For Windows (PowerShell)
go build -o getdiskspace.exe
```


### Run the tool to check disk usage for a specific path:

```bash
./getdiskspace usage --path /
```

or on Windows:

```powershell
.\getdiskspace.exe usage --path C:\
``` 

Using Makefile (Linux/WSL)
Run directly (builds and runs):

```bash
make run
``` 

Build executable only:

```bash
make build-linux
```