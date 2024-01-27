# Singleton

Ensure a class has only one instance, and provide a global point of access to it

# 单例模式

保证一个类仅有一个实例，并提供一个访问它的全局访问点。

### 优点

1. 内存中只有一个实例，减小内存开销
2. 避免对资源的多重占用

### 缺点

1. 没有接口，不能继承，与单一职责冲突，一个类只关心内部逻辑，而不关心外部怎么实例化