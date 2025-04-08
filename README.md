# litracer

A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files.

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
  -w, --workers int     Number of worker goroutines to use for parsing (default 1)
```
