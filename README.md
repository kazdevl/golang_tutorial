# golang_tutorial
golangを完全に理解するためのリポジトリ

## 参考URL
- [チュートリアル](https://go-tour-jp.appspot.com/basics/11)
- [言語仕様書](https://golang.org/ref/spec)
- [公式ドキュメント](https://golang.org/doc/)
- [標準パッケージ](https://golang.org/pkg/)
- [golangパッケージ](https://pkg.go.dev/)
- [goのロードマップ](https://github.com/Alikhll/golang-developer-roadmap/blob/master/i18n/ja-JP/ReadMe-ja-JP.md)
## 確認しておきたい内容
- Variables・Interface・Type・Strcutの基礎知識
    - sturctのポインタ型の、フィールドへのアクセスに関して、struct.Fieldという形式でアクセス可能なのは、Goコンパイラが、利便性のために暗黙的変換を行っているためである。実際には、(*struct).Fieladと解釈される。
    - 要確認: array・structは値型。Slice・Mapは参照型であることに注意
- Defer
    - deferステートメントで呼び出している関数は、その呼び出し元の関数の処理が終了時に実行される
    - 上記の関数の引数は、宣言時に評価される
- slice
    - 可変長配列
    - スライスは配列への参照のようなものです。スライスはどんなデータも格納しておらず、単に元の配列の部分列を指し示しています
    - コロンで区切られた二つのインデックス low と high の境界を指定することによってスライスが形成されます:
    ```
    a[low : high]
    これは最初の要素は含むが、最後の要素は除いた半開区間を選択します。
    ```
    - capの容量を超えて、スライスの長さを伸ばしたとき、元々のcapacityの倍capacitiyが取得される。cap(a)..4 → cap(a)..8 → cap(a)..16
- Pointer
    - 値のアドレス
    - poirterのpointerといった記述も可能
- Methods
    - 型にメソッド定義できる
    - 特別なレシーバー引数を持つ関数のこと
    - レシーバー引数の値を変更するようなメソッドでは、レシーバー引数をポインタ型にする必要がある
        - レシーバー引数は、そのメソッドを定義している型自身を指す
        - そうしない場合(値型の場合)、引数はコピーされた値を利用する(値のアドレスが異なる)ので、メソッドの内でのレシーバー引数の変更が、そのメソッドを持つ型の値に反映されない
- Goroutines: 深ぼっていく
    - ランタイムによって管理される軽量なスレッドのこと(厳密には違うらしい)
    - 並行処理と並列処理は別
        - [参考文献](https://freak-da.hatenablog.com/entry/20100915/p1)
        - 並行処理(Concurrency)
            - 複数の動作が、論理的に、順不同もしくは同時に起こりうる
            - CPUが1コアの時に、複数タスクを細切れにしてあたかも同時に動いてるかのようにしている処理
            - 同時にいくつかの異なるタスクを処理する
        - 並列処理(Parallelism)
            - 複数の動作が、物理的に、同時に起こること
            - 同じ処理を一斉に行う
            - 同時にいくつかの同じタスクを処理する
        - main Goroutineのプロセスが終了したら、全体の各Goroutineの処理も終了する
        - goroutinesが複数ある場合、どれが先にアクセスするか不明
            - 複数のgoroutineで共通の変数を取り扱うと、値の変更や参照が競合する
            - Do not communicate by sharing memory; instead, share memory by communicating
            - 解決手法1: 1つの変数には1つのgoroutineからアクセスする
            - 解決手法2: chanelを使って、goroutines間で値を送受信する
            - 解決手法3: syncパッケージで排他処理をする
        - chanel: 値を送受信できる通信路のようなもの、並行実行された関数間での値の送受信によく利用される
            - chanelは片方が準備できるまで、送受信は自動的にブロックされる(この結果同期的処理になる)
                - 送信時に、chanelのバファが埋まっていると、送信をブロック...chanelで受信して、送信できるようになるまで、送信をブロック
                - 受信時に、chanelのバッファが空だと、受信をブロック...chanelに値が送信されるまで受信をブロック
            - 上記により、明示的なロックや条件変数なしで、goroutine間で同期が取れる
            - バッファを持たせられる(指定なしだと、バッファは0)
            - 値の型を定義する
            - mutexなどの排他処理で共有メモリを問題なく扱うやり方もある
            - mainで、複数のchanelAとBを受信する処理を書いていて、goroutineAでchanleAに値を送信、goroutineBでchanelBに値を送信していて、
            goroutineAの処理がめっちゃ遅くて、goroutineBの処理がめっちゃ早いと、mainスレッドは、goroutineAの処理を受け取るまで、ブロックされるので、非効率。
            そこで、select case文で、ブロックされていないchanelの処理を随時実行する、ということが可能。
            - chanelを引数にするとき、受信専用chanel("<-chanel 型")や送信専用chanel("chanel<- 型")などにすることが可能...双方向chanelもある
                - 双方向→片側方向に関してはキャストの必要なし、しかし、逆は必要
            - Concurrencyの実現手法として、for selectパターンがある...goroutineで無限ループして、適宜結果をmainでselectで取得する(goroutineのメモリリークの可能性がある)
        - syncパッケージ
            - [参考文献1](https://qiita.com/t-mochizuki/items/80dc959b4221c53f3c40)
            - [参考文献2](https://golang.org/pkg/sync/)
            - [参考文献3](https://www.slideshare.net/takuyaueda967/gaegosync)
            - chanelだけを使っていると、コードが難解になる場合がある...送受信したいデータが単一型でない場合が多い。
            - データの競合が発生しないように、ロックを提供するパッケージ
            - sync.WaitGroup...複数のgoroutineを待機させる
                - DB系の処理に利用される
            - sync.Lock
            - sync.RLock
            - sync.Once
        - ErrGroup
            - [参考文献1](https://pkg.go.dev/golang.org/x/sync/errgroup#pkg-overview)
            - syncパッケージとcontextを利用して、共通のタスクのサブタスクで動作するgoroutineのグループに対して同期、エラー伝搬、contextのキャンセルを提供する.
            - func WithContext(ctx context.Context) (*Group, context.Context)
            - func (g *Group) Go(f func() error)
            - func (g *Group) Wait() error
        - コンテキスト
            - [参考文献1](https://qiita.com/marnie_ms4/items/985d67c4c1b29e11fffc)
            - [参考文献2](https://qiita.com/yoshinori_hisakawa/items/a6608b29059a945fbbbd)
            - [参考文献3](https://tutuz-tech.hatenablog.com/entry/2019/10/20/112353)
            - [参考文献4](https://blog.golang.org/context)
            - [参考文献5](https://golang.org/pkg/context/)
            - [参考文献6](https://pkg.go.dev/context)
            - [参考文献7](https://ayasuda.github.io/pages/what_is_context_at_go.html)
            - ゴールーチンをまたいだ処理のキャンセルを行う
            - 構造体のフィールドに保存しない
            - リクエスト起因のデータのみにする
            - Valueとして保存する場合のキーは外に公開しない
            - ```
                パッケージのContextはContextタイプを定義しています。Contextはデッドライン、キャンセルシグナル、その他のリクエストに対応した値をAPIの境界やプロセス間で伝達します。

                サーバーへの着信リクエストはContextを作成し、サーバーへの発信コールはContextを受け入れる必要があります。それらの間の関数呼び出しのチェーンは、Contextを伝播しなければならず、オプションとして、WithCancel、WithDeadline、WithTimeout、またはWithValueを使用して作成された派生Contextで置き換えることができます。あるContextがキャンセルされると、そのContextから派生したすべてのContextもキャンセルされます。

                WithCancel、WithDeadline、WithTimeoutの各関数は、Context（親）を受け取り、派生するContext（子）とCancelFuncを返します。CancelFuncを呼び出すと、子とその子がキャンセルされ、親の子への参照が削除され、関連するタイマーが停止されます。CancelFuncの呼び出しに失敗すると、親がキャンセルされるか、タイマーが切れるまで、子とその子がリークされます。go vetツールは、CancelFuncsがすべての制御フローのパスで使用されているかどうかをチェックします。

                コンテキストを使用するプログラムは、パッケージ間でインターフェイスの一貫性を保ち、静的解析ツールがコンテキストの伝播をチェックできるように、これらのルールに従うべきです。
            ```
    - goroutineに関しては並行処理
- testing
    - _testというsufixをファイル名につけることで、go test時にしかコンパイルされなくなる
        - [参考文献1](https://engineering.mercari.com/blog/entry/2018-08-08-080000/)
        - [参考文献2](https://future-architect.github.io/articles/20200601/)
        - [参考文献3](https://budougumi0617.github.io/2018/08/19/go-testing2018/)
        - mypkgと言うディレクトリにいる時でも、パッケージ名をmypkg_testとできて、別のパッケージとして扱える...従来、一つのディレクトリ内のコードは、1つのパッケージで構成されている必要がある
        - 非公開な変数や関数やタイプにアクセスしたい場合は、export_test.goと言うファイルで、mypkgと同じ、パッケージとして、非公開なものを公開する変数などを用意することで、テストで、ためしたい非公開なものにもアクセスできるし、export_test.goを挟むことで、公開範囲を適切に制限し、想定外の使われ方を省くことができる
    - testのデータなどは、testdataなどに入れるらしい...要確認
    - command
        - go test オプション
            - run
                - test対象を絞る
                - subテスト・subベンチマークも指定可能
            - v
                - 詳細を表示
            - short
                - shortのフラグを設定?
            - bench
                - ベンチマークの実行に必要
                - 関数の実行回数(有効なデータが得られるまで実行される)と1回の実行にかかった時間
            - benchmem
                - [参考文献](https://github.com/golang/tools/blob/master/benchmark/parse/parse.go#L28-L37)
                - メモリアロケートの回数が分かる
                - 実行ごとに割り当てられたメモリのサイズと1回の実行でメモリアロケーションが行われた回数が把握できる
            - その他諸々
                - [参考文献](https://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/)
             ```
                -test.bench regexp
                    run only benchmarks matching regexp
                -test.benchmem
                        print memory allocations for benchmarks
                -test.benchtime d
                        run each benchmark for duration d (default 1s)
                -test.blockprofile file
                        write a goroutine blocking profile to file
                -test.blockprofilerate rate
                        set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)
                -test.count n
                        run tests and benchmarks n times (default 1)
                -test.coverprofile file
                        write a coverage profile to file
                -test.cpu list
                        comma-separated list of cpu counts to run each test with
                -test.cpuprofile file
                        write a cpu profile to file
                -test.failfast
                        do not start new tests after the first test failure
                -test.list regexp
                        list tests, examples, and benchmarks matching regexp then exit
                -test.memprofile file
                        write an allocation profile to file
                -test.memprofilerate rate
                        set memory allocation profiling rate (see runtime.MemProfileRate)
                -test.mutexprofile string
                        write a mutex contention profile to the named file after execution
                -test.mutexprofilefraction int
                        if >= 0, calls runtime.SetMutexProfileFraction() (default 1)
                -test.outputdir dir
                        write profiles to dir
                -test.paniconexit0
                        panic on call to os.Exit(0)
                -test.parallel n
                        run at most n tests in parallel (default 2)
                -test.run regexp
                        run only tests and examples matching regexp
                -test.short
                        run smaller test suite to save time
                -test.testlogfile file
                        write test action log to file (for use only by cmd/go)
                -test.timeout d
                        panic test binary after duration d (default 0, timeout disabled)
                -test.trace file
                        write an execution trace to file
                -test.v
                        verbose: print additional output
                ```
    - Example
        - 処理の出力結果を比較する
        - コメントに、//Output: や、//Unorderd Output: ..順番を無視して出力検証
    - Test
        - 処理を試したい時に利用する
        - assertはなく、自分でエラーをカスタマイズして、動作を確認する
        - 成功時より、失敗の時が重要と考えているので、失敗時にたくさんの情報が出る
        - 並列実行などもあるらしい： [参考文献](https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/)
        - helper関数などもあるらしい: 要確認
    - Benchmark
        - 処理がどのくらいのスピードなのか、やどのくらいのメモリ消費量で行われるかを確認できる
        - *testing.Bで実行可能
        - 並列実行があるらしい
            - [参考文献1](https://qiita.com/marnie_ms4/items/8706f43591fb23dd4e64)
            - [参考文献2](https://golang.org/pkg/testing/)
            - benchmarkは一つの処理のベンチマークを図るものであり、複数の処理のベンチマークを並列に図るのではなく、一つのベンチマークを高速に図るために並列実行する(関数からも見てとれた)。...ベンチマークを取るために、めっちゃ処理を繰り返して行っているため。
            ```
            If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function; such benchmarks are intended to be used with the go test -cpu flag:
            ```
            - (*B).RunParallel
            ```
            RunParallel runs a benchmark in parallel. It creates multiple goroutines and distributes b.N iterations among them. The number of goroutines defaults to GOMAXPROCS. To increase parallelism for non-CPU-bound benchmarks, call SetParallelism before RunParallel. RunParallel is usually used with the go test -cpu flag.

            The body function will be run in each goroutine. It should set up any goroutine-local state and then iterate until pb.Next returns false. It should not use the StartTimer, StopTimer, or ResetTimer functions, because they have global effect. It should also not call Run.
            ```
            - (*PB).Next...実行するイテレーションがまだあるかどうかを報告する
    - Subtest & Subbenchmark
        - testやbenchmarkを階層化できる...これで複数ケースでの動作確認がやりやすい
        - *testing.T.Run()や*testing.B.Run()を利用して実行する
    - Main: 今後やっていく
        - なんか前処理と後処理とかが必要になってくるテストで使えるらしい
        - DBでの前処理でのinsertやDBでの後処理でのデータ消去
    - モック:
        - Bのメソッドを使うAのメソッドをtestしたいときで、Bのメソッドが外部APIなどを使っていて処理がばらつく、処理時間がめっちゃ時間かかる、実装がめっちゃ大変で内部処理が複雑である、といった時に、Bの振る舞いをするオブジェクト(機能的な実装はせずに、求めている型を持つ適当な戻り値と引数を設定している)を用意する。それがmockであり、AのメソッドがBのメソッドを呼び出すときの引数や回数が想定通りかを検証する。実際に、AのメソッドとBのメソッドがどのように連携するかを確かめるかのもの。テストの一部。
        - mockとstabの違い:
            - [参考文献1](https://craftsman-software.com/posts/38)
            - [参考文献2](https://qiita.com/k5trismegistus/items/10ce381d29ab62ca0ea6#:~:text=%E3%82%B9%E3%82%BF%E3%83%96%E3%81%A8%E3%83%A2%E3%83%83%E3%82%AF%E3%81%AE%E6%9C%80%E5%A4%A7,%E3%81%84%E3%81%A3%E3%81%A6%E3%82%88%E3%81%84%E3%81%A7%E3%81%97%E3%82%87%E3%81%86%E3%80%82)
            - [参考文献3](https://gotohayato.com/content/483/)
            - スタブとは、テストに必要だけどまだ実装出来ていないモジュールがある時に、そのモジュールの代わりにテストケースに沿った値を返してくれるオブジェクト
    - mockgenとgomock
        - mockgenは、インタフェースからmockを生成する
        - gomockは、mockを取り扱うライブラリ(当然だが、testingとかもかなりライブラリ内で利用されている)
        - [参考文献1](https://pkg.go.dev/github.com/golang/mock/gomock#Controller)
        - [参考文献2](https://github.com/golang/mock)
        - [参考文献3](https://www.asobou.co.jp/blog/web/gomock)
        - しっかりとソースコードを読み進める必要があるが、recoderの関数などを用いて設定した値をMockのメソッドで返している
    - APIサーバなどの、httpリクエスト関連のtest
        - 開発時に要確認
- goのディレクトリの意味の理解する！
    - [参考文献](https://future-architect.github.io/articles/20200528/)
    - testdata...コンパイルの対象外
- go generator
    - codeの自動生成ができるらしい
    - mockgenを大量に行う必要があるときに、これを使うことでかなり楽にmockファイルも作れそう
- GoDoc
    - 開発時に要確認
- GoCodeReview
    - 開発時に要確認

- 標準パッケージ
    - net
        - http
        - mail
    - os
    - reflect...ORMapperとの関連性がありそう
    - io
    - log
    - database/sql
        - [参考文献1](https://dev.mysql.com/doc/refman/5.6/ja/insert-speed.html)
            - Insertでは、単一のレコードを何度も挿入するより、まとめてinsertした方が良い。
            - SQLのInsertの構文的に、まとめてレコードを挿入するとが可能
            - Insertなどによる挿入にかかる時間とその割合
                - 接続: (3)
                - サーバーへのクエリーの送信: (2)
                - クエリーの解析: (2)
                - 行の挿入: (1*行のサイズ)
                - インデックスの挿入: (1*インデックス数)
                - クローズ: (1)
        - [参考文献2](https://www.ipa.go.jp/files/000017320.pdf)
            - SQLインジェクションはSQL文をアプリケーションから利用する場合、SQL文のリテラル部分をパラメータ化することが一般的です。パラメータ化された部分を実際の値に展開するとき、リテラルとして文法的に正しく文を生成しないと、パラメータに与えられた値がリテラルの外にはみ出した状態になり、リテラルの後ろに続く文として解釈されることになります。
            - 安全なSQLの呼び出し方
                - 文字列リテラルにはエスケープすべき文字をエスケープさせること(DBによって異なるケースがある)
                - 数値リテラルに対しては、数値以外の文字を入力させない

        - [参考文献3](https://golang.org/pkg/database/sql/#example_Stmt)
            - execはinterfaceであり、その実装はDriverに依存している
    - errors
    - time


## 豆知識
- interface
    - golangには明示的なimplementが存在しないため、大規模になってくるとどの構造体がどのインタフェースを実装しているかが一目で把握できない可能性がある(構造体を一挙に宣言している時など)
    - 上記の対策として、以下のように記述することで、aはAとBを実装していることが明記できるし、実際に実装できていないとコンパイルエラーになる
    ```
    var (
        _ A = a
        _ B = a
        ...
        ...
        ...
    )
    ```
    - 型はメソッドの任意のサブセットから、任意のインタフェースを実装するので、複数の異なるインタフェースを実装できる
        - よって、全ての型は空のインタフェースを実装しているとも言えるので、interface{}(0個のメソッドを実装した型)を全ての型が実装していると言える。
    - [参考文献](https://golang.org/ref/spec#Interface_types)
- struct
    - strcut{}{}はメモリを消費しない。
    - e.g.)システムに存在するアイテムのなかでユーザが持っていないものを確認したいとき、
    mapのなかに、空の構造体を使うことでメモリ消費を抑えた実装で実現することができる
        ```
        items := getAllItemsInGame() //戻り値: []struct{name string, value int}
        user_items := getUserItems() //戻り値: []struct{name string, value int}
        var hashmap_user_items map[string]struct{}
        for _, u_item := range user_items {
            hashmap_user_items[u_item.name] = sturct{}
        }
        // 全てのアイテムの中で、ユーザが持っていないものがあるかを確認する処理
        for _, item := range items {
            if _, ok := hashmap_user_items[item.name]; !ok {
                fmt.Printf("user don't have %s", item.name)
            }
        }
        ```
- 一般のアプリケーションでは、サーバーを起動し続けるためにmainゴールーチンが動きっぱなし、開放すべきものは開放することでメモリを効率的に扱えるので、使用しないchanelはcloseする
    - dbなどのコネクションに関しては何回も接続するのは圧倒的に非効率なので、defer db.Close()とすることで、mainゴールーチンが落ちた時(サーバが終了した時)のみデータベースとの接続を終了する
## 参考になる記事
- データベースのコネクションプールについて: [参考文献](https://please-sleep.cou929.nu/go-sql-db-connection-pool.html)
- GoroutineをWaitGroup(ErrGroup)で制御する: [参考文献](https://blog.toshimaru.net/goroutine-with-waitgroup/#goroutine--errgroup--context-%E3%82%92%E4%BD%BF%E3%81%86)
- Goroutineのエラーハンドリング: [参考文献](https://ludwig125.hatenablog.com/entry/2019/09/28/043334)
- sliceの理解: [参考文献](https://qiita.com/tenntenn/items/5229bce80ddb688a708a)
- Goのアプリケーションの実装の面で参考になる: [参考文献](https://github.com/camphor-/relaym-server/blob/master/main.go)
- mysqlの文字コードについて: [参考文献](https://yoneyore.hatenablog.com/entry/2020/11/26/232128)