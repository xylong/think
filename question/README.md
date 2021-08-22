1. defer：
    - 执行顺序：return>defer，defer>panic，panic>return
    - 链式调用
    - 循环调用