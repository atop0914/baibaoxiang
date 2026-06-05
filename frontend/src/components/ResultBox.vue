<template>
  <view class="result-box" v-if="visible">
    <view class="result-box__header">
      <text class="result-box__title">{{ title }}</text>
      <view class="result-box__copy" @click="handleCopy" v-if="content">
        <text class="result-box__copy-text">复制</text>
      </view>
    </view>
    <view class="result-box__content">
      <text class="result-box__text" :selectable="true">{{ content }}</text>
    </view>
  </view>
</template>

<script setup lang="ts">
interface Props {
  title?: string
  content?: string
  visible?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  title: '结果',
  content: '',
  visible: true
})

const handleCopy = () => {
  if (!props.content) return
  uni.setClipboardData({
    data: props.content,
    success: () => {
      uni.showToast({ title: '已复制', icon: 'success' })
    }
  })
}
</script>

<style scoped>
.result-box {
  background: #fff;
  border-radius: 12rpx;
  margin-top: 24rpx;
  overflow: hidden;
  border: 1rpx solid #EBEEF5;
}

.result-box__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16rpx 24rpx;
  background: #F5F7FA;
  border-bottom: 1rpx solid #EBEEF5;
}

.result-box__title {
  font-size: 24rpx;
  color: #666;
  font-weight: 500;
}

.result-box__copy {
  padding: 8rpx 20rpx;
  background: #4A90D9;
  border-radius: 8rpx;
}

.result-box__copy-text {
  font-size: 22rpx;
  color: #fff;
}

.result-box__content {
  padding: 24rpx;
  max-height: 600rpx;
  overflow-y: auto;
}

.result-box__text {
  font-size: 26rpx;
  color: #333;
  word-break: break-all;
  white-space: pre-wrap;
  font-family: 'Menlo', 'Monaco', 'Courier New', monospace;
  line-height: 1.6;
}
</style>
