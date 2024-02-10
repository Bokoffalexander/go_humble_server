# go_humble_server
Небольшой веб-сервер.
Получает из строки запроса значение переменной name 
и выводить это значение в окно браузера.

Запустить сервер:
./server

// запрос с строке: http://3vb.ru:6006/?name=Sasha&age=31
// ответ в браузере: My name is **Sasha**. Age is **31**.

// запрос с строке: http://3vb.ru:6006/city/?city=Rostov-on-Don
// ответ в браузере: My city is **Rostov-on-Don**.
