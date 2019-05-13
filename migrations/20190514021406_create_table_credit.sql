-- +goose Up
-- +goose StatementBegin
create table credit
(
    user_id int not null,
    chat_id bigint not null,
    length int default 1 not null,
    percents int default 1 not null,
    repair bool default false not null,
    created timestamp default current_timestamp,
    constraint credit_pk
        primary key (user_id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table credit;
-- +goose StatementEnd
