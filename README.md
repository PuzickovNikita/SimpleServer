# SimpleServer
Чтобы запустить сервис выполнить команду

    go run .\cmd\start.go <username> <password> <address> <database_name>

Все запросы указывают в параметрах имя таблицы, с которой производится взаимодействие и json в своем теле\

Все json имеют вид:

    {
        "key"
        "body"
    }

POST запрос получает json вида

примеры запросов:

    POST localhost:8080/PSQL/JSON?table=simpletable\
    Content-Type: application/json
    
    {
        "body" : "puzickov"
    }
<hr>

    GET localhost:8080/PSQL/JSON?table=simpletable\
    Content-Type: application/json
    
    {
        "key": 1
    }
<hr>

    GET localhost:8080/PSQL/JSON?table=simpletable\
    Content-Type: application/json
    
    {
    }
<hr>

    GET localhost:8080/PSQL/JSON?table=simpletable\
    Content-Type: application/json

    {
        "body": "puzickov"
    }
<hr>

    DELETE localhost:8080/PSQL/JSON?table=simpletable\
    Content-Type: application/json
    
    {
        "key": 1
    }

