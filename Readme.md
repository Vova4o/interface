# **Простенький пример чистой архитектуры**

1. В файле main создаем иммитацию подключения к ДБ и передаем его на уровень Sorage
2. В сервисный уровень передаем структуру уровня Storage
3. На уровень Handlers передаем сервисный слой который работает с внешними данными

В таком виде нет привязки к конкретным реализациям и улучшена масштабируемость и работа с данными.
В случае замены одного из слоев, нет необходимости менять все остальные слои.