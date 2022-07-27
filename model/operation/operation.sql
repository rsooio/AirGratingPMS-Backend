CREATE TABLE `operation` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `staffer_id` INT NOT NULL,
        -- INDEX `staffer_id_index`(`staffer_id`),
    `operation_name` VARCHAR(15) NOT NULL,
        -- INDEX `operation_name_index`(`operation_name`),
    `operation_table` VARCHAR(30) NOT NULL,
        -- INDEX `operation_table_index`(`operation_table`),
    `operation_id` INT NOT NULL,
        -- INDEX `operation_id_index`(`operation_id`),
    `value_before_operation` JSON NOT NULL,
    `value_after_operation` JSON NOT NULL,
    `operation_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;