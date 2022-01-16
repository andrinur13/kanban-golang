# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.21-MariaDB)
# Database: kanban-golang
# Generation Time: 2022-01-16 15:17:17 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table category
# ------------------------------------------------------------

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(256) NOT NULL DEFAULT '',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table task
# ------------------------------------------------------------

DROP TABLE IF EXISTS `task`;

CREATE TABLE `task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(256) NOT NULL DEFAULT '',
  `description` varchar(512) NOT NULL DEFAULT '',
  `status` tinyint(1) NOT NULL,
  `user_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `task` WRITE;
/*!40000 ALTER TABLE `task` DISABLE KEYS */;

INSERT INTO `task` (`id`, `title`, `description`, `status`, `user_id`, `category_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(2,'Nyuci Baju','Nyuci baju tiap minggu',0,4,1,'2022-01-16 17:48:04','2022-01-16 17:48:04','0000-00-00 00:00:00');

/*!40000 ALTER TABLE `task` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` varchar(256) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `password` varchar(512) NOT NULL DEFAULT '',
  `role` enum('admin','member') NOT NULL DEFAULT 'member',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `full_name`, `email`, `password`, `role`, `created_at`, `updated_at`, `deleted_at`)
VALUES
	(3,'Andri Hidayatulloh','andri@gmail.com','$2a$04$DXX492Zxekru0oTiw2KCi.VOydmAP5TyAEjGpx0Yptn3YKwW1S1Ri','member','2022-01-09 12:44:57','2022-01-09 15:43:24',NULL),
	(4,'Andri Nur H','andribis13@gmail.com','$2a$04$6J9lBHxvEuA1raIu6pnbXe34tmBaqDueVBorqcl3BQWacsuIt3Dcq','member','2022-01-09 13:14:56','2022-01-09 13:14:56',NULL),
	(5,'Andri Nur H','andribis11@gmail.com','$2a$04$4XBv/OSyKjP1P5i2hwyrJuXpt/vIEfTcCISwOm95QglWTUK6q15b.','member','2022-01-09 13:16:18','2022-01-09 13:16:18',NULL),
	(6,'Andri Nur H','andribis10@gmail.com','$2a$04$5YMXr4zd8BhnztGroAcMEOBY97bW3QQHkPKpYq3QHK0hgligyB3C2','member','2022-01-09 13:17:26','2022-01-09 13:17:26',NULL),
	(7,'Andri Nur H','andribis0@gmail.com','$2a$04$vRrOlsbwXLCYYKOxH.YnyOpT4Bys5hGohrsRPGlXmorOshttL78Nu','member','2022-01-09 13:17:41','2022-01-09 13:17:41',NULL),
	(8,'Andri Aja','andribis100@gmail.com','$2a$04$4BtUo1iIBv4rh2Jl4Y8TOelhHOirLuahO753/lRtFM6xLfjKW2db.','member','2022-01-09 13:17:50','2022-01-09 15:34:35',NULL);

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
