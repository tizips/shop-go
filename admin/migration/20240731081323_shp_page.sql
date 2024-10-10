-- +goose Up
-- +goose StatementBegin
create table `shp_page`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `code`            varchar(64)       not null comment 'CODE',
    `name`            varchar(255)      not null comment '名称',
    `is_system`       tinyint unsigned           default 0 comment '是否系统内置：1=是；2=否',
    `content`         text                       default null comment '内容',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-页面表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_page`;
-- +goose StatementEnd
