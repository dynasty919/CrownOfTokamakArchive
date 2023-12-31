CREATE TABLE `tbl_file` (
    `id_mysql` INT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `Author` VARCHAR(255) COMMENT '作者',
    `Title` VARCHAR(255) COMMENT '标题',
    `Content` TEXT COMMENT '文章内容',
    `PostTime` DATETIME COMMENT '发布时间',
    `Counter` INT COMMENT '后端计数器',
    `Id` VARCHAR(255) COMMENT '标题的Sha1哈希',
    PRIMARY KEY (`id_mysql`),
    UNIQUE KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='帽子姐回答表';