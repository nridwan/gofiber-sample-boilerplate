BEGIN;
DROP TABLE IF EXISTS user_tokens;
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS user_tokens_seq;
DROP SEQUENCE IF EXISTS users_seq;
COMMIT;