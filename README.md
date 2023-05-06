# アプリケーション構築課題
本リポジトリは、クラウドエース株式会社のアプリケーション構築課題のソースコードを格納するものです。

### 画面イメージ
![image](images/image.png)

### システム構成図
![システム構成図](images/システム構成図.png)

### api仕様書
[api仕様書](api/api_compose.yml)

### ER図
![ER図](images/DB.png)

### 技術スタック
- Golang 1.19
- Gorilla mux
- Firebase Authentication
- Vue3
- Vuetify3
- CloudRun
- CloudSQL for Postgres
- Docker

### ディレクトリ構成
/api にバックエンド、/web にフロントエンドソースコードを格納しています。

### バックエンドについて  
クリーンアーキテクチャを意識して実装しました。  
『依存関係の分離』を実現するため以下の構成を取っています。  
（以下、太字は /api/todos/pkg のディレクトリを示す）
- Enterprise Business Rules  
**/domain** ... APIでやりとりするデータ構造を定義
- Application Business Rules  
**/usecase** ... **/domain**で定義したデータ構造の取り扱いルールを定義
- Interface Adapters  
**/interface** ... こちらは実装途中です。が、現状は以下のように理解しています。  
  - **/usecase**と **/adapter** または **/infra**との橋渡しを行う。
  - 「あるべき姿」は以下のようになる。
    - ＜現状＞ 
      - **/adapter/handler.go**→ **/usecase/interactor.go**→ **/infra/db_connector.go**
    - <理想>
      - **/adapter/handler.go**→ （ **/interface**で/usecaseを呼び出し）→**/usecase/interactor.go**→ （ **/interface**でクエリ整形）→**/infra/db_connector.go**→（ **/interface**で返却値を生成）
- Frameworks & Drivers   
**/adapter** ... ルーティングフレームワークに依存する処理を記述  
**/infra** ... DBに依存する処理を記述  

参考：  
https://zenn.dev/tis1116/articles/6c5416e5d77dbf
https://blog.tai2.net/the_clean_architecture.html

### インフラについて
- CI/CD
  - devリポジトリにpushされると、CloudRunにデプロイされるよう構成。
- ローカル環境
  - docker-compose.ymlで管理。
  - 前提条件：サービスアカウントキーを付与されていること。
  - 手順：  
    以下のコマンドを実行する  
    ```docker-compose up --build -d```