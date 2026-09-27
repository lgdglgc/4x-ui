<div align="center">

# 🚀 4X-UI 控制面板

**基于官方最新稳定底座（v3.8.5+）深度定制的高性能 Xray / 多协议 Web 综合管理平台**

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="4X-UI" src="./media/3x-ui-light.png" width="300">
  </picture>
</p>

<p align="center">
  <a href="https://github.com/lgdglgc/4x-ui/releases"><img src="https://img.shields.io/github/v/release/lgdglgc/4x-ui?style=flat-square&color=2ea44f" alt="Release"></a>
  <a href="https://github.com/lgdglgc/4x-ui/actions"><img src="https://img.shields.io/github/actions/workflow/status/lgdglgc/4x-ui/release.yml.svg?style=flat-square" alt="Build Status"></a>
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat-square&logo=go" alt="Go Version"></a>
  <a href="https://www.gnu.org/licenses/gpl-3.0.html"><img src="https://img.shields.io/badge/License-GPL%20v3-blue.svg?style=flat-square" alt="License"></a>
  <img src="https://img.shields.io/badge/Base-v3.8.5%20Stable-success?style=flat-square" alt="Base">
  <img src="https://img.shields.io/badge/Frontend-React%2019%20%7C%20AntD%206%20%7C%20Vite%208-61dafb?style=flat-square" alt="Frontend Tech">
  <img src="https://img.shields.io/badge/Arch-amd64%20%7C%20arm64%20%7C%20s390x-orange?style=flat-square" alt="Architectures">
</p>

</div>

---

## 📑 目录

- [一、项目定位与核心理念](#一项目定位与核心理念)
- [二、4X-UI 专属深度定制与架构增强](#二4x-ui-专属深度定制与架构增强)
  - [2.1 Clash 订阅引擎重构与热更架构](#21-clash-订阅引擎重构与热更架构)
  - [2.2 批量入站规则导入机制](#22-批量入站规则导入机制)
  - [2.3 专有 Geo 路由分流与 AI 规则集成](#23-专有-geo-路由分流与-ai-规则集成)
  - [2.4 底层并发竞争与内存治理优化](#24-底层并发竞争与内存治理优化)
- [三、全协议生态与多进程管控模型](#三全协议生态与多进程管控模型)
  - [3.1 进程编排与托管模型](#31-进程编排与托管模型)
  - [3.2 支持的入站协议矩阵](#32-支持的入站协议矩阵)
  - [3.3 传输层与前沿安全伪装](#33-传输层与前沿安全伪装)
- [四、系统架构与数据流生命周期](#四系统架构与数据流生命周期)
  - [4.1 管理端请求生命周期](#41-管理端请求生命周期)
  - [4.2 独立订阅分发链路](#42-独立订阅分发链路)
  - [4.3 后台任务调度系统 (Cron Jobs)](#43-后台任务调度系统-cron-jobs)
  - [4.4 跨节点集群协同模型 (Runtime Abstraction)](#44-跨节点集群协同模型-runtime-abstraction)
- [五、存储层架构与双引擎支持](#五存储层架构与双引擎支持)
- [六、全维度安全风控与防滥用体系](#六全维度安全风控与防滥用体系)
- [七、快速安装与运维部署指南](#七快速安装与运维部署指南)
  - [7.1 一键安装脚本](#71-一键安装脚本)
  - [7.2 一键更新维护](#72-一键更新维护)
  - [7.3 无人值守自动化部署 (Cloud-Init)](#73-无人值守自动化部署-cloud-init)
  - [7.4 CLI 交互菜单与常用命令](#74-cli-交互菜单与常用命令)
  - [7.5 Docker 容器化编排](#75-docker-容器化编排)
- [八、核心环境变量配置规范](#八核心环境变量配置规范)
- [九、界面预览](#九界面预览)
- [十、开源协议与免责声明](#十开源协议与免责声明)

---

## 一、项目定位与核心理念

**4X-UI** 是一款专为大规模多协议网络代理节点管理、复杂流量路由与订阅分发而打造的下一代 Web 控制面板。

本项目以官方最新稳定版本 **v3.8.5** 为底座构建，彻底解决了原版在海量并发与集群协同场景下的多处数据竞争（Data Race）、锁竞争与多节点状态不一致问题。在保持强大底层协议兼容性的同时，针对国内中文网络生态与生产运维痛点进行了深度重构与定制开发：

1. **配置即代码与热更新优先**：配置变更优先采用 Xray gRPC API 差异化热加载（Hot Diff），最大程度减少核心进程重启，保障客户端长连接不中断。
2. **订阅交付体验现代化**：重构 Clash 订阅生成引擎，规范 YAML 语法与键值顺序，引入外部模板热加载与地区 Emoji 智能注入。
3. **集群化与分布式协同**：基于双向 mTLS 信道的 Master-Agent 架构，通过抽象的 `Runtime` 层统一纳管本地与多云边缘节点。
4. **极致易用与企业级运维**：提供全功能 CLI 管理工具、Fail2ban 防爆破联动、PostgreSQL 热迁移工具与无依赖 Docker 容器化方案。

> [!IMPORTANT]
> 本项目仅供网络技术学习、科研测试与个人合法运维使用，请严格遵守使用者所在国家及服务器部署所在地的法律法规。

---

## 二、4X-UI 专属深度定制与架构增强

```
┌────────────────────────────────────────────────────────────────────────┐
│                              4X-UI 核心架构                             │
├───────────────────┬───────────────────┬────────────────────────────────┤
│    Clash 订阅引擎  │    批量入站管理    │       专有 Geo 路由集成         │
│  - 模板外置热加载   │  - JSON 数组导入   │  - geosite_myai.dat (AI分流)   │
│  - 国旗 Emoji 注入 │  - 事务级校验     │  - geosite_ping.dat (测速优化) │
│  - 后缀智能净化    │  - 快速批量迁移   │  - 脚本自动拉取集成            │
├───────────────────┴───────────────────┴────────────────────────────────┤
│                    稳定内核底座 (v3.8.5+ 深度优化)                      │
│      Logger 并发锁治理  │  多节点心跳并行化  │  React 19 + AntD 6 极速前端 │
└────────────────────────────────────────────────────────────────────────┘
```

### 2.1 Clash 订阅引擎重构与热更架构

原版 3X-UI 生成的 Clash 配置往往排版混乱、字段丢失或因策略组硬编码导致客户端解析异常。4X-UI 对 `internal/sub/` 模块进行了彻底重构：

- **外部 YAML 模板热更新机制**：
  - 核心订阅服务解耦硬编码逻辑，支持直接热加载外部定制的 `default.yaml` 模板文件（位于 `/etc/x-ui/` 或子目录）。
  - 修改外部规则、策略组结构或分流逻辑时，**无需重新打包编译 Go 二进制**，即改即生效。
- **严格保持原始排版与键值顺序**（`internal/sub/clash_service.go`）：
  - 杜绝传统 YAML 库反序列化后再序列化导致的注释丢失、缩进错乱与键值字母排序问题。
  - 精准锚点注入节点列表，完整保留模板内所有注释文档、结构缩进与锚点引用。
- **智能节点美化与 Emoji 国旗自动注入**（`internal/sub/beautify.go`）：
  - 内置基于 ISO 国家代码与中文地名的智能地名识别字典。
  - 自动识别入站与客户端备注，并在节点名称前规范注入对应国家/地区的国旗 Emoji（例如 🇨🇳、🇭🇰、🇹🇼、🇯🇵、🇸🇬、🇺🇸、🇬🇧、🇩🇪、🇫🇷、🇰🇷 等）。
  - 自动正则扫描并剥离节点名称尾部多余的随机哈希、短 UUID 等无用标识符，保持订阅列表中节点名称清爽规范。
- **原生支持 Spider-X 与 Reality 扩展字段**：
  - 优化 Clash 节点配置输出结构，规范字段排布顺序，完美适配最新版 Clash Verge Rev、Mihomo 等核心对 `spider-x` 等 Reality 新特性的原生解析。
- **`proxy-groups` 智能防污染与去重**：
  - 优化策略组装配逻辑，在启用 `include-all-proxies` 或使用规则过滤时，避免向静态 `proxies` 数组硬编码重复节点，彻底避免 DIRECT 规则与本地节点污染。

### 2.2 批量入站规则导入机制

针对需要批量迁移节点或多机器协同备份的场景，4X-UI 在 `internal/web/controller/inbound.go` 中扩展了批量处理通道：
- **JSON 数组一键导入**：支持在后台直接粘贴多入站组合的 JSON 数组配置。
- **原子事务与冲突检查**：导入过程严格进行端口冲突扫描与必填项校验，确保配置批量入库过程安全可控。
- **运维提效**：解决过去上百个节点需人工逐一在 UI 界面点击创建的痛点，显著提升运维效率。

### 2.3 专有 Geo 路由分流与 AI 规则集成

为了适应现代复杂的 AI 服务（如 OpenAI、Claude、Gemini 等）访问环境：
- 内置针对大模型与主流生产力工具优化的专有规则集：**`geosite_myai.dat`** 与 **`geosite_ping.dat`**。
- `install.sh` 与 `update.sh` 脚本在部署阶段自动完成专有规则包的静默拉取与校验，无缝整合至 Xray 路由引擎中。

### 2.4 底层并发竞争与内存治理优化

基于官方 v3.8.5 的最前沿补丁：
- **日志器并发安全**：将日志器句柄改造为 `atomic.Pointer` 并引入互斥锁（Mutex）保护文件轮转（Rotate），彻底消除高并发请求下日志写入导致的 Data Race。
- **多节点并发同步**：重构多节点心跳与状态同步机制，支持高达 32 个节点同时并发调度，极大缓解大规模集群网络抖动时的调度堆积。
- **内存自适应释放**：集成针对 Go 运行时的 GOMEMLIMIT 软上限与定期操作系统内存返还机制（FreeOSMemory），避免在低配 VPS 上因 GC 延迟触发 OOM Killer。

---

## 三、全协议生态与多进程管控模型

### 3.1 进程编排与托管模型

4X-UI 主程序作为守护进程（Supervisor），负责生命周期编排与子进程管理：

| 进程名称 | 托管模块 | 说明 | 监听端口 |
| :--- | :--- | :--- | :--- |
| **4X-UI Panel** | `internal/web` | Web 核心，提供 REST/WS API、静态 SPA 服务与定时调度器 | 默认 `2053` |
| **Subscription Server** | `internal/sub` | 独立运行的高性能公共订阅服务（支持 Raw/Clash/JSON） | 可独立设置 `subPort` |
| **Xray-Core** | `internal/xray` | 主代理转发引擎，通过 gRPC API 接受动态流控与路由指令 | 依据各入站规则动态分配 |
| **mtg-multi** | `internal/mtproto` | MTProto 独立子进程，支持单进程为多客户端提供独立 FakeTLS 密钥与免断流热更 | 每个 MTProto 入站独立分配 |
| **tuic-server** | `internal/tuic` | TUIC v5 独立中继进程，配合 Go 原生 UDP Relay 监听并计量流量 | 每个 TUIC 入站独立分配 |

### 3.2 支持的入站协议矩阵

| 协议类型 | 核心能力与 4X-UI 亮点 | 适用客户端 |
| :--- | :--- | :--- |
| **VLESS** | 极简高效、无多余握手开销；全面支持 XTLS 与 REALITY 伪装，支持回落分流 | 全客户端通用 |
| **VMess** | 经典 VMess 协议，兼容 AEAD 加密，成熟稳定 | 全客户端通用 |
| **Trojan** | 伪装 HTTPS 流量，抗干扰能力极强；全面支持 gRPC / WebSocket / H2 传输 | 全客户端通用 |
| **Shadowsocks** | 经典 Shadowsocks，支持 2022 最新规范（AEAD 现代密码学算法） | 全客户端通用 |
| **AmneziaWG** | **面板内嵌用户态网络栈**（基于 gVisor），无需宿主机内核模块即可实现抗 DPI 混淆 | AmneziaWG 客户端 |
| **TUIC v5** | 基于 QUIC 协议的高性能代理，0-RTT 握手、BBR 拥塞控制，原生精确 UDP 流量计量 | Clash Verge Rev、Mihomo |
| **Hysteria2** | 基于定制 QUIC 的暴力抗丢包代理，在极恶劣网络环境下保持高吞吐量 | Sing-box、Mihomo |
| **MTProto** | 专为 Telegram 打造，支持每客户端独立 FakeTLS 密钥、广告频道标签与配额热更 | Telegram 官方客户端 |
| **WireGuard** | 原生 WireGuard 支持，具备极低资源消耗与高吞吐特征 | WireGuard 客户端 |
| **SOCKS / HTTP** | 标准局部代理协议，支持 Mixed 模式与 Dokodemo-door 端口映射转发 | 浏览器与系统代理 |
| **TUN** | 系统级虚拟网卡转发模式，实现全局流量截获与路由分流 | 网关与路由器 |

### 3.3 传输层与前沿安全伪装

- **传输载体**：TCP (Raw)、mKCP、WebSocket、gRPC、HTTPUpgrade、XHTTP。
- **安全伪装**：
  - **REALITY**：新一代借壳偷窥伪装技术，消除传统 TLS 证书特征，无需自备域名即可实现高强度抗封锁。
  - **TLS / XTLS**：支持自动 Let's Encrypt 证书签发与续期、自定义证书导入及 ALPN 多协议协商。
- **智能回落 (Fallbacks)**：支持单端口（例如 443）智能分流多种协议（如 VLESS + Trojan + Web 前端），有效隐藏服务器特征。

---

## 四、系统架构与数据流生命周期

### 4.1 管理端请求生命周期

```
浏览器 (React 19 / fetch)
  │  POST {basePath}/panel/api/inbounds/add
  ▼
Gin Engine (internal/web/web.go)
  ├── 1. 安全响应头注入 (SecurityHeaders)
  ├── 2. 请求体大小限制 (MaxBodyBytes 10 MiB)
  ├── 3. 会话校验与 CSRF 防御
  └── 4. 路由分发 -> Controller 校验 (internal/web/controller/inbound.go)
        ▼
     Service 业务处理 (internal/web/service/inbound.go)
        ├── 事务级持久化 -> GORM (internal/database)
        └── 调度抽象派发 -> runtime.Runtime (internal/web/runtime)
              ├─► 本地节点 (Local):
              │     internal/xray/hot_diff.go -> 计算差异 -> gRPC 动态注入 (无需重启核心)
              └─► 远程边缘节点 (Remote):
                    通过 mTLS 双向加密 HTTPS 接口派发至子节点
```

### 4.2 独立订阅分发链路

订阅服务部署在独立端口上运行，与管理后台完全解耦，具备极高的并发与抗压能力：
1. **客户端请求识别**：依据 HTTP 请求头 `User-Agent` 智能识别客户端类型（Clash、Sing-box、v2rayN、Shadowrocket 等）。
2. **动态配置装配**：依据客户端权限读取关联入站与主机节点，动态应用 Host 覆盖参数（SNI、路径等）。
3. **美化与规则注入**：进入 `internal/sub/clash_service.go`，执行节点名称 Emoji 国旗添加、后缀清除，并渲染至高质量 YAML 模板中。

### 4.3 后台任务调度系统 (Cron Jobs)

4X-UI 内部集成多达 **19 类** 秒级精度的异步后台调度作业（`internal/web/job/`）：
- **流量计量与统计**：每 5 秒从 Xray gRPC API 批量收集入站与客户端流量增量，写入数据库。
- **Fail2ban IP 限制巡检**：每 10 秒扫描连接活跃度，超出配额自动触发 iptables 封禁脚本。
- **多节点心跳探测**：每 5 秒轮询集群子节点状态，自动探测异常下线与自动恢复同步。
- **周期性配额重置**：按小时/按天/按月自动重置达到账单周期的客户端流量配额。
- **资源监控哨兵**：监测内存占用并定期执行内存整理（FreeOSMemory）。

### 4.4 跨节点集群协同模型 (Runtime Abstraction)

4X-UI 引入了高度抽象的 `Runtime` 接口：
- **`Local` 模式**：通过内存中的 gRPC 客户端直接与本地 Xray 交互。
- **`Remote` 模式**：Master 节点通过 HTTPS（支持 `verify`、`skip`、`pin` 及双向 `mTLS` 四种模式）与子节点通信。所有的入站创建、客户端变更均透明广播，使得多台跨国 VPS 可以由单个主面板统一编排调度。

---

## 五、存储层架构与双引擎支持

4X-UI 原生提供双数据库引擎支持，满足不同规模部署需求：

```
┌──────────────────────────────────────┐     ┌──────────────────────────────────────┐
│           SQLite (默认引擎)           │     │         PostgreSQL (集群引擎)         │
├──────────────────────────────────────┤     ├──────────────────────────────────────┤
│ • 单文件存储：/etc/x-ui/x-ui.db        │     │ • 专为海量客户端与多节点高并发打造    │
│ • 零配置，开箱即用，适合中小规模部署 │ ──► │ • 连接池管理与高可用支撑             │
│ • 极低系统资源占用                   │     │ • 完整 ACID 事务与高吞吐写入性能     │
└──────────────────────────────────────┘     └──────────────────────────────────────┘
                   │                                            ▲
                   └─────────── x-ui migrateDB 指令 ────────────┘
```

### 数据库平滑迁移

当节点规模扩大需切换至 PostgreSQL 时，无需停机导出 SQL，直接运行内置迁移工具：

```bash
x-ui migrateDB "postgres://user:password@127.0.0.1:5432/xui?sslmode=disable"
```

该命令会在不破坏现有 SQLite 数据库的前提下，自动完成全表结构同步与数据无损迁移。

---

## 六、全维度安全风控与防滥用体系

1. **客户端配额三维管控**：
   - 流量额度上限（总上行 + 下行）。
   - 服务到期时间（支持到期自动停机或周期续费）。
   - 上行/下行速率限制。
2. **防盗刷与连接数限制**：
   - **IP 连接数限制**：单客户端可绑定同时在线 IP 数量上限，联动系统级 Fail2ban 通过 iptables 实时拦截违规连接。
   - **白名单机制**：支持设置受信任的公共 IP 或特定跳板机，免除误封禁。
   - **HWID 硬件设备指纹锁**：基于客户端机器码限制在线设备数量，有效防止节点共享与账号倒卖。
3. **隧道主动健康巡检 (Tunnel Health Monitor)**：
   - 可配置定期通过指定代理向 Cloudflare 等节点发送探针。
   - 当探测失败达到预设阈值时，自动触发核心重启策略，实现无人值守的高可用容灾。

---

## 七、快速安装与运维部署指南

### 7.1 一键安装脚本

在 Linux 服务器（Debian / Ubuntu / CentOS / AlmaLinux / Rocky 等）终端中执行：

```bash
bash <(curl -Ls https://raw.githubusercontent.com/lgdglgc/4x-ui/main/install.sh)
```

- **安装指定版本**（例如 `v4.0.0`）：
  ```bash
  bash <(curl -Ls https://raw.githubusercontent.com/lgdglgc/4x-ui/main/install.sh) v4.0.0
  ```

- **安装最新 Dev 预览分支**：
  ```bash
  bash <(curl -Ls https://raw.githubusercontent.com/lgdglgc/4x-ui/main/install.sh) dev-latest
  ```

### 7.2 一键更新维护

如果已经安装 4X-UI，可随时通过更新脚本完成版本平滑升级：

```bash
bash <(curl -Ls https://raw.githubusercontent.com/lgdglgc/4x-ui/main/update.sh)
```

### 7.3 无人值守自动化部署 (Cloud-Init)

在批量部署云服务器或使用自动化配置工具（Terraform / Ansible）时，可通过设置环境变量 `XUI_NONINTERACTIVE=1` 启用全静默安装：

```bash
XUI_NONINTERACTIVE=1 bash <(curl -Ls https://raw.githubusercontent.com/lgdglgc/4x-ui/main/install.sh)
```

安装脚本会自动生成高强度的随机用户名、密码、端口与 BasePath，并将全部凭据固化输出至 `/etc/x-ui/install-result.env`。

### 7.4 CLI 交互菜单与常用命令

安装完成后，在终端直接输入 `x-ui` 即可呼出可视化命令行管理菜单：

```bash
x-ui
```

```
  4X-UI 控制面板管理脚本
  0. 退出脚本
————————————————
  1. 安装 4X-UI
  2. 更新 4X-UI
  3. 更新 4X-UI (开发版)
  4. 卸载 4X-UI
————————————————
  5. 重置用户名与密码
  6. 重置 Web 访问根路径
  7. 更改面板监听端口
  8. 查看当前配置参数
————————————————
  9. 启动 4X-UI
 10. 停止 4X-UI
 11. 重启 4X-UI (含 Xray)
 12. 仅重启 Xray 核心
 13. 查看运行状态
 14. 查看运行日志
 15. 查看 Fail2ban 封禁日志
————————————————
 16. 申请 SSL 证书 (ACME)
 17. 数据库迁移 (SQLite -> PostgreSQL)
 18. 更新 GeoIP / GeoSite 路由规则数据
```

#### 快捷执行参数表

| 指令 | 说明 |
| :--- | :--- |
| `x-ui start` | 启动面板服务 |
| `x-ui stop` | 停止面板服务 |
| `x-ui restart` | 重启面板与 Xray 核心 |
| `x-ui restart-xray` | 仅单独重启 Xray 核心 |
| `x-ui status` | 查看服务状态 |
| `x-ui settings` | 显示当前端口、BasePath 与管理员密码 |
| `x-ui log` | 查看面板运行日志 |
| `x-ui banlog` | 查看 Fail2ban IP 封禁记录 |
| `x-ui update` | 升级至最新稳定版 |
| `x-ui update-dev` | 升级至最新开发分支构建 |
| `x-ui update-all-geofiles` | 重新拉取最新的 Geo 路由规则数据 |
| `x-ui migrateDB` | 执行数据库在线平滑迁移 |

### 7.5 Docker 容器化编排

4X-UI 官方支持容器化部署。推荐使用 `docker-compose.yml`：

```yaml
services:
  4x-ui:
    image: ghcr.io/lgdglgc/4x-ui:latest
    container_name: 4x-ui
    restart: unless-stopped
    network_mode: host
    cap_add:
      - NET_ADMIN
      - NET_RAW
    volumes:
      - ./db/:/etc/x-ui/
      - ./cert/:/root/cert/
      - ./acme/:/root/.acme.sh/
      - ./logs/:/var/log/x-ui/
    environment:
      - XRAY_VMESS_AEAD_FORCED=false
      - XUI_ENABLE_FAIL2BAN=true
      - XUI_LOG_LEVEL=info
```

#### 启动服务

```bash
docker compose up -d
```

> **生产环境建议**：
> 1. 采用 `network_mode: host` 可以避免复杂的端口映射，确保 UDP（TUIC、Hysteria2）、XTLS 与 IPv6 正确监听。
> 2. `cap_add: [NET_ADMIN, NET_RAW]` 是 Fail2ban 联动宿主机防火墙执行封禁所必需的权限。

---

## 八、核心环境变量配置规范

4X-UI 支持通过环境变量全面定制运行参数（安装脚本默认写入 `/etc/default/x-ui`）：

| 环境变量 | 参数说明 | 默认值 |
| :--- | :--- | :--- |
| `XUI_PORT` | 面板监听服务端口 | `2053` |
| `XUI_INIT_WEB_BASE_PATH` | Web 访问根前缀（如 `/admin/`） | `/` |
| `XUI_DB_TYPE` | 存储引擎类型：`sqlite` 或 `postgres` | `sqlite` |
| `XUI_DB_DSN` | PostgreSQL 连接串（使用 postgres 时必填） | — |
| `XUI_DB_FOLDER` | SQLite 数据库持久化目录 | `/etc/x-ui` |
| `XUI_LOG_LEVEL` | 日志详细级别 (`debug`, `info`, `warning`, `error`) | `info` |
| `XUI_DEBUG` | 是否启用开发调试模式 | `false` |
| `XUI_ENABLE_FAIL2BAN` | 是否启用基于 Fail2ban 的 IP 限制自动封禁 | `true` |
| `XUI_GOGC` | Go 垃圾回收阈值参数（数值越低内存占用越小） | `75` |
| `XUI_MEMORY_RELEASE_INTERVAL`| 强制返还空闲物理内存给操作系统的间隔（分钟） | `10` |
| `XUI_TUNNEL_HEALTH_MONITOR` | 是否启用隧道健康嗅探与自动恢复 | `false` |
| `XUI_TUNNEL_HEALTH_URL` | 隧道健康嗅探目标探测地址 | `https://www.cloudflare.com/cdn-cgi/trace` |
| `XUI_TUNNEL_HEALTH_INTERVAL` | 探针执行间隔时间 | `30s` |

---

## 九、界面预览

<details>
<summary><b>点击展开查看 4X-UI 控制面板截图</b></summary>
<br>

<p align="center"><b>📊 系统运行状态仪表盘</b></p>
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/01-overview-dark.png">
    <img alt="系统概览" src="./media/01-overview-light.png">
  </picture>
</p>

<p align="center"><b>🌐 入站节点管理与批量导入</b></p>
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/02-inbounds-dark.png">
    <img alt="入站管理" src="./media/02-inbounds-light.png">
  </picture>
</p>

<p align="center"><b>👥 客户端配额与安全控制</b></p>
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/03-add-client-dark.png">
    <img alt="客户端设置" src="./media/03-add-client-light.png">
  </picture>
</p>

<p align="center"><b>🖥️ 集群多节点协同管理</b></p>
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/05-nodes-dark.png">
    <img alt="节点管理" src="./media/05-nodes-light.png">
  </picture>
</p>

<p align="center"><b>📑 内置 RESTful API 交互文档</b></p>
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/08-api-docs-dark.png">
    <img alt="API 文档" src="./media/08-api-docs-light.png">
  </picture>
</p>

</details>

---

## 十、开源协议与免责声明

- **开源协议**：本项目依据 [GNU General Public License v3.0 (GPL-3.0)](LICENSE) 协议进行分发。
- **技术基石**：致敬并感谢 [Xray-core](https://github.com/XTLS/Xray-core) 项目以及原版 [3X-UI](https://github.com/MHSanaei/3x-ui) 与 [X-UI](https://github.com/vaxilu/x-ui) 社区开发者的卓越付出。
- **合规免责声明**：本项目仅限用于网络通信技术研究、学术交流以及经过授权的私有基础设施运维。使用者因使用本软件引发的任何法律风险与纠纷，项目开发者及贡献者不承担任何直接或连带责任。

---

<div align="center">
  <sub>4X-UI · 致力于打造更高性能、更稳健优雅的代理控制面板</sub>
</div>
