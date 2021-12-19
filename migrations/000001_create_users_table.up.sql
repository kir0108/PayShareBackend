create table if not exists users
(
    id          bigserial not null,
    api_id      varchar   not null,
    api_name    varchar   not null,
    first_name  varchar,
    second_name varchar,
    image_url   varchar,
    constraint table_name_pk
    primary key (id)
    );

alter table users
    owner to postgres;

create unique index if not exists table_name_api_id_api_name_uindex
    on users (api_id, api_name);

create unique index if not exists table_name_id_uindex
    on users (id);
