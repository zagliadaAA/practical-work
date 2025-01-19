-- +goose Up
-- +goose StatementBegin
create table clients(
    ID bigserial primary key,
    name varchar(120) not null,
    b_date timestamp not null,
    phone_number varchar(11) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table clients;
-- +goose StatementEnd
