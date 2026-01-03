# go-projects
Collection of projects suggested by ChatGPT to learn Go

### ✅ Beginner Go Projects
1. **CLI calculator** - **DONE**
- takes input like: go run calc.go 2 + 5
- supports +, -, *, /
- optional: parentheses

👉 learn: os.Args, error handling, switches

2. **Word counter** - **DONE**

- Input: file path
- Output: counts lines, words, chars

👉 learn: working with files, bufio scanner, maps

3. **JSON Todo list (local only)** - **DONE**

- store todos in a JSON file
- add/remove/list from CLI

👉 learn: encoding/json, structs, slices

4. **Newsletter signup API** - **DONE**

- POST /signup store emails in memory or file

👉 learn: net/http, handlers, JSON responses

### 🔥 Intermediate Go Projects
5. **TCP chat server** - **DONE**

- clients connect via netcat or socat
- broadcast messages to all users

👉 learn: goroutines + channels

6. **Worker pool**

- queues tasks
- workers process tasks concurrently

👉 learn: concurrency patterns, rate limiting

7. **Simple Proxy Server** 

- listens on port 8080
- forwards requests to another host

👉 learn: io.Copy, networking primitives

8. **URL shortener**

- redirect logic
- persistent storage using BoltDB or SQLite

👉 learn: lightweight database + HTTP

### ⚙️ Advanced but small Go Projects
9. **Minimal container runtime** - **DONE**

- (not full Docker, just a toy)
- use namespaces + chroot
- run a command isolated

👉 learn: Linux syscalls + Go

10. **Simple key-value DB** - **DONE**

- append-only log file
- index in memory

👉 learn: I/O performance, WAL concepts