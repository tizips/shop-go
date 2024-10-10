-- +goose Up
-- +goose StatementBegin
create table `shp_sku`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `product_id`      varchar(64)       not null,
    `code`            char(128)         not null default '' comment '代码',
    `price`           int unsigned      not null default 0 comment '价格',
    `origin_price`    int unsigned      not null default 0 comment '原价',
    `cost_price`      int unsigned      not null default 0 comment '成本',
    `stock`           int unsigned      not null default 0 comment '库存',
    `warn`            int unsigned      not null default 0 comment '警告',
    `picture`         varchar(255)      not null default '' comment '图片',
    `is_default`      tinyint unsigned  not null default 0 comment '是否默认：1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`product_id`),
    key (`code`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商店-SKU 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_sku`;
-- +goose StatementEnd
