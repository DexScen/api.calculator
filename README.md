API для подсчета выражений

Пример работы с CURL на WINDOWS:

Правильное выражение

curl --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\\"expression\\": \\"2+2\\"}"
{"result":4}

Пример внутренней ошибки, при отправке без "

curl --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\\"expression\\": \\"2+2}"
{"error":"Internal server error"}

Пример ошибки с неправильным expression

curl --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\\"expression\\": \\"2+2a\\"}"
{"error":"Expression is not valid"}

Деление на 0

curl --location http://localhost:8080/api/v1/calculate --header "Content-Type: application/json" --data "{\\"expression\\": \\"2/0\\"}"
{"error":"Expression is not valid"}

Пример PostMan:

Правильное выражение

![image](https://github.com/user-attachments/assets/5e626412-2697-448b-be1d-fbeb5f9f9e21)


Пример внутренней ошибки, при отправке без "

![image](https://github.com/user-attachments/assets/12da9552-6e0b-4a65-8641-dc339e5c2090)


Пример ошибки с неправильным expression

![image](https://github.com/user-attachments/assets/c672740f-f897-44e5-914b-6e94f11445c6)


Деление на 0

![image](https://github.com/user-attachments/assets/428cb03c-2da8-4b45-8cb4-e5bed62de92b)


Запускается из корневой директории проекта api.calculator/
командой: go run ./cmd/main.go 

С наступающим новым годом !
