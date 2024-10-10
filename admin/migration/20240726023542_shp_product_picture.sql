-- +goose Up
-- +goose StatementBegin
create table `shp_product_picture`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `product_id`      varchar(64)       not null comment '产品ID',
    `url`             varchar(255)      not null default '' comment '链接',
    `order`           tinyint unsigned  not null default 0 comment '序号：正序',
    `is_default`      tinyint unsigned  not null default 0 comment '是否默认：1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`product_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-产品图片表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_product_picture`;
-- +goose StatementEnd
