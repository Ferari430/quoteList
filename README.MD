В данном репозиторие выполнена программа по тестовому заданию на тему 
Мини-сервис Цитатник  

Команды для тестов через curl:
    Добавление новой цитаты (POST /quotes):
curl -X POST http://localhost:8080/quotes \ -H "Content-Type: application/json" \ -d 
'{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}' 
 

ВНИМАНИЕ: Данная команда не запустилась через терминал на моем компьютере
(ошибка: Вместо \ необходимо писать /):, вместо нее я использовал:

curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

Следующие команды работают без ошибок:

    Получение всех цитат:
curl http://localhost:8080/quotes

    Получение случайной цитаты:
 curl http://localhost:8080/quotes/random

    Фильтр по авторам:
curl http://localhost:8080/quotes?author=Confucius 

    Удаление по ID:
curl -X DELETE http://localhost:8080/quotes/1 

Проект разбит на пэкэджи для удобства чтения кода. 
Слой repo вынесен отдельно, в нем находится интерфейс Storage, реализующий
CRUD операции, это сделано для того, чтобы in-memory хранилище можно было заменить на базу данных, например.
В силу либо моей неопытности, либо конструкции некоторых функций(например в in-memory.go) Функции не возращают error, что затрудняет их тестирование. Однако для функций Json (в pkg/res/res.go) и Decode (в pkg/req/decode.go) были все же написаны тесты!
Более того, использовался validator/v10 из 
github.com/go-playground/validator/v10.com,  который под капотом проверяет отправленные данные в сервис, а именно: в payload.go напротив полей структуры Quote установлены соответствующие теги, требующие передачи этих полей.

Спасибо за уделенное время, пожалуйста, оставьте обратную связь! 
TG: @Airfool
