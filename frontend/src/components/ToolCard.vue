<template>
  <view class="tool-card" @click="handleClick">
    <view class="tool-card__icon" :style="{ background: bgColor }">
      <text class="tool-card__icon-text">{{ icon }}</text>
    </view>
    <view class="tool-card__info">
      <text class="tool-card__name">{{ name }}</text>
      <text class="tool-card__desc">{{ desc }}</text>
    </view>
    <text class="tool-card__arrow">›</text>
  </view>
</template>

<script setup lang="ts">
interface Props {
  icon: string
  name: string
  desc: string
  path: string
  bgColor?: string
}

const props = withDefaults(defineProps<Props>(), {
  bgColor: '#4A90D9'
})

const handleClick = () => {
  if (props.path) {
    uni.navigateTo({
      url: props.path
    })
  }
}
</script>

<style scoped>
.tool-card {
  display: flex;
  align-items: center;
  background: #fff;
  border-radius: 16rpx;
  padding: 28rpx 24rpx;
  margin-bottom: 20rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.05);
  transition: transform 0.15s ease;
}

.tool-card:active {
  transform: scale(0.98);
}

.tool-card__icon {
  width: 88rpx;
  height: 88rpx;
  border-radius: 20rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.tool-card__icon-text {
  font-size: 44rpx;
}

.tool-card__info {
  flex: 1;
  margin-left: 24rpx;
  display: flex;
  flex-direction: column;
}

.tool-card__name {
  font-size: 30rpx;
  font-weight: 600;
  color: #333;
  line-height: 1.4;
}

.tool-card__desc {
  font-size: 24rpx;
  color: #999;
  margin-top: 6rpx;
  line-height: 1.4;
}

.tool-card__arrow {
  font-size: 36rpx;
  color: #ccc;
  flex-shrink: 0;
  margin-left: 16rpx;
}
</style>
