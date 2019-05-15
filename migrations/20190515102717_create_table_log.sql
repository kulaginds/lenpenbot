-- +goose Up
-- +goose StatementBegin
create table log
(
    id bigserial not null
        constraint log_pk
            primary key,
    text text not null,
    created TIMESTAMP default CURRENT_TIMESTAMP
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
drop table log;
-- +goose StatementEnd
