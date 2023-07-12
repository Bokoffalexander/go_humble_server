/*
Небольшой веб-сервер.
Получает из строки запроса значение переменной name 
и выводить это значение в окно браузера.
*/
package main // компилятор Гоу когда компилирует программу он сначала ищет пакет мэйн

import (
	"fmt" // для вывода в браузер
	"log" // будет выводить информацию в консоль
	"net/http" // предоставляет реализации HTTP-клиента и сервера
)

// запрос с строке: http://127.0.0.1:8000/?name=Sasha&age=31
// ответ в браузере: My name is Sasha. Age is 31.
func greetings(w http.ResponseWriter, r *http.Request) { 
	// w - интерфейс, записывает ответ(вывод) в браузер.
	// r - http запрос, который будет приходить на наш веб-сервер
	name := r.URL.Query().Get("name") // значение из URL строчки т.е. строки запроса браузера
	age := r.URL.Query().Get("age") // обьект URL, метод Query(), функция Get("age")
	fmt.Fprintf(w, "<h1>My name is %s. Age is %v.</h1>", name, age) // выводит значение в какой-либо интерфейс
}

// http://127.0.0.1:8000/city/?city=Rostov-on-Don
// My city is Rostov-on-Don.
func city(w http.ResponseWriter, r *http.Request) { 
	// w - интерфейс, записывает ответ(вывод) в браузер.
	// r - http запрос, который будет приходить на наш веб-сервер
	city := r.URL.Query().Get("city") // значение из URL строчки т.е. строки запроса браузера
	fmt.Fprintf(w, "<h1>My city is %s.</h1>", city) // выводит значение в какой-либо интерфейс
}

func main() { // компилятор ищет функцию мэйн
	http.HandleFunc("/", greetings) // обработичик по такому адресу
	http.HandleFunc("/city/", city) // обработичик по такому адресу
	log.Println("http://localhost:8000") // вывод информации в консоль
	http.ListenAndServe(":8000", nil) // запускает дефолтный листенер http-сервера
}

