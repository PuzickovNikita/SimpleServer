# SimpleServer
Чтобы запустить сервис выполнить команду  

go run .\cmd\start.go <username> <password> <adress> <database name>  
    
POST запрос получает json вида  
    
{  
    "table",  
    "body"  
}  
и добавляет в указанную таблицу указанный body  
GET запрос получается в параметрах имя таблицы, ключ и возвращает по нему json  
{  
    "key"  
    "body"  
}  
DELETE запрос получает в параметрах имя таблицы, ключ и возвращает в теле ответа информацию о кол-ве удаленных рядов<br/>
  
примеры запросов:  
POST localhost:8080/PSQL/JSON  
Content-Type: application/json  
  
{  
"table": "simpletable",  
  "body" : "puzickov"  
}  
  
GET localhost:8080/PSQL/JSON?table=simpletable&key=3  

DELETE localhost:8080/PSQL/JSON?table=simpletable&key=1  
