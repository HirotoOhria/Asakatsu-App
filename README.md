# 目次

- [朝活App](#asakatsu-app)
    - [概要](#gaiyou)
- [CloudFunctions](#cloud-functions)
    - [新しい関数の追加](#add-new-function)
        - [Makefire](#makefile)
        - [デプロイ関数](#deploy-function)
        - [ローカル実行用の main メソッド](#local-exec-main-method)
        - [ハンドラーメソッド](#handler-method)
    - [デプロイ](#deploy)
        - [最初のデプロイ](#first-deploy)
        - [2回目以降のデプロイ](#second-deploy)
- [振り返り](#look-back)
    - [吾輩はWebアプリケーションである。テストはまだない。](#no-test)
    - [Clean Architecture](#clean-architecture)

<a id="asakatsu-app"></a>

# 朝活App

<a id="gaiyou"></a>

## 概要

朝活Appは、朝活部の価値をより高めるためのアプリケーションです。

<a id="cloud-functions"></a>

# CloudFunctions

<a id="add-new-function"></a>

## 新しい関数の追加

<a id="makefile"></a>

### Makefile

- 既存の `Makefile` をコピーし、`firebase/functions/cmd/${function-name}/Makefile` を作成する
- `Makefile` 内の変数の値を書き換える
    - `function_name`: CloudFunctins の関数名
    - `trigger_topic_name`: PubsubMessage の名前（PubsubMessage トリガーの関数の場合）
    - `cmd_directory_name`: `firebase/functions/cmd/${function-name}` の `${function]name}` 部分の名前
    - `runtime_go_version`: 基本的に書き換える必要はない

<a id="#deploy-function"></a>

### デプロイ関数

- Cloudfunctions にデプロイする関数を `firebase/functions/src/cloud_dunctions.go` に追加する
    - デプロイ関数からは、ハンドラーのメソッドを呼び出す

<a id="local-exec-main-method"></a>

### ローカル実行用の main メソッド

- `firebase/functions/src/app/${function-name}/main.go` を作成し、ローカル実行用の main メソッドを定義する
    - ローカル実行用のメソッドからは、ハンドラーのメソッドを呼び出す

<a id="handler-method"></a>

### ハンドラーメソッド

- `firebase/functions/src/handler/handler.go` に、ハンドラーメソッドを追加する
    - ハンドラーメソッドからは、ユースケースのメソッドを呼び出す

<a id="deploy"></a>

## デプロイ

<a id="first-deploy"></a>

### 最初のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make create_function
```

<a id="second-deploy"></a>

### 2回目以降のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make deploy 
```

<a id="look-back"></a>

# 振り返り

<a id="no-test"></a>

## 吾輩はWebアプリケーションである。テストはまだない。

- 反省しています
    - 正直、プロトタイプ気味の小規模アプリケーションでは手動確認でも十分では
    - しかし、API が 3つを超えたあたりからテストがないときつくなる気がする

<a id="clean-architecture"></a>

## Clean Architecture

- どう考えてもオーバースペック
    - 初期開発がしづらい
    - 重厚でコードの記述量が増える
        - 小規模アプリケーションには不向きでは
