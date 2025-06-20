# Автоудалитель автодобавляемых товаров из акций Ozon

Озон периодически автодобаляет товары в акции и их приходится постоянно удалять вручную из разделов стабильно раз\два в неделю.
Это тратит время. Чтобы его сэкономить, был разработан инструмент, который делает это автоматически.

## Требования
Нужен Linux\MacOS и установленный go. На Windows должно работать, но не тестировал

## Установка
1. Скачать последний релиз с [этой страницы](https://github.com/execut/ozon-promotions-auto-deleter/releases/latest) для своей архитектуры
1. Распаковать в любую папку
1. Создать файл с настройками `config.yml`, скопировав `config.example.yml`
1. Заполнить его своими данными:
   1. **cookie** - все кукисы, которые браузер передаёт при скачивании одного из ваших отчётов. Их нужно периодически обновлять на актуальные.
  Можно подсмотреть во вкладке консоли разработчика браузера Chrome Network -> кликнуть на произвольный запрос, который скачивал отчёт и скопировать там значение заголовка Cookie
   1. **companyID** - номер компании. Можно подсмотреть тоже в заголовках браузера, либо на странице https://id.ozon.ru/ во вкладке
  Ozon Seller -> Развернуть через кнопку рядом с почтой -> цифры после строчки "Seller ID: sc-"
   1. **organizationID** - номер организации. Можно посмотреть в заголовках браузера
1. Запустить исполняемый файл `ozon-promotions-auto-deleter`
1. При успешном запуске должно отобразиться сообщение: `Успешно очишено X акций`. Где X - общее количество акций
1. После этого все товары, которые Озон обещает автодобавить, удаляются из акций.

## Для разработчиков:
```
git clone 
go run main.go
```