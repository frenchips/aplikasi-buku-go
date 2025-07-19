-- +migrate Up
-- +migrate StatementBegin
create table buku (
	id serial primary key,
	title varchar(64) not null,
	category_id int not null,
	description varchar(256) not null,
	image_url varchar(4000) not null,
	release_year int not null,
	price int not null,
	total_page int not null,
	thickness int not null,
	create_at timestamp not null,
	create_by varchar(64) not null,
	modified_at timestamp,
	modified_by varchar(64),
	foreign key (category_id) references category(id)
)

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS buku;
-- +migrate StatementEnd