-- +goose Up
-- +goose StatementBegin
create table `hr_organization`
(
    `id`          varchar(64)       not null,
    `platform`    smallint unsigned not null      default 0 comment '平台',
    `clique_id`   varchar(64)                     default null comment '集团ID',
    `brand_id`    int unsigned      not null      default 0 comment '品牌ID',
    `parent_id`   varchar(64)                     default null comment '父级ID',
    `name`        varchar(64)       not null      default '' comment '名称',
    `valid_start` date              not null comment '有效期：开始',
    `valid_end`   date              not null comment '有效期：结束',
    `user`        varchar(32)       not null null default '' comment '联系人',
    `telephone`   varchar(32)       not null      default '' comment '联系电话',
    `province_id` int unsigned      not null      default 0 comment '省ID',
    `city_id`     int unsigned      not null      default 0 comment '市ID',
    `area_id`     int unsigned      not null      default 0 comment '区ID',
    `address`     varchar(255)      not null      default '' comment '详细地址',
    `longitude`   double(9, 6)      not null      default 0 comment '坐标：经度',
    `latitude`    double(8, 6)      not null      default 0 comment '坐标：纬度',
    `description` varchar(255)      not null      default '' comment '描述',
    `is_enable`   tinyint unsigned  not null      default 0 comment '启用：1=是；2=否',
    `created_at`  timestamp         not null      default CURRENT_TIMESTAMP,
    `updated_at`  timestamp         not null      default CURRENT_TIMESTAMP,
    `deleted_at`  timestamp                       default null,
    primary key (`id`),
    key (`clique_id`),
    key (`brand_id`),
    key (`parent_id`),
    key (`valid_start`),
    key (`valid_end`),
    key (`province_id`),
    key (`city_id`),
    key (`area_id`)
) auto_increment = 10000
  default collate = utf8mb4_unicode_ci comment '人资-组织表';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists `hr_organization`;
-- +goose StatementEnd
