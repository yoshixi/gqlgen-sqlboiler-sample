-- migrate:up
CREATE TABLE users (
    id uuid PRIMARY KEY default gen_random_uuid() NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL
);

-- migrate:down
DROP TABLE users;


