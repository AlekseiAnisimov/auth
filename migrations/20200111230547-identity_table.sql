-- +migrate Up
CREATE TABLE identity (
	id integer(11) AUTO_INCREMENT,
	login varchar(32) NOT NULL,
	email varchar(64) NOT NULL,
	phone varchar(32) NOT NULL,
	password varchar(40) NOT NULL,
	create_date datetime NOT NULL,
	update_date datetime,
	PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE identity;