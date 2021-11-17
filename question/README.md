1. defer：
    - 执行顺序：return>defer，defer>panic，panic>return
    - 链式调用
    - 循环调用
2. take_up_time(统计多个协程耗时)
3. channel_lock(channel模拟🔒)