# # Homework

>参考 Hystrix 实现一个滑动窗口计数器
>以上作业，要求提交到 GitHub 上面，Week06 作业提交地址：
>https://github.com/Go-000/Go-000/issues/81
```plain
#学号: G20200607010680
#班级: 4班
#作业链接: https://github.com/hi20160616/Go-000/tree/main/Week06/homework
```
# # Hystrix

## 2.1. 滑动窗口

At 40'00":[https://u.geekbang.org/lesson/68?article=326838](https://u.geekbang.org/lesson/68?article=326838)

### 2.1.1. Reference

* TerryMao suggest
    * RollingCounter:[https://github.com/go-kratos/kratos/blob/master/pkg/stat/metric/rolling_counter.go](https://github.com/go-kratos/kratos/blob/master/pkg/stat/metric/rolling_counter.go)
    * 限流算法:"github.com/go-kratos/kratos/pkg/ratelimit/bbr"
        * [https://github.com/go-kratos/kratos/tree/master/pkg/ratelimit/bbr](https://github.com/go-kratos/kratos/tree/master/pkg/ratelimit/bbr)
* So simple and easy to understand explain:[https://blog.csdn.net/renhui1993/article/details/72123455?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-4.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-4.control](https://blog.csdn.net/renhui1993/article/details/72123455?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-4.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-4.control)
* [http://timd.cn/sliding-window/](http://timd.cn/sliding-window/)
* [https://github.com/Netflix/Hystrix/wiki/How-it-Works](https://github.com/Netflix/Hystrix/wiki/How-it-Works)

![图片](https://uploader.shimo.im/f/1rP9zoqwKzrg3Uye.png!thumbnail?fileGuid=3RVWY8CqhKcwrYG9)

### 2.1.2. Slide window

1. Hystrix通过滑动窗口来对数据进行“平滑”统计
2. 默认情况下
    1. 一个滑动窗口包含10个桶（Bucket）
    2. 每个桶时间宽度是1秒，负责1秒的数据统计
    3. *滑动窗口包含的总时间以及其中的桶数量都是可以配置的*

![图片](https://uploader.shimo.im/f/4qEU7jVd8zMoHarR.png!thumbnail?fileGuid=3RVWY8CqhKcwrYG9)

* 上图的每个小矩形代表一个桶
* 每个桶都记录着1秒内的四个指标数据
    * 成功量、失败量、超时量和拒绝量
* 10个桶合起来是一个完整的滑动窗口
* 计算一个滑动窗口的总数据需要将10个桶的数据加起来
3. 滑动窗口和桶的设计特别讲究技巧，需要尽可能做到性能、数据准确性两方面的极致.
4. 桶的数据统计简单来说可以分为两类:
    1. 简单自增计数器，比如请求量、错误量等
    2. 并发最大值，比如一段时间内的最大并发量
5. Hystrix对不同的事件使用不同的数组index（即枚举的顺序），这样对于某个桶（即某一秒）的指定类型的数据，总能从数组中找到对应的简单自增或最大并发值对象来进行自增或更新操作
6. 滑动窗口由多个桶组成，业界一般的做法是将数组做成环，Hystrix中也类似

PS: 吴恩达机器学习经典名课


