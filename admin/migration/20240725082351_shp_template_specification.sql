-- +goose Up
-- +goose StatementBegin
create table `shp_template_specification`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(32)       not null default '' comment '名称',
    `label`           varchar(32)       not null default '' comment '标签',
    `options`         text                       default null comment '选项',
    `is_enable`       tinyint unsigned  not null default 0 comment '是否启用：1=是；2=否；',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-规格模板表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_template_specification`;
-- +goose StatementEnd
