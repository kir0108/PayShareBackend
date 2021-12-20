create table if not exists participants
(
    id      bigserial not null,
    room_id bigint    not null,
    user_id bigint    not null,
    constraint participants_pk
    primary key (id),
    constraint participants_rooms_id_fk
    foreign key (room_id) references rooms
    on delete cascade,
    constraint participants_users_id_fk
    foreign key (user_id) references users
    on delete cascade
    );

alter table participants
    owner to postgres;

create unique index if not exists participants_id_uindex
    on participants (id);

create unique index if not exists participants_user_id_room_id_uindex
    on participants (user_id, room_id);
