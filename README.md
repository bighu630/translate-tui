#### 快捷键

> 打开程序后程序捕捉到的快捷键

| key           | 作用               |
| ------------- | ------------------ |
| c-\           | 翻译From框中的内容 |
| enter         | 翻译From框中的内容 |
| c-p           | 翻译剪贴板中的内容 |
| c-y/ 鼠标右键 | 复制选中框格的内容 |
| c-d           | 图片识别翻译       |

### 效果图

![tui翻译效果图](https://s1.ax1x.com/2023/08/06/pPAYaWj.png)

#### 配置方法

```sh
git clone https://github.com/bighu630/translate-tui
cd translate-tui
mv config.toml.example config.toml
go build
```

#### 依赖

- gnome-screenshot : 截图翻译依赖,for x11
- spectacle : kde截图软件
- grim + slurp : Hyprland Sway 截图软件

#### 获取腾讯翻译KEY

参考这篇文章中[腾讯翻译API](https://blog.csdn.net/weixin_44253490/article/details/126365385)部分

拿到APIKEY后将对应字段复制到config.toml中

#### 食用建议

1 绑定一个快捷键，例如 ctrl+shift+s 打开软件

2 该软件基于tui，需要一个终端作为允许基础,可以使用st这类终端，可以指定终端大小

```sh
st -g 80x20 -A 0.7 -c float -e '/data/code/go/translate-tui/tui'
```

#### for kde

如果您使用kde，并且使用上述命令作为tui的终端,您可以在设置 -> 窗口管理 -> 窗口规则 里面添加如下规则

![窗口规则](https://i.imgur.com/IjDzrRf.jpeg)
