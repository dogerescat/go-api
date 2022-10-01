DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
	`id` INT AUTO_INCREMENT,
	`namea` CHAR,
	`email` CHAR,
	`password` CHAR,
	PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;;

DROP TABLE IF EXISTS `todos`;

CREATE TABLE IF NOT EXISTS `todos` (
	`id` INT AUTO_INCREMENT,
	`title` CHAR,
	`is_dane` BOOLEAN,
	`user_id` INT,
	PRIMARY KEY (`id`)
)DEFAULT CHARSET=utf8 COLLATE=utf8_bin;;
