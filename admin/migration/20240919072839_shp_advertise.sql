-- +goose Up
-- +goose StatementBegin
create table `shp_advertise`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `page`            varchar(20)       not null default '' comment '页面：home=首页',
    `block`           varchar(64)       not null default '' comment '广告位：new_product=新品',
    `title`           varchar(255)      not null default '' comment '标题',
    `target`          varchar(10)       not null default '' comment '打开方式：blank=新窗口；self=该窗口',
    `url`             varchar(255)      not null default '' comment '链接',
    `thumb`           varchar(255)      not null default '' comment '图片',
    `order`           tinyint unsigned  not null default 0 comment '排序（正序）',
    `is_enable`       tinyint unsigned  not null default 0 comment '是否启用；1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-广告表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_advertise`;
-- +goose StatementEnd
