# reserve-bitcoin-lambda

## 概要
AWS Lambda(Golang)を用いて、bitflyerのAPIをコールします。
ビットコインなどの暗号資産の積立購入を行うためのサーバーレスアプリケーションです。

以下のUdemyの講座で使用しているソースコードになります。 
https://www.udemy.com/course/bitcoin-lambda-golang/

## アーキテクチャ
![アーキテクチャ](https://raw.githubusercontent.com/Kohei-Sato-1221/reserve-bitcoin-lambda/images/architecture.png)


## 使い方
```
//ビルド
make build

//ローカルで実行
sam local invoke
```
