# Supply Chain X


## Test Environment

The application can be tested with several nodes with a test environment composed of the following:
- `test_env/` contains several files to test the application
- `test_validator.Dockerfile` is used to represent a validator already set in the provided genesis. This docker image contains a predefined private key for the validator
- `test.Dockerfile` is used to represent other daemon nodes in the network, not directly validator but that can apply to become a validator
- `docker-compose.yml` used to lauch the test environment with the command `docker-compose up`

`test_env/` contains the following files:
- `node_key.json` a node key file copied into the initial validator node to have a predefined node id for the validator
- `testkey.json` a predefined private key for the validator, so that we can determine it in the genesis file
- `config.toml` a customized config file for the nodes. This config define the validator node as a persistent peer with its node id
- `genesis.json` a predefined genesis file with a validator and five accounts

The five accounts in the genesis file are the following:

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

The accounts can be imported locally with the commands

```
scxcli keys add regulatorA --recover
scxcli keys add regulatorB --recover
scxcli keys add Acme --recover
scxcli keys add HahaTransport --recover
scxcli keys add HahaPhone --recover
```

