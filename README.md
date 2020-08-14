# 游戏服务器
  1. 使用zinx框架搭建了基础的网络通信服务器
  2. 在服务器的基础框架上进行业务逻辑处理
     注册，登录，选择人物，建立房间，邀请好友请求的 业务处理 
     并且采用帧同步实现各个用户之间的数据同步
  1. 服务器模块：使用zinx框架搭建基础网络通信服务器
  2. 数据库模块：使用gorm框架，实现数据库访问
  3. 链接管理模块：使用map存储并管理用户链接状态
                开始游戏状态/正常状态，结束状态
  4. 消息模块：采用TLV数据格式以及protobuf进行数据的序列化与反序列化实现客户端与服务器之间的数据沟通
  5. 业务请求模块：封装链接信息，对请求进行封装，加入任务协程池进行处理。
  6. 消息处理路由模块：使用数组存储消息类型与注册的处理方法 之间的关系
  7. 房间管理模块：使用map存储管理房间状态信息，创建，战斗，结束