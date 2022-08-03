CREATE TABLE `enterprise` (
    `id` INT NOT NULL AUTO_INCREMENT,
        PRIMARY KEY (`id`),
    `name` VARCHAR(30) NOT NULL,
        UNIQUE KEY `name_unique_index`(`name`),
    `address` VARCHAR(100) NOT NULL DEFAULT "",
        -- INDEX `addrass_index`(`address`),
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300) NOT NULL DEFAULT "",
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `deleted_enterprise` (
    `id` INT NOT NULL,
        PRIMARY KEY (`id`),
    `name` VARCHAR(30) NOT NULL,
        UNIQUE KEY `name_unique_index`(`name`),
    `address` VARCHAR(100) NOT NULL DEFAULT "",
        -- INDEX `addrass_index`(`address`),
    `insert_time` TIMESTAMP NULL,
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(300) NOT NULL DEFAULT "",
        -- FULLTEXT INDEX `remark_fulltext_index`(`remark`),
    `version` INT NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;