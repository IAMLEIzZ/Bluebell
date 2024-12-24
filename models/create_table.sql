DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `user_id` bigint(20) NOT NULL,
 `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
 `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
 `email` varchar(64) COLLATE utf8mb4_general_ci,
 `gender` tinyint(4) NOT NULL DEFAULT '0',
 `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
 `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`),
 UNIQUE KEY `idx_username` (`username`) USING BTREE,
 UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`
(
    `id`             int(11)                                 NOT NULL AUTO_INCREMENT,
    `community_id`   int(10) unsigned                        NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time`    bigint                            NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_community_id_delete_time` (`community_id`, `delete_time`),
    UNIQUE INDEX `idx_community_name_delete_time` (`community_name`, `delete_time`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    COMMENT = '社区表：存储社区信息';
INSERT INTO `community`
VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10', 0);
INSERT INTO `community`
VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00', 0);
INSERT INTO `community`
VALUES ('3', '3', 'PUBG', '大吉大利，今晚吃鸡。', '2018-08-07 08:30:00', '2018-08-07 08:30:00', 0);
INSERT INTO `community`
VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00',0);

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `id`           bigint(20)                               NOT NULL AUTO_INCREMENT COMMENT '自增主键，唯一标识每条帖子记录',
    `post_id`      bigint(20)                               NOT NULL COMMENT '帖子ID，用于业务中的帖子唯一标识',
    `title`        varchar(128) COLLATE utf8mb4_general_ci  NOT NULL COMMENT '帖子标题',
    `summary` varchar(120) COLLATE utf8mb4_general_ci NOT NULL  COMMENT '帖子摘要',
    `author_id`    bigint(20)                               NOT NULL COMMENT '作者的用户ID，用于关联用户表',
    `community_id` bigint(20)                               NOT NULL COMMENT '所属社区ID，用于关联社区表',
    `status`       tinyint(4)                               NOT NULL DEFAULT '1' COMMENT '帖子状态：1-正常，0-隐藏或删除',
    `create_time`  timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP COMMENT '帖子创建时间，默认当前时间',
    `update_time`  timestamp                                NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '帖子更新时间，每次更新时自动修改',
    `delete_time`  bigint                               NULL DEFAULT 0 COMMENT '逻辑删除时间，NULL表示未删除',

    PRIMARY KEY (`id`) COMMENT '主键索引',

    UNIQUE INDEX `idx_post_id_delete_time` (`post_id`, `delete_time`) COMMENT '联合索引：帖子ID和删除时间确保未删除的帖子ID唯一',

    INDEX `idx_author_id` (`author_id`) COMMENT '普通索引：按作者ID查询帖子',

    INDEX `idx_community_id` (`community_id`) COMMENT '普通索引：按社区ID查询帖子'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
    COMMENT = '帖子表：存储用户发布的帖子及其状态';