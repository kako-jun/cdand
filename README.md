[日本語はこっち (Let's try reading in Japanese.)](https://github.com/kako-jun/cdand/blob/master/README_ja.md)

# :file_folder: cdand

[![Build Status](https://travis-ci.org/kako-jun/cdand.svg?branch=master)](https://travis-ci.org/kako-jun/cdand)

`cdand` is a simple CLI command.

It can launch any commands(e.g. `git`, `yarn`) in your target directory without `cd`.

It's written in GOlang, so it will work on many operating systems.

　

## Description

### Demo

![demo](https://raw.githubusercontent.com/kako-jun/cdand/master/assets/screen_1.gif)

### VS.

There are many commands that expect you already changed the current directory before launch.

For example, `pwd`.

```sh
$ pwd

/media/removable/SD Card/current

$ pwd subdir/subsubdir

/media/removable/SD Card/current
```

　

By using `cdand`, you can type simply.

```sh
$ cdand subdir/subsubdir pwd

/media/removable/SD Card/current/subdir/subsubdir
```

If you don't use `cdand`, there is no way to get expected results without changing the current directory.

And one more command will be required to restore the current directory.

```sh
$ cd subdir/subsubdir; pwd
$ cd ../..
```

By adding `cdand` to PATH, you can say good-bye to `cd ../..`.

　

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

　

"Why is `cdand` useful?" Here are the samples.

#### Examples

##### e.g. No need to call `cd` before executing Git commands.

```sh
$ cdand your/git/repository git status
```

##### e.g. No need to call `cd` before executing Node commands.

```sh
$ cdand your/node/project yarn
```

　

In other words, you don't need to care the difference between options of each command.

e.g.

- `-C` option of `git`
- `-prefix`, `-cwd` option of `node`

They are hard to remember, aren't they?

　

How is following?

```sh
$ (cd subdir/subsubdir; ls)
```

It's a cool way for who likes typing many brackets.

　

After you typed `cd`, if you saw a troublesome future, let's append `and`.

After that, use TAB key to auto complete the target directory name.

```sh
$ cdand .. git pull
```

It's also useful to keep indoors(subdirectory).

　

##### e.g. appendix

`cdand` can be called recursively.

```sh
$ cdand subdir cdand .. ls
```

Ummmm...are there any merits? (Probably nothing.)

　

#### Unsupported

##### Coloring on stdout will be disabled.

e.g. `ls -G`

##### For incomplete commands, you may get unexpected results.

`vim` can launch, but not shown in terminal.

`less` can launch, but cannnot paging.

The stdout from interactive command is shown at the same time.

##### The target of pipe and redirect is `cdand` itself.

For example, in case you type as below

```sh
$ cdand subdir cat my_secrets.txt | grep treasure
```

At first,

```sh
$ cdand subdir cat my_secrets.txt
```

runs.

Next,

```sh
$ grep treasure
```

runs for the first result.

In pipe case, the above order may not be a problem.

　

But in redirect case, a problem occurs.

```sh
$ cdand subdir cat my_secrets.txt > my_will.txt
```

`my_will.txt` isn't created in `subdir` but in the current directory.

　

If you want to create it in `subdir`, you need to create a script as below.

```sh
cat my_secrets.txt > my_will.txt
```

And put the above script in `subdir`.

Lastly,

```sh
$ cdand subdir foo.sh
```

But it probably has no benefit... :sob:

　

### Coding

```golang
import "github.com/kako-jun/cdand/cdand-core"

cdand.Exec("subdir/subsubdir", "ls", "-alF")
```

### Contributing

I'm always looking for new contributing.

- Adding features
- Overwriting with better code
- Finding (and fixing) bugs
- Improving translation

　

## Authors

kako-jun

- :octocat: https://github.com/kako-jun
- :notebook: https://gist.github.com/kako-jun
- :house: https://llll-ll.com
- :bird: https://twitter.com/kako_jun_42

### :lemon: Lemonade stand

Please open the following link, and enter `kako.hydrajin@gmail.com` into 'To' box.

　**[:hearts: Donate](https://www.amazon.com/Amazon-Amazon-com-eGift-Cards/dp/BT00DC6QU4)**

- Tell me your impressions in 'Message' box.
- The sender is hidden.
- From $1

　

## License

This project is licensed under the MIT License.

See the [LICENSE](https://github.com/kako-jun/cdand/blob/master/LICENSE) file for details.

## Acknowledgments

- [Go](https://golang.org/)
- and you
