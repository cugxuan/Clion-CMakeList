在写 C++ 的题目的时候经常会遇到这样的问题，写了多个 cpp 文件，在 clion 中编译报错不能同时存在多 main 函数
这里写了一个小程序优雅地解决这个问题

# 多个 main 函数的报错
在 clion 中写完一题，想写下一题，结果发现 main 函数不能运行

![none-main](http://image.cugxuan.cn/Software/clion/none-main.png)

# 解决方法

需要在 clion 工程下 `CMakeLists.txt` 中加入对应的 `add_executable(1.cpp 1.cpp)` 即可

我编写了一个程序读取文件然后自动添加，大家可以直接下载
- [Win](https://github.com/cugxuan/Clion-CMakeList/releases/download/1.0/Generate_Clion_Win.exe)
- [Mac](https://github.com/cugxuan/Clion-CMakeList/releases/download/1.0/Generate_Clion_Mac)

# 使用方法
终端执行或者双击打开，win 用户下载后直接打开即可
Mac 如果没有执行权限，在终端加入执行权限
`$ chmod a+x Generate_Clion_Mac`

在写完 cpp 文件后，放入 clion 项目运行即可，记得第一次打开自动 Enable CMakeLists.txt 的修改，否则可能需要 Rebuild
![generate](http://image.cugxuan.cn/Software/clion/generate.png)