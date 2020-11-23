# GoBilibiliExpHelper

![test](https://github.com/xmmmmmovo/GoBilibiliExpHelper/workflows/Test%20And%20Lint/badge.svg) [![codecov](https://codecov.io/gh/xmmmmmovo/GoBilibiliExpHelper/branch/main/graph/badge.svg?token=0OW3ZT44OU)](https://codecov.io/gh/xmmmmmovo/GoBilibiliExpHelper)

## 参考项目

**[bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)**

[metowolf](https://github.com/metowolf)/**[BilibiliHelper](https://github.com/metowolf/BilibiliHelper)**

[JunzhouLiu](https://github.com/JunzhouLiu)/**[BILIBILI-HELPER](https://github.com/JunzhouLiu/BILIBILI-HELPER)**

## 功能

- [x] **每天自动登录，获取经验**
- [ ] **每天自动观看、分享、投币视频** *（支持指定想要支持的up主，优先选择配置的up主的视频，不配置则随机选取视频）*
- [x] **每天漫画自动签到**
- [x] **每天自动直播签到，领取奖励** 
- [ ] **每天自动使用直播中心银瓜子兑换B币，避免浪费**
- [ ] **每月自动使用快过期的B币券为自己充电** 
- [ ] **每个月自动领取5张B币券和大会员权益** 
- [ ] **银瓜子兑换硬币**
- [ ] **登录模式**
- [ ] **自动化更新docker镜像**
- [ ] **分布式(消息队列，K8S)**

## 编译

### Linux

> make build
>
> ./build/main

### Windows

> go mod download
>
> go build -i -o build/main GoBilibiliExpHelper