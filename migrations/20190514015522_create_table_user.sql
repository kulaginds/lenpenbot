-- +goose Up
-- +goose StatementBegin
create table "user"
(
    user_id int not null,
    chat_id bigint not null,
    constraint user_chat_pk
        primary key (user_id, chat_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table "user";
-- +goose StatementEnd
