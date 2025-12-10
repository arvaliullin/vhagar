# Зависимости проекта Vhagar

Данный документ содержит информацию о всех библиотеках, инструментах и образе, используемых в devcontainer.

**Условные обозначения:**
- ✅ **В сборке** - библиотека установлена в образе и указана в go.mod, может входить в состав исполняемого файла
- ❌ **Только для разработки** - инструмент/библиотека используется только при разработке/тестировании, не попадает в исполняемый файл

## Базовый образ

| Компонент | Версия | Источник | Описание |
|-----------|--------|----------|----------|
| Base Image | `registry.astralinux.ru/library/astra/ubi18-golang121:1.8.4` | `registry.astralinux.ru` | Базовый образ на основе Astra Linux UBI 18 с предустановленным Go 1.21 |

## Системные инструменты (APT пакеты)

| Пакет | Версия | Источник | В сборке | Описание |
|-------|--------|----------|----------|----------|
| git | `1:2.43.0-1ubuntu7.3+b3` | main-repository | ❌ | Система контроля версий |
| curl | `7.88.1-10+deb12u15.astra1` | main-repository | ❌ | Утилита для передачи данных по URL |
| wget | `1.21.3-1.astra1+b3` | main-repository | ❌ | Утилита для загрузки файлов из интернета |
| make | `4.3-4.1+b4` | main-repository | ❌ | Утилита для управления сборкой |
| zsh | `5.9-4+b6` | main-repository | ❌ | Расширенная оболочка командной строки |
| vim | `2:9.1.1230-2.astra10` | main-repository | ❌ | Текстовый редактор |
| ca-certificates | `20230311+b9` | main-repository | ❌ | Сертификаты CA для SSL/TLS |
| sudo | - | main-repository | ❌ | Утилита для выполнения команд от имени суперпользователя |
| openssh-client | - | main-repository | ❌ | SSH клиент для удаленного доступа |

**APT репозитории:**
- `https://download.astralinux.ru/astra/frozen/1.8_x86-64/1.8.4/main-repository` (в таблицах указано как `main-repository`)
- `https://download.astralinux.ru/astra/frozen/1.8_x86-64/1.8.4/extended-repository` (в таблицах указано как `extended-repository`)

## Go и инструменты разработки

| Инструмент | Версия | Источник | В сборке | Описание |
|------------|--------|----------|----------|----------|
| Go | `1.21.10` | extended-repository | ❌ | Язык программирования Go |
| golang-1.21 | `1.21.10-1.astra3` | extended-repository | ❌ | Пакет Go 1.21 |
| gopls | `v0.20.0` | `go install golang.org/x/tools/gopls@latest` | ❌ | Go Language Server для IDE |
| delve (dlv) | `1.21.2` | `go install github.com/go-delve/delve/cmd/dlv@v1.21.2` | ❌ | Отладчик для Go |
| golangci-lint | `2.7.2` | `https://github.com/golangci/golangci-lint` | ❌ | Линтер для Go кода |
| mockgen | `1.6.0` | extended-repository (golang-github-golang-mock-dev) | ❌ | Генератор моков для Go (часть golang-mock) |

## Docker инструменты

| Инструмент | Версия | Источник | В сборке | Описание |
|------------|--------|----------|----------|----------|
| Docker | `26.1.0` | `https://download.docker.com/linux/static/stable/x86_64/` | ❌ | Контейнеризация приложений |
| Docker Compose | `v2.29.2` | `https://github.com/docker/compose/releases` | ❌ | Оркестрация многоконтейнерных приложений |

## Go библиотеки (зависимости проекта)

### Основные зависимости (go.mod)

| Библиотека | Версия | Источник | В сборке | Описание |
|------------|--------|----------|----------|----------|
| gopkg.in/yaml.v3 | `3.0.0` | extended-repository (`golang-gopkg-yaml.v3-dev`)<br>GitHub: `https://gopkg.in/yaml.v3` | ✅ | Парсер YAML для Go |
| github.com/streadway/amqp | `1.0.0` | extended-repository (`golang-github-streadway-amqp-dev`)<br>GitHub: `https://github.com/streadway/amqp` | ✅ | Клиент AMQP 0.9.1 для RabbitMQ |
| github.com/lib/pq | `1.10.9` | extended-repository (`golang-github-lib-pq-dev`)<br>GitHub: `https://github.com/lib/pq` | ✅ | Драйвер PostgreSQL для database/sql |
| github.com/golang/mock | `1.6.0` | extended-repository (`golang-github-golang-mock-dev`)<br>GitHub: `https://github.com/golang/mock` | ❌ | Фреймворк для создания моков (только для тестов и генерации) |
| github.com/go-chi/chi | `1.5.4` | extended-repository (`golang-github-go-chi-chi-dev`)<br>GitHub: `https://github.com/go-chi/chi` | ✅ | Легковесный HTTP роутер |

### Установленные через APT (dev пакеты)

| Пакет | Версия APT | Версия библиотеки | Источник | В сборке | Описание |
|-------|------------|-------------------|----------|----------|----------|
| golang-gopkg-yaml.v3-dev | `3.0.1-3+b1` | `3.0.1` | extended-repository | ✅ | YAML поддержка для Go |
| golang-github-streadway-amqp-dev | `0.0~git20200716.e6b33f4-3+b1` | `e6b33f4` (2020-07-16) | extended-repository | ✅ | Go клиент для AMQP 0.9.1 |
| golang-github-lib-pq-dev | `1.10.9-2+b1` | `1.10.9` | extended-repository | ✅ | Чистый Go драйвер PostgreSQL |
| golang-github-golang-mock-dev | `1.6.0-2+b3` | `1.6.0` | extended-repository | ❌ | Фреймворк для моков в Go |
| golang-github-go-chi-chi-dev | `5.2.0-1+b1` | `5.2.0` | extended-repository | ✅ | Легковесный идиоматичный роутер для Go HTTP сервисов |
| golang-golang-x-mod-dev | `0.19.0-1+b1` | `0.19.0` | extended-repository | ❌ | Модули Go (x/mod) - инструмент разработки |
| golang-golang-x-tools-dev | `1:0.25.0+ds-1` | `0.25.0` | extended-repository | ❌ | Инструменты Go (x/tools) - инструмент разработки |
| golang-golang-x-tools | `1:0.25.0+ds-1` | `0.25.0` | extended-repository | ❌ | Исполняемые инструменты Go |

## VS Code расширения

| Расширение | ID | Описание |
|------------|----|----------|
| Go | `golang.go` | Официальное расширение Go для VS Code |
| Docker | `ms-azuretools.vscode-docker` | Расширение для работы с Docker |

## Переменные окружения

| Переменная | Значение | Описание |
|------------|----------|----------|
| GOTOOLCHAIN | `auto` | Автоматический выбор версии Go toolchain |
| GOPATH | `/usr/share/gocode` | Путь к рабочему пространству Go |
| DOCKER_HOST | `unix:///var/run/docker.sock` | Сокет Docker для подключения |
| GIT_EDITOR | `vim` | Редактор по умолчанию для Git |

## Пробрасываемые порты

| Порт | Назначение |
|------|------------|
| 5672 | RabbitMQ AMQP |
| 15672 | RabbitMQ Management UI |

## Примечания

- Все Go зависимости установлены через APT пакеты в `/usr/share/gocode/src/`
- В `go.mod` используются директивы `replace` для указания локальных путей к зависимостям
- Oh My Zsh установлен для улучшения работы с zsh
- Docker и Docker Compose установлены вручную из официальных релизов

## Детали использования библиотек в сборке

### Библиотеки, входящие в состав исполняемого файла (✅ В сборке):

Все библиотеки, указанные в `go.mod` и установленные через APT, могут входить в состав исполняемого файла при использовании в коде проекта:

- **gopkg.in/yaml.v3** - парсер YAML для работы с конфигурационными файлами
- **github.com/streadway/amqp** - клиент для работы с RabbitMQ (AMQP 0.9.1)
- **github.com/lib/pq** - драйвер PostgreSQL для работы с базой данных
- **github.com/go-chi/chi** - HTTP роутер для создания веб-сервисов

*Примечание: Проект находится в стадии разработки. Текущие примеры использования библиотек находятся в директории `examples/`, но все указанные библиотеки доступны для использования в основном коде проекта.*

### Инструменты разработки (❌ Только для разработки):

- **github.com/golang/mock** - используется только для генерации моков и тестирования
- **golang-golang-x-mod-dev**, **golang-golang-x-tools-dev** - инструменты для разработки, не попадают в сборку
- **gopls**, **delve**, **golangci-lint**, **mockgen** - инструменты разработки, не попадают в сборку

