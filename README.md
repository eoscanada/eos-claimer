Never forget to claimrewards
----------------------------

Just claim it, claim it, claim it before it's gone.



Setup permissions
-----------------

Use a throw-away key for this action.

```
cleos set account permission PRODUCERACCT claimer '{"threshold":1,"keys":[{"key":"YOUR_PUB_KEY","weight":1}]}' "active" -p PRODUCERACCT@active
cleos set action permission PRODUCERACCT eosio claimrewards claimer
```

With Kubernetes, inject secret with:

```
kubectl create secret generic claimer --from-literal=privkey=PRIVATEKEY --from-literal=pubkey=YOUR_PUB_KEY
```

Boot and profit.




Install
-------

With:

```
go get -u -v github.com/eoscanada/eos-claimer
```


FAQ
---

Q: Isn't 10 secs a bit aggressive ?

A: Not really. Could even be lowered. It'll bump against the front
node receiving your transaction and fail each time until the time
comes and is ready for your claim. The transaction will not propagate
because it fails the 24h assertion.  Do it against your nodes though,
not ours.


Q: Wut's that `set action permission` thing ?

A: This calls `eosio::linkauth` and assigns a named permission to that
particular action. This means you can now sign a transaction that
authorizes *that specific action* on *that specific contract* with the
permission `claimer`.  This is a really powerful feature of EOS.IO
blockchains that isn't shining as it should right now.  Mind you, with
release 1.0, the `get_account` calls do not give you a portrait of
your linked auths and actions-to-permissions mapping.  I'm sure that's
coming soon though.  With a setup like this here, and the design of
the `claimrewards` action.. you could almost give the associated
private key to your enemies.. they could claimrewards for you. BUT HEY
DON'T DO THAT.. you know, just in case.

LICENSE
-------

MIT
