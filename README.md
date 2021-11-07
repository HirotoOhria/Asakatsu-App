# 目次

- [朝活App](#asakatsu-app)
    - [概要](#overview)
- [構成](#composition)
    - [バックエンド](#backend)
        - [三層アーキテクチャ](#mvc)
        - [DDD](#ddd)
        - [Clean Architecture](#clean-architecture)
        - [Cloud Function用](#composition-cloud-function)
        - [ローカル用](#composition-local)
        - [その他](#compotion-other)
    - [フロントエンド](#forntend)
        - [Flux](#flux)
        - [Atomic Design](#atomic-design)
- [開発](#decelopment)
    - [Firebase Hosting](#firebase-hosting)
    - [CloudFunctions](#cloud-functions)
        - [新しい関数の追加](#add-new-function)
            - [Makefire](#makefile)
            - [デプロイ関数](#deploy-function)
            - [ローカル実行用の main メソッド](#local-exec-main-method)
            - [ハンドラーメソッド](#handler-method)
        - [デプロイ](#deploy)
            - [最初のデプロイ](#first-deploy)
            - [2回目以降のデプロイ](#second-deploy)
- [開発予定の機能](#feature)
- [振り返り](#look-back)
    - [吾輩はWebアプリケーションである。テストはまだない。](#no-test)
    - [Clean Architecture の採用](#adopt-clean-architecture)

<a id="asakatsu-app"></a>

# 朝活App

<a id="overview"></a>

## 概要

朝活Appは、朝活部の価値をより高めるためのアプリケーションです。

<a id="composition"></a>

# 構成

<a id="infrastructure"></a>

## インフラ

<a id="backend"></a>

## バックエンド

<a id="mvc"></a>

### 三層アーキテクチャ

- 三層アーキテクチャを採用しています
- `src/handler/executor.go`
    - ハンドラーの実行者です
    - 依存関係を注入し、ハンドラーを実行します
- `src/handler/functions_handler`
    - ハンドラーを配置します
    - ハンドラーは、外部からのリクエスト情報を受け取り、レスポンスを書き込みます
    - 内部的にインプットを作成し、ユースケースを呼び出します
    - ローカル実行用に、アウトプットを返却します
- `src/usecase`
    - ユースケースを配置します
    - インプットを受け取ります
    - 仕様を満たすために、リポジトリ、ドメインオブジェクトを操作します
- `src/domain/api_io`
    - APIのインプット/アウトプットを配置します
    - APIのインプット/アウトプットは、APIの入力情報/返却情報を表現します
- `src/domain/repository`
    - リポジトリを配置します
    - リポジトリは、外部システムとのやり取りを請け負います
- `src/domain/entity`
    - DBのエンティティを配置します
    - エンティティは、DBのレコードを表現します

<a id="ddd"></a>

### DDD

- DDDを採用しています
- `src/domain/domain_object`
    - ドメインオブジェクトを配置します
    - 業務ロジックは基本的にドメインオブジェクトに定義します

<a id="clean-architecture"></a>

### Clean Architecture

- Clean Architecture を採用しています
    - 依存関係は低レイヤーに向かっていき、逆流しません
- `src/injector`
    - インジェクターを配置します
    - インジェクターは、各層の依存関係を注入し、構造体を返却します

<a id="composition-cloud-function"></a>

### CloudFunction用

- `src/cloud_functins`
    - CloudFunction の関数を配置します
    - CloudFunction の本番環境には、ここに定義されている関数がデプロイされます

<a id="composition-local"></a>

### ローカル用

- `src/app/**/main.go`
    - main 関数を配置します
    - ローカルからはこの main 関数を実行します
- `cmd/**/Makefile`
    - Makefileを配置します
    - Makefileには、開発に必要なコマンドを記述します
        - デプロイ、ローカル実行用のコマンドなど

<a id="composition-other"></a>

### その他

- `src/infrastructure/client`
    - 各種SDKのクライントの初期設定を配置します
- `src/util`
    - Utilを配置します
    - Utilには、内部状態に依存しない単純なロジックを配置します

<a id="frontend"></a>

## フロンエンド

<a id="flux"></a>

### Flux

- 状態管理には、Fluxを採用しています
- 仕様パッケージは `Vuex` です

<a id="atomic-design"></a>

### Atomic Design

- アトミックデザインを採用しています

<a id="development"></a>

# 開発

<a id="firebase-hosting"></a>

## Firebase Hosting

<a id="cloud-functions"></a>

## CloudFunctions

<a id="add-new-function"></a>

### 新しい関数の追加

<a id="makefile"></a>

#### Makefile

- 既存の `Makefile` をコピーし、`firebase/functions/cmd/${function-name}/Makefile` を作成する
- `Makefile` 内の変数の値を書き換える
    - `function_name`: CloudFunctins の関数名
    - `trigger_topic_name`: PubsubMessage の名前（PubsubMessage トリガーの関数の場合）
    - `cmd_directory_name`: `firebase/functions/cmd/${function-name}` の `${function]name}` 部分の名前
    - `runtime_go_version`: 基本的に書き換える必要はない

<a id="#deploy-function"></a>

#### デプロイ関数

- Cloudfunctions にデプロイする関数を `firebase/functions/src/cloud_dunctions.go` に追加する
    - デプロイ関数からは、ハンドラーのメソッドを呼び出す

<a id="local-exec-main-method"></a>

#### ローカル実行用の main メソッド

- `firebase/functions/src/app/${function-name}/main.go` を作成し、ローカル実行用の main メソッドを定義する
    - ローカル実行用のメソッドからは、ハンドラーのメソッドを呼び出す

<a id="handler-method"></a>

#### ハンドラーメソッド

- `firebase/functions/src/handler/handler.go` に、ハンドラーメソッドを追加する
    - ハンドラーメソッドからは、ユースケースのメソッドを呼び出す

<a id="deploy"></a>

### デプロイ

<a id="first-deploy"></a>

#### 最初のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make create_function
```

<a id="second-deploy"></a>

#### 2回目以降のデプロイ

- `FetchActivitiesFromSlackBatch` の場合

```shell
$ cd firebase/functions/cmd/fetch_activities_from_slack

$ make deploy 
```

<a id="feature"></a>

# 開発予定の機能

- インフラ
    - [ ] Docker化
- フロントエンド
- バックエンド
    - [ ] インターフェースの追加
    - [ ] テストの追加

<a id="look-back"></a>

# 振り返り

<a id="no-test"></a>

## 吾輩はWebアプリケーションである。テストはまだない。

- 反省しています
    - 正直、プロトタイプ気味の小規模アプリケーションでは手動確認でも十分では
    - しかし、API が 3つを超えたあたりからテストがないときつくなる気がする

<a id="adopt-clean-architecture"></a>

## Clean Architecture の採用

- どう考えてもオーバースペック
    - 初期開発がしづらい
    - 重厚でコードの記述量が増える
        - 小規模アプリケーションには不向きでは
