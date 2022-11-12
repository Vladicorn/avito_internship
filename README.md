# avito_internship
##### Перед запуском изменить в ./pkg/DBconnect.go
```go
const (
	host     = "127.0.0.1"
	port     = 5432 // Default port
	user     = "postgres"
	password = "987123"
	dbname   = "postgres"
)
```
#### Для подключения к базе
 
##### Схема для создания базы юзеров
```sql
CREATE TABLE public.avito_users
(
    id serial,
    balance real,
    PRIMARY KEY (id)
);
```
##### Схема для создания базы хранения транзакций
```sql
CREATE TABLE avito_balance
(
    id serial,
    id_user integer,
    id_service integer,
    id_order integer,
    price real,
    start boolean,
    finish boolean,
    error boolean,
    time date,
    PRIMARY KEY (id),
    FOREIGN KEY (id_user)
        REFERENCES avito_users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);
```
##### Для запуска docker-compose up
##### Для просмотра API, IP:8000/swagger
## Релизованы следующие методы
:white_check_mark: Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.  
:white_check_mark: Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.  
:white_check_mark: Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.  
:white_check_mark: Метод получения баланса пользователя. Принимает id пользователя.  
:white_check_mark: Реализован сценарий разрезервирования денег, если услугу применить не удалось. Поле Valid.  
:white_check_mark: Бухгалтерия раз в месяц просит предоставить сводный отчет по всем пользователем, с указанием сумм выручки по каждой из предоставленной услуги для расчета и уплаты налогов.