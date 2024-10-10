-- +goose Up
-- +goose StatementBegin
create table shp_specification
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `product_id`      varchar(64)       not null comment '产品ID',
    `parent_id`       int unsigned      not null comment '父级ID',
    `name`            varchar(120)      not null default '' comment '名称',
    `order`           tinyint unsigned  not null default 0 comment '序号：正序',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`product_id`),
    key (`parent_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商店-产品规格表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists shp_specification;
-- +goose StatementEnd
