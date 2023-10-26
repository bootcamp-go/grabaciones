-- DDL
DROP DATABASE IF EXISTS `movies_test_db`;
CREATE DATABASE `movies_test_db`;

USE `movies_test_db`;

CREATE TABLE `movies` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `year` int(11) NOT NULL,
  `director` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;