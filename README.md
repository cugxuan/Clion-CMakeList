参见博客：
https://blog.cugxuan.cn/2019/11/01/Software/multi-main-func-run-in-clion/


在写 C++ 的题目的时候经常会遇到这样的问题，写了多个 cpp 文件，在 clion 中编译报错不能同时存在多 main 函数

~~这里写了一个小程序优雅地解决这个问，非常简单，就是就是读字符串写文件~~

直接修改一下 CMake 自动遍历文件即可

# 多个 main 函数的报错

在 clion 中写完一题，想写下一题，结果发现 main 函数不能运行

![none-main](http://image.cugxuan.cn/Software/clion/none-main.png)

# 正确的解决方法

在牛客用户[AAnonymous](https://www.nowcoder.com/profile/214695)的告知下，直接修改 CMake 即可，自己写了一个傻傻的方法。在后面加入一段即可，以我的 Project 的 CMakeList.txt 为例

```
cmake_minimum_required(VERSION 3.15)
project(JZ_offer)

set(CMAKE_CXX_STANDARD 14)

# 遍历项目根目录下所有的 .cpp 文件
file (GLOB files *.cpp)
foreach (file ${files})
    string(REGEX REPLACE ".+/(.+)\\..*" "\\1" exe ${file})
    add_executable (${exe} ${file})
    message (\ \ \ \ --\ src/${exe}.cpp\ will\ be\ compiled\ to\ bin/${exe})
endforeach ()
```
