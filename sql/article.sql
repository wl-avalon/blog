CREATE TABLE `article_detail` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '文章ID',
  `tag` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '文章标签',
  `content` TEXT COMMENT '文章内容',
  `browser_count` BIGINT(11) NOT NULL DEFAULT 0 COMMENT '浏览数',
  `create_time` varchar(40) NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `main_category` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '主类别ID',
  `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类别名称',
  `create_time` varchar(40) NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uuid`(`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sub_category` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '副类别ID',
  `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类别名称',
  `create_time` varchar(40) NOT NULL DEFAULT '' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uuid`(`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;