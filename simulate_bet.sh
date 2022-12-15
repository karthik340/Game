#!/bin/sh
name="client"
token="token"
clientCount=20
rounds=100

sleep 360 # sleep for 6 minutes as first block takes 5 minutes to get initialized

for ((i=1;i<=rounds;i++));do

  for ((i=1;i<=clientCount;i++));do
    result=$name$i # forming client+number
    bet=$i$token # forming number+token

    export client=$(gamed keys show $result -a) # extract address into client variable

    gamed tx lottery place-bet 5token $bet --from $client --gas auto  --yes & # place bet in parallel
  done

  sleep 300 # sleep for 5 minutes as there is 5 minutes delay between blocks

done

for ((i=1;i<=rounds;i++));do # printing all 100 blocks
  gamed q block $i # query block at height i
done

./balance.sh # print balances after 100 rounds

for ((i=0;i<rounds;i++));do # print winners of all 100 rounds
  gamed q lottery get-winner-by-round $i
done


