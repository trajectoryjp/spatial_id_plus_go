# 空間IDライブラリ

## 概要
任意の座標を空間IDに変換するライブラリです。 \
利用するためには別途、外部ライブラリのインストールが必要です(後述)。 \
提供機能は以下の通りです。
  - 任意の座標と座標を結ぶ線を中心軸とした円柱状の空間IDを取得する機能
  
空間ID仕様については[Digital Architecture Design Center 3次元空間情報基盤アーキテクチャ検討会 会議資料](https://www.ipa.go.jp/dadc/architecture/pdf/pj_report_3dspatialinfo_doc-appendix_202212_1.pdf)を参照して下さい。


# 事前インストールが必要な外部ライブラリ
外部ライブラリとしてAzul3Dを使用しています。
Azul3Dの動作の前提としてODEライブラリが必要になるため、事前にインストールが必要です。

インストール手順は下記です。

## ODEのインストール手順
ODEはC++の物理エンジンです。

[公式サイト](http://www.ode.org/)

Azul3DではODEをWrapして衝突判定に用いています。そのため、Azul3Dの前提ライブラリとしてインストールします。

1. ODEのソースを取得します。
[最新版のソース](https://bitbucket.org/odedevs/ode/downloads/ode-0.16.2.tar.gz)
1. ファイルを解凍して配置します。
1. 配置先をカレントにして下記コマンドでインストールします。
```
$ cd ode-0.16.2
$ ./configure --enable-double-precision --enable-shared
$ make
$ sudo make install
```
 - トラブルシューティング
Azul3Dのパッケージをimportしたプログラムの実行時に下記のメッセージが出た場合
```
error while loading shared libraries: libode.so.8: cannot open shared object file: No such file or directory
```
1. 「/etc/ld.so.conf」を編集し、「/usr/local/lib」をファイル末尾に追加します。
2. 下記、コマンドを実行します。
```
$ sudo /sbin/ldconfig
```


## 外部ライブラリ
- 外部ライブラリ
  - ODE
    - バージョン:0.16.2
    - 確認日:2023/3/8
    - 用途:円柱と空間ボクセルの衝突確認に使用します。
