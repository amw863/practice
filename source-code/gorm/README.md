两个观察者模式：把处理行初始化为一个列表，然后执行时根据
### version 1.22.5
### version 1.9.16
- 入口 ```gorm.Open```
- 返回 ```gorm.DB```
  - SQLCommon[interface]
    - sql.Result[database/sql]
    - sql.Row[database/sql]
  - logger[interface]
  - search[struct]
    - gorDB
    - searchPreload
  - parent[gorm.DB]
  - callbacks[struct]
    - Callback
      - logger
      - Scope
        - search
        - gorDB
        - Field
          - StructField
            - Relationship
              - JoinTableHandlerInterface[interface]
                - JoinTableForeignKey[struct]
                - JoinTableHandlerInterface[struct]
                - gorDB
      - CallbackProcessor
        - Callback
        - Scope
  - dialect[interface]
    - SQLCommon
    - StructField
- dialect 针对不同的SQL
- callback 各种操作，利用init，依赖 scope 注入

- step one: register callback func for curd
- step two: get conn with config
- step three: exec sql proxy with sql