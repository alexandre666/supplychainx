# Testing the application

The application is composed of the following:
- `scxd` the daemon node that runs the blockchain
- `scxcli` a cli client that interacts with the blockchain and stores accounts

The demonstration uses Docker containers to represent daemon nodes while the cli is installed and used locally.

You can install the application with the following instructions:

```
cd $GOPATH/src/github.com
mkdir ltacker
cd ltacker
git clone https://github.com/ltacker/supplychainx.git
cd supplychainx
make
```

To run the nodes, run in the same directory:

```
docker-compose up
```

## Test Environment

The application can be tested with several nodes with a test environment composed of the following:
- `test_env/` contains several files to test the application
- `test_validator.Dockerfile` is used to represent a validator already set in the provided genesis. This docker image contains a predefined private key for the validator
- `test.Dockerfile` is used to represent other daemon nodes in the network, not directly validator but that can apply to become a validator
- `docker-compose.yml` used to launch the test environment with the command `docker-compose up`

`test_env/` contains the following files:
- `node_key.json` a node key file copied into the initial validator node to have a predefined node id for the validator
- `testkey.json` a predefined private key for the validator, so that we can determine it in the genesis file
- `config.toml` a customized config file for the nodes. This config define the validator node as a persistent peer with its node id
- `genesis.json` a predefined genesis file with a validator and six accounts

The six accounts in the genesis file are the following:

```
[
    {
        "name": "regulatorA",
        "type": "local",
        "address": "cosmos1g7zvz09eyawjqhc95fyjdtj2paqmscfxn2wt57",
        "pubkey": "cosmospub1addwnpepqgr49jezdhezj2wq4hsc7s69lkj55vts0w78kufj3jn8ee0ygzkygrs87j5",
        "mnemonic": "slot alien drive pool siege smile whip desert resist decide parent sign bone add chronic reason practice pulse rally glad hover apart walnut budget"
    },
    {
        "name": "regulatorB",
        "type": "local",
        "address": "cosmos1d6ygahd78ftmm9c2u9n9lhu6m3d39l74qwdklr",
        "pubkey": "cosmospub1addwnpepqfkaeh8e5s4733wr9xmentlluv8xv57x9kk59vcmepuvta9zrlyfwv9yzuc",
        "mnemonic": "attitude kiwi assist stock tonight hybrid subject large scheme transfer sibling era core they nothing toss ocean lemon twelve depend circle human alone thunder"
    },
    {
        "name": "regulatorC",
        "type": "local",
        "address": "cosmos19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv0dfaxv",
        "pubkey": "cosmospub1addwnpepqtuhx5dxyp88dzn0hdwzxe83075sk3frwj4eq7dzywztm4tu7uvxx99dngn",
        "mnemonic": "pretty prefer sponsor margin large divorce bubble silly prefer symptom rather draft similar bubble object pass book recall join say detect define cave cause"
    },
    {
        "name": "Acme",
        "type": "local",
        "address": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
        "pubkey": "cosmospub1addwnpepqddjprqqydn843h89n7tcuzj6uydhdw8x6d8qwewyeg48nyndat9qn2z4tt",
        "mnemonic": "fun body weasel minimum artefact mandate impulse elbow captain surge dove before true pause lecture spirit seek into carry cream theory pupil rhythm bonus"
    },
    {
        "name": "HahaTransport",
        "type": "local",
        "address": "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a",
        "pubkey": "cosmospub1addwnpepqdqpecs49ae3wtxde5umf405qg4nrsggfttgfr6egtahjfkm8zd7urw3npc",
        "mnemonic": "dragon blanket library fly casino treat mix angle uphold engine main normal present upper sadness dice pulp favorite ritual rate water hazard empower what"
    },
    {
        "name": "HahaPhone",
        "type": "local",
        "address": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
        "pubkey": "cosmospub1addwnpepqtrllghfm9jpgn9xdkg5l9z7dq25q4hsqwgjylk0l35nmjq90tnswmmx9l5",
        "mnemonic": "nasty dream summer bullet vicious aspect local gate finger slab leave canvas police uphold gym portion host three barely polar purchase rubber feel unlock"
    }
]
```

The accounts can be imported locally with the commands:
```
scxcli keys add regulatorA --recover
scxcli keys add regulatorB --recover
scxcli keys add regulatorC --recover
scxcli keys add Acme --recover
scxcli keys add HahaTransport --recover
scxcli keys add HahaPhone --recover
```
The above mnemonics must be provided.

## PoA module

The following section demonstrates commands that can be used to interact with the validators through the PoA module. 

With the `docker-compose up` command, there is initially only the node `supplychainx_validator_node_1` which runnning as a validator.

We can use PoA transactions to add new validator.

First we must get the consensus public key of a node with the following command (for the node 1 in this example):
```
> docker exec supplychainx_node1_1 scxd tendermint show-validator
```

We can use this key to apply to become a validator:

```
> scxcli tx poa apply --from regulatorB <cons_pubkey> --identity RegulatorB --details "A regulator from Korea"
```

We can observe this application with the `applications` query:

```
> scxcli query poa applications
[
  {
    "subject": {
      "operator_address": "cosmosvaloper1d6ygahd78ftmm9c2u9n9lhu6m3d39l7496erns",
      "consensus_pubkey": "cosmosvalconspub1zcjduepqpdp3c466zhdr0glusu2uknwur7ypgcqnv04wwxvlk6gc5c75dl8ql528x7",
      "description": {
        "moniker": "",
        "identity": "RegulatorB",
        "website": "",
        "security_contact": "",
        "details": "A regulator from Korea"
      }
    },
    "approvals": "0",
    "totals": "0",
    "voter": null
  }
]
```

There is only one validator, there only its vote is necesary to approve this application:

```
> scxcli tx poa vote-application --from regulatorA cosmosvaloper1d6ygahd78ftmm9c2u9n9lhu6m3d39l7496erns approve
```

We can reject the application by replacing `approve` by `reject`

We can observe the appended validator with the `validators` command:

```
> scxcli query poa validators
[
  {
    "operator_address": "cosmosvaloper1g7zvz09eyawjqhc95fyjdtj2paqmscfxk767cd",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqs53kvv8mkecwjl8zxu09fl6x83vnsncxc0ultle7y5ygf70nrqnqsxslwv",
    "description": {
      "moniker": "validator",
      "identity": "",
      "website": "",
      "security_contact": "",
      "details": ""
    }
  },
  {
    "operator_address": "cosmosvaloper1d6ygahd78ftmm9c2u9n9lhu6m3d39l7496erns",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqpdp3c466zhdr0glusu2uknwur7ypgcqnv04wwxvlk6gc5c75dl8ql528x7",
    "description": {
      "moniker": "",
      "identity": "RegulatorB",
      "website": "",
      "security_contact": "",
      "details": "A regulator from Korea"
    }
  }
]
```

Let's apply with the regulator C:
```
> docker exec supplychainx_node2_1 scxd tendermint show-validator
> scxcli tx poa apply --from regulatorC <cons_pubkey> --identity RegulatorC --details "A regulator from China"
> scxcli query poa applications
[
  {
    "subject": {
      "operator_address": "cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l",
      "consensus_pubkey": "cosmosvalconspub1zcjduepq57tnjguj96vhuyzrgc4p0d2nuudsaeu7m3y0hamhshgacsu5wvascjkzgq",
      "description": {
        "moniker": "",
        "identity": "RegulatorC",
        "website": "",
        "security_contact": "",
        "details": "A regulator from China"
      }
    },
    "approvals": "0",
    "totals": "0",
    "voter": null
  }
]
```

We can observe the state of the application being updated when a validator approves it:

```
> scxcli tx poa vote-application --from regulatorA cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l approve
> scxcli query poa applications
[
  {
    "subject": {
      "operator_address": "cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l",
      "consensus_pubkey": "cosmosvalconspub1zcjduepq57tnjguj96vhuyzrgc4p0d2nuudsaeu7m3y0hamhshgacsu5wvascjkzgq",
      "description": {
        "moniker": "",
        "identity": "RegulatorC",
        "website": "",
        "security_contact": "",
        "details": "A regulator from China"
      }
    },
    "approvals": "1",
    "totals": "1",
    "voter": [
      "cosmosvaloper1g7zvz09eyawjqhc95fyjdtj2paqmscfxk767cd"
    ]
  }
]
```

The vote from regulatorB is necessary to append the validator, let's vote:
```
> scxcli tx poa vote-application --from regulatorB cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l approve
> scxcli query poa validators
[
  {
    "operator_address": "cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l",
    "consensus_pubkey": "cosmosvalconspub1zcjduepq57tnjguj96vhuyzrgc4p0d2nuudsaeu7m3y0hamhshgacsu5wvascjkzgq",
    "description": {
      "moniker": "",
      "identity": "RegulatorC",
      "website": "",
      "security_contact": "",
      "details": "A regulator from China"
    }
  },
  {
    "operator_address": "cosmosvaloper1g7zvz09eyawjqhc95fyjdtj2paqmscfxk767cd",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqs53kvv8mkecwjl8zxu09fl6x83vnsncxc0ultle7y5ygf70nrqnqsxslwv",
    "description": {
      "moniker": "validator",
      "identity": "",
      "website": "",
      "security_contact": "",
      "details": ""
    }
  },
  {
    "operator_address": "cosmosvaloper1d6ygahd78ftmm9c2u9n9lhu6m3d39l7496erns",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqnc7ll3usecdd8g9vyfe0xl346xettv78l7fw36cllq5gvvadv7ps67dtr5",
    "description": {
      "moniker": "",
      "identity": "RegulatorB",
      "website": "",
      "security_contact": "",
      "details": "A regulator from Korea"
    }
  }
]
```

For an unknown reason, RegulatorB doesn't get along with RegulatorC, it decides to create a kick proposal to remove this validator

```
> scxcli tx poa propose-kick --from regulatorB cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l
> scxcli tx poa vote-kick-proposal --from regulatorB cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l approve
> scxcli query poa kick-proposals
[
  {
    "subject": {
      "operator_address": "cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l",
      "consensus_pubkey": "cosmosvalconspub1zcjduepq57tnjguj96vhuyzrgc4p0d2nuudsaeu7m3y0hamhshgacsu5wvascjkzgq",
      "description": {
        "moniker": "",
        "identity": "RegulatorC",
        "website": "",
        "security_contact": "",
        "details": "A regulator from China"
      }
    },
    "approvals": "1",
    "totals": "1",
    "voter": [
      "cosmosvaloper1d6ygahd78ftmm9c2u9n9lhu6m3d39l7496erns"
    ]
  }
]
```

RegulatorA doesn't have any problem with RegulatorC, therefore it rejects the kick proposal:

```
> scxcli tx poa vote-kick-proposal --from regulatorA cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l reject
> scxcli query poa kick-proposals
null
```

Upset, RegulatorB decides to leave the validator set by itself:

```
> scxcli tx poa leave-validator-set --from regulatorB
> scxcli query poa validators
[
  {
    "operator_address": "cosmosvaloper19ywsweqdwd2nwyaqlvdz5xcp8vfpp0qv2eag2l",
    "consensus_pubkey": "cosmosvalconspub1zcjduepq57tnjguj96vhuyzrgc4p0d2nuudsaeu7m3y0hamhshgacsu5wvascjkzgq",
    "description": {
      "moniker": "",
      "identity": "RegulatorC",
      "website": "",
      "security_contact": "",
      "details": "A regulator from China"
    }
  },
  {
    "operator_address": "cosmosvaloper1g7zvz09eyawjqhc95fyjdtj2paqmscfxk767cd",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqs53kvv8mkecwjl8zxu09fl6x83vnsncxc0ultle7y5ygf70nrqnqsxslwv",
    "description": {
      "moniker": "validator",
      "identity": "",
      "website": "",
      "security_contact": "",
      "details": ""
    }
  }
]
```

## Supply chain

The following section demonstrates commands that can be used to interact with the supply chain.

First, let's the authority appending the necessary organizations:
```
> scxcli tx scx append-organization --from regulatorA cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e Acme --organization-description "A company making everything"

> scxcli tx scx append-organization --from regulatorA cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a HahaTransport --organization-description "A company transporting everything"

> scxcli tx scx append-organization --from regulatorA cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt HahaPhone  --organization-description "A company making cool phones"
```

We can display all the organizations with the command:
```
> scxcli query scx organizations
[
  {
    "address": "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a",
    "name": "HahaTransport",
    "description": "A company transporting everything",
    "approved": true
  },
  {
    "address": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
    "name": "HahaPhone",
    "description": "A company making cool phones",
    "approved": true
  },
  {
    "address": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "name": "Acme",
    "description": "A company making everything",
    "approved": true
  }
]
```


Let's create the products manufactured by Acme:
```
> scxcli tx scx create-product  --from Acme batteryX --product-description "A long lasting battery"

> scxcli tx scx create-product  --from Acme processorX --product-description "A good processor"
```

Display a product:
```
> scxcli query scx product batteryX 
{
  "name": "batteryX",
  "description": "A long lasting battery",
  "manufacturer": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
  "count": "0"
}
```

Let's create the product manufactured by HahaPhone
```
> scxcli tx scx create-product  --from HahaPhone phoneX --product-description "A revolutionary phone"
```

Let's create an unit of BatteryX
```
> scxcli tx scx create-unit --from Acme batteryX --unit-details "Made in China"
{
  "height": "0",
  "txhash": "A485E2B7EE8521AE53BD29CA303E78FA38DAE8042AA53354B73A438DFEEE75FC",
  "raw_log": "[]"
}
```

The reference of the unit can be retrieve from the TxHash:
```
> scxcli query tx A485E2B7EE8521AE53BD29CA303E78FA38DAE8042AA53354B73A438DFEEE75FC
```
It can be retrieved in the events:
```
{
  "type": "create_unit",
    "attributes": [
    {
      "key": "module",
      "value": "scx"
    },
    {
      "key": "reference",
      "value": "e939fc9494b6803a1c6c89db27398215"
    },
    {
      "key": "product",
      "value": "batteryX"
    }
  ]
},
```
```
> scxcli query scx unit e939fc9494b6803a1c6c89db27398215
{
  "reference": "e939fc9494b6803a1c6c89db27398215",
  "product": "batteryX",
  "details": "Made in China",
  "components": null,
  "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
  "holder_history": null,
  "component_of": ""
}
```

Let's create other units:
```
> scxcli tx scx create-unit --from Acme batteryX --unit-details "Made in China"
> scxcli tx scx create-unit --from Acme batteryX --unit-details "Made in China"
```

We can retrieve all the units of a product:
```
> scxcli query scx product-units batteryX
[
  {
    "reference": "e939fc9494b6803a1c6c89db27398215",
    "product": "batteryX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "holder_history": null,
    "component_of": ""
  },
  {
    "reference": "b9732b722b797fe6e5c16a0021c989e9",
    "product": "batteryX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "holder_history": null,
    "component_of": ""
  },
  {
    "reference": "99ec72066affe7424dddd4046c159a38",
    "product": "batteryX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "holder_history": null,
    "component_of": ""
  }
]
```

Let's create some processors
```
> scxcli tx scx create-unit --from Acme processorX --unit-details "Made in China"
> scxcli tx scx create-unit --from Acme processorX --unit-details "Made in China"

> scxcli query scx product-units processorX
[
  {
    "reference": "65297cbe8b4d462bb8eb70a1941a2156",
    "product": "processorX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "holder_history": null,
    "component_of": ""
  },
  {
    "reference": "d36002e32dfc0cc372b89eed3c9bdea3",
    "product": "processorX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "holder_history": null,
    "component_of": ""
  }
]
```

Let's transfer a BatteryX to HahaPhone passing through HahaTransport
```
> scxcli tx scx transfer-unit --from Acme e939fc9494b6803a1c6c89db27398215 cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a

> scxcli tx scx transfer-unit --from HahaTransport e939fc9494b6803a1c6c89db27398215 cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt
```

We can observe the new holder of the unit:
```
> scxcli query scx unit e939fc9494b6803a1c6c89db27398215
{
  "reference": "e939fc9494b6803a1c6c89db27398215",
  "product": "batteryX",
  "details": "Made in China",
  "components": null,
  "holder": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
  "holder_history": [
    "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a"
  ],
  "component_of": ""
}
```

We can get the trace of the unit with the trace command to get more details:
```
> scxcli query scx unit-trace e939fc9494b6803a1c6c89db27398215
[
  {
    "address": "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "name": "Acme",
    "description": "A company making everything",
    "approved": true
  },
  {
    "address": "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a",
    "name": "HahaTransport",
    "description": "A company transporting everything",
    "approved": true
  },
  {
    "address": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
    "name": "HahaPhone",
    "description": "A company making cool phones",
    "approved": true
  }
]
```

Let's transfer a ProcessorX to HahaPhone passing through HahaTransport
```
> scxcli tx scx transfer-unit --from Acme 65297cbe8b4d462bb8eb70a1941a2156 cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a

> scxcli tx scx transfer-unit --from HahaTransport 65297cbe8b4d462bb8eb70a1941a2156 cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt
```

HahaPhone holds a BatteryX and a ProcessorX, it can manufacture a PhoneX composed of these two components:
```
> scxcli tx scx create-unit --from HahaPhone phoneX --unit-details "Made in Japan" --unit-components "e939fc9494b6803a1c6c89db27398215,65297cbe8b4d462bb8eb70a1941a2156"
{
  "height": "0",
  "txhash": "CD59A91B3B6DBB741423A9CA97DD66F84BFFE2318C6429CCEF65C869FD00A5E3",
  "raw_log": "[]"
}

> scxcli query tx CD59A91B3B6DBB741423A9CA97DD66F84BFFE2318C6429CCEF65C869FD00A5E3
{
  "type": "create_unit",
    "attributes": [
    {
      "key": "module",
      "value": "scx"
    },
    {
      "key": "reference",
      "value": "92df7a71a87874b2b29a56ad2b949ed8"
    },
    {
      "key": "product",
      "value": "phoneX"
    }
  ]
}

> scxcli query scx unit 92df7a71a87874b2b29a56ad2b949ed8
{
  "reference": "92df7a71a87874b2b29a56ad2b949ed8",
  "product": "phoneX",
  "details": "Made in Japan",
  "components": [
    "e939fc9494b6803a1c6c89db27398215",
    "65297cbe8b4d462bb8eb70a1941a2156"
  ],
  "holder": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
  "holder_history": null,
  "component_of": ""
}
```

We can now observe that the BatteryX `e939fc9494b6803a1c6c89db27398215` is a component of the new PhoneX `92df7a71a87874b2b29a56ad2b949ed8`:
```
> scxcli query scx unit e939fc9494b6803a1c6c89db27398215
{
  "reference": "e939fc9494b6803a1c6c89db27398215",
  "product": "batteryX",
  "details": "Made in China",
  "components": null,
  "holder": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
  "holder_history": [
    "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
    "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a"
  ],
  "component_of": "92df7a71a87874b2b29a56ad2b949ed8"
}
```

We can get the details of the components composing a phoneX with the `unit-components` command:
```
> scxcli query scx unit-components 92df7a71a87874b2b29a56ad2b949ed8
[
  {
    "reference": "e939fc9494b6803a1c6c89db27398215",
    "product": "batteryX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
    "holder_history": [
      "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
      "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a"
    ],
    "component_of": "92df7a71a87874b2b29a56ad2b949ed8"
  },
  {
    "reference": "65297cbe8b4d462bb8eb70a1941a2156",
    "product": "processorX",
    "details": "Made in China",
    "components": null,
    "holder": "cosmos1e45uknzep6ra0l9y9p2097maswh95drrj7usqt",
    "holder_history": [
      "cosmos1mmhdzqnrfhk8q0z83jn8ylc727jcewq5x5rm3e",
      "cosmos1fxkdqzgmnylga0l4q4kreegltgaxhfq8r9dd8a"
    ],
    "component_of": "92df7a71a87874b2b29a56ad2b949ed8"
  }
]
```
