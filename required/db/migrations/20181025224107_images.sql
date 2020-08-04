
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table images
(
  id         int unsigned auto_increment
    primary key,
  image_name varchar(255)  not null
  comment '画像名',
  image_path varchar(1024) not null
  comment '画像パス',
  created_at datetime      not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table images;

