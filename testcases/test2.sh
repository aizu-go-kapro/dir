#!/bin/sh
#expect -c を使ったテスト用のスクリプト
export HOGE="hoge"
readonly SRC="abcdefghijklmnopqrstu"
readonly ENV=${HOGE}
echo $SRC $DEST $ENV >./testcases/out.txt
expect -c "
spawn grep "abcdefghijklmnopqrstu" ./testcases/out.txt
expect \"abcdefghijklmnopqrstu hoge\"
exit
"
rm ./testcases/out.txt
