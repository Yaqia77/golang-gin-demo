CREATE TABLE user(
    id bigint  AUTO_INCREMENT,
    username varchar(36) NOT NULL,
    password varchar(128) DEFAULT '',
    UNIQUE name_index (username),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci;