-- +goose Up
-- +goose StatementBegin
create table `shp_cart`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null comment '用户ID',
    `product_id`      varchar(64)       not null comment '产品ID',
    `sku_id`          varchar(64)       not null comment 'SkuID',
    `code`            char(128)         not null default '' comment '代码',
    `specifications`  text              not null comment '规格',
    `name`            varchar(255)      not null default '' comment '名称',
    `picture`         varchar(255)      not null default '' comment '图片',
    `price`           int unsigned      not null default 0 comment '价格',
    `quantity`        int unsigned      not null default 0 comment '数量',
    `is_invalid`      tinyint unsigned           default 0 comment '是否失效：1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`user_id`),
    key (`product_id`),
    key (`sku_id`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-购物车表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_cart`;
-- +goose StatementEnd
