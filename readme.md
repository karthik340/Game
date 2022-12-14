# Game
**Game** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve 
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Table of Contents
* Environment
* Problem statement explained
* Code explained (How solution is approached)
* Configuration
* How to play lottery game (walk through)
* Strategies

## Environment
* Ignite CLI version: v0.22.1
* OS: darwin
* arch: arm64
* go version:  go version go1.19.3 darwin/amd64

## Problem statement explained
### Rules of Lottery 
* Anyone can enter the lottery as long as they have enough funds
* Winner is chosen at the end of block, if lottery has 10 or more valid lottery transactions
* if there weren't enough transactions, lottery continues
* once a winner is chosen, pay
* out is sent and next lottery cycle begins

### Enter Lottery transaction
* Valid only when sender has enough funds to cover lottery fee + minimal bet
* Only 1 enter lottery transaction is valid per account per lottery
* Transaction Fields:
  * Lottery Fee (in tokens, e.g. 5token )
  * Bet Size (in tokens, e.g. 100token )
    * User can only bet between 1 to 100 tokens
### Lottery Blocks
* After each lottery zero out transaction counter. counter increases for each new acceptable lottery transaction
* when the counter hits 10, at the end of the block, calculate the lottery on all previous transactions (could be more than 10)
* If the counter doesn’t reach 10, continue in the next block
* the chosen block proposer can't have any lottery transactions with itself as a sender, if this is the case, then the lottery won’t fire this block, and continue on the next one
* if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution

### Choosing a winner
* At the end of the lottery, on the block end, append the data of the transactions (retaining their order) , then hash the data to get the result.
* Take the lowest 16 bits of the resulting hash and do a modulo on the number of lottery transactions in the block to determine the winner!

```
winner_index = (hash_result ^ 0xFFFF) % number_of_transactions_in_block
```
* Example: if winner_index is 2, the sender of the third transaction is the winner

### Payout Calculation
* If the winner placed the highest bet the entire pool is paid to the winner (including fees)
* if the winner paid the lowest bet, no reward is given, lottery total pool is carried over
* All other results: winner is paid the sum of all bets (without fees) in the current lottery only
* Payment is from the lottery pool

### Demo
* Configure the chain to have 20 clients each with 500 tokens, and a block every 5 minutes
* Build & run the chain
* Place the following bets for all clients:
  * client1: 1token
  * client2: 2token
  * client3: 3token
  * ...
  * client20: 20token
* Repeat 100 times or until all clients run out of funds
* Show the blocks, results and balances

## Code explained (How solution is approached)
### Data Modelling
```
RoundKey = "Round-value-"
```
* round key used to store the current round and it is incremented at the end of block after picking winner
```
TxnCounterKey = "TxnCounter-value-"
```
* transaction counter key is used to store number of bets in ongoing round
* It is incremented each time new bet comes and stored in that bet


* We need a structure called bet which holds betSize, Fee, winStatus, transaction number
```

message Bet {
  string sender = 1;
  cosmos.base.v1beta1.Coin fee = 2
  [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin"];

  cosmos.base.v1beta1.Coin bet = 3
  [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin"];

  bool status = 4;

  uint64 txNum =5;
}
```
* sender tells who sent the bet
* win status tells if bet won lottery
* txNum can be used during payout
* Key for storing bet structure is
```
 (BetKeyPrefix + round + sender) -> Bet
```
* BetKeyPrefix helps in distinguishing from other keys
* whenever a transaction comes we need to check if user is sending another transaction in the same lottery round, so if we use round as key then we don't need to traverse all bets in all rounds.
* If we use sender also in key, then we don't need to traverse all senders in the round.
* using this key we can also fetch, how every user performed in particular round.
* conclusion :  data is stored in a such a way that it decreases search space

### Core Logic
* user sends the fees and bet size in transaction
```
message MsgPlaceBet {
  string creator = 1;
  cosmos.base.v1beta1.Coin fee = 2
  [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin"];

  cosmos.base.v1beta1.Coin bet = 3
  [(gogoproto.nullable) = false, (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coin"];
}
```
* This hits the handler defined under app module
* this routes to the function registered under MsgServer
* Algorithm
```
func handler(){
1. validate bet
2. check if sender has already placed bet in current round
if yes {
2.1 return his previous fee and bet to him 
2.2 send money from module account (pool) to sender 
2.3. use same txNum as bet exists already
2.4. update with new betSize and store in kv store
} else {
2.1 send money from module account (pool) to sender 
2.2 Add bet to kv store
2.3 increment txnCounter
}
```
* EndBlocker Algorithm
```
func EndBlocker(){
1. If proposer has a bet in the lottery{
  return
}
2. if txnCount < 10{
 return
}
3. serialize bets
4. append round numer to serialized data so that there will be randomness even if transactions are same as previous round
5. take last 16 bits of Keccak(serialized data) % txnCounter
6. bet whose transaction number matches above result is the winner
7. check if winner bet is highest or lowest
8. if highest, pay him totalBet + totalFees collected in round
9. if lowest, pay nothing
10. inother cases pay only totalBet collected in round
11. increment round
12. set transaction counter to zero
13. update winner status in the bet struct
}
```
## Configuration
* default config files generated by ignite can be updated with help of config.yml
* update config.yml with 20 clients and give each one 500tokens, which updates genesis.json in home/.game/config/genesis.json
* update timeout_commit to 5 minutes in order to create delay of 5 minutes in between blocks
* timeout_commit gets updated with 5 minutes in home/.game/config/config.toml
* timeout_commit creates delay in consensus after committing the block

##  How to play lottery game (walk through)
* to start chain run: ignite chain serve 
* wait 5 minutes as it takes 5 minutes to generate genesis block
* place bet using
* export alice=$(gamed keys show alice -a)
* gamed tx lottery place-bet 5token 100token --from $alice --gas auto  --yes
* gamed query bank balances $alice
```
ignite chain serve 
export alice=$(gamed keys show alice -a)
gamed tx lottery place-bet 5token 100token --from $alice --gas auto  --yes
gamed query bank balances $alice
```
* steps to place bet by 20 clients of sizes 1,2,3 ... 20 for 100 times 
* as soon as you start chain , run ./simulate_bet.sh in other terminal
* simulate_bet.sh sleeps for 6 minutes initially to wait for genesis block to be generated
* then 20 clients place bets and sleeps for 5 minutes as it takes 5 minutes to move to next block, this happens for 100 times
* prints balances, 100 blocks, rounds results
* use ./balance.sh to print all 20 clients balances
* to get block by height : gamed q block 1 (height)
```
./simulate_bet.sh
./balance.sh // use to know balances exclusively
```

## Strategies
* Assuming uniform random bet from all other clients, what is the best strategy for client1?
  * as bet sizes distributed uniformly and also randomly,lets divide 100 bet sizes among other 19 clients
  * ceil(100/19) = 6 
  * there is a chance of 1 bet in every consecutive 6 numbers
  * probability of winning for any client will be 1/20 as lottery picked in random fashion
  * inorder get tokens even in picked round we should not place lowest bet
  * it's better to place lower bet and earn others bets to maximize profit than place highest bet(100) just to earn fee of 100
  * let's say that on average other 19 clients place bet of 50tokens (19 * 50 = 950) based on uniform random distribution
  * client1 places bet of 7 tokens as there is a chance of any other client paces bet size of (1-6) tokens
  * It costs  (5 fee + 7 bet)*20 = 240 tokens for client1 for 20 rounds
  * if it wins 1 in 20 then it earns 950
  * profit (950-240) = 710 which is profit of 295%
* Assuming uniform random bet from all other clients, and client1 behaves in the strategy mentioned (1.), what is the best strategy for client2?
  * client2 can reduce his risk as he knows client1 will place bet of 7 tokens in every round
  * client2 can place bet of 8 tokens and never loses when he is picked up in lottery
  * or he can take more risk than client1 to increase profit and place bet of 5 tokens in every round
* What is the Nash equilibrium?
  * In Nash equilibrium, strategies of players are considered optimal given strategies of other players, results in a stable state where no player can gain advantage by changing their strategy
  * So nobody wants to place lowest bet in this game, so every one places maximum bet which is 100 tokens
  * when a winner picked up he gets (100+5)*20 = 2500tokens
  * every client spends (100+5)*20 = 2500tokens
  * So there won't be any gain or loss of money in Nash equilibrium
  * If everyone places same bet, then it is considered winner placed highest be in both code and strategy

## Changes in app.go
* give permission for module account to mint and burn tokens, so that we can store tokens in module account
* Add bank keeper and authentication keeper for lottery module 
