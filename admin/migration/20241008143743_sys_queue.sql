-- +goose Up
-- +goose StatementBegin
create table `sys_queue`
(
    `id`         int unsigned     not null auto_increment,
    `queue`      varchar(64)      not null default '' comment '队列',
    `message`    text                      default null comment '消息',
    `error`      text                      default null comment '错误',
    `is_omitted` tinyint unsigned not null default 0 comment '是否忽略：1=是；2=否',
    `is_tried`   tinyint unsigned not null default 0 comment '是否重试：1=是；2=否',
    `is_handled` tinyint unsigned not null default 0 comment '是否处理：1=是；2=否',
    `created_at` timestamp        not null default CURRENT_TIMESTAMP,
    `updated_at` timestamp        not null default CURRENT_TIMESTAMP,
    `deleted_at` timestamp                 default null,
    primary key (`id`)
) auto_increment = 1000
  default collate = utf8mb4_unicode_ci comment ='系统-失败队列表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `sys_queue`;
-- +goose StatementEnd
