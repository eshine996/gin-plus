# gin-plus

gin工程化示例

## 概述：

### 设计原则

1.分层设计。</br>
1.各模块可单元测试。尤其是service层</br>

在gin的基础上，在原有路由，函数的注册方式上，增加一种路由注册的方式,方式来源与goframe框架。 </br>
orm使用gorm，配置文件库使用viper。配置文件加载，数据库连接初始化封装在ginp内部。 </br>
命令行启动使用cobra</br>

## 目录简介</br>

## 已完成的

1.在gin的基础上增加一种新的路由注册方式，注册对象。思路来源于goframe。
2.将gorm封装到dao层，同时也暴露了orm的增删改查的接口，使用者不需要对orm初始化等操作。
3.目录/代码分层设计，思路来源于goframe。

## 进行中的

1.命令行启动/配置文件使用cobra，viper俩件套。

## 要做的：</br>

1.错误处理/具备栈回溯的错误日志 todo</br>
2.日志处理/具备统一的日志格式，统一gorm和gin的日志 todo</br>
3.api文档/准备借鉴goframe的方式，openAPI3+swagger 来生成文档 todo</br>
4.ctl工具/一键迁移数据库或从数据库生成模型，一键生成各层生成代码 todo</br>
