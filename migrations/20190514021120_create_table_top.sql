-- +goose Up
-- +goose StatementBegin
create table top
(
    chat_id bigint not null,
    type varchar(5) default 'all',
    message text not null,
    constraint top_pk
        primary key (chat_id, type)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table top;
-- +goose StatementEnd
