# Основы технологий хранения данных

### Проект. Быков Илья, 6383

### 0. Установка и запуск

Для запуска требуется:
```
  go >= 1.10.4
  vue >= 3.55
  pgAdmin >= 4.0
```

Запуск

```
    server: go build && ./backend
    client: npm run serve
```
### 1. Таблицы

Имеем 4 таблицы 

1. Касса (id (BIGSERIAL), Фильм (VARCHAR), Внутренние сборы (INTEGER), Международные сборы (INTEGER))

Создание таблицы:

```
    CREATE TABLE IF NOT EXISTS boxoffice(
            id BIGSERIAL PRIMARY KEY NOT NULL,
    		movie VARCHAR(255) NOT NULL,
    		domestic_sales INTEGER,
    		international_sales INTEGER
        );
```

Таблица:

![table1](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q1.png)
