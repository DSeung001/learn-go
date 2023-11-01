USE short_url;

DROP TABLE IF EXISTS url;

CREATE TABLE url
(
    id        INT AUTO_INCREMENT NOT NULL,
    alias_url VARCHAR(128)       NOT NULL,
    full_url  VARCHAR(128)       NOT NULL,
    PRIMARY KEY (`id`)
);