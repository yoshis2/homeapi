-- +migrate Up
create table twitters (
    id         int unsigned auto_increment primary key,
    tweet      varchar(160) not null comment 'ツイートメッセージ',
    created_at datetime     not null,
    updated_at datetime     not null
);

-- +migrate Down
drop table twitters;
