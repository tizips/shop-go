-- +goose Up
-- +goose StatementBegin
create table `hr_children`
(
    `id`              int unsigned not null auto_increment,
    `organization_id` varchar(64)  not null default '' comment '组织ID',
    `child_id`        varchar(64)           default null comment '子级ID',
    `created_at`      timestamp    not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp             default null,
    primary key (`id`),
    key (`organization_id`),
    key (`child_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '人资-组织关系表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `hr_children`;
-- +goose StatementEnd
