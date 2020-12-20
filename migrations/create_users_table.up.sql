CREATE TABLE users (
  user_id integer unique,
  chat_id integer unique,
  first_name varchar,
  last_name varchar,
  username varchar not null,
  start_date timestamp not null
);
