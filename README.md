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

Вывести год выпуска и студию фильмов, чья общая выручка превосходит указанное значение

Код:

```
    SELECT title,year,studio,(domestic_sales+international_sales) "+
    		"FROM movies JOIN boxoffice "+
    		"ON movies.title = boxoffice.movie "+
    		"WHERE domestic_sales+international_sales > $1
```

Результат

![q1](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q1.png)

#### 2.2 Запрос 2

Описание:

Показать директоров, у которых более одного фильма в таблице Фильмы

Код:

```
    SELECT director, count(*) 
    from movies 
    group by director 
    having count(*)>1

```

Результат

![q2](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q2.png)

#### 2.3 Запрос 3

Описание:

Показать общую прибыль студий от фильмов из таблицы Фильмы. Сортировать по убыванию

Код:

```
    SELECT studio, sum(domestic_sales+international_sales) 
    from movies 
    left join boxoffice 
    on movies.title = boxoffice.movie 
    group by movies.studio 
    order by  sum(domestic_sales+international_sales) DESC

```

Результат

![q3](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q3.png)

#### 2.4 Запрос 4

Описание:

Показать директоров и их возраст, работающих со студией Universal Studios. Сортировать по старшинству

Код:

```
    SELECT DISTINCT director, age from movies 
    inner join directors 
    on movies.director = directors.name and movies.studio = 'Universal Studios' 
    order by age DESC

```

Результат

![q4](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q4.png)

#### 2.5 Запрос 5

Описание:

Показать директоров-мужчин, снявших менее 2х фильмов и пользующихся почтой gmail

Код:

```
    select director, cnt, gender, email 
    from (
        select director, count(*) as cnt 
        from movies 
        group by director 
        having count(*)<2) 
    movies 
    join directors 
    on movies.director = directors.name and gender = 'муж.' and email LIKE '%@gmail.com
```

Результат

![q5](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q5.png)

#### 2.6 Запрос 6

Описание:

Показать самую молодую студию, снявшую более одного фильма из таблицы Фильмы

Код:

```
    select studio, year, cnt 
    from (
        select studio, count(*) as cnt 
        from movies 
        group by studio 
        having count(*)>1) 
    movies join studios 
    on movies.studio = studios.name 
    where year = (
        SELECT MAX(year) 
        FROM (
            select studio, count(*) as cnt 
            from movies 
            group by studio 
            having count(*)>1) 
        movies join studios on movies.studio = studios.name)
```

Результат

![q6](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q6.png)

#### 2.7 Запрос 7

Описание:

Показать процент фильмов студий, показанных в таблице, от общего числа фильмов студий для каждой студии

Код:

```
    select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  
    from (
        select studio, count(*) as cnt 
        from movies group by studio) 
    movies join studios 
    on movies.studio = studios.name

```

Результат

![q7](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q7.png)

#### 2.8 Запрос 8

Описание:

Показать общую длительность фильмов, снятых студиями, открытыми после указанного года

Код:

```
    select studio, len  
    from (
        select studio, sum(length) as len 
        from movies group by studio) 
    movies join studios 
    on movies.studio = studios.name 
    where studios.year > $1

```

Результат

![q8](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q8.png)

#### 2.9 Запрос 9

Описание:

Показать средний год выпуска фильмов, длительность которых выше указанной

Код:

```
    select cast(avg(year) as char(4)) 
    from movies 
    where length > $1

```

Результат

![q9](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q9.png)

#### 2.10 Запрос 10

Описание:

Выбрать возраст второго по старшинству директора

Код:

```
    select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  
    from (
        select studio, count(*) as cnt 
        from movies group by studio) 
    movies join studios 
    on movies.studio = studios.name

```

Результат

![q10](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q10.png)

#### 2.11 Запрос 11

Описание:

Показать самый дорогой фильм Кристофера Нолана

Код:

```
    select movie, (domestic_sales+international_sales)  
    from (boxoffice 
        inner join movies 
        on boxoffice.movie = movies.title 
        and movies.director='Кристофер Нолан' 
        and (domestic_sales+international_sales) = (
            select max(domestic_sales+international_sales)  
            from (boxoffice 
                inner join movies 
                on boxoffice.movie = movies.title 
                and movies.director='Кристофер Нолан'
            )
        )
    )
```

Результат

![q11](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q11.png)

#### 2.12 Запрос 12

Описание:

Показать студии, которые имеют фильмы длиной более 120 минут. Показать количество таких фильмов у каждой студии

Код:

```
    select studio, count(*) 
    from movies 
    where length > 120 
    group by studio
```

Результат

![q12](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q12.png)

#### 2.13 Запрос 13

Описание:

Посчитать процентное соотношение студий в таблице Фильмы

Код:

```
    select studio, cast ((cast(cnt as float) / (select count(*) from movies) * 100) as char(4))  
    from (
        select studio, count(*) as cnt 
        from movies 
        group by studio) movies
```

Результат

![q13](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q13.png)

#### 2.14 Запрос 14

Описание:

Посчитать количество вхождений буквы "е" в столбце Название таблицы Студии

Код:

```
     SELECT SUM(ROUND ( LENGTH(name)- LENGTH( REPLACE ( name, 'e', '') )) ) 
     FROM studios 

```

Результат

![q14](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q14.png)

#### 2.15 Запрос 15

Описание:

Код:

```
    select title, year, length 
    from movies 
    where year between 1900 and 1999 and length between 100 and 125
```

Результат

![q15](https://github.com/BykovIlya/Basics_of_storage_technology/blob/master/images/q15.png)
