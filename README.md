## Flow live Reload

![preview](https://github.com/user-attachments/assets/a4d204ca-83e0-4c7b-8da7-d8d9cbb87ea7)

---

Paste the below `.flow.toml` file and the binary in the root of your project.

```toml
# .flow.toml

root = "."

# amount of time to wait before triggering new build. default=500
debounce = 300

[build]
  # bin stores generated binary
  bin = "tmp/main"

  # the command which will be executed to generate binary
  cmd = "go build -o tmp/main main.go && ./tmp/main"

  # extensions to include
  include_ext = ["go", "html"]

  # patterns to exclude in filenames, this may include test or auto-generated files
  exclude_regex = ["*_test.go", ".*_templ.go"]

  # directories to exclude (optional)
  exclude_dir = []
```
