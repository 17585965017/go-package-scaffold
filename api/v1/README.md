# 接口设计规范

## **RESTful API 接口设计规范以及技巧**

- 资源请求使用名词， 并且使用复数形式

  ```
  GET /shop/product/v1/products/get/1  // ❌
  GET /shop/product/v1/products/1  // ✅
  ```

- 尽量将请求转化成为的资源的属性，或者使用请求方法 method 来区分

  ```
  PSOT /shop/account/v1/users/1?status=disable
  PUT /shop/account/v1/users/1/disable
  DELETE /shop/account/v1/users/1/disable
  ```

- URI 路径使用小写字母（除版本号外）和中横线

  ```
  GET /product-search/v1/SearchProduct?KEYWORD=xxx  ❌
  GET /product-search/v1/search-product?keyword=xxx ✅
  ```

  为什么不推荐使用中横线

- 避免使用层级较深的 URI， 超过了 2 层就尽量使用查询参数代替

  ```
  GET /schoo-manager/region/1/schools/2/studdents/3  ❌
  GET /schoo-manager/studdents?id=3&scholl_id=2%region=1  ✅
  ```

- 经常使用的，复杂查询标签化，降低使用和维护成本

  ```
  GET /shop/account/v1/user?status=disable&sort=update_time.desc  ❌
  GET /shop/account/v1/user/recently-disabled (查询标签化)  ✅
  ```

- 针对批量请求方式

  - 直接使用 post, body 传递多个资源参数，而不是发送多个 delete 请求
  - 通过携带多个资源 id 并用分隔符隔开的方式

## **接口的安全性和幂等性**

安全性： 接口所指向的资源是只读的

幂等性： 操作一次的和多次 结果的是一样的

## **RESTful API 接口设计规范以及技巧**

- 资源请求使用名词， 并且使用复数形式

  ```
  GET /shop/product/v1/products/get/1  // ❌
  GET /shop/product/v1/products/1  // ✅
  ```

- 尽量将请求转化成为的资源的属性，或者使用请求方法 method 来区分

  ```
  PSOT /shop/account/v1/users/1?status=disable
  PUT /shop/account/v1/users/1/disable
  DELETE /shop/account/v1/users/1/disable
  ```

- URI 路径使用小写字母（除版本号外）和中横线

  ```
  GET /product-search/v1/SearchProduct?KEYWORD=xxx  ❌
  GET /product-search/v1/search-product?keyword=xxx ✅
  ```

  为什么不推荐使用中横线

- 避免使用层级较深的 URI， 超过了 2 层就尽量使用查询参数代替

  ```
  GET /schoo-manager/region/1/schools/2/studdents/3  ❌
  GET /schoo-manager/studdents?id=3&scholl_id=2%region=1  ✅
  ```

- 经常使用的，复杂查询标签化，降低使用和维护成本

  ```
  GET /shop/account/v1/user?status=disable&sort=update_time.desc  ❌
  GET /shop/account/v1/user/recently-disabled (查询标签化)  ✅
  ```

- 针对批量请求方式

  - 直接使用 post, body 传递多个资源参数，而不是发送多个 delete 请求
  - 通过携带多个资源 id 并用分隔符隔开的方式

------

## **接口的安全性和幂等性**

安全性： `**接口所指向的资源是只读的**`

幂等性： `**操作一次的和多次 结果的是一样的**`

|        | 安全性 | 幂等性 |
| ------ | ------ | ------ |
| POST   | ❌      | ❌      |
| GET    | ✅      | ✅      |
| DELETE | ❌      | ✅      |
| PUT    | ❌      | ✅      |

------

## 接口的错误码设计

设计原则：

- 不推荐使用全局错误码
- 按照项目，服务，模块，错误类型依次编号
- 模块错误码 不能 超过 99 个

错误码例子：

- 06100325
  - 06 表示 06 号项目组
  - 10 代码是 10 号服务
  - 03: 03 号模块
  - 25: 错误码为 25

------

## 接口的兼容设计

在 接口开发的时候就`添加版本目录`，新的需求来的时候，不要在原有的接口中更改，而是心开发接口改掉业务逻辑，并且在后期对接完成并且没有使用的是时候在将其删除
