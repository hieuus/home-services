start transaction;

create table "user"
(
    id              bigserial primary key,
    code            varchar(32)               not null,
    full_name       varchar(128)              not null,
    email           varchar(128)              not null,
    phone_number    varchar(11)               not null,
    password        text                      not null,
    salt            text                      not null,
    is_verify_phone boolean     default false not null,
    is_verify_email boolean     default false not null,
    created_at      timestamptz default now() not null,
    updated_at      timestamptz default now() not null
);

alter table "user"
    add constraint user_code_uindex unique (code);



commit;