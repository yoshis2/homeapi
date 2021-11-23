# Home API

## 環境構築

### git クローン

```shell
git clone git@github.com:yoshis2/homeapi.git
```

### configの設定

infrastructure/config の階層にある```develop.env.sample``` のファイルを  
 ```develop.env```に名前を変更する。```.sample```を取り除くだけです。

### TwitterAPIの必要なキーを以下のURLから取得する

```develop.env```ファイルの下４つのパラメータは以下のURLから取得する  
https://developer.twitter.com/en/apps

### Google Cloud Platformの設定ファイル取得

Google Cloud Platformの管理画面から以下のファイルを取得します

- serviceaccount-key.json

取得したファイルはinfrastructure/configに格納します

### ログファイルの作成

以下のログファイルを作成してください

```shell
mkdir required/logs
cd required/logs
touch access.log
```

### ビルドと実行

configの準備ができたらビルドして実行することができます。実行手順は以下の通り

```shell
make build
make serve
```

## HomeAPIの概要

### 目的

自宅で必要な情報をAPIとして作成しています。クリーンアーキテクチャによる設計で作成されているため、クリーンアーキテクチャの設計に興味がある人はご覧ください。

### 機能

- 部屋の温度を一定の時間で取得しデータを収集する。
- MySQLからTwitterのつぶやき
- Firebase Firestoreへの参照、更新、追加、削除機能

## URL

- api url  
https://www.seldnext.com  

- swagger  
https://www.seldnext.com/swagger/index.html

## 開発環境

- Go バージョン  1.14
- Docker環境
- MySQL8
- Redis
- Firebase Firestore

## sql-migrateのファイル作成

### development環境の実行

```shell
make in
cd required
sql-migrate create createTwitters
sql-migrate up
```

### production環境の実行

```shell
make in
cd required
sql-migrate new -env="production" create_users
sql-migrate up -env="production"
```
