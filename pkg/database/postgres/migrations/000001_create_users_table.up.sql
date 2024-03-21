CREATE TABLE if NOT EXISTS "users"
(
    user_id       uuid DEFAULT gen_random_uuid() primary key,
    login         varchar(255) unique not null,
    email         varchar(255) unique not null,
    hash_password varchar(255)        not null
);

CREATE TABLE if NOT EXISTS "chats"
(
    chat_id uuid DEFAULT gen_random_uuid() primary key
);

CREATE TABLE if NOT EXISTS "chat_users"
(
    chat_id uuid,
    user_id uuid,
    CONSTRAINT fk_chat_id
        FOREIGN KEY(chat_id)
        REFERENCES chats(chat_id),
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
        REFERENCES users(user_id)
);

CREATE TABLE if NOT EXISTS "messages"
(
    message_id SERIAL PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    chat_id uuid,
    CONSTRAINT fk_chat_id
        FOREIGN KEY(chat_id)
        REFERENCES chats(chat_id)
);
