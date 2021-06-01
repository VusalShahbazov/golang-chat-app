CREATE TABLE messages
(
    id           BIGINT(11) UNSIGNED AUTO_INCREMENT PRIMARY KEY NOT NULL,
    body         VARCHAR(255) not null,
    from_user_id BIGINT(11) UNSIGNED,
    room_id      BIGINT(11) UNSIGNED,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL
);

ALTER TABLE messages
    ADD CONSTRAINT fk_from_user_id
        FOREIGN KEY (from_user_id) REFERENCES users (id);

ALTER TABLE messages
    ADD CONSTRAINT f_message_room_id
        FOREIGN KEY (room_id) REFERENCES rooms (id);