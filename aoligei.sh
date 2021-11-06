#!/bin/sh
#notice: 有些题目不是leetcode上的原题，以打开链接为准

find . -type f -name "*.go" \
    | grep -v "define" \
    | awk -F '[/|.]' '{print $3,$4,"https://leetcode-cn.com/problems/"$4"/","https://leetcode.com/problems/"$4"/"}'
