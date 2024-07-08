# example
**example** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/example@latest! | sudo bash
```
`username/example` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## Instructions to test 
Run `ignite chain build` to rebuild your blockchain with the changes.
Start your blockchain: `ignite chain serve`

format : `[blockchain application binary executable] [tx: transaction] [module] [action] [Postid] [flag specifiying account initiating tx] [flag specifying blockchain network's chain ID]`

Create a Post:
`exampled tx example create-post hello world --from alice --chain-id example`

List Posts:
`exampled q example list-post`

Get a specific post:
format : `exampled q example show-post [id]`
eg. `exampled q example show-post 0`

Update Post:
format : `exampled q example update-post [id] "Updated Title" "Updated Content" `
eg. `exampled tx example update-post 0 "Hello" "Cosmos" --from alice --chain-id example`

Delete Post:
`exampled tx example delete-post 0 --from alice --chain-id example`

Example of the command above: 
![image](https://github.com/rouzhen/miniblockchain/assets/119156668/78486ad0-1b86-4b6e-9b8a-a5667df2f36f)

![image](https://github.com/rouzhen/miniblockchain/assets/119156668/b2d3d371-3c50-4e9e-b8c9-07ea1209b9dd)

![image](https://github.com/rouzhen/miniblockchain/assets/119156668/233979a3-50a6-4674-b99c-7fa73fe45408)



