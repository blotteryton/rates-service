# CoinsMarketCup TONCoin rates fetcher

Сервис получения курсов токена TON coin (`toncoin`) с биржи
CoinsMarketCup в валютах `RUB` и `USD`.

## Запуск

Для запуска сервиса предварительно необходимо получить `API KEY` для доступа к API CoinsMarketCup. Далее этот ключ необходимо передать в переменную окружения `API_KEY`.

### Пример запуска docker-контейнера

```sh
docker run -e API_KEY=YOUR-API-KEY -p 8080:8080 -it cmc_fetcher
```

### Локальный запуск

Для локального запуска (например, при разработке/доработке) необходимо в корень проекта положить файл `.env` примерно такого содержания

```
#.env file
API_KEY=YOUR-API-KEY
```

После этого выполнить команду

```bash
make run
```

или

```bash
go run cmd/main.go
```

## Получение курса TonCoin

### Получение курса в RUB

```http
GET http://localhost:8080/api/v1/rates/toncoin/RUB
Content-Type: application/json
```

В ответ придет структура

```json
{
  "rate": 123.45
}
```

### Получение курса в USD

```http
GET http://localhost:8080/api/v1/rates/toncoin/USD
Content-Type: application/json
```

В ответ придет структура

```json
{
  "rate": 1.2345
}
```

## Дополнительная информация

Дополнительные инструменты/хелперы можно посмотреть в `Makefile`

**Важно!** Необходимо кешировать ответ запроса, т.к. каждый запрос расходует лимит операций к API CoinsMarketCup. Рассчитать кеш можно исходя из тарифного плана полученного `API-KEY`. 1 запрос к сервису - один запрос к API CoinsMarketCup.
