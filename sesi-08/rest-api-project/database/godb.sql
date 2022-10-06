-- create new database
CREATE DATABASE `godb`

-- godb.people definition

CREATE TABLE `people` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_people_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- godb.people data
INSERT INTO godb.people (created_at,updated_at,deleted_at,first_name,last_name) VALUES
	 ('2022-10-06 13:07:57','2022-10-06 13:07:57',NULL,'Husniyah','Lisan'),
	 ('2022-10-06 13:08:20','2022-10-06 13:10:44',NULL,'Carmellia','Maxwell'),
	 ('2022-10-06 13:08:46','2022-10-06 13:08:46','2022-10-06 13:11:20','Parath','Handson');