<script setup lang="ts">
import { onLaunch, onShow, onHide } from "@dcloudio/uni-app";

onLaunch(() => {
  console.log("百宝箱 App Launch");
});

onShow(() => {
  console.log("百宝箱 App Show");
});

onHide(() => {
  console.log("百宝箱 App Hide");
});
</script>

<style>
/* 全局样式 */
page {
  background-color: #F5F7FA;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC',
    'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial,
    sans-serif;
  font-size: 28rpx;
  color: #333;
  line-height: 1.6;
}

/* 全局工具类 */
.container {
  padding: 24rpx;
}

.card {
  background: #fff;
  border-radius: 16rpx;
  padding: 32rpx;
  margin-bottom: 24rpx;
  box-shadow: 0 2rpx 12rpx rgba(0, 0, 0, 0.06);
}

.btn-primary {
  background: #4A90D9;
  color: #fff;
  border: none;
  border-radius: 12rpx;
  padding: 20rpx 40rpx;
  font-size: 28rpx;
  text-align: center;
}

.btn-primary:active {
  background: #3A7BC8;
}

.text-secondary {
  color: #999;
  font-size: 24rpx;
}

.divider {
  height: 1rpx;
  background: #EBEEF5;
  margin: 24rpx 0;
}

/* 安全区域适配 */
.safe-bottom {
  padding-bottom: env(safe-area-inset-bottom);
}
</style>
