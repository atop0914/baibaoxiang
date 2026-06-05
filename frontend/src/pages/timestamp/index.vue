<template>
  <view class="page">
    <NavBar title="时间戳转换" :show-back="true" />
    <view class="container">
      <view class="coming-soon">
        <text class="coming-soon__icon">⏰</text>
        <text class="coming-soon__title">时间戳转换</text>
        <text class="coming-soon__desc">该功能将在 Day 4 开发实现</text>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import NavBar from '@/components/NavBar.vue'
</script>

<style scoped>
.page {
  min-height: 100vh;
  background: #F5F7FA;
}

.container {
  padding: 24rpx;
}

.coming-soon {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 120rpx 40rpx;
  background: #fff;
  border-radius: 16rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.06);
}

.coming-soon__icon {
  font-size: 96rpx;
  margin-bottom: 24rpx;
}

.coming-soon__title {
  font-size: 36rpx;
  font-weight: 600;
  color: #333;
  margin-bottom: 12rpx;
}

.coming-soon__desc {
  font-size: 26rpx;
  color: #999;
}
</style>
