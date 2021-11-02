# GIT学习

## 1. 全局知识
* 版本库又名仓库，英文名repository
* 版本控制只能跟踪文本文件的变动，而不能跟踪二进制文件
* HEAD为当前版本，HEAD^为上一个版本，HEAD^^为上上一个版本
* 工作区就是本地文件夹，版本库就是文件夹下的.git文件夹
* add为提交到暂存区，commit为将暂存区所有内容提交到分支
* 修复分支命名fix-----，新功能分支feature-----
* 远程仓库默认名称为origin
  
## 2. 创建本地仓库
    git init（在你想要的文件夹下）

## 3. 提交到本地仓库
    git add filename…
  	git commit -m “comment”

## 4. 查看git日志 
	git log（由近到远显示）
	git log --graph（分支合并图）
	git reflog 查看历史命令，可以用于回到未来（已回退）的版本

## 5. 查看状态 git status(文件，文件夹在工作区，暂存区的状态)
	Changes to be committed:表示已经从工作区add到暂存区的file
	Changes not staged for commit:表示工作区，暂时区都存在的file（文件或文件夹），在工作区进行修改或删除，但是没有add到暂存区
	Untracked files:表示只在工作区有的file（文件或文件夹），也就是在暂时区没有该file

## 6. 回退版本 git reset --hard commid_id（HEAD^、或者commit id）
	git reset HEAD filename 撤销暂存取的修改回到工作区
## 7. 撤销工作区的修改  git checkout -- filename
	一种是自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；
	一种是已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。

## 8. 删除版本库文件 
	1，先删除本地
	2，git rm filename
	3，git commit

## 9. 关联远程库
    git remote add origin git@server-name:path/repo-name.git（或者直接复制创建仓库后的提示）
    关联后第一次推送：git push -u origin master
    查看远程库：git remote -v
    删除远程库与本地关联：git remote rm origin
## 10. 分支操作
    查看所有分支：git branch
    切换分支：git checkout  branchname
    切换并创建分枝：git checkout -b branchname
    创建分支：git branch branchname
    创建本地分支和远程分支并关联：git checkout -b local_branch origin/local_branch
    删除本地分支：git branch -d branchname
    删除本地没有与主版本合并的分支：git branch -D branchname

## 11. 合并
    git  merge branchname（将其与当前分支合并）
    合并冲突后会在发生冲突的文件出现：
    <<<<<<< HEAD
    Creating a new branch is quick & simple.
    =======
    Creating a new branch is quick AND simple.
    >>>>>>> feature1
    Git用<<<<<<<，=======，>>>>>>>标记出不同分支的内容
    修改冲突后再次add， commit即可
    
    禁用快速合并（移动指针的合并），将两个分支合并到一个commit
	  git merge --no-ff -m “common” branchname
## 12. stash命令，紧急修复bug
    Git stash用于开发途中紧急修改某个bug
	因为没有被add或者commit的修改是会显示在各个分支里面的，我们要暂存当前未完成的开发，而在上一个版本上进行修复bug
	1. git stash（保存当前状态）
	2. 切回主分支且新建修复bug分支，修复bug，add，commit
	3. 切回主分支，合并修复bug的分支
	4. 切回开发分支，git stash pop（恢复内容，切删除stash记录）
	5.将修复同步到开发分支，git cherry-pick 4c805e2（修复bug时commit的编号）

## 13. 推送到远程分支
    git push origin local_branch(将本地分支推送到对应的远程分支上)
## 14. 拉取
    使用git pull不带参数的默认前提是本地分支已经与远程分支相关联
    如果没有关联则需指定参数：git pull <remote> <branch>
    或者将其关联：git branch --set-upstream-to=origin/<branch> local_branch

## 15. rebase操作
    把分叉的提交历史“整理”成一条直线
    git pull --rebase 以最新仓库的代码为基础，使得commit成为一条线，而不是多条线合并（本地和远程两条线合并）

## 16. 标签
    git tag tagname <commit_id>，也可以为指定的commit打标签
    git tag -a <tagname> -m "blablabla..."指定标签信息
    git tag 查看标签
    git show <tagname>查看指定标签
    git tag -d <tagname>可以删除一个本地标签；

## 17. .gitignore
    文件内放置的就是需要忽略的匹配规则
    匹配规则和linux文件匹配一样
    以斜杠“/”开头表示目录；
    以星号“*”通配多个字符；
    以问号“?”通配单个字符
    以方括号“[]”包含单个字符的匹配列表；
    以叹号“!”表示不忽略(跟踪)匹配到的文件或目录；
    如：
        # 忽略public下的所有目录及文件
        /public/*
        #不忽略/public/assets，就是特例的意思，assets文件不忽略
        !/public/assets
        # 忽略具体的文件
        index.php
        # 忽略所有的php
        *.php
        # 忽略 a.php b.php
        [ab].php

## 18. git配置文件
    当前用户的Git配置文件放在用户主目录下的一个隐藏文件.gitconfig中
