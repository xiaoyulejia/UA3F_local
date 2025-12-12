# UA3F_local

[English](README_EN.md) | 简体中文

> 本项目 fork from [SunBK201/UA3F](https://github.com/SunBK201/UA3F)

UA3F_local 仅作少量调整，修改部分日志逻辑确保兼容 Windows 和 macOS 环境。

本项目的目的是在不使用软路由或其他硬件的情况下突破校园网限制。

> ⚠️ 感谢原作者开源，本项目仅为测试性项目，请勿在大陆地区使用。

## 🚀 部署 - 公共部分

1. 下载你熟悉的 clash 客户端，这里以 [clash-verge-rev](https://github.com/clash-verge-rev/clash-verge-rev/releases) 作为演示
2. 下载对应的 UA3F_local 客户端到电脑，并且请存放在一个没有中文路径的地方
3. 启动 UA3F_local 服务，请在对应的文件夹下打开 cmd / 终端窗口输入下列命令执行：

```bash
# Windows
.\ua3f-windows-amd64.exe -m SOCKS5 -p 1080 -f "FFF"

# macOS & Linux
./ua3f -m SOCKS5 -b 127.0.0.1 -p 1080 -f "FFF"
```

### 无代理需求

1. 使用 [无代理 Clash 配置](clash/ua3f-socks5-cn.yaml) 作为配置文件，复制文件内容
2. 在 clash 中选择配置新配置文件，类型选择 local
3. 右键新配置文件选择编辑文件，将复制的内容粘贴进去
4. 在 clash 左侧代理选项卡中选择 Global-ua3f
5. 检查运行的 UA3F_local 是否有正常日志输出
6. 测试是否成功：[测试网站](http://ua.233996.xyz)，请检查服务端 UA 是否为 FFF

### 有代理需求

1. 使用 [有代理 Clash 配置](clash/ua3f-socks5-global.yaml) 作为配置文件，复制文件内容
2. 在 clash 中选择配置新配置文件，类型选择 local
3. 右键新配置文件选择编辑文件，将复制的内容粘贴进去
4. 将你原本的订阅连接填写在文件中对应的位置
5. 在 clash 左侧代理选项卡中选择 CN-ua3f，其他的代理选项按照你的需求选择
6. 检查运行的 UA3F_local 是否有正常日志输出
7. 测试是否成功：[测试网站](http://ua.233996.xyz)，请检查服务端 UA 是否为 FFF

### 手机端部署
1. clash-(左侧)设置-局域网设置  打开这个选项
2. 手机打开 wifi 配置页面-高级设置-代理模式-手动-设置 ip 为路由器为电脑下发的 ip 地址, 端口请按照 clash 设置填写, 我提供的为 7897
3. 测试是否成功：[测试网站](http://ua.233996.xyz)，请检查服务端 UA 是否为 FFF




## ⚙️ 使用方法

相关命令行启动参数：

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `-m <mode>` | 服务模式，支持 HTTP、SOCKS5、TPROXY、REDIRECT | SOCKS5 |
| `-b <bind addr>` | 自定义绑定监听地址 | 127.0.0.1 |
| `-p <port>` | 端口号 | 1080 |
| `-l <log level>` | 日志等级，可选：debug | info |
| `-x` | 重写策略，支持 GLOBAL、DIRECT、RULES | GLOBAL |
| `-f <UA>` | 自定义 UA | FFF |
| `-r <regex>` | 自定义正则匹配 User-Agent，为空表示所有 User-Agent 都会被重写 | 空 |
| `-s` | 部分替换，仅替换正则匹配到的部分 | 无 |
| `-z` | 重写规则，json string 格式，仅在 RULES 重写策略模式下生效 | 无 |
| `-o ttl,tcpts,ipid` | 启用 TTL、TCP Timestamp、IP ID 伪装功能 | 无 |

> 默认日志位置：`/var/log/ua3f.log`

### 服务模式说明

UA3F 支持 5 种不同的服务模式，各模式的特点和使用场景如下：

| 服务模式 | 工作原理 | 是否依赖 Clash 等 | 兼容性 | 性能 | 能否与 Clash 等伴生运行 |
|----------|----------|-------------------|--------|------|-------------------------|
| **HTTP** | HTTP 代理 | 是 | 高 | 低 | 能 |
| **SOCKS5** | SOCKS5 代理 | 是 | 高 | 低 | 能 |
| **TPROXY** | netfilter TPROXY | 否 | 中 | 中 | 能 |
| **REDIRECT** | netfilter REDIRECT | 否 | 中 | 中 | 能 |
| **NFQUEUE** | netfilter NFQUEUE | 否 | 低 | 高 | 能 |

### 重写策略说明

UA3F 支持 3 种不同的重写策略：

| 重写策略 | 重写行为 | 重写 Header | 适用服务模式 |
|----------|----------|-------------|--------------|
| **GLOBAL** | 所有请求均进行重写 | User-Agent | 适用于所有服务模式 |
| **DIRECT** | 不进行重写，纯转发 | 无 | 适用于所有服务模式 |
| **RULES** | 根据重写规则进行重写 | 自定义 | 适用于 HTTP/SOCKS5/TPROXY/REDIRECT |

## 📋 Clash 配置建议

详见 [Clash 配置](docs/Clash.md)

## 🙏 References & Thanks

- [UA2F](https://github.com/Zxilly/UA2F)
- [uaProxy](https://github.com/huhu415/uaProxy)
- [xmurp-ua](https://github.com/CHN-beta/xmurp-ua)
- [Clash](https://github.com/Dreamacro/clash)
- [mihomo](https://github.com/MetaCubeX/mihomo)
- [UA3F](https://github.com/SunBK201/UA3F)