# acat v1.0.0
### Advanced Cat (Written in Go)

A fast, hacker-friendly CLI utility that visualizes invisible/control characters in text streams. Built for debugging, security testing, payload inspection, and detecting hidden characters in files or piped data.

---

## 📌 Features

- Visualize spaces as visible characters (default: `.`)
- Show newline as literal `\n` while preserving actual line breaks
- Highlight control characters:
  - `\t` (tab)
  - `\r` (carriage return)
  - `\0` (null byte)
- Smart color handling (TTY-aware)
- Works seamlessly with pipes (stdin/stdout)
- File input support
- Line number support (`-l`)
- Raw mode (`--raw`) — behaves like normal `cat`
- Control-character-only mode (`--only`)
- Customizable visible space character
- Efficient streaming using `bufio` (handles large files)

---

## ⚙️ Options

| Option        | Description                                      |
|---------------|--------------------------------------------------|
| `-f <file>`   | Read input from file                             |
| `-nc`         | Disable color output                             |
| `-l`          | Show line numbers                                |
| `--raw`       | Raw mode (no transformations, like `cat`)        |
| `--only`      | Show only control characters                     |
| `-c <char>`   | Customize visible space character (default: `.`) |

---

## 🧪 Examples

### Basic usage
```bash
root@root:~$ acat file.txt
-.Show.newline.as.literal.`\n`.while.preserving.actual.line.breaks\n
-.Highlight.control.characters:\n
..-.`\t`.(tab)\n
..-.`\r`.(carriage.return)\n
..-.`\0`.(null.byte)\n
-.Smart.color.handling.(TTY-aware)\n
-.Works.seamlessly.with.pipes.(stdin/stdout)\n
-.File.input.support\n
-.Line.number.support.(`-l`)\n
-.Raw.mode.(`--raw`).—.behaves.like.normal.`cat`\n
-.Control-character-only.mode.(`--only`)\n
-.Customizable.visible.space.character\n
-.Efficient.streaming.using.`bufio`.(handles.large.files)
```
---

### Installation:

#### Install dependency:
```bash
go get golang.org/x/term
```
#### Clone repository:
```bash
git clone https://github.com/Shaikh-Khizer/acat.git
cd acat
```
#### Initialize module:
```bash
go mod init acat
go mod tidy
```
#### Build
```bash
go build -o acat
```
#### Make global path (optional):
```bash
sudo mv acat /usr/local/bin/
```

### Run:
```bash
acat -h
```
---

## 📄 License

This project is licensed under the MIT License.  
See the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Shaikh Khizer**  
Computer Science Student | Penetration Tester
