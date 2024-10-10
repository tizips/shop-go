-- +goose Up
-- +goose StatementBegin
create table `shp_seo`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '平台',
    `clique_id`       varchar(64)                default null comment '集团ID',
    `organization_id` varchar(64)                default null comment '组织ID',
    `channel`         varchar(20)       not null default '' comment '渠道：product=产品；page=单页面；category=栏目',
    `channel_id`      varchar(64)       not null comment '渠道 ID',
    `title`           varchar(255)      not null default '' comment '标题',
    `keyword`         varchar(255)      not null default '' comment '关键词',
    `description`     varchar(255)      not null default '' comment '描述',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`clique_id`),
    key (`organization_id`),
    key (`channel_id`),
    key (`deleted_at`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='商城-SEO 表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `shp_seo`;
-- +goose StatementEnd
