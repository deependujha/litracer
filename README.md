# litracer

A tool for converting Lightning AI's LitData logs to Chrome-compatible trace files.

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
