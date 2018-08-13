CREATE TABLE `article_detail` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '文章ID',
  `title` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` TEXT COMMENT '文章内容',
  `browser_count` BIGINT(11) NOT NULL DEFAULT 0 COMMENT '浏览数',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `main_category` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '主类别ID',
  `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类别名称',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uuid`(`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `sub_category` (
  `id`  BIGINT(11)NOT NULL AUTO_INCREMENT COMMENT '主键自增ID',
  `uuid` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '副类别ID',
  `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类别名称',
  `create_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_uuid`(`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;