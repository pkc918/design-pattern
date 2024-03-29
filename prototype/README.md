# 原型模式

原型模式用于创建重复的对象，同时又能保证性能。这种模式实现了一个原型接口，该接口用于创建当前对象的克隆。

### 优点

1. 性能提高
2. 逃避构造函数的约束

### 缺点

1. 配备克隆方法需要对类的功能进行通盘考虑，这对全新的类不是很难，但对于已有的类不一定很容易，特别当一个类引用不支持串行化的间接对象，或者引用含有循环结构的时候。
2. 必须实现 Cloneable 接口。

### 使用场景
通过 new 产生一个对象需要非常繁琐的数据准备或访问权限，则可以使用原型模式。