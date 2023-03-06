-- CREATE DATABASE `sandbox`;

USE `sandbox`;

CREATE TABLE `codes` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `code` int DEFAULT NULL,
                         `user_id` int DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `expire_at` timestamp NULL DEFAULT NULL,
                         `action` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `data` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `is_used` tinyint(1) NOT NULL DEFAULT '0',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `comments` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `post_id` int DEFAULT NULL,
                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2051 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `posts` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int DEFAULT NULL,
                         `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `body` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `password` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `role` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `is_active` tinyint(1) DEFAULT '0',
                         `google_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;