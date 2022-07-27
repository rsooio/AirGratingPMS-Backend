CREATE TABLE `technology` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `name` VARCHAR(30) NOT NULL,
        UNIQUE KEY `name_unique_index`(`enterprise_id`, `workshop_id`, `name`),
    `info` VARCHAR(300),
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300),
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;