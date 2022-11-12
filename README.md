# avito_internship
 
 

Настройки базы
	host     = "127.0.0.1"
	port     = 5432 // Default port
	user     = "postgres"
	password = "qwerty"
	dbname   = "postgres"
Схема для создания базы юзеров
CREATE TABLE public.avito_users
(
    id serial,
    balance real,
    PRIMARY KEY (id)
);
Схема для создания базы хранения транзакций
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