## Redis的数据结构
redis为什么那么快？
除了它是内存数据库之外，还有一个重要因素--它实现的数据结构，使得在增删改查操作时高效的处理
- redis有五种基础数据类型由六种数据结构实现
    - string 并不是c的 char*来实现 而是封装了一个结构 叫做sds (Simple Dynamic String)