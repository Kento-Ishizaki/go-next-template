-- CreateTable
CREATE TABLE `todos` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `text` varchar(256) NOT NULL,
    `done` bool NOT NULL,
    `user_id` int(11) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `deleted_at` DATETIME(3) NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
