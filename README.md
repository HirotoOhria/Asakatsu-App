# 朝活App

## 概要

朝活Appは、朝活部の価値をより高めるためのアプリケーションです。

# CloudFunctions

## 新しい関数の追加

### Makefile

- 既存の `Makefile` をコピーし、`firebase/functions/cmd/${function-name}/Makefile` を作成する
- `Makefile` 内の変数の値を書き換える
    - `function_name`: CloudFunctins の関数名
    - `trigger_topic_name`: PubsubMessage の名前（PubsubMessage トリガーの関数の場合）
    - `cmd_directory_name`: `firebase/functions/cmd/${function-name}` の `${function]name}` 部分の名前
    - `runtime_go_version`: 基本的に書き換える必要はない

### デプロイ関数

- Cloudfunctions にデプロイする関数を `firebase/functions/src/cloud_dunctions.go` に追加する
    - デプロイ関数からは、ハンドラーのメソッドを呼び出す

### ローカル実行用の main メソッド

- `firebase/functions/src/app/${function-name}/main.go` を作成し、ローカル実行用の main メソッドを定義する
    - ローカル実行用のメソッドからは、ハンドラーのメソッドを呼び出す

### ハンドラーメソッド

- `firebase/functions/src/handler/handler.go` に、ハンドラーメソッドを追加する
    - ハンドラーメソッドからは、ユースケースのメソッドを呼び出す

## デプロイ

### 最初のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make create_function
```

### 2回目以降のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make deploy 
```