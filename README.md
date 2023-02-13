# Streaming Practice

- [Streaming Practice](#streaming-practice)
  - [Goal](#goal)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Quick Start](#quick-start)
  - [References](#references)

---

## Goal

- 熟悉 Stream processing
- POC 練習

---

## Getting Started

### Prerequisites

- Go v1.18 ~
- OBS (Open Broadcaster Software)
- ffmpeg

### Quick Start

1. 啟動 Streaming server
```sh
go run main.go demo
```

2. 啟動 OBS 向本機的 Streaming server 推送直播畫面
```
server: http://localhost/live
token: test
```

3. 透過 ffplay 拉取直播畫面
```sh
ffplay rtmp://localhost/live/test
```

---

## References

- [[Doc] Monibuca Streaming Framework](https://m7s.live/)
- [[Doc] Homebrew ffmpeg](https://formulae.brew.sh/formula/ffmpeg)