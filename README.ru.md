![Go](https://img.shields.io/badge/Go-1.26.6-blue)
[![License](https://img.shields.io/badge/License-GPL--3.0-green)](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20MacOS-red)
![Last commit](https://img.shields.io/github/last-commit/Tima-games/TextStats)

[Russian version] | [[English version](README.md)]


# CLI-утилита Text Stats

Простая CLI-утилита, написанная на Go, для анализа текста.

Я сделал этот проект для обучения Go, Git'у, GitHub'у и просто по фану, поэтому не относитесь к нему серьёзно.

## Возможности
- Подсчёт строк
- Подсчёт слов
- Подсчёт символов (без пробелов)
- Подсчёт пробелов
- Чтение из файла

## Использование
```Bash
./textstats-[ваша версия/архитектура]
```
**Введите свой текст**

**Завершите ввод Control + D (Linux/MacOS)**

или

```Bash
./textstats-[ваша версия/архитектура] text.txt
```
**"text.txt"** - любой текстовый файл. 

*Корректнно* поддерживаемые форматы: **любой UTF-8 текстовый файл** 

*Например:* .txt .md .go .py .json ...
## Пример
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
## Примечания
- Слова разделяются пробельными символами
- Поддерживается Unicode

## Установка
Скачайте готовый бинарник со [страницы релизов](https://github.com/Tima-games/TextStats/releases).

или

### Сборка из исходников
```bash
git clone https://github.com/Tima-games/TextStats.git
cd TextStats
go build -o textstats
```

## План развития
- [x] Подсчёт строк
- [x] Подсчёт слов
- [x] Подсчёт символов
- [x] Подсчёт пробелов
- [x] README
- [x] Чтение из файла
- [ ] Поддержка Windows
- [ ] GUI
- [ ] Встроенный редактор

...и многое другое

## Лицензия
[GPL-3.0 license](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
