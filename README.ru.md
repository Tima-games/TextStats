![Go](https://img.shields.io/badge/Go-1.25.5-blue)
[![License](https://img.shields.io/badge/License-GPL--3.0-green)](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20MacOS-red)
![Last commit](https://img.shields.io/github/last-commit/Tima-games/TextStats)

[Russian version] | [[English version](README.md)]


# CLI-утилита Text Stats

Простая CLI-утилита, написанная на Go, для анализа текста.

Я сделал этот проект для обучения и просто по фану, поэтому не относитесь к нему серьёзно.

## Возможности
- Подсчёт строк
- Подсчёт слов
- Подсчёт символов (без пробелов)
- Подсчёт пробелов

## Использование
```Bash
./textstats
```
**Введите свой текст**

**Завершите ввод Control + D (Linux/MacOS)**
## Пример
```
Hello world

In your string(s) 1 lines
In your string(s) 2 words
In your string(s) 1 spaces
In your string(s) 10 letters
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
- [ ] Чтение из файла
- [ ] Поддержка Windows

...и многое другое

## Лицензия
[GPL-3.0 license](https://github.com/Tima-games/TextStats/blob/main/LICENSE)
