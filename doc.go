/*
 * Opay 提供一套简单在线支付系统。
 *
 * 特点
 * 1. 完全面向接口开发
 * 2. 支持超时自动撤销处理订单
 *
 * 使用方法
 * 1. 注册资产账户操作接口实例
 * 2. 实现订单接口
 * 3. 注册订单类型对应的操作接口实例
 * 4. 新建服务实例 var opay=NewOpay(db, 5000)
 * 5. 开启服务协程 go opay.Serve()
 * 6. 推送订单 done, err:=opay.Push(Request{})
 * 7. 使用 <-done 等待订单处理结束
 */

package opay
