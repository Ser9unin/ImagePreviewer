# Превьювер изображений
## Общее описание
Сервис предназначен для изготовления preview (создания изображения
с новыми размерами на основе имеющегося изображения).

## Архитектура
Сервис представляет собой web-сервер (прокси), загружающий изображения,
масштабирующий/обрезающий их до нужного формата и возвращающий пользователю.

## Основной обработчик
http://localhost:8000/fill/300/200/images.wallpaperscraft.com/image/single/beaver_cute_art_127732_1366x768.jpg

<- микросервис -><- размеры превью -><--------- URL исходного изображения --------------------------------->

В URL выше:
- http://localhost:8000/fill/300/200/ - endpoint нашего сервиса,
в котором 300x200 - это размеры финального изображения.
- https://images.wallpaperscraft.com/image/single/beaver_cute_art_127732_1366x768.jpg - адрес исходного изображения;

в API сервиса добавляется URL исходного изображения, утилита скачивает его, изменяет до необходимых размеров и возвращает.

## Конфигурация
Основной параметр конфигурации сервиса - разрешенный размер LRU-кэша.
Изменяется в файле `.env`, по-умолчанию установлено значение `3`.
Поскольку размер места для кэширования ограничен, то для удаления редко используемых изображений применен алгоритм **"Least Recent Used"**.

## Развертывание
Развертывание микросервиса можно произвести комадной `make run` в директории с проектом. (внутри `docker compose up`)
