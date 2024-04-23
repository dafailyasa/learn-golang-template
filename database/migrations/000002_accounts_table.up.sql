CREATE TABLE `accounts` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) DEFAULT NULL,
  `balance` bigint DEFAULT NULL,
  `currency` varchar(50) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_accounts_balance_currency` (`user_id`,`currency`),
  KEY `idx_accounts_user_id` (`user_id`)
)