# 百宝箱 — 多功能开发者工具集微信小程序

## 项目概述

「百宝箱」是一款集成多种高频开发者/实用工具的微信小程序，后端使用 Golang (Gin) + SQLite，前端使用 uni-app (Vue3)。

## 技术栈

| 层 | 技术 | 说明 |
|---|------|------|
| 前端 | uni-app + Vue3 + Vite | 一套代码出小程序+H5 |
| 后端 | Go 1.21+ / Gin | 轻量高性能 REST API |
| 数据库 | SQLite3 (go-sqlite3) | 零配置，单文件数据库 |
| 部署 | 用户自备服务器 | Nginx 反代 + systemd |

## 功能清单（10 个工具）

| # | 工具 | 路由 | API 端点 | 优先级 |
|---|------|------|----------|--------|
| 1 | JSON 格式化/校验 | /pages/json-formatter | POST /api/tools/json-format | P0 |
| 2 | Base64 编解码 | /pages/base64 | POST /api/tools/base64 | P0 |
| 3 | 时间戳转换 | /pages/timestamp | POST /api/tools/timestamp | P0 |
| 4 | 二维码生成 | /pages/qrcode | POST /api/tools/qrcode | P1 |
| 5 | 颜色转换 | /pages/color | POST /api/tools/color | P1 |
| 6 | 文本处理 | /pages/text | POST /api/tools/text | P1 |
| 7 | 单位换算 | /pages/unit | 纯前端计算 | P1 |
| 8 | UUID 生成 | /pages/uuid | POST /api/tools/uuid | P2 |
| 9 | 正则测试 | /pages/regex | 纯前端计算 | P2 |
| 10 | Markdown 预览 | /pages/markdown | 纯前端渲染 | P2 |

## 项目结构

```
baibaoxiang/
├── PLAN.md              # 本文件
├── PROGRESS.md          # 进度追踪
├── docs/                # 文档
├── backend/             # Go 后端
│   ├── go.mod
│   ├── go.sum
│   ├── main.go          # 入口
│   ├── config/          # 配置
│   ├── database/        # SQLite 初始化
│   ├── models/          # 数据模型
│   ├── handlers/        # 路由处理器
│   │   ├── json_handler.go
│   │   ├── base64_handler.go
│   │   ├── timestamp_handler.go
│   │   ├── qrcode_handler.go
│   │   ├── color_handler.go
│   │   ├── text_handler.go
│   │   ├── uuid_handler.go
│   │   └── history_handler.go
│   ├── middleware/       # 中间件
│   │   ├── cors.go
│   │   └── logger.go
│   └── utils/           # 工具函数
└── frontend/            # uni-app 前端
    ├── package.json
    ├── vite.config.js
    ├── src/
    │   ├── App.vue
    │   ├── main.js
    │   ├── manifest.json
    │   ├── pages.json     # 路由配置
    │   ├── api/           # API 调用封装
    │   │   └── index.js
    │   ├── pages/         # 页面
    │   │   ├── index/         # 首页工具列表
    │   │   ├── json-formatter/
    │   │   ├── base64/
    │   │   ├── timestamp/
    │   │   ├── qrcode/
    │   │   ├── color/
    │   │   ├── text/
    │   │   ├── unit/
    │   │   ├── uuid/
    │   │   ├── regex/
    │   │   └── markdown/
    │   ├── components/    # 公共组件
    │   │   ├── ToolCard.vue
    │   │   ├── NavBar.vue
    │   │   └── ResultBox.vue
    │   ├── utils/         # 工具函数
    │   └── static/        # 静态资源
    └── uni_modules/       # uni-app 插件
```

## 每日开发任务（14天）

### Day 1 — Go 后端骨架搭建
- [ ] 初始化 Go module (`go mod init baibaoxiang`)
- [ ] 安装依赖：gin, go-sqlite3, go-qrcode, uuid
- [ ] 编写 main.go：启动 Gin 服务，注册路由
- [ ] 编写 database/sqlite.go：SQLite 初始化 + 建表（tool_usage 历史记录表）
- [ ] 编写 middleware/cors.go + logger.go
- [ ] 编写健康检查接口 GET /api/ping
- [ ] 测试：启动服务，curl /api/ping 返回 200

### Day 2 — uni-app 前端骨架 + 首页
- [ ] 初始化 uni-app Vue3 项目（Vite）
- [ ] 配置 pages.json：首页路由 + tabBar（可选）
- [ ] 编写首页 pages/index/index.vue：工具卡片网格布局
- [ ] 编写 components/ToolCard.vue：图标+名称+描述卡片
- [ ] 编写 api/index.js：封装 uni.request 的 GET/POST 方法
- [ ] 配置 manifest.json：AppID、服务器地址
- [ ] 测试：小程序模拟器能显示首页工具列表

### Day 3 — JSON 格式化工具（前后端联调）
- [ ] 后端：handlers/json_handler.go
  - POST /api/tools/json-format
  - 功能：美化(pretty)、压缩(minify)、校验(valid)
  - 返回格式化后的 JSON 或错误信息
- [ ] 前端：pages/json-formatter/index.vue
  - 输入框（粘贴 JSON）
  - 操作按钮：美化 / 压缩 / 校验
  - 结果展示区（带语法高亮或等宽字体）
  - 一键复制
- [ ] 记录使用历史到 SQLite
- [ ] 测试：各种 JSON 输入，包括非法 JSON

### Day 4 — Base64 编解码 + 时间戳转换
- [ ] 后端：handlers/base64_handler.go
  - POST /api/tools/base64
  - 功能：encode / decode，支持文本
- [ ] 后端：handlers/timestamp_handler.go
  - POST /api/tools/timestamp
  - 功能：当前时间戳、时间戳→日期、日期→时间戳
  - 支持秒级和毫秒级
- [ ] 前端：pages/base64/index.vue（输入、编码/解码切换、复制）
- [ ] 前端：pages/timestamp/index.vue（时间戳输入、日期选择器、实时当前时间戳）
- [ ] 测试

### Day 5 — 二维码生成
- [ ] 后端：handlers/qrcode_handler.go
  - POST /api/tools/qrcode
  - 输入文本/URL，生成二维码图片（PNG base64 或直接返回图片流）
  - 使用 github.com/skip2/go-qrcode
- [ ] 前端：pages/qrcode/index.vue
  - 输入框 + 生成按钮
  - 二维码图片展示
  - 长按保存到相册
- [ ] 测试：各种长度的文本、URL

### Day 6 — 颜色转换
- [ ] 后端：handlers/color_handler.go
  - POST /api/tools/color
  - 功能：HEX ↔ RGB ↔ HSL 互转
  - 返回所有格式 + 预览色值
- [ ] 前端：pages/color/index.vue
  - 输入任意格式颜色值
  - 自动识别并转换为其他格式
  - 颜色预览块
  - 拾色器（可选，用 uni-app picker 或手动实现）
- [ ] 测试：边界值（纯黑、纯白、透明）

### Day 7 — 文本处理
- [ ] 后端：handlers/text_handler.go
  - POST /api/tools/text
  - 功能：字数统计、行数统计、大小写转换、去重、排序（按行）
- [ ] 前端：pages/text/index.vue
  - 大文本输入框
  - Tab 切换不同功能
  - 实时统计（字数、字符数、行数、单词数）
  - 处理结果 + 复制
- [ ] 测试：中英文混合、空文本、超长文本

### Day 8-9 — 单位换算
- [ ] 前端纯计算，不需要后端 API
- [ ] 前端：pages/unit/index.vue
  - 分类 Tab：长度、重量、温度、面积、体积、速度
  - 每个分类下的单位列表
  - 输入数值，实时换算所有单位
  - 换算公式参考
- [ ] 数据：编写 utils/unitData.js，包含所有单位和换算系数
- [ ] 测试：各类别边界值

### Day 10 — UUID 生成
- [ ] 后端：handlers/uuid_handler.go
  - POST /api/tools/uuid
  - 功能：生成 v4 UUID，支持批量（1-100个）
  - 使用 github.com/google/uuid
- [ ] 前端：pages/uuid/index.vue
  - 数量选择（1-100）
  - 格式选项（带横线/不带/大写/小写）
  - 一键生成 + 全部复制
- [ ] 测试

### Day 11-12 — 正则测试
- [ ] 前端纯计算（JavaScript 正则引擎）
- [ ] 前端：pages/regex/index.vue
  - 正则表达式输入框
  - 修饰符选择（g/i/m/s）
  - 测试文本输入
  - 匹配结果高亮展示
  - 匹配详情（分组捕获）
  - 常用正则模板库（手机号、邮箱、身份证、URL等）
- [ ] 测试：复杂正则、分组、零宽断言

### Day 13 — Markdown 预览
- [ ] 前端纯渲染
- [ ] 集成 markdown-it 或 towxml（小程序 markdown 渲染）
- [ ] 前端：pages/markdown/index.vue
  - 左侧/上方：Markdown 编辑区
  - 右侧/下方：实时预览区
  - 常用语法快捷工具栏
  - 复制源码 / 复制 HTML
- [ ] 测试：标题、列表、代码块、表格、图片

### Day 14 — 收尾优化 + 上线准备
- [ ] UI 统一美化：配色、字体、间距一致性
- [ ] 历史记录页面：查看使用过的工具和输入
- [ ] 首页增加最近使用、热门工具排序
- [ ] 小程序隐私协议、用户协议页面
- [ ] 错误处理：网络异常、服务不可用的友好提示
- [ ] 小程序体验版测试
- [ ] 准备提交审核材料（截图、描述、类目）

## API 规范

### 请求格式
```
POST /api/tools/{tool}
Content-Type: application/json
Body: { "action": "xxx", "input": "xxx", ... }
```

### 响应格式
```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

错误响应：
```json
{
  "code": 400,
  "message": "Invalid JSON input",
  "data": null
}
```

### 历史记录表
```sql
CREATE TABLE IF NOT EXISTS tool_usage (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tool_name TEXT NOT NULL,
    input_summary TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```
