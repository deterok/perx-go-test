# Тестовое задание на позицию Go-разработчика
Напишите REST-сервис, который обслуживает генерацию и выдачу уникальных ключей, а также их последующее погашение. Пример такого сервиса: подтверждение платежей банком по SMS.

# Функции сервиса
1. Выдача уникальных ключей. По обращению к сервису происходит выдача ключа клиенту
2. Погашение ключа. Помечает ключ как использованный. Повторно погасить ключ нельзя. Нельзя погасить ключ, если он не был предварительно выдан клиенту.
3. Проверка ключа. Возвращает информацию информацию о ключе: не выдан, выдан, погашен.
4. Информация о количестве оставшихся, не выданных ключах

# Требования к генерации ключей
- Ключи состоят из букв латинского языка в нижнем и верхнем регистре и цифр.
- Длина ключа - 4 символа
- Уникальность. Один и тот же ключ не должен быть выдан более 1 раза.

# Требования к реализации и передаче решения
- Можно использовать любое хранилище данных
- Время на решение тестового задания - 1-3 дня
- Результат присылать в виде ссылки на github (или аналоги)
