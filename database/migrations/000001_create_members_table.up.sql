BEGIN;

CREATE TABLE IF NOT EXISTS members (
    member_id char(36) NOT NULL comment 'uuid',

	# 确认本人唯一性，并不一定要手机号
    mobile varchar(11) NOT NULL DEFAULT '' comment '手机号',
#     email varchar(30) NOT NULL DEFAULT '' comment '邮箱',
    password varchar(255) NOT NULL DEFAULT '' comment '表单登录密码',

    source varchar(30) NOT NULL DEFAULT '' comment '注册来源自定义',

#     可选通过关注微信公众号建立唯一性
#     open_id varchar(60) NOT NULL DEFAULT '' comment '微信OpenId',
#     subscribe tinyint(1) UNSIGNED NOT NULL DEFAULT 0 comment '是否已关注,0-未关注,1-已关注',

    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime DEFAULT NULL,

    PRIMARY KEY (`member_id`),
    KEY `members_phone_index` (`phone`),
    KEY `members_source_index` (`source`),
    KEY `members_deleted_at_index` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

COMMIT;