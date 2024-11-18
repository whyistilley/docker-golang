-- get_users()
CREATE FUNCTION get_users()
RETURNS SETOF users
AS
$$
    SELECT * FROM users;
$$
LANGUAGE SQL IMMUTABLE STRICT;

-- get_user_by_id()
CREATE FUNCTION get_user_by_id(IN a UUID)
RETURNS users
AS
$$
    SELECT * FROM users WHERE id = $1 LIMIT 1;
$$
LANGUAGE SQL IMMUTABLE STRICT;

-- get_user_by_email()
CREATE FUNCTION get_user_by_email(IN email TEXT)
RETURNS users
AS
$$
    SELECT * FROM users WHERE email = $1 LIMIT 1;
$$
LANGUAGE SQL IMMUTABLE STRICT;

-- insert_user()
CREATE FUNCTION insert_user(IN name TEXT, IN email TEXT)
RETURNS users
AS
$$
    INSERT INTO users (name, email) VALUES ($1, $2) RETURNING *;
$$
LANGUAGE SQL STRICT;

-- delete_user_by_id()
CREATE FUNCTION delete_user_by_id(IN id UUID)
RETURNS users
AS
$$
    DELETE FROM users WHERE id = $1 RETURNING *;
$$
LANGUAGE SQL STRICT;

-- delete_user_by_email()
CREATE FUNCTION delete_user_by_email(IN email TEXT)
RETURNS users
AS
$$
    DELETE FROM users WHERE email = $1 RETURNING *;
$$
LANGUAGE SQL STRICT;

-- update_user_by_id()
CREATE FUNCTION update_user_by_id(IN id UUID, IN name TEXT)
RETURNS users
AS
$$
    UPDATE users SET name = $2 WHERE id = $1 RETURNING *;
$$
LANGUAGE SQL STRICT;

-- update_user_by_email()
CREATE FUNCTION update_user_by_email(IN email TEXT, IN name TEXT)
RETURNS users
AS
$$
    UPDATE users SET name = $2 WHERE email = $1 RETURNING *;
$$
LANGUAGE SQL STRICT;