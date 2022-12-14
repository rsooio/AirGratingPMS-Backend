CREATE TABLE `product` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `product_set_id` INT NOT NULL,
        INDEX `product_set_id_index`(`product_set_id`),
    `technology_id` INT NOT NULL,
    `length` VARCHAR(10) NOT NULL,
    `width` VARCHAR(10) NOT NULL,
    `unit_price` VARCHAR(10) NOT NULL,
    `quantity` INT NOT NULL,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(200) NOT NULL DEFAULT "",
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;