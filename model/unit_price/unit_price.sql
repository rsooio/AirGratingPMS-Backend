CREATE TABLE `unit_price` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `technology_id` INT NOT NULL,
        INDEX `technology_id_index`(`technology_id`),
    `client_id` INT NOT NULL,
        INDEX `client_id_index`(`client_id`),
        UNIQUE KEY `unit_price_unique_key`(`enterprise_id`, `workshop_id`, `client_id`, `technology_id`),
    `unit_price` DECIMAL(6,2) NOT NULL,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300) NOT NULL DEFAULT "",
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;