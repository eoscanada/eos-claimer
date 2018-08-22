千万别忘了 claimrewards！
----------------------------

Just claim it, claim it, claim it before it's gone.



设置权限
-----------------

请用一次性密钥进行此操作。

```
cleos set account permission PRODUCERACCT claimer '{"threshold":1,"keys":[{"key":"YOUR_PUB_KEY","weight":1}]}' "active" -p PRODUCERACCT@active
cleos set action permission PRODUCERACCT eosio claimrewards claimer
```

用 Kubernetes，输入 secret：

```
kubectl create secret generic claimer --from-literal=privkey=PRIVATEKEY --from-literal=pubkey=YOUR_PUB_KEY
```
启动后等着获益吧。




安装
-------

请用:

```
go get -u -v github.com/eoscanada/eos-claimer
```


常见问题
---


问：10秒是不是有点太短了？

答：不是。再少点也没问题。它会撞到前面接收你交易的节点，
一直失败到你的获得奖励时间。交易不会广播，因为它不会通过24小时的断言。
不过，对着你自己的节点这样做，别弄我们的。


问： `set action permission` 是什么玩意儿？

答：它会调用 `eosio::linkauth` 并为这个 action 分配一个命名权限。 
这意味着您现在可以签署一个交易，用 `claimer` 权限来授权 *这个合约* 上的 *这个 action*。 
这是个EOS.IO区块链上还没有得到应有的注意力的，非常强大的功能。
不过请注意，1.0版本的 `get_account` 指令不会给你一个具象化的链接身份验证和action到权限映射。 
不过我肯定马上就会有的。 有了这样的设置和
`claimrewards` action 的设计..你甚至可以把相关的私钥给你的敌人......
他们可以替你 claimrewards。 **但是别这么做啊**......你知道吧，以防万一。

证书
-------

MIT
