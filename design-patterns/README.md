### 一、开篇
1. 养成写高质量代码的习惯（什么是高质量代码）
2. 设计模式是如何写出读性、可扩展、可维护的代码

### 二、为什么要学设计模式
1. 面试
2. 职业素养
3. 提高复杂代码的设计开发能力
4. 读源码（高手也是这样玩的）

### 三、好坏代码的指标&怎么写出高质量的代码
1. 灵活性、
   1. 可扩展性、可维护性、可读性、可理解性、易修改、复用性、可测试性、模块化 、高内聚低耦合、高效、高性能、安全性、兼容性、易用性、整洁、清晰、简单、直接 少即是多、文档详细、分层清晰、正确性、鲁棒性、健壮性、可用性、可伸缩性、稳定性、优雅
2. 代码好坏的评价是主观的
3. 常用：
   1. 可维护性（修bug、调整老代码和添加新功能）：在不引起bug的情况下能够快速添加或者修改代码。
   2. 可读性：编码规范、命名、注释、函数长短、分层、模块、高内聚低耦合
   3. 可扩展性：不修改或者少修改原代码的基础上添加新的代码
   4. 灵活性：（有点可维护+可扩展的味道）
   5. 简洁性：KISS（保持简单的代码，我喜欢这个）
   6. 复用性：减少重复代码
   7. 可测试性：不好测的代码就是有问题的代码

### 四、面向对象、设计模式、设计原则、编码规范和重构
1. 面向对象的特性：有很多特性（封装、抽象、继承和多态），可以实现很多设计思路，是设计原则和设计模式的基础
2. 9 种设计原则：指导设计的经验总结，可以指导是否用某种设计模式。
3. 23 种设计模式：是针对开发过程中常见的问题总结的方案和思路，提高可扩展性
4. 编码规范：提高可读性
5. 重构：对以上4中思想的运用

### 五、什么是面向对象和面向对象语言
1. 面向对象是一种编程范式，以对象为组织代码单元，并将封装、抽象、继承、多态
2. 支持面向对象特性的语言

### 六、继承、封装和多态解决什么问题
1. 封装：访问控制权，防止任意修改，隐藏
2. 抽象：隐藏方法的具体实现，调用者只用关系有哪些方法，不关注方法实现
3. 继承：代码复用，过度继承过导致代码可读性降低，耦合度加重
4. 多态：提升代码复用

### 七、面向对象vs面向过程
1. oop 更易应对大规模软件开发
2. 更易复用、扩展、维护
3. 更智能

### 八、接口和抽象类
1. 抽象类有属性和抽象方法，出于代码复用，抽离除子类中公共的is-a
2. 接口只有抽象方法，出于解耦 契约 has-a

### 九、为什么要基于接口编程
1. 依赖抽象不理来实现
   1. 函数名不暴露任何实现细节
   2. 封装细节
   3. 为实现定义抽象接口
2. 抽象意识、封装意识、接口意识
3. 只有一种实现方法，未来不可能被替换的情况下就没必要基于接口编程
4. 越是不稳定的系统越要在扩展性和维护性上下功夫。

### 十三、面向对象分析(OOA)
1. 非业务开发要有组件化、框架化和抽象化的意识
2. 全面了解业务，全面满足需求，初步方案+反复优化
3. 使用进化算法的思想，提出一个MVP（最小可行性产品），逐步迭代改进


### 十四、面向对象设计(OOD)+面向读编程(OOP)
1. 面向对象设计
   1. 划分职责识别出有哪些类
      1. 罗列需求中的名词作为类的备选（或者根据需求把功能点列出来，再去进行职能分类，对于大规模的需求先分模块）
   2. 定义类的属性和方法
   3. 定义类与类之间的关系
      1. 泛化（继承）
      2. 实现（实现接口）
      3. 聚合（包含，但是不依赖）
      4. 组合（包含，依赖）
      5. 关联（聚合+组合）
      6. 依赖
   4. 组装类并提供入口

十五、单一职责
1. SRP内涵：不要设计大而全的类，要设计粒度小、功能单一的类。
2. 如何判断功能是否单一
   1. 这是个主观性比较强的问题
   2. 可以先写个粗粒度的类，根据业务背景再进行拆分。这就是所谓的重构
   3. 经验性技巧：
      1. 类中代码行数，函数或者属性过多
      2. 类中依赖其他类较多
      3. 私有方法过多
      4. 类名比较难定
      5. 类中大量方法集中操作那几个属性
3. 也并不是拆的约细越好

### 十六、对扩展开放、对修改关闭
1. 重点原则，23种设计模式主要围绕该原则展开，开闭原则讲的就是扩展性问题
2. 应该是在原有的功能上扩展，而不是在原有的基础上修改
3. 核心逻辑是没有破坏原代码的正常运行，并有破坏原有的单元测试
4. 多态、依赖注入、基于接口而非实现编程（装饰、策略、模板、责任链、状态）
5. 时刻有抽象、封装的意识, 预留扩展点
   1. 怎么预留扩展点？
      1. 面向业务的开发需要多了解业务
      2. 面向技术的开发需要了解你的软件将会被如何使用
      3. 预留扩展也不能到处预留，要有取舍
   2. 有些情况下扩展和可读性有冲突

### 十七、里氏替换
1. 内涵：子对象可以替换父对象，并且保证原来逻辑的不变性和正确性
2. 关键点：子类在设计的时候要遵守父类的行为协定，协定包涵输入、输出、异常的约定，甚至包涵特殊逻辑的约定。

### 十八、接口隔离原则
1. 客户端（调用者）不应该强迫依赖它不需要的接口。
   1. 一组接口组合
   2. 单个API或者函数
   3. OOP中的接口概念
2. 和单一职责模式很像，但是不一样的是一方面是接口设计，另一方面是思考角度不同，从调用者的角度去理解

### 十九、依赖反转
1. 控制反转(IoC, Inversion of control),框架提供一个可扩展的代码骨架，用来组装对象、管理整个执行流程，开发只需要再扩展点写自己的业务代码，就可以利用框架驱动整个业务
2. 控制指的是对程序执行流程的控制，反转是指从原本由程序员的自己控制整个程序的执行，使用框架后，整个流程由框架控制。
3. 依赖注册(DI, Dependency Inject)
   1. 不用new 方式在内部初始化类，而是将依赖的类再外部创建好之后通过构造函数、函数参数的方式传递给类使用
   2. 提升代码的可扩展性？？怎么提升的？？可以灵活的替换类？？所以依赖的就不能是类而是接口了？？
4. 依赖注册是编写可测试性代码的最有效手段
5. DI Framework 
6. 依赖反转原则(DIP)：
   1. 高层模块不要依赖低层模块
   2. 高层模块和低层模块应该通过抽象相互依赖
   3. 抽象不要依赖细节，细节要依赖抽象
7. 高层vs低层: 
   1. 调用者属于高层，被调用者属于低层
   2. 平时业务开发是高层模块依赖低层模块没问题
   3. 该原则主要指导框架层面

### 二十、KISS&YANGI
1. KISS：keep it simple stupid, 保持简洁
   1. 不要使用同事可能不懂的技术实现代码
   2. 不要重复造轮子，善于使用已有的工具类库
   3. 不要过渡优化
   4. code review 时同事对你的代码有疑问就说明你的代码不够"简单"，需要优化
2. YANGI: you ain't Gonna need it, 不要过渡设计
   1. 不要设计当前用不到的功能
   2. 不要设计当前用不到的代码
   3. 当前不需要就不做
   4. 做和留扩展还是不一样的，扩展还是得留
3. DRY：dont repeat your self, 不要重复
4. LoD：law of Demeter, 最小知识

### 二十一、DRY
1. 重复形式
   1. 实现逻辑重复
   2. 功能语义重复
   3. 代码执行重复
2. 代码复用性
   1. 代码复用：表示一种行为，尽量复用已存在的代码。
   2. 代码复用性：表示一段代码可以被复用的特性或能力，编写代码的时候尽量让代码可以复用。
   3. DRY原则：不要写重复的代码。
3. 提高复用性
   1. 减少耦合
   2. 满足单一职责
   3. 模块化
   4. 业务与非业务分离
   5. 通用代码下沉
   6. 继承、封装、多态
   7. 应用模板模式

### 二十二、LoD 迪米特法则
0. 很重要实用的一个原则
1. 高内聚：相近的功能放到同一个类，不相近的代码放不要放到相同的类
2. 松耦合：类与类之间的依赖关系简单、清晰。即使二者有依赖，但是一个的修改不会引起或者很少引起另一个类的修改
3. 粒度小、职责单一、结构简单、清晰
4. 不应该有直接依赖的类之间不要有依赖，有依赖关系的类之间，尽量依赖有必要的接口

### 二十三、业务系统如何做需求开发
1. 需求分析
   1. 技术应该深度到产品（DDD），具备产品思维
   2. 学会"借鉴"（看看大厂怎么做的）// 抄个基础版的
   3. 了解自身的业务需求
2. 代码设计实现
   1. 合理地将功能划分到不同模块，做到模块层面的高内聚、低耦合
   2. 设计模块与模块之间的交互关系
      1. 同步接口调用：简单直接
      2. 异步接口调用：解耦
   3. 设计模块的接口、数据库和业务模型
3. 系统上线维护
4. 系统设计
   1. 数据库
   2. 接口
   3. 业务逻辑（MVC）
5. 分层的作用
   1. 提高代码复用
   2. 隔离变化【下层变化，上层无感知】
   3. 隔离关注点【各个层只关注自身的层的变化】
   4. 提高代码可测
   5. 应对系统的复杂性
6. BO/VO/Entity
   1. 对应各自的层Service/Control/Repository
   2. 存在的意义是各个层有不同的需求，可以数据不同
   3. 基于贫血模型

### 二十三、针对非业务的通用框架开发
1. 需求分析
   1. 功能性需求分析
   2. 非功能性需求分
      1. 易用性：是否可插拔、易集成、个业务松耦合、接口是否灵活、文档是否清晰。
      2. 性能：性能消耗是否足够低，低延迟、低内存
      3. 扩展性：不修改，或者少修改添加新功能，并保障旧功能正常
      4. 容错性：所有运行时错误和非运行时错误进行捕捉
      5. 通用性：提高代码复用性
   3. TDD和最小原型
      1. 聚焦一个最简单的应用场景，设计一个简单的原型。
   4. 小步快跑、逐步迭代：学会结合具体的需求，做合理的预判、假设和取舍

### 二十四、重构
0. 定义：在不改变外部可见行为情况下的一种对软件内部结构的改善，目的是为了可读性、可扩展和可维护
1. 为什么重构WHY：
   1. 保障代码质量
   2. 系统演进的方向
   3. 避免过渡设计
   4. 初级维护、高级设计、资深重构
2. 重构的对象WHAT:
   1. 大规模
      1. 系统、模块、代码结构、类与类
      2. 工具：分层、模块化、解耦、抽象、可复用。【设计思想、设计原则和模式】
      3. 影响面大，易引入bug
   2. 小规模
      1. 函数、类：超大函数，命名规范、注释规范、提取重复代码
      2. 影响面小，引入bug的概率小，时间比较少
3. 时机WHEN：
   1. 平时没事把烂代码主动重构
   2. 做其他需求的时候随便重构
   3. 把重构作为开发的一部分，时刻有重构意识
4. 怎么重构HOW：
   1. 大型重构，做好计划，分步进行，做好兼容，
   2. 小重构：随时

### 二十五、重构不出错
1. 单元测试
   1. 集成测试是一种端到端的测试，测试功能是否可用
   2. 单元测试是对类或者函数的测试，判断是否符合预期
2. 好处：
   1. 单元测试和codeReview 保障代码质量
   2. 发现设计中的问题：基本上难以使用单元测试的代码都有问题
   3. 单测是对集成测试的补充
   4. 单测本身就是重构（设计）
   5. 阅读单测可以快速熟悉代码
   6. 单测是TDD的改进
3. 针对各种输入、异常、边界条件的测试用例，并将以上翻译成代码

### 二十六、代码可测性
1. 手段
   1. mock
   2. 依赖注入
   3. 二次封装
2. 反例
   1. 未决行为：时间、随机数
   2. 全局变量
   3. 静态方法
   4. 复杂继承
   5. 高耦合
   
### 二十七、快速解耦
1. 怎么判断是否需要解耦
   1. 看是代码会不会牵一发而动全身
   2. 画出模块和模块之间的依赖，根据依赖关系图的复杂性判断是否需要解耦
2. 如何解耦
   1. 封装、抽象
   2. 中间层
   3. 模块化
   4. 设计思想和原则【高内聚低耦合+SOLID+KISS+YANIGO+DRY+LOD】

### 二十八、编程规范
1. 命名
   1. 长度：常用的缩写、作用域较小的缩写、作用域大的尽量用长的
   2. 利用上下文简化命名：
   3. 命名可读【不要生僻】、可搜索【IDE友好】
   4. 接口和抽象类(I & abstract, golang er)
2. 注释
   1. 写什么：做什么、为什么这样做、怎么做
   2. 一般类和函数都要写注释，而且详细写比较好，函数内部怎根据命名去推测
3. 类、函数大小
   1. 函数不要超过一屏显示器
   2. 类不要让自己觉得头大
4. 行的长度：一般不要超过一行显示器
5. 善于用空行分割
6. 缩进：不要tab缩进，不同IDE tab的宽度不同
7. 大括号
8. 类中成员的顺序：
9. 把代码分割成更小的单元
10. 避免参数过多
    1. 拆成多个函数
    2. 将参数封装成对象【不要用数组当参数，最好封装成对象】
11. 不要用参数控制逻辑？？（WHY）违反了单一职责原则？？switch呢？私有函数，影响范围有限，两个函数经常被调用则酌情考量
12. 函数要单一职责
13. 移除嵌套过深的层次
14. 学会用解释性变量？？？什么叫解释性变量，逻辑更清晰的变量

### 二十二、单例
1. 单例是指在进程中唯一
2. 如果实现线程唯一？利用map
3. 集群唯一？把对象存储到外部文件，存取时加锁

### 二十四、[工厂模式](https://www.cnblogs.com/ricklz/p/15399178.html#%E7%AE%80%E5%8D%95%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8Fsimple-factory)
1. 简单工厂：if ... else ... 先初始化好，根据参数输出不同的类
2. 工厂方法：利用接口
   1. 简单工厂 vs 工厂方法：当创建的逻辑对象比较复杂，不只是简单new一下时推荐使用工厂方法，否则推荐使用简单工厂
3. 抽象工厂：
   1. 也是利用接口，让一个工厂负责创建多个不同类型的对象
4. 创建逻辑比较复杂的大工程时考虑使用工厂模式，将创建和使用分离
   1. 类似配置解析，代码中存在动态生成类
   2. 或者不需要动态生成类，但是单个对象创建过程比较复杂
5. 工厂模式和依赖注入
   1. DI容器就是一个大的工厂，负责在启动到时候根据配置事先创建好对象。当应用程序需要哪些对象时，直接从容器中获取，容器负责整个应用对象的创建
   2. DI的核心功能
      1. 配置解析
      2. 对象创建，利用反射机制，动态创建类
      3. 生命周期管理
      
### 二十五、建造者模式 builder
1. 当一个类的参数超过 4 个且有些是可选参数则可以考虑建造者模式
2. 建造者解决的问题：参数过多，中间隐含顺序、默认值、可读性、依赖等问题
   1. 解决方法：set
   2. 建造者模式：把类当做参数

### 二十六、原型模式 proto
1. 对象的创建成本比较大，同一个类的不同对象之间差别不大，利用已有的对象进行复制的方式创建对象，以达到节省时间的目的
   1. 对象的创建成本大：对象中的数据需要经过复杂的计算才能得到或者需要重RPC、网络、数据库文件系统等非常慢的IO中读取
2. 深、浅拷贝
   1. 浅拷贝只会复制对象中基本数据类型数据和引用对象的内存地址，不会递归复制引用对象，以及对于对象的引用
   2. 深拷贝得到的一份完全独立的对象
   3. 浅拷贝会共享部分数据

-- ```以上为创建模式：封装创建过程，解耦创建和使用```
-- ```以下为结构模式：解决特定场景下的问题```

### 二十七、代理模式【高频】 proxy
1. 感觉是在业务外层又包了一层，可以实现业务无侵入，更换代理后业务无需变动
   1. 面向接口或
   2. 者面向继承(没有接口的情况下)
   3. 动态代理（基于反射的原理）？？？没看懂
2. 应用场景
   1. 业务系统非功能性开发：监控、统计、鉴权、限流、事务、幂等、日志
   2. RPC、缓存

### 二十八、桥接模式【低频】
1. 将抽象的实现解耦，各自可以灵活扩展???没看懂

### 二十九、装饰器 decorator
1. 和原始类继承同样的父类
2. 装饰器模式是对类功能的增强<核心作用点>??和代理模式区别在哪呢？代理模式不修改业务公共实现额外功能，装饰器模式增加的是相同的功能

### 三十、适配器 Adapter
1. 将不兼容的接口转换为兼容的接口（类似USB）
2. 类适配器和对象适配器
   1. 类适配器：基于继承
   2. 对象适配器：基于组合
3. 应用场景：
   1. 接口设计有缺陷（包含大量静态方法、不好测试）
   2. 统一多个类的接口设计（某个功能实现依赖多个外部系统）
   3. 替代外部系统
   4. 兼容老版本
   5. 适配不同的数据
4. 代理 桥接 装饰器 适配器比较
   1. 代理：在不改变原始类的接口的条件下，为原始类顶一个代理，主要是控制访问，而非加强功能
   2. 桥接：将部分接口和实现分离，从而较为容易、也相对独立地加以改变
   3. 装饰器：在不改变原始类的情况下，对原始类的功能加以增强，并支持多个装饰器嵌套使用
   4. 适配器：是一种时候补救策略，适配器原始类的接口不同，而代理模式、装饰器模式提供的都和源氏类相同的接口

### 三十一、门面模式 Facade
1. 解决接口粒度问题
2. 为子系统提供一组统一的接口，定义一组高接口让子系统更容易用
   1. 解决易用性问题:封装系统的底层实现，隐藏系统的复杂性，提供一组更加简单易用、更高层的接口
   2. 解决性能问题：减少io
   3. 解决分布式问题：
3. 接口设计原则：尽量保持复用性，特殊情况允许提供门面接口提高易用性


### 三十二、组合模式 composite【针对特殊的数据结构，不是很常用】
1. 将一组对象组织成树形结构，以表示一种部分-整体的层次结构，能通过树的遍历算法解决的场景


### 三十三、享元模式【挺好用的】
1. 共享单元简称享元，目的是复用对象，节省内存，前提是复用的对象是不可变对象。
2. 棋局demo？节省了颜色和名称，但是对于位置还是得初始化，核心的把不变的部分抽离出来，根据简单工厂模式获取
3. 享元模式在文本编辑器中的使用？？怎么用？怎么用来节省内存？？
   1. 记录文子和格式，格式包括：字体，颜色，大小
   2. 意思是把这些当做一个类抽离出来？
4. 感觉享元模式是基于内存的缓存
5. 享元模式 vs 单例 缓存 对象池
   1. 单例是一个类只有一个对象，享元是一个类可以有多个对象，享元的目前是节省内存，单例是限制对象个数
   2. 享元"缓存"的是对象，目的是为了复用，缓存的目的是为了提高访问速度而非复用
   3. 对象池是为了避免重复创建，节省时间，而享元是共享，节省的是空间。

-------------------- 行为模式-----------------------------
- 创建型：解决对象创建的问题，把创建和使用解耦（单例、简单工厂、工厂、抽象工厂、建造者、原型）
- 结构型：类或者对象的组装问题，把不同功能的代码解耦（代理、装饰器、适配器、桥接、门面、组合、享元）
- 行为型：类和对象之间的交互问题，把不同的行为解耦（观察者、模板、策略、责任链、状态、迭代、访问者、备忘录、命令、解释器、中介）

### 三十四、观察者模式【重点模式】
1. 也称发布订阅模式：在对象之间定义一个一对多的依赖，当一个对象的状态改变的时候，所有的对象有自动收到通知。
2. 解决的问题：将观察者和被观察者解耦

### 三十五、模板模式 【常用】TemplateMethod
1. 主要解决复用和扩展问题
2. 模板模式是在一个方法中定义好一个算法骨架，并将某些步骤推迟到子类中实现。
3. 感觉就是抽象类，
4. 模板vs回调
   1. 同步回调和模板基本一致，都是在一个大的算法骨架中，自由替换其中的某个步骤，起到代码复用和扩展的目的。异步回调更新观察者模式
   2. 回调是一种双向调用的关系，回调基于组合、模板基于继承

### 三十六、策略模式【常用】strategy
0. 定义一族算法类，将每个算法分别封装起来，让它们可以互相替换，使算法的变化独立于使用他们的客户端。
1. 解决什么问题呢：解耦策略的定义、创建、使用
   1. 定义：一个策略接口和一组策略的实现
   2. 创建：策略工厂类
   3. 使用：运行时动态确定
2. 应用场景呢：简单好用，减少 if-else/switch-case


### 三十七、责任链模式【常用】不用修改任何代码就需要配置式开发
1. 将请求和发送解耦，让多个接受对象都有机会处理这个请求，将这些接受对象串成一条链，并沿着这条链传递这个请求直到某个对象处理改对象。
2. 多个处理器依次处理请求
3. 场景：过滤器和拦截器

### 三十八、状态模式
1. 不常用，类似组合模式，在特殊场景下，游戏、工作流。
2. 状态机：状态(State)、事件(Event)、动作(Action)
3. 实现方式：
   1. 分支逻辑：if-else...
   2. 查表法：搞成一个二维数组
   3. 状态模式：通过将事件触发的状态转移和动作执行，拆分到不同的状态类中，来避免分支判断逻辑
4. 游戏状态机比较复杂推荐使用查表法，而电商状态比较少，转移比较简单，但是事件触发的逻辑比较复杂，推荐使用状态机模式


### 三十九、迭代器模式 iterator or cursor 【感觉坑还不少】
1. 迭代对象【数组、链表、树、图、跳表】，迭代器将对象的遍历操作从集合类中分离出来，放到迭代器中，职责更单一
2. iterator/ConcreteIterator/ConcreteCollection/Collection
3. 迭代器的意义
   1. 对于数组、链表for就够了，但是对于树、图比较复杂的数据结构由客户端实现遍历容易出错，应对复杂性的方案就是拆分
   2. 将游标指向当前位置，存储到迭代器类中，可以创建多个迭代器
   3. 容器和迭代器都基于接口
4. 迭代器遍历过程中增删都会有坑
   1. 正确处理方案是，遍历时进行增或删报错

### 四十、访问者模式【不推荐使用】
### 四十一、备忘录模式（snapshot）
1. 场景：防止丢失、撤销、恢复
2. 在不违背封装原则的前提下，捕获一个对象的内部状态，并在这个对象之外保存这个状态，以便之后恢复对象为先前状态

### 四十二、命令模式
1. 将请求命令封装成一个对象
2. 应用场景 & 解决方案，设计模式之间的主要区别在于意图，即：应用场景。

### 四十三、解释器模式
1. 根据语法规则解读"句子"的解释器

### 四十四、中介模式
1. 定义了一个对象，来组装一组对象之间的交互，将这组对象之间的交互委派给中介对象、来避免对象之间的直接交互


设计思想和原则是内功、设计模式是招式