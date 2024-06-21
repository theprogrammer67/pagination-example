CREATE TABLE IF NOT EXISTS client_accounts (
	id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
	name TEXT NOT NULL DEFAULT '',
	login TEXT NOT NULL DEFAULT '',
	password TEXT NOT NULL DEFAULT '',
	email TEXT NOT NULL DEFAULT '',
	role TEXT NOT NULL DEFAULT 'user'
);

CREATE OR REPLACE FUNCTION get_account_list(
	IN a_limit INT,
	IN a_offset INT
)
RETURNS TABLE (
  list client_accounts,
  full_count BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    RETURN QUERY
		SELECT a, count(*) OVER() AS full_count FROM client_accounts AS a
		ORDER BY a.id  ASC
		LIMIT a_limit
		OFFSET a_offset;			
END;
$$;

INSERT INTO client_accounts (name, login, password, email, role)
VALUES
    ('John', 'john', "123", 'john.doe@example.com', 'user'),
    ('Jane', 'jane', "123", 'jane.smith@example.com', 'user'),
    ('Bob', 'bob', "123", 'bob.johnson@example.com', 'admin');