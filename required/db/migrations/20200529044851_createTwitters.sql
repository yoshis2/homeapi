
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table twitters (
    id int unsigned auto_increment primary key,
    message varchar(160) not null,
    created_at datetime,
    updated_at datetime
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table twitters;