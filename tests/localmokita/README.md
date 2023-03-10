# LocalMokita

LocalMokita is a complete Mokita testnet containerized with Docker and orchestrated with a simple docker-compose file. LocalMokita comes preconfigured with opinionated, sensible defaults for a standard testing environment.

LocalMokita comes in two flavors:

1. No initial state: brand new testnet with no initial state. 
2. With mainnet state: creates a testnet from a mainnet state export

## Prerequisites

Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

## 1. LocalMokita - No Initial State

The following commands must be executed from the root folder of the Mokita repository.

1. Make any change to the mokita code that you want to test

2. Initialize LocalMokita:

```bash
make localnet-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.mokitad` folder

3. Start LocalMokita:

```bash
make localnet-start
```

> Note
>
> You can also start LocalMokita in detach mode with:
>
> `make localnet-startd`

4. (optional) Add your validator wallet and 9 other preloaded wallets automatically:

```bash
make localnet-keys
```

- These keys are added to your `--keyring-backend test`
- If the keys are already on your keyring, you will get an `"Error: aborted"`
- Ensure you use the name of the account as listed in the table below, as well as ensure you append the `--keyring-backend test` to your txs
- Example: `mokitad tx bank send lo-test2 moki1cyyzpxplxdzkeea7kwsydadg87357qnahakaks --keyring-backend test --chain-id LocalMokita`

5. You can stop chain, keeping the state with

```bash
make localnet-stop
```

6. When you are done you can clean up the environment with:

```bash
make localnet-clean
```

## 2. LocalMokita - With Mainnet State

Running LocalMokita with mainnet state is resource intensive and can take a bit of time.
It is recommended to only use this method if you are testing a new feature that must be thoroughly tested before pushing to production.

A few things to note before getting started. The below method will only work if you are using the same version as mainnet. In other words,
if mainnet is on v8.0.0 and you try to do this on a v9.0.0 tag or on main, you will run into an error when initializing the genesis.
(yes, it is possible to create a state exported testnet on a upcoming release, but that is out of the scope of this tutorial)

Additionally, this process requires 64GB of RAM. If you do not have 64GB of RAM, you will get an OOM error.

### Create a mainnet state export

1. Set up a node on mainnet (easiest to use the [get.mokita.zone](https://get.mokita.zone) tool). This will be the node you use to run the state exported testnet, so ensure it has at least 64GB of RAM.

```sh
curl -sL https://get.mokita.zone/install > i.py && python3 i.py
```

2. Once the installer is done, ensure your node is hitting blocks.

```sh
source ~/.profile
journalctl -u mokitad.service -f
```

3. Stop your Mokita daemon

```sh
systemctl stop mokitad.service
```

4. Take a state export snapshot with the following command:

```sh
cd $HOME
mokitad export 2> state_export.json
```

After a while (~15 minutes), this will create a file called `state_export.json` which is a snapshot of the current mainnet state.

### Use the state export in LocalMokita

1. Copy the `state_export.json` to the `LocalMokita/state_export` folder within the mokita repo

```sh
cp $HOME/state_export.json $HOME/mokita/tests/LocalMokita/state_export/
```

6. Ensure you have docker and docker-compose installed:

```sh
# Docker
sudo apt-get remove docker docker-engine docker.io
sudo apt-get update
sudo apt install docker.io -y

# Docker compose
sudo apt install docker-compose -y
```

7. Build the `local:mokita` docker image:

```bash
make localnet-state-export-init
```

The command:

- Builds a local docker image with the latest changes
- Cleans the `$HOME/.mokitad` folder

3. Start LocalMokita:

```bash
make localnet-state-export-start
```

> Note
>
> You can also start LocalMokita in detach mode with:
>
> `make localnet-state-export-startd`

When running this command for the first time, `local:mokita` will:

- Modify the provided `state_export.json` to create a new state suitable for a testnet
- Start the chain

You will then go through the genesis initialization process. This will take ~15 minutes.
You will then hit the first block (not block 1, but the block number after your snapshot was taken), and then you will just see a bunch of p2p error logs with some KV store logs.
**This will happen for about 1 hour**, and then you will finally hit blocks at a normal pace.

9. On your host machine, add this specific wallet which holds a large amount of moki funds

```sh
MNEMONIC="bottom loan skill merry east cradle onion journey palm apology verb edit desert impose absurd oil bubble sweet glove shallow size build burst effort"
echo $MNEMONIC | mokitad keys add wallet --recover --keyring-backend test
```

You now are running a validator with a majority of the voting power with the same state as mainnet state (at the time you took the snapshot)

10. On your host machine, you can now query the state-exported testnet:

```sh
mokitad status
```

11. Here is an example command to ensure complete understanding:

```sh
mokitad tx bank send wallet moki1nyphwl8p5yx6fxzevjwqunsfqpcxukmtk8t60m 10000000umoki --chain-id testing1 --keyring-backend test
```

12. You can stop chain, keeping the state with

```bash
make localnet-state-export-stop
```

13. When you are done you can clean up the environment with:

```bash
make localnet-state-export-clean
```

Note: At some point, all the validators (except yours) will get jailed at the same block due to them being offline.

When this happens, it may take a little bit of time to process. Once all validators are jailed, you will continue to hit blocks as you did before.
If you are only running the validator for a short time (< 24 hours) you will not experience this.

## LocalMokita Accounts

LocalMokita is pre-configured with one validator and 9 accounts with ION and MOKI balances.

| Account   | Address                                                                                                | Mnemonic                                                                                                                                                                   |
| --------- | ------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| lo-val    | `moki1phaxpevm5wecex2jyaqty2a4v02qj7qmlmzk5a`<br/>`mokivaloper1phaxpevm5wecex2jyaqty2a4v02qj7qm9v24r6` | `satisfy adjust timber high purchase tuition stool faith fine install that you unaware feed domain license impose boss human eager hat rent enjoy dawn`                    |
| lo-test1  | `moki1cyyzpxplxdzkeea7kwsydadg87357qnahakaks`                                                          | `notice oak worry limit wrap speak medal online prefer cluster roof addict wrist behave treat actual wasp year salad speed social layer crew genius`                       |
| lo-test2  | `moki18s5lynnmx37hq4wlrw9gdn68sg2uxp5rgk26vv`                                                          | `quality vacuum heart guard buzz spike sight swarm shove special gym robust assume sudden deposit grid alcohol choice devote leader tilt noodle tide penalty`              |
| lo-test3  | `moki1qwexv7c6sm95lwhzn9027vyu2ccneaqad4w8ka`                                                          | `symbol force gallery make bulk round subway violin worry mixture penalty kingdom boring survey tool fringe patrol sausage hard admit remember broken alien absorb`        |
| lo-test4  | `moki14hcxlnwlqtq75ttaxf674vk6mafspg8xwgnn53`                                                          | `bounce success option birth apple portion aunt rural episode solution hockey pencil lend session cause hedgehog slender journey system canvas decorate razor catch empty` |
| lo-test5  | `moki12rr534cer5c0vj53eq4y32lcwguyy7nndt0u2t`                                                          | `second render cat sing soup reward cluster island bench diet lumber grocery repeat balcony perfect diesel stumble piano distance caught occur example ozone loyal`        |
| lo-test6  | `moki1nt33cjd5auzh36syym6azgc8tve0jlvklnq7jq`                                                          | `spatial forest elevator battle also spoon fun skirt flight initial nasty transfer glory palm drama gossip remove fan joke shove label dune debate quick`                  |
| lo-test7  | `moki10qfrpash5g2vk3hppvu45x0g860czur8ff5yx0`                                                          | `noble width taxi input there patrol clown public spell aunt wish punch moment will misery eight excess arena pen turtle minimum grain vague inmate`                       |
| lo-test8  | `moki1f4tvsdukfwh6s9swrc24gkuz23tp8pd3e9r5fa`                                                          | `cream sport mango believe inhale text fish rely elegant below earth april wall rug ritual blossom cherry detail length blind digital proof identify ride`                 |
| lo-test9  | `moki1myv43sqgnj5sm4zl98ftl45af9cfzk7nhjxjqh`                                                          | `index light average senior silent limit usual local involve delay update rack cause inmate wall render magnet common feature laundry exact casual resource hundred`       |
| lo-test10 | `moki14gs9zqh8m49yy9kscjqu9h72exyf295afg6kgk`                                                          | `prefer forget visit mistake mixture feel eyebrow autumn shop pair address airport diesel street pass vague innocent poem method awful require hurry unhappy shoulder`     |
