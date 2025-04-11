# Litracer

A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files.

<img width="1439" alt="Screenshot 2025-04-08 at 12 35 15 PM" src="https://github.com/user-attachments/assets/cfa919c8-2ba2-4a7a-b054-d94ccb3e15b1" />

---

## Install

- For linux (deb)

```bash
⚡ ~ uname -m # displays machine hardware name
# lightning-studio: x86_64

# download relevant release (prefer latest version)
⚡ ~ wget https://github.com/deependujha/litracer/releases/download/v0.0.4/litracer_0.0.7_linux_amd64.deb

⚡ ~ sudo dpkg -i litracer_0.0.7_linux_amd64.deb
```

- Download using go

```bash
go install github.com/deependujha/litracer@latest
```

---

## Usage

```bash
litracer <log_file>
```

- Available options/flags:

```txt
Usage:
  litracer [flags]

Flags:
  -h, --help            help for litracer
  -o, --output string   Path to the output trace file (default "litdata_trace.json")
  -s, --sink int        Sink limit: number of lines to write at once (default 100)
  -w, --workers int     Number of worker goroutines to use for parsing (default 1)
```
