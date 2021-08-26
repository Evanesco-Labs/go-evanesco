## Evanesco Main Chain

The goal of Evanesco Main Chain is to bring programmability and interoperability to Evanesco Ecology. In order to embrace the existing popular community and advanced technology, it will bring huge benefits by staying compatible with all the existing smart contracts on Ethereum and Ethereum tooling. And to achieve that, the easiest solution is to develop based on go-ethereum fork, as we respect the great work of Ethereum very much.

Evanesco Main Chain starts its development based on go-ethereum fork. So you may see many toolings, binaries and also docs are based on Ethereum ones.

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)
[![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.com/invite/VnYXBSF)

**The Evanesco Main Chain** will be:

- **A self-sovereign blockchain**: Provides security and safety with miner and elected validators.
- **EVM-compatible**: Supports all the existing Ethereum tooling along with faster finality and cheaper transaction fees.
- **Interoperable**: Comes with efficient native dual chain communication; Optimized for scaling high-performance dApps that require fast and smooth user experience.
- **Distributed with on-chain governance**: GPOW brings in decentralization and community participants. As the native token, EVA will serve as both the gas of smart contract execution and tokens for staking.

Cross-chain transfer and other communication are possible due to native support of interoperability. CrossWrapper and on-chain contracts are developed to support that. 

More details in [White Paper](https://evanesco.org/assets/whitepaper.pdf).

## Key features

### Grandpa over PoW
Although Proof-of-Work (PoW) has been approved as a practical mechanism to implement a decentralized network, it is not friendly to the environment and also requires a large size of participants to maintain the security and consistency.

To combine PoW and Grandpa for consensus, Evanesco Main Chain implement a novel consensus engine called GPoW that:

GPoW consensus includes two layers of consensus mechanisms, which are nested, influence each other and play different roles. GPoW algorithm not only provides almost real-time, asynchronous and safe finality similar to GRANDPA algorithm, but also can fairly distribute new tokens according to PoW, enabling a wider range of communities to go in for the construction of the whole ecology.
The basic steps of the GPoW consensus are:
1.	When the whole network starts, the network miners start to run PoW algorithm and transmit data based on the privacy cascade communication protocol.
a.In the case of transaction or routing data, public or cascaded private transmission is set based on transaction settings
b.In the case of a PoW block, it is publicly broadcast to nearby network nodes
c.On average, the network miner calculates a PoW block and broadcasts it every 10 minutes

2.	The two-layer Sorter network packages the broadcast transactions into blocks, and determines the final consistency of the whole chain according to GRANDPA protocol.
a.In the case of transaction data, the block-generating person is obtained according to the drawing algorithm, and the block is generated and determined as final (second level)
b.In the case of a PoW block, the most suitable block is determined according to the content of the block broadcast by the block-generating person and the finality is determined (10 minutes)


Now we are at the stage of **α-testnet**, Evanesco α Chain introduces a system of 7 validators with POA consensus that can support Privacy Account and EVM-compatible privacy-middleware. It is very easy to test the functionality of Evanesco.


## Native Token

EVA will run on Evanesco Main Chain in the same way as ETH runs on Ethereum so that it remains as `native token` for Evanesco. This means,
EVA will be used to:

1. pay `gas` to deploy or invoke Smart Contract on Evanesco Main Chain
2. perform cross-chain operations, such as transfer token assets across Evanesco Main Chain and Ethereum.

## Building the source

Many of the below are the same as or similar to go-ethereum.

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/install-and-build/installing-geth).

Building `eva` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make eva
```

or, to build the full suite of utilities:

```shell
make all
```

## Running `eva`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `eva` instance.

### Hardware Requirements

The hardware must meet certain requirements to run a full node.
- VPS running recent versions of Mac OS X or Linux.
- 100G of SSD storage for Avis testnet and another 10G if you start miner on full node.
- 2 gigabytes of memory (RAM) for Avis testnet and another 4 gigabytes of memory (RAM) if you start miner on full node.
- A broadband Internet connection with upload/download speeds of at least 10 megabyte per second

### Run Full node to Join Avis TestNet

#### 1. Preparation
First, go to the source code directory `go-evaneso` and follow the instruction above to build `eva`.
Then, make a new directory. Here we name this new directory `avisnode`.
```shell
mkdir avisnode
```
Copy these 4 files into `avisnode` with this command: 
```shell
cp ./build/bin/eva ./avisnode
cp ./verifykey.txt ./avisnode
cp ./avis.json ./avisnode
cp ./avis.toml ./avisnode
```
Make a new directory to store blockchain data with this command:
```shell
mkdir data
```

#### 2. Generate Account
Generate a new account with this command, and remember the password you entered for this account:
```shell
./eva --datadir data account new
```

#### 3. Init Genesis Block
Init genesis block with this command:
```shell
./eva --datadir data init ./avis.json
```

####4. Start Full Node 
Start up full node with this command:
```shell
./eva --datadir ./data --syncmode 'full' --port 30303 --rpc --rpcaddr '0.0.0.0' --rpccorsdomain "*" --rpcport 8545 --rpcapi 'personal,eth,net,web3,txpool,miner,clique' --ws --ws.addr '0.0.0.0' --ws.port 7777 --ws.api 'personal,eth,net,web3,txpool,miner,clique' --zkpvkpath ./verifykey.txt --config ./avis.toml
```

This will connect your node to the Avis TestNet, logs like the following will be printed:
```shell
INFO [08-25|20:25:23.922] Block synchronisation started 
INFO [08-25|20:25:24.861] Downloader queue stats                   receiptTasks=0 blockTasks=0 itemSize=592.71B throttle=8192
INFO [08-25|20:25:24.863] Imported new chain segment               blocks=11 txs=0 mgas=0.000 elapsed=1.809ms mgasps=0.000 number=495 hash=f90dce..88cb1c dirty=0.00B      ignored=1
INFO [08-25|20:25:25.630] Imported new chain segment               blocks=1  txs=0 mgas=0.000 elapsed="293.044µs" mgasps=0.000 number=496 hash=a9e7a8..46b0bc dirty=0.00B
INFO [08-25|20:25:31.130] Imported new chain segment               blocks=1  txs=0 mgas=0.000 elapsed="462.746µs" mgasps=0.000 number=497 hash=2beba7..83b028 dirty=0.00B
INFO [08-25|20:25:37.128] Imported new chain segment               blocks=1  txs=0 mgas=0.000 elapsed="216.098µs" mgasps=0.000 number=498 hash=fef60d..a92615 dirty=0.00B
INFO [08-25|20:25:43.612] Imported new chain segment               blocks=1  txs=0 mgas=0.000 elapsed="725.186µs" mgasps=0.000 number=499 hash=0043f8..05f977 dirty=0.00B
```
The http RPC port is 8545 and the WebSocket rpc port is 7777. Try use RPC request to check block number with this command:
```shell
curl --location --request POST 'localhost:8545/' \
--header 'Content-Type: application/json' \
--data-raw '{
	"jsonrpc":"2.0",
	"method":"eth_blockNumber",
	"params":[],
	"id":83
}'
```

####5. Start Full Node with Miner
Before starting to mine, you also need to download a ZKP prove key file `provekey.txt`. This is a unique ZKP prove key, and miner have to load this ZKP prove key to start GPow working.

Please download from IPFS, IPFS CID:

Copy this file to the `avisnode` directory, and start full node and miner with this command:
```shell
./eva --datadir ./data --syncmode 'full' --port 30303 --rpc --rpcaddr '0.0.0.0' --rpccorsdomain "*" --rpcport 8545 --rpcapi 'personal,eth,net,web3,txpool,miner,clique' --ws --ws.addr '0.0.0.0' --ws.port 7777 --ws.api 'personal,eth,net,web3,txpool,miner,clique' --zkpminer --zkppkpath ./provekey.txt --zkpvkpath ./verifykey.txt --config ./avis.toml
```

This command will start miner with the account you just created and send reward to this address if your mining work has the best score.

Set flag `--zkpkeypath` the path of your keyfile, if you want to derive miner address from this keyfile.

Set flag `----zkpcoinbase` the coinbase address, if you want to receive mining rewards to this address.  

### Configuration

As an alternative to passing the numerous flags to the `eva` binary, you can also pass a
configuration file via:

```shell
$ eva --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ eva --your-favourite-flags dumpconfig
```

### Programmatically interfacing `eva` nodes

As a developer, sooner rather than later you'll want to start interacting with `eva` and the
Evanesco network via your own programs and not manually through the console. To aid
this, `eva` has built-in support for a JSON-RPC based APIs, as on Ethereum, ([standard APIs](https://eth.wiki/json-rpc/API))

These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `eva`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

* `--http` Enable the HTTP-RPC server
* `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
* `--http.port` HTTP-RPC server listening port (default: `8545`)
* `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
* `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
* `--ws` Enable the WS-RPC server
* `--ws.addr` WS-RPC server listening interface (default: `localhost`)
* `--ws.port` WS-RPC server listening port (default: `8546`)
* `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
* `--ws.origins` Origins from which to accept websockets requests
* `--ipcdisable` Disable the IPC-RPC server
* `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
* `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `eva` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
Evanesco Main Chain nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**

## Contribution

Thank you for considering to help out with the source code! We welcome contributions
from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to evanesco, please fork, fix, commit and send a pull request
for the maintainers to review and merge into the main code base. If you wish to submit
more complex changes though, please check up with the core devs first on [our discord channel](https://discord.com/invite/VnYXBSF)
to ensure those changes are in line with the general philosophy of the project and/or get
some early feedback which can make both your efforts much lighter as well as our review
and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

* Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting)
  guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
* Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary)
  guidelines.
* Pull requests need to be based on and opened against the `master` branch.
* Commit messages should be prefixed with the package(s) they modify.
    * E.g. "eth, rpc: make trace configs optional"


## License

The evanesco library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The evanesco binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
