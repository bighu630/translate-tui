### 效果图

!!! 新增图片识别翻译

![tui翻译效果图](https://s1.ax1x.com/2023/08/06/pPAYaWj.png)

#### 配置方法

```sh
git clone https://github.com/bighu630/translate-tui
cd translate-tui
mv config.toml.example config.toml
go build
```

#### 获取腾讯翻译KEY

参考这篇文章中[腾讯翻译API](https://blog.csdn.net/weixin_44253490/article/details/126365385)部分

拿到ABI后将对应字段复制到config.toml中

#### 快捷键

> 打开程序后程序捕捉到的快捷键

| key   | 作用               |
| ----- | ------------------ |
| enter | 翻译From框中的内容 |
| c-p   | 翻译剪贴板中的内容 |
| c-y   | 复制翻译后的内容   |
| c-i   | 图片识别翻译       |
