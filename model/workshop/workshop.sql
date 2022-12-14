CREATE TABLE `workshop` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
        INDEX `enterprise_id_index`(`enterprise_id`),
    `name` VARCHAR(30) NOT NULL,
        UNIQUE KEY `name_index`(`enterprise_id`, `name`),
    `address` VARCHAR(100) NOT NULL DEFAULT "",
        -- INDEX `address_index`(`address`),
    `phone_number` VARCHAR(20) NOT NULL DEFAULT "",
        -- INDEX `phone_number_index`(`phone_number`),
    `manager_id` INT NOT NULL DEFAULT 0,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300) NOT NULL DEFAULT "",
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;