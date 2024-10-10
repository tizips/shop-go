-- +goose Up
-- +goose StatementBegin
create table `shp_category`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `level`           varchar(4)                 default null comment '分类层级：lv_1=一级分类；lv_2=二级分类；lv_3=三级分类',
    `parent_id`       int unsigned      not null default 0 comment '父级ID',
    `name`            varchar(64)       not null default '' comment '名称',
    `order`           tinyint unsigned  not null default 0 comment '序号：正序',
    `is_enable`       tinyint unsigned  not null default 0 comment '启用：1=是；2=否',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`organization_id`),
    key (`parent_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '商城-栏目表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_category`;
-- +goose StatementEnd
