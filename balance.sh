#!/bin/sh
name="client"
end=20
for ((i=1;i<=end;i++));do
  res=$name$i
  export res=$(gamed keys show $res -a)
  echo $res
  gamed query bank balances $res

done

