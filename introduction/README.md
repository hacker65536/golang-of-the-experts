# 序章 プロダクト開発の前に習得しておきたい6機能
  
##  0.1.1 Goにおける型


https://go.dev/ref/spec#Types
### 組み込み型一覧 (primtive)

import をしなくても使える


| 種類 | 型名 |
|--------|----|
| 整数  | int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr,byte,rune |
| 浮動小数点数    | float32,float64 |
| 複素数     | complex64,complex128 |
| 文字列     | string |
| 真偽値     | bool |
| エラー     | error |


型名は予約後ではないので、変数名として使うことができるが混乱をさけるために使わないほうが良い  
https://go.dev/play/p/39uG5TFg3ju


### コンポジット型 (composite)

複数のデータを一つの集合として表す型
struct,arry,slice,map,cannel(chan)

#### 構造体  
https://go.dev/play/p/mzZjR_Vh4Ol


#### 配列
型とデータ数は一度決めたら固定  
https://go.dev/play/p/gYFVhCmhGiA


#### スライス
スライスは型とデータ数を持たないために可変である  
https://go.dev/play/p/nblFiIojR9V


#### マップ
キーと値は別々の型を使うことができる、ただしキーは比較演算子で比較できる型でなければならない  
https://go.dev/play/p/LNtBCLF0XJy

### ユーザ定義型 (type definitions)


ユーザ定義型

```go
type MyDuration time.Duration
                // ↑ 基底型 (underlying type)
```

https://go.dev/play/p/1lVVC0OFrLA

https://zenn.dev/sryoya/articles/6a8ae7daa20aa7  
sliceを追加するときの挙動は一定ではないので、パフォーマンスのことも考慮する必要がある




##  0.1.2 スライスの操作


https://go.dev/play/p/fVsbIKFjQ4L

悩んだ時は型のポインタを指定するとよい


slice tricks

https://ueokande.github.io/go-slice-tricks/  
https://github.com/golang/go/wiki/SliceTricks  
https://go.dev/ref/spec#Slice_expressions

copyでsliceの複製をする

https://go.dev/play/p/t7HtpC5Z9jF

appendでスライス同士を連結

https://go.dev/play/p/7hg1-MeFPLh

append,copyでスライス内の要素を削除

https://go.dev/play/p/8nfz18dEZlU

スライスを逆順に並び替える

https://go.dev/play/p/5SBJ5XUYjZl

スライスの要素を偶数のみでフィルタリングする

https://go.dev/play/p/soU62BNmFp6


スライスを任意の要素数に分割する

https://go.dev/play/p/S6gz5maCScE


##  0.1.3 マップの操作

キーはユニークである、順序は保証されない  
キーは比較演算子で等しく比較できる型であるひつようがある

### 値の初期化

https://go.dev/play/p/gR68CgC4e7a

### 値及びキーの存在の有無の取得

値のみを取得する/値、キーの存在濃霧を取得する

https://go.dev/play/p/CiuWvYWB0T7


ゼロ値のマップから値と、キーの存在の有無を取得する

https://go.dev/play/p/cYuIgCYaGXc


マップを利用して重複排除処理を実装する

https://go.dev/play/p/ymvIAQJrnR_O

##  0.1.4 構造体の設計

型名の先頭が大文字か小文字禍で参照できる範囲が異なる

```go
	// パッケージ外からも参照できる
	type Export struct {
		// パッケージ外からも参照できる
		Name string
		// パッケージ外からは参照できない
		age int
	}

	// パッケージ外から参照できない
	type unexport struct {
		// パッケージ外からも参照できる
		Name string
		// パッケージ外からは参照できない
		age int
	}
```
[カウンターの実装](0.1.19/)


公開されているフィールドのみをJSONへ変換する

[code](./0.1.20/)

##  0.1.5 メソッドの呼び出し
##  0.1.6 インタフェースの基礎
##  0.1.7 インタフェースの設計
##  0.1.8 型情報を利用した処理
