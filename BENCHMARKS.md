# Benchmarks

Measurements taken on Linux (Arch, kernel 6.18.2), AMD Ryzen, Go 1.24.0.

## Go version

| Metric | Value |
|---|---|
| Binary size | 27 MB |
| Startup time (initialize response) | ~10 ms |
| All tests | 0.03s |

Startup measured by piping an MCP `initialize` request via stdin and timing the response:

```bash
time echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"bench","version":"1.0"}}}' | ./google-workspace-mcp-go 2>/dev/null > /dev/null
```

Three runs: 12ms, 9ms, 10ms.

## Python version (original)

The [original Python server](https://github.com/taylorwilsdon/google_workspace_mcp) (v1.11.5) requires Python 3.10+ and uv/pip with ~11 direct dependencies (FastMCP, FastAPI, google-api-python-client, etc.).

Startup measured with the same method (pipe MCP `initialize` via stdin, same machine):

Three runs: 3.943s, 3.145s, 3.022s (~3s average).

## Reproduce

```bash
# Build
go build -o google-workspace-mcp-go .

# Binary size
ls -lh google-workspace-mcp-go

# Startup time
time echo '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"bench","version":"1.0"}}}' | ./google-workspace-mcp-go 2>/dev/null > /dev/null

# Idle RSS (start server, then check /proc/$PID/status)
echo '...' | ./google-workspace-mcp-go &
PID=$!; sleep 1; grep VmRSS /proc/$PID/status; kill $PID
```
