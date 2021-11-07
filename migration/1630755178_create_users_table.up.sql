BEGIN;
-- SQLINES LICENSE FOR EVALUATION USE ONLY
DROP SEQUENCE IF EXISTS users_seq;
CREATE SEQUENCE users_seq;

CREATE TABLE users (
  id bigint check (id > 0) NOT NULL DEFAULT NEXTVAL ('users_seq'),
  username varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) DEFAULT NULL,
  created_at timestamptz(0) NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  CONSTRAINT users_username_unique UNIQUE  (username)
)  ;

-- SQLINES LICENSE FOR EVALUATION USE ONLY
DROP SEQUENCE IF EXISTS user_tokens_seq;
CREATE SEQUENCE user_tokens_seq;

CREATE TABLE user_tokens (
  id bigint check (id > 0) NOT NULL DEFAULT NEXTVAL ('user_tokens_seq'),
  user_id bigint check (user_id > 0) NULL,
  hash varchar(255) NOT NULL,
  expired_at timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
 ,
  CONSTRAINT user_tokens_user_id_foreign FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
)  ;

CREATE INDEX user_tokens_user_id_foreign ON user_tokens (user_id);

-- SQLINES DEMO *** PRODUCTION, hash value is "password"
INSERT INTO users values (DEFAULT, 'username', '$2a$12$C.rSdRpXicGcNknv37GQ9e49YXc3w4gJSb.48cyfjCkWYT4wOrd3O', 'User Name', DEFAULT, DEFAULT);
COMMIT;