# acc utils

[atcoder-cli](https://github.com/Tatamo/atcoder-cli)で生成されるjsonファイルを使用して、問題を提出するページを開いたりするツールです

linuxとmacで動作します

# 使い方

### submitコマンド

`~/.config/acc_utils/config.toml`を設定してから利用してください

設定しないと候補も出ません

```bash
acc_utils submit [提出するファイルがあるディレクトリ]
```

なおURLはディレクトリに紐づいているので、ディレクトリ名だけを指定すればいいです

##### 設定方法

設定ファイルに提出したいファイルのファイル名、そしてそのファイルをコピーするコマンドを設定ファイルに書き込んでください

##### 設定例

```toml ~/.config/acc_utils/config.toml
[[configs]]
file = "main.py" # ファイル名
cmd = "/bin/cat main.py | xclip" # コピーするコマンド デフォルトのシェルで実行されます

[[configs]]
file = "main.cpp"
cmd = "oj-bundle main.cpp -I ~/cpp_library | xlip" # またこのような事もできます
```

### OEISコマンド

これは設定もcontest.acc.jsonも不要です
指定されたコマンドの結果をOEISで検索できます

実行されるコマンドは整数nを受けとり1つの整数を出力する形でないといけません

```bash
acc_utils oeis [実行するコマンド(必須)] -s [初期値(任意)] -e [終了値(任意)]
```

# インストール方法

[リリース](https://github.com/hidehic0/acc_utils/releases/latest)にあるバイナリを落してパスを通してある所に置いてください

おすすめは`/usr/local/bin`です

またarchlinuxユーザーの方はAURパッケージを[作成](https://aur.archlinux.org/packages/acc_utils-bin)したのでそちらも使えます
