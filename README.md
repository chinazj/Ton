## Ton

构建在tekton基础之上的CI工具。

### tekton

- 云原生，轻量级CI工具。资源占用量极少。
- 高度灵活性，同样也足够复杂。
- 谷歌背后大厂支持。
- 支持git跨域构建。多语言构建

### Flow

Flow 里面有两个关键:

1. Task代表步骤，基本分为`build`,`Test`,`deploy`
2. Step代表Task里面的小步，`build1`,`build2`,`build3`

### 附加功能

代码点

1. 更多的资源监控
2. pprof debug
3. error 封装
4. event 事件
5. sync

功能点

1. result
2. 

