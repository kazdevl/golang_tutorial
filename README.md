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
        - https://freak-da.hatenablog.com/entry/20100915/p1
        - 並行処理(Concurrency)
            - 複数の動作が、論理的に、順不同もしくは同時に起こりうる
            - CPUが1コアの時に、複数タスクを細切れにしてあたかも同時に動いてるかのようにしている処理
            - 同時にいくつかの異なるタスクを処理する
        - 並列処理(Parallelism)
            - 複数の動作が、物理的に、同時に起こること
            - 同じ処理を一斉に行う
            - 同時にいくつかの同じタスクを処理する
        - main Goroutineのプロセスが終了したら、全体の各Goroutineの処理も終了する
        - chanel: 値を送受信できる通信路のようなもの、並行実行された関数間での値の送受信によく利用される
            - chanelは片方が準備できるまで、送受信は自動的にブロックされる
            - 上記により、明示的なロックや条件変数なしで、goroutine間で同期が取れる
            - mutexなどの排他処理で共有メモリを問題なく扱うやり方もある
    - goroutineに関しては並行処理
- testing
    - _testというsufixをファイル名につけることで、go test時にしかコンパイルされなくなる
        - https://engineering.mercari.com/blog/entry/2018-08-08-080000/
        - https://future-architect.github.io/articles/20200601/
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
                - https://github.com/golang/tools/blob/master/benchmark/parse/parse.go#L28-L37
                - メモリアロケートの回数が分かる
                - 実行ごとに割り当てられたメモリのサイズと1回の実行でメモリアロケーションが行われた回数が把握できる
            - その他諸々
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
        - 並列実行などもあるらしい： (参考文献)[https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/]
        - helper関数などもあるらしい: 要確認
    - Benchmark
        - 処理がどのくらいのスピードなのか、やどのくらいのメモリ消費量で行われるかを確認できる
        - *testing.Bで実行可能
        - 並列実行があるらしい: (参考文献)[https://qiita.com/marnie_ms4/items/8706f43591fb23dd4e64]
    - Subtest & Subbenchmark
        - testやbenchmarkを階層化できる...これで複数ケースでの動作確認がやりやすい
        - *testing.T.Run()や*testing.B.Run()を利用して実行する
    - Main: 今後やっていく
        - なんか前処理と後処理とかが必要になってくるテストで使えるらしい
        - DBでの前処理でのinsertやDBでの後処理でのデータ消去
    - モック:
        - Bのメソッドを使うAのメソッドをtestしたいときで、Bのメソッドが外部APIなどを使っていて処理がばらつく、処理時間がめっちゃ時間かかる、実装がめっちゃ大変で内部処理が複雑である、といった時に、Bの振る舞いをするオブジェクト(機能的な実装はせずに、求めている型を持つ適当な戻り値と引数を設定している)を用意する。それがmockであり、AのメソッドがBのメソッドを呼び出すときの引数や回数が想定通りかを検証する。実際に、AのメソッドとBのメソッドがどのように連携するかを確かめるかのもの。テストの一部。
        - mockとstabの違い:
            - (参考文献1)[https://craftsman-software.com/posts/38]
            - (参考文献2)[https://qiita.com/k5trismegistus/items/10ce381d29ab62ca0ea6#:~:text=%E3%82%B9%E3%82%BF%E3%83%96%E3%81%A8%E3%83%A2%E3%83%83%E3%82%AF%E3%81%AE%E6%9C%80%E5%A4%A7,%E3%81%84%E3%81%A3%E3%81%A6%E3%82%88%E3%81%84%E3%81%A7%E3%81%97%E3%82%87%E3%81%86%E3%80%82]
            - (参考文献3)[https://gotohayato.com/content/483/]
            - スタブとは、テストに必要だけどまだ実装出来ていないモジュールがある時に、そのモジュールの代わりにテストケースに沿った値を返してくれるオブジェクト
    - mockgenとgomock
        - mockgenは、インタフェースからmockを生成する
        - gomockは、mockを取り扱うライブラリ(当然だが、testingとかもかなりライブラリ内で利用されている)
        - (参考文献1)[https://pkg.go.dev/github.com/golang/mock/gomock#Controller]
        - (参考文献2)[https://github.com/golang/mock]
        - (参考文献3)[https://www.asobou.co.jp/blog/web/gomock]
        - しっかりとソースコードを読み進める必要があるが、recoderの関数などを用いて設定した値をMockのメソッドで返している
    - APIサーバなどの、httpリクエスト関連のtest
        - 開発時に要確認
- goのディレクトリの意味の理解する！
    - 参考文献: https://future-architect.github.io/articles/20200528/
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
    - errors
    - time