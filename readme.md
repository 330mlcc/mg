# 常见问题

## 1.import的的github地址出现红色波浪线，提示找不到对应的地址，即网上的“go can't find import: "github.com/** 错误”错误。

    原因：goland导入的项目的位置不在gopath的工作空间的src目录下，因此idea无法工作。


### 设置gopath

    指定go项目的所有的工作空间(gopath地址，存放各种go项目工程)


## 添加git

    git add main.go
    git commit -m “first commit”
    git remote add origin https://github.com/fwhezfwhez/csdn_blog_use.git
    git push -u origin master