# 🚀 **Abort**

**Abort** is a lightweight and efficient Go library for gracefully halting system processes. Designed with simplicity and reliability in mind, Abort provides developers with an intuitive API to manage process termination tasks. 💡

---

## ✨ **Features**

- ⚡ **Quick Process Termination**: Instantly halt system processes by ID or name.
- 🤝 **Safe and Graceful**: Ensures processes are terminated cleanly.
- 💻 **Cross-Platform**: Works seamlessly on Windows, Linux, and macOS.
- 📦 **Minimal Dependencies**: Keep your project lightweight and efficient.
- 🔒 **Secure Handling**: Avoid unintended terminations with robust checks.

---

## 📥 **Installation**

Install the Abort library using `go get`:

```bash
go get github.com/yourusername/abort
```

---

## 🚀 **Getting Started**

Here’s how you can get started with **Abort** in your Go project:

### Import the Library

```go
package main

import (
 "fmt"
 "github.com/mrcurlybrown/abort"
)
```

### Example: Terminate a Process by ID

```go
func main() {
 pid := 1234 // Replace with the PID of the process you want to terminate
 err := abort.ByPID(pid)
 if err != nil {
  fmt.Printf("❌ Failed to abort process %d: %v\n", pid, err)
 } else {
  fmt.Printf("✅ Successfully aborted process %d\n", pid)
 }
}
```

### Example: Terminate a Process by Name

```go
func main() {
 name := "example-process" // Replace with the process name
 err := abort.ByName(name)
 if err != nil {
  fmt.Printf("❌ Failed to abort process '%s': %v\n", name, err)
 } else {
  fmt.Printf("✅ Successfully aborted process '%s'\n", name)
 }
}
```

---

## 🛠️ **API Reference**

### `abort.ByPID(pid int) error`

Halts a process by its PID (Process ID).

- **Parameters**:
  - `pid`: The process ID of the target process.
- **Returns**:
  - `error`: Returns an error if the process termination fails.

### `abort.ByName(name string) error`

Halts a process by its name.

- **Parameters**:
  - `name`: The name of the target process.
- **Returns**:
  - `error`: Returns an error if the process termination fails.

---

## 🌐 **Platform Support**

| Platform    | Status   |
|-------------|----------|
<!-- | Windows     | ✅ Pending | -->
| Linux       | ✅ Supported |
<!-- | macOS       | ✅ Pending | -->

---

## 📞 **Support**

Need help? Have questions? Feel free to reach out:

- 💌 Email: <mrcurlybrown@gmail.com>
- 🐦 Twitter: [@yourusername](https://twitter.com)
- 📂 Issues: [GitHub Issues](https://github.com/mrcurlybrowm/abort/issues)

---

Made with ❤️ and Go! 🐹
