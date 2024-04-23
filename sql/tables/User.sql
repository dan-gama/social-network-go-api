CREATE TABLE User (
    id          int auto_increment primary key,
    email       varchar(120) not null unique,
    name        varchar(120) not null,
    password    varchar(80) not null,
    createdAt   timestamp default current_timestamp()
) ENGINE=INNODB;