
-- +migrate Up
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

-- +migrate Down
drop table images;
