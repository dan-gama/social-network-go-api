CREATE TABLE User (
    id          int auto_increment primary key,
    name        varchar(120) not null,
    email       varchar(120) not null unique,
    password    varchar(20) not null unique,
    createdAt   timestamp default current_timestamp()
) ENGINE=INNODB;