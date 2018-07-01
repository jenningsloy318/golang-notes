OpenTracing——概念与术语

OpenTracing中的概念与术语， 基本上都是从Dapper论文中提取的。包括Trace、Span、Inter-Span Reference、Annotation等术语

**Spans**: 表示定义的一个执行单元的粗细粒度，包括执行单元的开始时间和执行时长。通过把一条链路的所有Spans按照一定的规则排列，形成时间序列图。

**Operation Name**: 每个span都需要一个操作名，要求：简单、可读性高。即看到这个操作名，大概就知道这个span所在的执行单元做了什么事情。例如：可以采用rpc方法名、函数名、自定义执行单元的命名。当一个执行单元无法用简单的语言表达时，那么具体描述可以使用Tags.

**Inter-Span Reference**: 因为span代表执行单元，那么执行单元之间存在一定的是否依赖关系。如：嵌入的执行单元，两个独立的执行单元等等。

[Inter-Span Reference](https://github.com/opentracing/specification/blob/master/specification.md)中的ChildOf和FollowsFrom两个概念，我觉得还是有些模糊不清。我理解的含义如下：

1. ChildOf： 表示执行单元的嵌入，也就是说：执行单元之间有比较强的结果依赖；
2. FollowsFrom: 表示两个执行单元相对独立，不是强依赖。

官方文档：

ChildOf references: A Span may be the ChildOf a parent Span. In a ChildOf reference, the parent Span depends on the child Span in some capacity. All of the following would constitute ChildOf relationships:

1. A Span representing the server side of an RPC may be the ChildOf a Span representing the client side of that RPC
2. A Span representing a SQL insert may be the ChildOf a Span representing an ORM save method
3. Many Spans doing concurrent (perhaps distributed) work may all individually be the ChildOf a single parent Span that merges the results for all children that return within a deadline

FollowsFrom references: Some parent Spans do not depend in any way on the result of their child Spans. In these cases, we say merely that the child Span FollowsFrom the parent Span in a causal sense. There are many distinct FollowsFrom reference sub-categories, and in future versions of OpenTracing they may be distinguished more formally.

这里面的ChildOf含义：在某些程度上, 父级span依赖于子span。模糊概念，同时又举了三个例子：

  1. rpc所在的服务端是客户端的子级，则rpc server = ChildOf(rpc client); client -> server
  2. sql所在的服务端是客户端的子级，则sql server = ChildOf(sql client); client -> server
  3. 相互独立的各个span，是同一个span的子级。例如：下订单请求，在新增订单方法中，可能会同时又扣库存、新增订单、获取商品信息等，这些事同级的，但是与新增订单方法是一父多子关系。

第一点和第二点可以归为一类，微服务之间的调用都是ChildOf关系；在一个执行单元中的所有嵌入执行单元是FollowsFrom关系。可能在spans之间存在的FollowsFrom和ChildOf两个概念本身，没有清晰的边界定义，取决于业务开发者本身，带一些主观。官网中的一些Spans之间的关系时序图，其实没有什么意义。直接理解一点：span是一个执行单元，至于这个执行单元的粗细粒度，取决于业务需要。

**Log**: Log不能在span之间传递，它的生命周期在执行单元中，和分布式日志系统的概念完全不同，它只是做一些事件日志，例如：发生error时的错误日志，非常轻量级。用途：Trace dashboard查询和问题追踪

**Tags**： Tags不能在span之间传递，在Span操作名无法满足时，使用Tags操作存储简单数据，例如：某个orderid=123的订单，可以使用tags存储orderid: 123的标签，这样我们去定位问题时，直接使用tags搜索，就可以找到追踪到具体时间、调用时长，然后再在分布式日志系统中，根据分布式跟踪系统中得到的时间、服务和订单号，追踪业务细节。

**Baggage**: 中文：行李。它是Span之间的上下文信息携带者，通过SpanContext存储，例如：存储traceid等信息。notice：不要在把大量的信息存储在SpanContext中，因为当业务量过大时，网络传输量会爆炸式增长，会造成大的业务性能和时间消耗。造成服务严重抖动，非业务因素影响了业务的高可用。分布式跟踪系统的一个设计目标：消耗低

**Inject & Extract**: SpanContext通过Inject和Extract方法，通过指定的key注入到header头部或者通过指定的key从header取出SpanContext
OpenTracing interface

**Span interface**必须实现以下功能：

  1. Get the span's SpanContext;
  2. Finish; span执行单元结束
  3. Set Tag{Key: value}；key：string，value：string|布尔值|数字类型
  4. Add a new event log；增加一个log事件。
  5. Set a Baggage item。
  6. Get a Baggage item。

**Tracer interface**必须实现以下功能：

  1. Start a new span。notice：可以从上下文中获取SpanContext，然后利用span之间的关系，创建span。
  2. Inject a SpanContext into Carrier。
  3. Extract a SpanContext from Carrier。

Global & No-op Tracer

每一个平台的OpenTracing API库(opentracing-go, opentracing-java等)，都必须实现一个空的Tracer，No-op Tracer的实现必须不会出错，并且不会有任何副作用。这样在业务方没有指定collector服务、storage、和初始化全局tracer时，但是rpc组件，orm组件或者其他组件加入了探针。这样全局默认是No-op Tracer实例，则对业务不会有任何影响。