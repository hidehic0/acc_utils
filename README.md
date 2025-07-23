# acc utils

[atcoder-cli](https://github.com/Tatamo/atcoder-cli)で生成されるjsonファイルを使用して、問題を提出するページを開いたりするツールです
現在はxdg-openが使える環境でのみ動作します

# 使い方

### submitコマンド

`~/.config/acc_utils/config.toml`を設定してから利用してください
設定しないと候補も出ません

```bash
acc_utils submit [提出するファイルがあるディレクトリ]
```

なおURLはディレクトリに紐づいているので、ディレクトリ名だけを指定すればいいです

##### 設定方法

```toml
[[configs]]
file = "main.py" # ファイル名
cmd = "/bin/cat main.py | wl-copy" # コピーするコマンド デフォルトのシェルで実行されます
```

# インストール方法

[リリース](https://github.com/hidehic0/acc_utils/releases/latest)にあるバイナリを落してパスを通してある所に置いてください
