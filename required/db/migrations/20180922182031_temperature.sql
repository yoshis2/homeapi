
-- +migrate Up
create table temperatures
(
  id         int unsigned auto_increment
    primary key,
  temp       varchar(4) not null,
  humi       varchar(4) not null,
  created_at datetime   not null
)
  charset = utf8;



-- +migrate Down
DROP TABLE `temperatures`;
