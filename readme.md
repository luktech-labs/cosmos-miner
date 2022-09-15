# cosmosminer
**cosmosminer** is a crypto mining blockchain game built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
This is a POC built in order to demo how Cosmos SDK is working

# Objects stored
***
MinerTemplate:
* index (unique identifier of the object)
* name (descriptive name of the miner)
* price (price of the miner)
* claimTime (minimum amount of time to claim mining rewards since last claimed)
* prodPerSecond (amount of mining reward units produced per second)

***
Miner:
* index (unique identifier of the object)
* creator (owner of the miner)
* lastClaimed (last time when the mining rewards were claimed for this miner)
* templateIndex (unique identifier of the miner template)

***
Balance:
* map[accountAddress] amount (storing the amount of rewards claimed by each account address)

***
MinerInfo:
* templateNextID (incremental id to uniquely identify miner templates)
* minerNextID (incremental id to uniquely identify miners)

# Messages
***
*  create-miner [miner-template-id] (used to create/buy a miner by a user)
*  create-miner-template [name] [price] [prod-per-second] [min-claim-time] (used to create a miner template by an admin)
*  mine [miner-id] (used to claim mining rewards by a user)


## Get started

First install IgniteCLI

```
curl https://get.ignite.com/cli@! | bash
```

Now you serve the chain
```
$ ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development. Please note that it also builds the chain cli binary: `cosmos-minerd`

2 account will be created: `alice` & `bob` which will also receive a lot faucet tokens.

Let's see how many tokens alice has:

```
$ cosmos-minerd query bank balances cosmos1eaurd46ulwa3yj8ukwm0zl7e06lpz25f7drhrx
balances:
- amount: "100000000"
  denom: stake
- amount: "20000"
  denom: token
```

As it's the first time running this chain let's first create a miner(rig) template.
Let's choose some demo data like "asic1" as the name with a price of 100 tokens which will produce 1 unit per second with a min claim time of 20 seconds.

```
 $ cosmos-minerd tx cosmosminer create-miner-template "asic1" 100 1 20 --from cosmos1eaurd46ulwa3yj8ukwm0zl7e06lpz25f7drhrx --chain-id cosmosminer
 ...
 txhash: 5307FB8AAABCCD05905DC72ADD7816C101A85468640114B3AF14AED078734098
```

Now let's get the template id by listing the templates using the query module:

```
$ cosmos-minerd query cosmosminer list-stored-miner-template
storedMinerTemplate:
- claimTime: "20"
  index: "1"
  name: asic1
  price: "100"
  prodPerSecond: "1"
```

Now that we have a template, let's also create a miner using that template index (1)

```
cosmos-minerd tx cosmosminer create-miner 1 --from cosmos1eaurd46ulwa3yj8ukwm0zl7e06lpz25f7drhrx --chain-id cosmosminer
...
txhash: 3390991EF15BD67FAC4228290B08279E97E3EFB0E9D27389DB7AA51A64CF4A81
```

As the price of the miner is 100 tokens, if we will query the balance again, we should have 20000(original) - 100(price) = 19900 tokens

```
cosmos-minerd query bank balances cosmos1eaurd46ulwa3yj8ukwm0zl7e06lpz25f7drhrx
balances:
- amount: "100000000"
  denom: stake
- amount: "19900"
  denom: token
```

Now that we have a miner, let's claim some rewards by using the miner id 1:

```
cosmos-minerd tx cosmosminer mine 1 --from cosmos1eaurd46ulwa3yj8ukwm0zl7e06lpz25f7drhrx --chain-id cosmosminer
...
txhash: 525891C6887EEAA9C4D3A06F87EE8D524622915C51E3580FFECA1ECDB9F8BC90
```

Now let's check our mining balance

```
cosmos-minerd query cosmosminer list-miner-balance
minerBalance:
- amount: "226"
  index: ""
```







## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
