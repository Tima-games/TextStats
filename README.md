![Go](https://img.shields.io/badge/Go-1.25.5-blue)
[![License](https://img.shields.io/badge/License-GPL--3.0-green)](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20MacOS-red)
![Last commit](https://img.shields.io/github/last-commit/Tima-games/TextStats)

[English version] | [[Russian version](README.ru.md)]

# CLI-utility Text Stats

Simple CLI tool written in Go to analyze text.

I made this project for learning and just for fun, so don't be too serious.

## Features
- Count lines
- Count words
- Count letters (excluding spaces)
- Count spaces

## Usage
```Bash
./textstats
```
**Type your text**

**Finish input with Control + D (Linux/MacOS)**
## Example
```
Hello world

In your string(s) 1 lines
In your string(s) 2 words
In your string(s) 1 spaces
In your string(s) 10 letters
```
## Notes
- Words are separated by whitespace
- Unicode is supported

## Installation
Download binary from [Releases page](https://github.com/Tima-games/TextStats/releases).

or

### Build from source
```bash
git clone https://github.com/Tima-games/TextStats.git
cd TextStats
go build -o textstats
```

## Roadmap
- [x] Count lines
- [x] Count words
- [x] Count letters
- [x] Count spaces
- [x] README
- [ ] Read from file
- [ ] Windows support

...and much more

## License
[GPL-3.0 license](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
