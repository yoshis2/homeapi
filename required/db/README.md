# homeapi データベースMigragion

## 概要
Go言語製のmigrationであるliask/gooseでテーブルを管理する。

## 前提条件
MySQLが起動していること


## 手順
### インストール

```bash
go get github.com/liamstask/goose/cmd/goose
```

#### migration 実行するためのファイルを生成

```bash
goose create ファイル名 sql
```

#### migration 実行
```bash
goose up
```

#### 状態確認

```bash
goose status
```

## 例

エラーメッセージテーブル作成用ファイル作成
（すでに用意されているのでこの作業は不要）
```bash
goose create create_payment_error_messages sql
```

エラーメッセージテーブル作成を実行
```bash
goose up
```

エラーメッセージテーブルが作成されたかどうか確認
```bash
goose status
```

作成されていると以下のメッセージが出ます。
```
goose: status for environment 'development'
    Applied At                  Migration
    =======================================
    Wed Aug 15 07:23:32 2018 -- 20180815162246_create_payment_error_messages.sql
```


## 参考

https://github.com/liamstask/goose/src/master/
https://qiita.com/K_ichi/items/b9362e3a3c5688e494e2
http://shusatoo.net/programming/golang/goose-mysql-migration/

