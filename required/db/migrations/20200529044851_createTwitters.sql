
-- +migrate Up
create table twitters (
    id int unsigned auto_increment primary key,
    message varchar(160) not null,
    created_at datetime,
    updated_at datetime
);


-- +migrate Down

drop table twitters;
