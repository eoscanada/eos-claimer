Never forget to claimrewards
----------------------------

Just claim it, claim it, claim it before it's gone.



Setup permissions
-----------------

Use a throw-away key for this action.

cleos set account permission eoscanadacom claimer '{"threshold":1,"keys":[{"key":"EOS7NFuBesBKK5XHHLtzFxm7S57Eq11gUtndrsvq3Mt3XZNMTHfqc","weight":1}]}' "active" -p eoscanadacom@active
cleos set action permission eoscanadacom eosio claimrewards claimer

Inject secret with:

kubectl create secret generic claimer --from-literal=privkey=PRIVATEKEY --from-literal=pubkey=EOS7NFuBesBKK5XHHLtzFxm7S57Eq11gUtndrsvq3Mt3XZNMTHfqc

Boot and profit.


LICENSE
-------

MIT or wut ?
