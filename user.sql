create schema player collate utf8mb4_0900_ai_ci;
create table user
(
    id          int auto_increment,
    user_name   varchar(20)                  not null,
    pass_word   varchar(50) default '123456' not null,
    phone       varchar(20)                  null,
    sex         varchar(1)  default '0'      null,
    salt        varchar(50)                  null,
    create_time datetime                     null,
    update_time datetime                     null,
    qq          varchar(20)                  null,
    last_ip     varchar(50)                  null,
    token       varchar(50)                  null,
    status      varchar(3)  default '0'      null,
    constraint user_id_uindex
        unique (id),
    constraint user_user_name_uindex
        unique (user_name)
);

alter table user
    add primary key (id);