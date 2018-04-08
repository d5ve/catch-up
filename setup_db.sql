CREATE DATABASE IF NOT EXIST catchup;

USE catchup;

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
    `catchup_id` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `date` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `vote` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `option_id` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `voter` varchar(255) NOT NULL,
    `ynm` char(1) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

