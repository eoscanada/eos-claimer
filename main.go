package main

import (
	"fmt"
	"log"
	"os"
	"time"

	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/system"
)

// Define EOS_CLAIMER_KEY (with the private key that is allowed to
// call `eosio::claimrewards`)
//
// Define EOS_CLAIMER_PERMISSION if the permission of the
// EOS_CLAIMER_OWNER account is not `active` (because you linkauth'd
// the action to another permission.
//
// Define EOS_CLAIMER_OWNER which is your BP account.
//
// Define EOS_CLAIMER_URL to a reachable endpoint where to send
// transactions.

func main() {
	api := eos.New(os.Getenv("EOS_CLAIMER_URL"))

	keyBag := eos.NewKeyBag()
	for _, key := range []string{
		os.Getenv("EOS_CLAIMER_KEY"),
	} {
		if err := keyBag.Add(key); err != nil {
			log.Fatalln("Couldn't load private key:", err)
		}
	}

	api.SetSigner(keyBag)
	api.Debug = true

	act := system.NewClaimRewards(eos.AccountName(os.Getenv("EOS_CLAIMER_OWNER")))
	if perm := os.Getenv("EOS_CLAIMER_PERMISSION"); perm != "" {
		act.Authorization[0].Permission = eos.PermissionName(perm)
	}

	var sleep time.Duration
	for {
		time.Sleep(sleep)
		sleep = 3 * time.Minute

		actionResp, err := api.SignPushActions(act)
		if err != nil {
			fmt.Println("ERROR calling :", err)
		} else {
			fmt.Println("RESP:", actionResp)
		}
	}
}
