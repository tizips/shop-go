-- +goose Up
-- +goose StatementBegin
create table `hr_organization_thumb`
(
    `id`              int unsigned not null auto_increment,
    `organization_id` varchar(64)  not null default '' comment '组织ID',
    `clique_id`       varchar(64)           default null comment '集团ID',
    `url`             varchar(255) not null default '' comment 'URL',
    `created_at`      timestamp    not null default CURRENT_TIMESTAMP,
    `updated_at`      timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp             default null,
    primary key (`id`),
    key (`organization_id`),
    key (`clique_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '人资-组织封面图表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `hr_organization_thumb`;
-- +goose StatementEnd
