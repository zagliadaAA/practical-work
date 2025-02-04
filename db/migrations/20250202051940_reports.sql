-- +goose Up
-- +goose StatementBegin
create table reports(
                        ID bigserial primary key,
                        doctor_name varchar(120) not null,
                        diagnosis varchar(6) not null,
                        created_at timestamptz  not null,
                        updated_at timestamptz not null,
                        id_client integer references clients(id) unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table clients;
-- +goose StatementEnd