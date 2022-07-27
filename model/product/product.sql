CREATE TABLE `product` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `enterprise_id` INT NOT NULL,
    `workshop_id` INT NOT NULL,
        INDEX `workshop_id_index`(`enterprise_id`, `workshop_id`),
    `product_set_id` INT NOT NULL,
        -- INDEX `product_set_id_index`(`product_set_id`),
    `technology_id` INT NOT NULL,
    `length` DECIMAL(6,2) NOT NULL,
    `width` DECIMAL(6,2) NOT NULL,
    `unit_price` DECIMAL(6,2) NOT NULL,
    `quantity` INT NOT NULL,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(200),
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;