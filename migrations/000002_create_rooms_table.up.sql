create table if not exists rooms
(
    id        bigserial             not null,
    room_name varchar               not null,
    room_date varchar               not null,
    owner_id  bigint                not null,
    close     boolean default false not null,
    constraint rooms_pk
    primary key (id),
    constraint rooms_users_id_fk
    foreign key (owner_id) references users
    on delete cascade
    );

alter table rooms
    owner to postgres;

create unique index if not exists rooms_id_uindex
    on rooms (id);
