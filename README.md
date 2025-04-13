## Flow live Reload

Paste the below `.flow.toml` file and the binary in the root of your project.

```toml
# .flow.toml

root = "."

[build]
  # bin stores generated binary
  bin = "tmp/main"

  # the command which will be executed to generate binary
  cmd = "go build -o tmp/main main.go && ./tmp/main"

  # extensions to include (optional)
  include_ext = []

  # directories to exclude (optional)
  exclude_dir = []
```
