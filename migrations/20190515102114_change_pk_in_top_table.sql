-- +goose Up
-- +goose StatementBegin
alter table top drop constraint top_pk;

alter table top
    add constraint top_pk
        primary key (chat_id, type, created);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table top drop constraint top_pk;

alter table top
    add constraint top_pk
        primary key (chat_id, type);
-- +goose StatementEnd
