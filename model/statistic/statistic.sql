CREATE TABLE `statistic` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `time` TIMESTAMP NOT NULL,
        -- INDEX `time_index`(`time`),
    `version` INT NOT NULL,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;