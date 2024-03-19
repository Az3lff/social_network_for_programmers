CREATE TABLE if NOT EXISTS "users"
(
    user_id       uuid DEFAULT gen_random_uuid() primary key,
    login         varchar(255) unique not null,
    email         varchar(255) unique not null,
    hash_password varchar(255)        not null
)