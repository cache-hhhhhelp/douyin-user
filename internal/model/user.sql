create table user
(
    user_id          bigint auto_increment,
    username         varchar(64)  not null,
    password         varchar(128) not null,
    primary key (user_id)
);