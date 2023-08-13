CREATE DATABASE IF NOT EXISTS `demo`;
USE `demo`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `created_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `deleted_at` timestamp           DEFAULT NULL COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'Jerry', '', 1, '2022-04-01 10:00:00', '2022-04-01 10:00:00'),
       (2, 'Tom', '', 2, '2022-04-01 10:00:00', '2022-04-01 10:00:00');

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`       varchar(128)        NOT NULL default '' COMMENT '标题',
    `content`     varchar(255)        NOT NULL COMMENT '评论内容',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';