-- +goose Up
-- +goose StatementBegin
create table `shp_order_address`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `order_id`        varchar(64)       not null comment '订单ID',
    `user_id`         varchar(64)       not null comment '用户id',
    `first_name`      varchar(64)       not null default '' comment 'first name',
    `last_name`       varchar(64)       not null default '' comment 'last name',
    `company`         varchar(255)      not null default '' comment '公司',
    `country`         varchar(255)      not null default '' comment '国家',
    `prefecture`      varchar(255)      not null default '' comment '州府',
    `city`            varchar(255)      not null default '' comment '城市',
    `street`          varchar(255)      not null default '' comment '街道',
    `detail`          varchar(255)      not null default '' comment '详细地址',
    `postcode`        varchar(64)       not null default '' comment '邮编',
    `phone`           varchar(64)       not null default '' comment '电话',
    `email`           varchar(120)      not null default '' comment '邮箱',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`order_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-订单收货地址表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_order_address`;
-- +goose StatementEnd
