#!/bin/sh

echo -e "\033[35m 准备更新代码 \033[0m"
# echo -e "\033[32m -------------------------------  git status  ------------------------------- \033[0m"
# git status
# echo -e "\033[34m -------------------------------  git add .  -------------------------------- \033[0m"
git add .
# echo -e "\033[35m -------------------------------  git commit -m '先下拉=>后上传'  ------------- \033[0m"
git commit -m '优化后台'
# echo -e "\033[33m -------------------------------  git pull origin master  ------------------- \033[0m"
git pull origin master
# echo -e "\033[36m -------------------------------  git push origin master  ------------------- \033[0m"
git push origin master

echo -e "\033[35m 结束 \033[0m"

# read
