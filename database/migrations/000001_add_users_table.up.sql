CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `password` longtext NOT NULL, 
  `email` varchar(191) NOT NULL,
  `name` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`)
);
