# 概要
指定したn番目のフィボナッチ数を返すREST API。

# 技術スタック
- 言語: Go
- フレームワーク: Gin
- サーバ: AWS EC2 Ubuntu22.04
- リバースプロシキサーバ: Nginx

# APIエンドポイント
`http://ec2-13-236-52-212.ap-southeast-2.compute.amazonaws.com/fib?n=99`

クエリパラメータnの値を変更することでn番目のフィボナッチ数を返す。

# ディレクトリ構造
- controller/ # Modelへの橋渡しを担当。クライアントから送られたURIをルーティング。
    - error_controller.go # エラー時のレスポンスを返す関数がまとまっている。
    - handle_controller_test.go # handle_controllerのテストコード。
    - handle_controller.go # レスポンスを返す処理。
    - router.go # リクエストされたURIルーティングする。
- model/ # ビジネスロジックを担当。
    - calc_fib_test.go # calc_fibのテストコード。
    - calc_fib.go # フィボナッチ数の計算を行う。
- go.mod # モジュールのインポートパスとバージョン情報。
- go.sum # 依存先モジュールのハッシュを記録。
- main.go # エントリーポイント。

# 主要な関数とロジック
- CalcFib(num int) *big.Int 
  - ソースファイル: model/calc_fib.go
  - ロジック: 引数numが1または2のとき、1を返す。numが2以上の整数のとき、初期値1のpre_fib, cur_fibを1回のループで`pre_fib=cur_fib, cur_fib=cur_fib+pre_fib`を代入し、num-1回繰り返すことでn番目のフィボナッチ数を*big.Int型で返す。numが大きくなると返り値も指数関数的に大きくなるため、より大きな計算を扱えるよう`"math/big"`ライブラリを使用。
- GetRouter() *gin.Engine
  - ソースファイル: controller/router.go
  - ロジック: Ginフレームワークを使用してルーターを設定し、各エンドポイントに関数を関連付ける。ルーターに対して「ルートが見つからない場合」と「メソッドが許可されていない場合」のハンドリング方法を設定している。
- ShowFibNumber(c *gin.Context)
  - ソースファイル: controller/handle_controller
  - ロジック: HTTPリクエストが /fib エンドポイントに対して送信された際に呼び出される。クライアントからのリクエストを受けてフィボナッチ数を計算し、正常なJSONレスポンスとして返すためのロジックを提供する。0以下または250000を超える場合、リクエストが無効であると判断し、BadRequestを返す。これは、計算不可能または膨大な結果を返してしまうため。

# テスト
- controller/handle_controller_test.go
  - 説明: ShowFibNumber()関数が異なるケースに対して正しく動作するかを検証する。各テストケースはクエリパラメータを変えてリクエストを送信し、期待されるHTTPステータスコードとレスポンスの内容が正しく返されるかをテストする。
- model/calc_fib_test.go
  - 説明: CalcFib()関数が異なる入力ケースに対して正しく動作するかを検証する。各テストケースは、異なる入力値numに対して CalcFib()関数を呼び出し、期待される計算結果と実際の結果を比較してテストする。テストケースは、正の整数、負の整数、0など、さまざまな入力ケースをカバーし、関数の振る舞いを網羅的にテストする。

実行例
``` 
$ go test ./...
?       github.com/KaiYanagisawa/fib_api        [no test files]
ok      github.com/KaiYanagisawa/fib_api/controller     (cached)
ok      github.com/KaiYanagisawa/fib_api/model  (cached)
```

# 使用方法

```
$ curl -X GET -H "Content-Type: application/json" "http://ec2-13-236-52-212.ap-southeast-2.compute.amazonaws.com/fib?n=10" 
  {"result":55}
```
