-- +migrate Up
-- +migrate StatementBegin
create table users(
	id serial primary key,
	username varchar(64) not null,
	password varchar(256) not null,
	created_at timestamp not null,
	created_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64)
)
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS gorp_migrations;
DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
