-- データベースの作成
create database handson;

-- 作成したデータベースの確認
show databases;

-- データベースの選択
use handson;

-- テーブルの作成 IDをプライマリキーとして設定
create table handson.metrics (
    id int primary key auto_increment,
    metrics_name varchar(100),
    value int
);

-- 作成したテーブルの確認
show tables;

-- テーブルのスキーマの確認
show columns from handson.metrics;

-- メトリクス用のデータを追加
insert into handson.metrics(metrics_name, value) value ('sample_metrics', 1);

-- 追加したデータの確認
select * from handson.metrics;

-- 追加したデータの値を更新
update handson.metrics set value=2 where id=1;