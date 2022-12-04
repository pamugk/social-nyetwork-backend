CREATE TABLE public.message(
    id UUID NOT NULL,
	created timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	last_modified timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	exist bool NOT NULL DEFAULT true,

	sender bigint NOT NULL,
	recipient bigint NOT NULL,
	text VARCHAR(1000) NOT NULL,
	original_id UUID NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (sender) REFERENCES public."user"(id),
    FOREIGN KEY (recipient) REFERENCES public."user"(id)
);

CREATE UNIQUE INDEX message_users_idx ON public.message(sender, recipient);

---- create above / drop below ----

DROP INDEX message_users_idx;

DROP TABLE public.message;
