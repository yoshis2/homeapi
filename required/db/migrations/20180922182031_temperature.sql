
-- +migrate Up
create table thermometers
(
  id          int unsigned auto_increment primary key,
  temperature float unsigned not null comment '温度',
  humidity    float unsigned not null comment '湿度',
  created_at  datetime not null
);

-- +migrate Down
DROP TABLE `temperatures`;
