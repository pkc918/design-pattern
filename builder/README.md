# 建造者模式

建造者模式就是使用一个一个简单的对象一步一步构建成一个复杂的对象

### 意图
将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示

### 优点
1. 分离构建过程和表示，使得构建过程更加灵活，可以构建不同的表示。
2. 可以更好地控制构建过程，隐藏具体构建细节。
3. 代码复用性高，可以在不同的构建过程中使用相同的构建者。

### 缺点
1. 如果产品属性少，建造者模式可能导致代码冗余。
2. 建造者模式增加了系统的类和结构体的数量。