# Небольшой веб-сервер.
# Получает из строки запроса значения переменных
# Выводит ответ сервера в окно браузера.

Скомпилировал написанный файл: go build server.go
Запустил в фоновом режиме: ./server &

Перейдите по сслыкам ниже и изменяйте значение переменной
Отправим http запросы с методом get:

// запрос в строке: http://3vb.ru:6006/?name=Sasha&age=31
// ответ в браузере: My name is Sasha. Age is 31.

// запрос с строке: http://3vb.ru:6006/city/?city=Rostov-on-Don
// ответ в браузере: My city is Rostov-on-Don.
