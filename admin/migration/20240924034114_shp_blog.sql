-- +goose Up
-- +goose StatementBegin
create table `shp_blog`
(
    `id`              varchar(64)       not null,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `name`            varchar(255)      not null comment '名称',
    `thumb`           varchar(255)      not null comment '图片',
    `posted_at`       date              not null comment '发布日期',
    `is_top`          tinyint unsigned  not null default 0 comment '是否置顶：1=是；2=否',
    `summary`         varchar(1000)     not null default '' comment '简介',
    `content`         text                       default null comment '内容',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-博客表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_blog`;
-- +goose StatementEnd
