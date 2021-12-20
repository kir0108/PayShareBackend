create table if not exists purchases
(
    id          bigserial           not null,
    room_id     bigint              not null,
    p_name      varchar             not null,
    description jsonb  default '{}' not null,
    owner_id    bigint              not null,
    cost        bigint default 0    not null,
    constraint purchases_pk
        primary key (id),
    constraint purchases_rooms_id_fk
        foreign key (room_id) references rooms
            on delete cascade
);

alter table purchases
    owner to postgres;

create unique index if not exists purchases_id_uindex
    on purchases (id);
