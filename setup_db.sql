START TRANSACTION;

DROP DATABASE catch_up;

CREATE DATABASE catch_up;

USE catch_up;

CREATE TABLE `catch_up` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `name` varchar(255) NOT NULL,
    `details` varchar(255) NOT NULL,
    `start_date` datetime NOT NULL,
    `end_date` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `option` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `catch_up_id` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `date` datetime NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`catch_up_id`) REFERENCES catch_up (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- TODO: foreign key constraints.
CREATE TABLE `vote` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `option_id` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `voter` varchar(255) NOT NULL,
    `ynm` char(1) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`option_id`) REFERENCES `option` (id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE USER 'catch_up'@'%' IDENTIFIED BY 'catch_up_password';
GRANT SELECT, INSERT, UPDATE, DELETE, CREATE TEMPORARY TABLES ON catch_up.* TO 'catch_up'@'%';

COMMIT;
