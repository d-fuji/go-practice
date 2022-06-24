# Goの個人学習

## 6/24
[本日のコード](https://github.com/d-fuji/go-practice/blob/272f092aab77edbccbdf513daccd9819d0b102ff/code/2022/06/23/main.go)  

[A Tour of Go](https://go-tour-jp.appspot.com/list)でGoの超基本的な文法を一通り把握した  
- DMMだと新卒研修のドキュメントになってるの羨ましいな

[Go での並列処理 - 最初の一歩](youtube.com/watch?v=3OOYON47aQ0)を一通り見た、ふーんといった感じ  
- 単純にGoルーチンを実行しただけでは処理が終了する前にプログラムが終了して何も出力されないということを知った  
- 上記を回避するために、limitやsync、waitを使うんだなというざっくりとした理解  

[Go言語による並行処理](https://www.oreilly.co.jp/books/9784873118468/)を購入した  

シングルスレッドのノンブロッキングI/Oであるnodejsでも並列実行とほぼ同じこと（待ち時間を有効活用することが目的だと思っている）を実現できると思うが、Goの方がコンパイル言語なので単純にパフォーマンスが優れているのか？（limitやsync、waitとかを使うのが面倒くさいので開発者のレベルによってはnodejsの方が良さそう、方に関してはTSがあるので引き分けか？）
