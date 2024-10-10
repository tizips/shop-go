-- +goose Up
-- +goose StatementBegin
create table `sys_user_bind_role`
(
    `id`              int unsigned      not null auto_increment,
    `platform`        smallint unsigned not null default 0 comment '层位',
    `organization_id` varchar(64)                default null comment '组织ID',
    `user_id`         varchar(64)       not null default '' comment '用户ID',
    `role_id`         int unsigned      not null default 0 comment '角色ID',
    `created_at`      timestamp         not null default CURRENT_TIMESTAMP,
    `deleted_at`      timestamp                  default null,
    primary key (`id`),
    key (`organization_id`),
    key (`user_id`),
    key (`role_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '系统-用户绑定角色表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sys_user_bind_role`;
-- +goose StatementEnd
