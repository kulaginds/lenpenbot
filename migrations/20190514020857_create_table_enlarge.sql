-- +goose Up
-- +goose StatementBegin
create table enlarge
(
    user_id int not null,
    chat_id bigint not null,
    length int default 1 not null,
    created TIMESTAMP default CURRENT_TIMESTAMP,
    constraint enlarge_pk
        primary key (user_id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table enlarge;
-- +goose StatementEnd
