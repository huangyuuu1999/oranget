REATE TABLE `sites` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
	`sitename` VARCHAR(32) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',
	`uri` VARCHAR(128) NULL DEFAULT NULL COMMENT 'url' COLLATE 'utf8mb4_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
	`visit_num` INT(10) UNSIGNED NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `sitename` (`sitename`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=14
;
