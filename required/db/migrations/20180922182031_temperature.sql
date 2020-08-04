
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table temperatures
(
  id         int unsigned auto_increment
    primary key,
  temp       varchar(4) not null,
  humi       varchar(4) not null,
  created_at datetime   not null
)
  charset = utf8;



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `temperatures`;
