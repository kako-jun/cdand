[English](https://github.com/kako-jun/cdand)

# :file_folder: cdand

[![Build Status](https://travis-ci.org/kako-jun/cdand.svg?branch=master)](https://travis-ci.org/kako-jun/cdand)

`cdand` は、シンプルなコマンドラインツールです

カレントディレクトリにしか効果の及ばないコマンド（`git`、`yarn` など）を、カレントディレクトリを変更することなく実行します

Goで書かれているため、多くのOSで動作します

　

## Description

### Demo

![demo](https://raw.githubusercontent.com/kako-jun/cdand/master/assets/screen_1.gif)

### VS.

「カレントディレクトリを移動してからでなければ、実行できないコマンド」は多くあります

例えば、`pwd`

```sh
$ pwd

/media/removable/SD Card/current

$ pwd subdir/subsubdir

/media/removable/SD Card/current
```

　

`cdand` を使うと、簡潔に書けます

```sh
$ cdand subdir/subsubdir pwd

/media/removable/SD Card/current/subdir/subsubdir
```

使わない場合、カレントディレクトリが変化してしまうため、戻るコマンドが必要になります

```sh
$ cd subdir/subsubdir; pwd
$ cd ../..
```

`cdand` にPATHを通すことで、`cd ../..` にお別れできます

　

## Installation

### Requirements

- Operating System

    - Windows
    - macOS
    - Linux

### Download binaries

- Windows: [cdand.zip](https://github.com/kako-jun/cdand/raw/master/bin/windows/cdand.zip)
- macOS: [cdand.dmg](https://github.com/kako-jun/cdand/raw/master/bin/mac/cdand.dmg)
- Linux ( `chmod u+x cdand` required)

    - x64: [cdand_amd64.tar.gz](https://github.com/kako-jun/cdand/raw/master/bin/linux/cdand_amd64.tar.gz)
    - ARM: [cdand_arm64.tar.gz](https://github.com/kako-jun/cdand/raw/master/bin/linux/cdand_arm64.tar.gz)
    - Raspberry Pi: [cdand_armv7l.tar.gz](https://github.com/kako-jun/cdand/raw/master/bin/linux/cdand_armv7l.tar.gz)

### go get

```sh
$ go get github.com/kako-jun/cdand
```

　

## Features

### Usage

```sh
$ cdand subdir/subsubdir ls -alF

drwxrwxr-x.  2 kako-jun kako-jun 4096  Apr  2 04:20 ./
drwxr-xr-x. 10 kako-jun kako-jun 4096  Apr  2 04:20 ../
-rw-rw-r--.  1 kako-jun kako-jun    0  Apr  2 04:20 my_secrets.txt
```

　

「なぜこれが便利なのか……？」の例を、以下に挙げます

#### Examples

##### e.g. Gitリポジトリのディレクトリに `cd` する必要がない

```sh
$ cdand your/git/repository git status
```

##### e.g. Nodeプロジェクトのディレクトリに `cd` する必要がない

```sh
$ cdand your/node/project yarn
```

　

つまり、実行ディレクトリを変えるオプションを、コマンドごとに覚える必要がなくなります

カレントディレクトリ以外で実行するオプションとして、

- `git` には `-C` オプションが
- `node` には `-prefix`、`-cwd` オプションが

それぞれあります

でも覚えにくいです

　

```sh
$ (cd subdir/subsubdir; ls)
```

という書き方もあります

でも面倒です

　

`cdand` ならば `cd` と打ち、「あ……めんどいな……」と思ったら `and` と打ち足せばイイだけです

その後、ディレクトリ名を打つ時には `cd` と同じくTABでの補完が効くため、違和感なく高速に打てます

```sh
$ cdand ../sister_project npm install
```

親方向に使っても便利

　

##### e.g. おまけ

`cdand` 自身を入れ子にして呼べます

```sh
$ cdand subdir cdand .. cdand subdir cdand .. ls
```

特にメリットはありませんが、可能です

　

#### Unsupported

##### 端末の文字色変更の効果は失われる

`ls -G` でも色はつきません

##### 1コマンドで完結しないコマンドは呼べない

- `vim` は起動しますが、端末に表示されません
- `less` は起動しますが、ページ送りはできません
- 対話型コマンドも対話できず、結果がまとめて表示されます

##### パイプ、リダイレクトの対象は、`cdand` コマンド自身である

例えば、

```sh
$ cdand subdir cat my_secrets.txt | grep treasure
```

と書いた場合、

```sh
$ cdand subdir cat my_secrets.txt
```

した結果をカレントディレクトリで

```sh
$ grep treasure
```

するという意味になります

パイプの場合、それでも特に結果は変わらないでしょう

　

しかし、リダイレクトの場合、結果が変わります

```sh
$ cdand subdir cat my_secrets.txt > my_will.txt
```

を実行すると、`my_will.txt` が作られるのは `subdir` 内でなく、カレントディレクトリ内です

　

どうしても `subdir` 内に作りたい場合は

```sh
cat my_secrets.txt > my_will.txt
```

という内容の `foo.sh` ファイルを作って `subdir` 内に置き、

```sh
$ cdand subdir foo.sh
```

とすれば可能です
（あまり意味はない気もしますが……）

　

### Coding

```golang
import "github.com/kako-jun/cdand/cdand-core"

cdand.Exec("subdir/subsubdir", "ls", "-alF")
```

### Contributing

Pull Requestを歓迎します

- `cdand` をより便利にする機能の追加
- より洗練されたGoでの書き方
- バグの発見、修正
- もっと良い英訳、日本語訳があると教えたい

など、アイデアを教えてください

　

## Authors

kako-jun

- :octocat: https://github.com/kako-jun
- :notebook: https://gist.github.com/kako-jun
- :house: https://llll-ll.com
- :bird: https://twitter.com/kako_jun_42

### :lemon: Lemonade stand

寄付を頂けたら、少し高い猫エサを買おうと思います

下のリンクから、Amazonギフト券（Eメールタイプ）を送ってください

「受取人」欄には `kako.hydrajin@gmail.com` と入力してください

　**[:hearts: Donate](https://www.amazon.co.jp/Amazon%E3%82%AE%E3%83%95%E3%83%88%E5%88%B8-1_JP_Email-Amazon%E3%82%AE%E3%83%95%E3%83%88%E5%88%B8-E%E3%83%A1%E3%83%BC%E3%83%AB%E3%82%BF%E3%82%A4%E3%83%97-Amazon%E3%83%99%E3%83%BC%E3%82%B7%E3%83%83%E3%82%AF/dp/B004N3APGO/)**

- 「メッセージ」欄を使って、感想を伝えることもできます
- 送り主が誰かは分かりません
- ¥15 から送れます

　

## License

This project is licensed under the MIT License.

See the [LICENSE](https://github.com/kako-jun/cdand/blob/master/LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/)
- and you
