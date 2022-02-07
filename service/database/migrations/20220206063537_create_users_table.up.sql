CREATE TABLE users
(
    id           BIGINT AUTO_INCREMENT,
    username     VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL,
    password     VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    bio          text         NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP    NULL,
    PRIMARY KEY (id),
    UNIQUE KEY users_email_uindex (email),
    UNIQUE KEY users_username_uindex (username)
);