CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS task;
CREATE TABLE IF NOT EXISTS task
(
    id_task        INT(11)    NOT NULL AUTO_INCREMENT,
    description    TEXT(1000) NOT NULL,
    status         bool       NOT NULL DEFAULT FALSE,
    created_date   DATETIME  NOT NULL  DEFAULT CURRENT_TIMESTAMP,
    completed_date DATETIME,

    PRIMARY KEY (id_task)
    )
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARSET = utf8;