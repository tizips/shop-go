-- +goose Up
-- +goose StatementBegin
create table `shp_shipping`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(64)       not null comment '名称',
    `money`           int unsigned      not null default 0 comment '费用',
    `query`           varchar(255)      not null default '' comment '查询地址',
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
  default collate = utf8mb4_unicode_ci comment ='商城-快递公司表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_shipping`;
-- +goose StatementEnd
