USE `testdb`;
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `content` text COLLATE utf8mb4_bin,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin