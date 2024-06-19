# Тестируем 

## 1
```
cd cmd
go run main.go
```

## 2 Проверяем инсерт
```
curl --location 'localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "100 Go Mistakes and How to Avoid Them",
    "author": "Teiva Harsanyi",
    "year_published": 2022,
     "price":  60,
     "category": "education"
}'

curl --location 'localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "100 Go Mistakes and How to Avoid Them",
    "author": "Teiva Harsanyi",
    "year_published": 2022,
     "price":  60,
     "category": "education"
}'
```

## 3 Проверяем чтение всех
```
curl --location 'localhost:8080/books'
```

## 4 Проверяем чтение по id

```
curl --location 'localhost:8080/books/1'
curl --location 'localhost:8080/books/2'
```

## 5 Проверяем Обновление

```
curl --location --request PUT 'localhost:8080/books/1' \
--header 'Content-Type: application/json' \
--data '{
    "price": 160
}'
```

## 6 Проверяем удаление

```
curl --location --request DELETE 'localhost:8080/books/2' \
--header 'Content-Type: application/json' \
--data '{
    "price": 60
}'
```

# Dockerize

```
docker-compose -f "docker-compose.yml" up -d --build
```