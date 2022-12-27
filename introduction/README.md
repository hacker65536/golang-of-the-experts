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

struct,arry,slice,map,cnanne,chan

構造体
https://go.dev/play/p/mzZjR_Vh4Ol


配列
型とデータ数は一度決めたら固定
https://go.dev/play/p/gYFVhCmhGiA


スライス
スライスは型とデータ数を持たないために可変である
https://go.dev/play/p/nblFiIojR9V


マップ
キーと値は別々の型を使うことができる、ただしキーは比較演算子で比較できる型でなければならない
https://go.dev/play/p/LNtBCLF0XJy

### ユーザ定義型 (type definitions)


ユーザ定義型

```go
type MyDuration time.Duration
                // ↑ 基底型 (underlying type)
```

https://go.dev/play/p/5y7QES8P83D


##  0.1.2 スライスの操作


https://go.dev/play/p/fVsbIKFjQ4L

##  0.1.3 マップの操作
##  0.1.4 構造体の設計
##  0.1.5 メソッドの呼び出し
##  0.1.6 インタフェースの基礎
##  0.1.7 インタフェースの設計
##  0.1.8 型情報を利用した処理
