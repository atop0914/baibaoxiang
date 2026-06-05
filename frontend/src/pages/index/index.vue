<template>
  <view class="page">
    <!-- 自定义导航栏 -->
    <view class="header" :style="{ paddingTop: statusBarHeight + 'px' }">
      <view class="header__content">
        <view class="header__title-row">
          <text class="header__icon">🧰</text>
          <text class="header__title">百宝箱</text>
        </view>
        <text class="header__subtitle">多功能开发者工具集</text>
      </view>
    </view>
    <view :style="{ height: (statusBarHeight + 100) + 'px' }"></view>

    <!-- 工具列表 -->
    <view class="tool-list container">
      <view class="section-title">
        <text class="section-title__text">🔧 开发工具</text>
      </view>

      <ToolCard
        v-for="tool in tools"
        :key="tool.name"
        :icon="tool.icon"
        :name="tool.name"
        :desc="tool.desc"
        :path="tool.path"
        :bg-color="tool.bgColor"
      />
    </view>

    <!-- 底部信息 -->
    <view class="footer">
      <text class="footer__text">百宝箱 v1.0.0</text>
      <text class="footer__text">持续更新中...</text>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { onShow } from '@dcloudio/uni-app'
import ToolCard from '@/components/ToolCard.vue'

const statusBarHeight = ref(0)
const systemInfo = uni.getSystemInfoSync()
statusBarHeight.value = systemInfo.statusBarHeight || 0

interface ToolItem {
  icon: string
  name: string
  desc: string
  path: string
  bgColor: string
}

const tools = ref<ToolItem[]>([
  {
    icon: '📋',
    name: 'JSON 格式化',
    desc: '美化、压缩、校验 JSON 数据',
    path: '/pages/json-formatter/index',
    bgColor: '#E8F5E9'
  },
  {
    icon: '🔐',
    name: 'Base64 编解码',
    desc: '文本的 Base64 编码与解码',
    path: '/pages/base64/index',
    bgColor: '#E3F2FD'
  },
  {
    icon: '⏰',
    name: '时间戳转换',
    desc: '时间戳与日期格式互转',
    path: '/pages/timestamp/index',
    bgColor: '#FFF3E0'
  },
  {
    icon: '📱',
    name: '二维码生成',
    desc: '将文本或链接生成二维码',
    path: '/pages/qrcode/index',
    bgColor: '#F3E5F5'
  },
  {
    icon: '🎨',
    name: '颜色转换',
    desc: 'HEX、RGB、HSL 颜色格式互转',
    path: '/pages/color/index',
    bgColor: '#FCE4EC'
  },
  {
    icon: '📝',
    name: '文本处理',
    desc: '字数统计、大小写转换、去重排序',
    path: '/pages/text/index',
    bgColor: '#E0F7FA'
  },
  {
    icon: '📐',
    name: '单位换算',
    desc: '长度、重量、温度等单位转换',
    path: '/pages/unit/index',
    bgColor: '#F1F8E9'
  },
  {
    icon: '🆔',
    name: 'UUID 生成',
    desc: '批量生成 UUID v4',
    path: '/pages/uuid/index',
    bgColor: '#FFFDE7'
  },
  {
    icon: '🔍',
    name: '正则测试',
    desc: '正则表达式在线测试与匹配',
    path: '/pages/regex/index',
    bgColor: '#EFEBE9'
  },
  {
    icon: '📄',
    name: 'Markdown 预览',
    desc: 'Markdown 实时编辑与预览',
    path: '/pages/markdown/index',
    bgColor: '#E8EAF6'
  }
])

onShow(() => {
  console.log('首页显示')
})
</script>

<style scoped>
.page {
  min-height: 100vh;
  background: #F5F7FA;
}

.header {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999;
  background: linear-gradient(135deg, #4A90D9 0%, #357ABD 100%);
}

.header__content {
  padding: 20rpx 32rpx 24rpx;
}

.header__title-row {
  display: flex;
  align-items: center;
}

.header__icon {
  font-size: 48rpx;
  margin-right: 16rpx;
}

.header__title {
  font-size: 40rpx;
  font-weight: 700;
  color: #fff;
}

.header__subtitle {
  font-size: 24rpx;
  color: rgba(255, 255, 255, 0.8);
  margin-top: 4rpx;
}

.container {
  padding: 24rpx;
}

.section-title {
  margin-bottom: 20rpx;
}

.section-title__text {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
}

.tool-list {
  padding-bottom: 24rpx;
}

.footer {
  text-align: center;
  padding: 40rpx 0 60rpx;
}

.footer__text {
  display: block;
  font-size: 22rpx;
  color: #bbb;
  line-height: 1.8;
}
</style>
