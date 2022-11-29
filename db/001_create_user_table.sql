CREATE TYPE "gender" AS ENUM ('FEMALE', 'MALE');

CREATE TABLE public."user" (
	id bigserial NOT NULL,
	created timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	last_modified timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	active bool NOT NULL DEFAULT true,
	
	login varchar(255) NOT NULL,
	"password" bytea NOT NULL,
	full_name varchar(1500) NOT NULL,
	
	birthday date NOT NULL,
	gender public.gender NULL,
	
	phone varchar(16) NULL,
	email varchar(320) NULL,
	
	about varchar(1000) NULL,
	country char(2) NOT NULL,
	country_region varchar(6) NOT NULL,
	currency char(3) NOT NULL,
	language varchar(35) NOT NULL,
	timezone varchar(50) NOT NULL,
	
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX user_login_idx ON "user"(login) WHERE active;

---- create above / drop below ----

DROP INDEX user_login_idx;

DROP TABLE "user";

DROP TYPE "gender";
