- sum(求和)
- array_cross_merge(交叉合并数组)
- crawler(简单网页爬虫，defer、panic、select、goto)
- read_line(多协程利用互斥锁按顺序、按行读取文件)
- groutine_num(协程数量)
    - 限制协程数量基本方式
    - 分批、周期性执行
- produce_consume(生产消费模式)
- go_func(封装协程和channel，即简单任务模式)


```mermaid
    graph LR
    A[生产者] -->B(channel)
        B --> C[消费者]
```      