DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
	`id` INT AUTO_INCREMENT,
	`name` CHAR(20),
	`email` CHAR(20),
	`password` CHAR(20),
	PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;;

DROP TABLE IF EXISTS `todos`;

CREATE TABLE IF NOT EXISTS `todos` (
	`id` INT AUTO_INCREMENT,
	`title` CHAR(20),
	`is_dane` BOOLEAN,
	`user_id` INT,
	PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;;
