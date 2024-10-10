-- +goose Up
-- +goose StatementBegin
create table `shp_banner`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(120)      not null default '' comment '名称',
    `description`     varchar(1000)     not null default '' comment '描述',
    `button`          varchar(120)      not null default '' comment '按钮',
    `picture`         varchar(255)      not null default '' comment '图片',
    `target`          varchar(10)       not null default '' comment '打开方式：blank=新窗口；self=该窗口',
    `url`             varchar(255)      not null default '' comment '链接',
    `order`           tinyint unsigned  not null default 0 comment '序号：正序',
    `is_enable`       tinyint unsigned  not null default 0 comment '是否启用：1=是；2=否；',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-轮播表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_banner`;
-- +goose StatementEnd
