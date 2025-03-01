# dvopsctl - DevOps CLI Tool

`dvopsctl` is a command-line interface written in Go, designed to simplify Dev & DevOps tasks. It automates common actions, making daily work easier for developers and DevOps engineers.

---

## Installation

### Clone the Project

Start by retrieving the source code from GitHub:
```sh
git clone https://github.com/MyLittleParis/dvopsctl.git && \
cd dvopsctl
```

### Build the Tool

Make sure you have Go installed (>=1.18),
```sh
go version
```
If not look at https://go.dev/doc/install.  

Then run the following command:
```sh
go build -o dvopsctl
```
This will generate an executable `dvopsctl` in the current directory.

### Add `dvopsctl` to PATH

This step is essential to run `dvopsctl` from any directory in your terminal without specifying its full path. Otherwise, you would always need to type `./dvopsctl` from its directory. Adding the binary to `PATH` allows for smoother and more intuitive usage, just like any system command.

To use `dvopsctl` from any location in the terminal, add its path to `PATH`:

#### On Linux / macOS
```sh
mv dvopsctl /usr/local/bin/
```
Or add the project directory directly to `PATH`:
```sh
echo 'export PATH="$PATH:$(pwd)"' >> ~/.bashrc
source ~/.bashrc
```

#### On Windows (PowerShell)
```powershell
$env:Path += ";$PWD"
```
Or manually add the folder containing `dvopsctl.exe` to Windows environment variables.

---

## Usage

Once installed, you can use the following commands:
```sh
dvopsctl server -open
```
This command search in your project if variable `SERVER_NAME` exists and open it in your browser:

---

## Development

If you want to test or modify `dvopsctl` without compiling it, you can run it directly with:
```sh
go run main.go server -open
```

---

## License

This project is licensed under the MIT License. You are free to modify and redistribute it under the terms of this license.

---

## Next Steps

We plan to add new features, including:
** Docker container management
** Kubernetes interaction
** AWS integration
** Monitoring and observability

---

**Contribute!**

If you have improvement ideas or want to report an issue, feel free to open an issue on GitHub! 🚀
