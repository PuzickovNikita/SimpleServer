# SimpleServer
Чтобы запустить сервис выполнить команду<br/>
go run .\cmd\start.go <username> <password> <adress> <database name><br/>
POST запрос получает json вида <br/>
{<br/>
    "table",<br/>
    "body"<br/>
}<br/>
и добавляет в указанную таблицу указанный body<br/>
GET запрос получается в параметрах имя таблицы, ключ и возвращает по нему json<br/>
{<br/>
    "key"<br/>
    "body"<br/>
}<br/>
DELETE запрос получает в параметрах имя таблицы, ключ и возвращает в теле ответа информацию о кол-ве удаленных рядов<br/>
<br/>
примеры запросов:<br/>
POST localhost:8080/PSQL/JSON<br/>
Content-Type: application/json<br/>
<br/>
{<br/>
"table": "simpletable",<br/>
  "body" : "puzickov"<br/>
}<br/><br/>

GET localhost:8080/PSQL/JSON?table=simpletable&key=3<br/>

DELETE localhost:8080/PSQL/JSON?table=simpletable&key=1<br/>
