CREATE TABLE `products` (
  `id` varchar(36) NOT NULL,
  `name` varchar(191) NOT NULL,
  `stock` bigint NOT NULL DEFAULT '0',
  `price` bigint NOT NULL,
  `images` json NOT NULL,
  `isPublished` tinyint(1) DEFAULT '0' COMMENT '0 = false 1 = true',
  `user_id` varchar(36) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_users_products` (`user_id`),
  CONSTRAINT `fk_users_products` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
)