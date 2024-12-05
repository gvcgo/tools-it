## It-tools是什么？

是一个基于go wails开发的小工具，主要用于程序员的日常工作中，提升工作效率。
目前只有一个功能，就是可视化的代码统计(CLOC)，快速的统计项目中的各种代码，了解代码的分布情况。

## 如何使用？

去wails[官网](https://wails.io/zh-Hans/)查看文档，安装好wails，然后拉取it-tools项目，使用wails build命令编译即可得到可执行文件。

以winddows下的编译为例：

```bash
 wails build -f -ldflags="-s -w"  -platform linux -o it-tools.exe -webview2 embed
```

可以在build/bin/目录下找到可执行文件it-tools.exe。

## 如何添加新的功能？

你需要了解Vue3, ElementUI-Plus, tailwindcss, go等。前端代码在frontedn中，后端代码在backend中。
前端基本布局已做好，新功能只需要在侧边栏新增item，新功能页面放在src/views/your_new_folder下即可。

## 截图展示

![主界面](https://github.com/moqsien/tools-it/blob/main/docs/it_tools_welcome.png)

![代码统计查询参数](https://github.com/moqsien/tools-it/blob/main/docs/cloc_query.png)

![代码统计结果](https://github.com/moqsien/tools-it/blob/main/docs/cloc_result.png)
