-- +goose Up
-- +goose StatementBegin
create table `shp_payment_channel`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(120)      not null default '' comment '名称',
    `channel`         varchar(10)       not null default '' comment '支付渠道：paypal=贝宝',
    `key`             varchar(255)      not null default '' comment 'KEY',
    `secret`          varchar(255)      not null default '' comment '密钥',
    `is_debug`        tinyint unsigned  not null default 0 comment '是否调试：1=是；2=否',
    `ext`             text              not null comment '扩展信息',
    `order`           tinyint unsigned  not null default 0 comment '序号：正序',
    `is_enable`       tinyint unsigned  not null default 0 comment '是否启用：1=是；2=否；',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment ='商城-支付渠道表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_payment_channel`;
-- +goose StatementEnd
