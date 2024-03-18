CREATE TABLE `products` (
  `id` varchar(36) NOT NULL,
  `name` varchar(191) NOT NULL,
  `stock` bigint DEFAULT '0',
  `images` json NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_products` (`user_id`),
  CONSTRAINT `fk_users_products` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)