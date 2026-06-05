<template>
  <view class="nav-bar" :style="{ paddingTop: statusBarHeight + 'px' }">
    <view class="nav-bar__content">
      <view class="nav-bar__left" @click="handleBack" v-if="showBack">
        <text class="nav-bar__back">‹</text>
      </view>
      <view class="nav-bar__left" v-else>
        <view style="width: 60rpx;"></view>
      </view>
      <text class="nav-bar__title">{{ title }}</text>
      <view class="nav-bar__right">
        <slot name="right"></slot>
      </view>
    </view>
  </view>
  <!-- 占位元素 -->
  <view :style="{ height: (statusBarHeight + 44) + 'px' }"></view>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Props {
  title?: string
  showBack?: boolean
}

withDefaults(defineProps<Props>(), {
  title: '百宝箱',
  showBack: true
})

const statusBarHeight = ref(0)

// 获取状态栏高度
const systemInfo = uni.getSystemInfoSync()
statusBarHeight.value = systemInfo.statusBarHeight || 0

const handleBack = () => {
  uni.navigateBack({
    fail: () => {
      uni.switchTab({ url: '/pages/index/index' })
    }
  })
}
</script>

<style scoped>
.nav-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 999;
  background: #4A90D9;
}

.nav-bar__content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 44px;
  padding: 0 16rpx;
}

.nav-bar__left {
  width: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.nav-bar__back {
  font-size: 48rpx;
  color: #fff;
  font-weight: 300;
  line-height: 1;
}

.nav-bar__title {
  font-size: 32rpx;
  font-weight: 600;
  color: #fff;
}

.nav-bar__right {
  width: 60rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
