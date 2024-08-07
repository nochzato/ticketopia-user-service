CREATE TABLE users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	full_name varchar(255) NOT NULL,
	username varchar(50) UNIQUE NOT NULL,
	password varchar(255) NOT NULL,
	email varchar(255) UNIQUE NOT NULL,
	created_at timestamptz NOT NULL DEFAULT now()
);
