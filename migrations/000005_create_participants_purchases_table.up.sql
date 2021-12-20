create table if not exists participants_purchases
(
    id             bigserial             not null,
    purchase_id    bigint                not null,
    participant_id bigint                not null,
    paid           boolean default false not null,
    constraint participants_purchases_pk
    primary key (id),
    constraint participants_purchases_participants_id_fk
    foreign key (participant_id) references participants
    on delete cascade,
    constraint participants_purchases_purchases_id_fk
    foreign key (purchase_id) references purchases
    on delete cascade
    );

alter table participants_purchases
    owner to postgres;

create unique index if not exists participants_purchases_id_uindex
    on participants_purchases (id);
