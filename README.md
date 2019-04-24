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

#### 1.1 Касса 

Касса (id (BIGSERIAL), Фильм (VARCHAR), Внутренние сборы (INTEGER), Международные сборы (INTEGER))

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

![table1](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/table1.png)

Добавление записи:

```
    INSERT INTO boxoffice(movie,domestic_sales,international_sales) VALUES($1,$2,$3) returning id
```

![modal1](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/modal1.png)

Обновление записи:

```
    update boxoffice set (movie, domestic_sales, international_sales)=($1,$2,$3) where id=$4
```

![modal2](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/modal2.png)

Удаление записи:

```
    delete from boxoffice where id=$1
```

![modal3](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/modal3.png)


#### 1.2 Директоры

Директоры (id (BIGSERIAL), ФИО (VARCHAR), Возраст (INTEGER), Пол (VARCHAR), Почта (VARCHAR))

Создание таблицы:

```
    CREATE TABLE IF NOT EXISTS directors(
            id BIGSERIAL PRIMARY KEY NOT NULL,
    		name VARCHAR(255) NOT NULL,
    		age INTEGER,
    		gender VARCHAR(255) NOT NULL,
    		email VARCHAR(255) NOT NULL
        );
```

Таблица:

![table2](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/table2.png)

Добавление, обновление и удаление реализованы аналогично 1.1

#### 1.3 Фильмы

Фильмы (id (BIGSERIAL), Название (VARCHAR), Директор (VARCHAR), Год (INTEGER), Длительность (INTEGER), Студия (VARCHAR))

Создание таблицы:

```
    CREATE TABLE IF NOT EXISTS movies(
            id BIGSERIAL PRIMARY KEY NOT NULL,
    		title VARCHAR(255) NOT NULL,
    		director VARCHAR(255) NOT NULL,
    		year INTEGER,
    		length INTEGER,
    		studio VARCHAR(255) NOT NULL
        );
```

Таблица:

![table3](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/table3.png)

Добавление, обновление и удаление реализованы аналогично 1.1

#### 1.4 Студии 

Студии (id (BIGSERIAL), Название (VARCHAR), Год основания (INTEGER), Общее количество фильмов (INTEGER))

Создание таблицы:

```
    CREATE TABLE IF NOT EXISTS studios(
            id BIGSERIAL PRIMARY KEY NOT NULL,
    		name VARCHAR(255) NOT NULL,
    		year INTEGER,
    		all_films INTEGER
        );
```

Таблица:

![table4](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/table4.png)

Добавление, обновление и удаление реализованы аналогично 1.1

### 2. Запросы

#### 2.1 Запрос 1

Описание:

Код:

```

```

Результат

![q1](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q1.png)

#### 2.2 Запрос 2

Описание:

Код:

```

```

Результат

![q2](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q2.png)

#### 2.3 Запрос 3

Описание:

Код:

```

```

Результат

![q3](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q3.png)

#### 2.4 Запрос 4

Описание:

Код:

```

```

Результат

![q4](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q4.png)

#### 2.5 Запрос 5

Описание:

Код:

```

```

Результат

![q5](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q5.png)

#### 2.6 Запрос 6

Описание:

Код:

```

```

Результат

![q6](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q6.png)

#### 2.7 Запрос 7

Описание:

Код:

```

```

Результат

![q7](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q7.png)

#### 2.8 Запрос 8

Описание:

Код:

```

```

Результат

![q8](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q8.png)

#### 2.9 Запрос 9

Описание:

Код:

```

```

Результат

![q9](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q9.png)

#### 2.10 Запрос 10

Описание:

Код:

```

```

Результат

![q10](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q10.png)

#### 2.11 Запрос 11

Описание:

Код:

```

```

Результат

![q11](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q11.png)

#### 2.12 Запрос 12

Описание:

Код:

```

```

Результат

![q12](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q12.png)

#### 2.13 Запрос 13

Описание:

Код:

```

```

Результат

![q13](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q13.png)

#### 2.14 Запрос 14

Описание:

Код:

```

```

Результат

![q14](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q14.png)

#### 2.15 Запрос 15

Описание:

Код:

```

```

Результат

![q15](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q15.png)
