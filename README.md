![Go](https://img.shields.io/badge/Go-1.26-blue)
[![License](https://img.shields.io/badge/License-GPL--3.0-green)](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20MacOS-red)
![Last commit](https://img.shields.io/github/last-commit/Tima-games/TextStats)

[English version] | [[Russian version](README.ru.md)]

# CLI-utility Text Stats

Simple CLI tool written in Go to analyze text.

I made this project for learning Go, Git, GitHub and just for fun, so don't be too serious.

## Features
- Count lines
- Count words
- Count letters (excluding spaces)
- Count spaces
- Read from file
- Flags (in β version)

## Usage
```Bash
./textstats-[your version/arch] [Flag(s)]
```
**Type your text**

**Finish input with Enter & Control + D (Linux/MacOS)**

or

```Bash
./textstats-[your version/arch] text.txt
```
**"text.txt"** - any text file. 

*Correctly* supported formats: **any UTF-8 text file** 

*For example:* .txt .md .go .py .json ...

### Flags (available in beta)

*All supported flags:*

**-h, --help - Shows help message and exit**

**-v, --version - Shows version and exit**  

**-l, --lines - Shows only lines count**

**-w, --words - Shows only words count** 

**-s, --spaces - Shows only spaces count**

**-c, --letters - Shows only letters count**

**And you also can use it with files.**


## Examples
```
Hello world

In your string(s) 1 lines
In your string(s) 2 words
In your string(s) 1 spaces
In your string(s) 10 letters
```
```
text.txt

In your string(s) 6 lines
In your string(s) 29 words
In your string(s) 19 spaces
In your string(s) 78 letters
```
```
-l
Hello world

In your string(s) 1 lines
```
```
-v
TextStats v1.5.0 (11-04-26 release)
```
```
-c example.md

In your string(s) 6561 letters
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
- [x] Read from file
- [x] Flags
- [ ] Windows support
- [ ] GUI
- [ ] Built-in text editor

...and much more

## License
[GPL-3.0 license](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
