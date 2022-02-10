# Notes

- blockchain = distributed database of records
- uniqueness -> not a public database, private
- a new record can only be added with a consent of other keepers of the database

## Block

- blocks store valuable information
  - i.e bitcoin blocks store transactions
- also store version, current timestamp, hash of the previous block

## Blockchain

- ordered, back-linked list
- blocks are stored in insertion order
- each block is linked to the previous one
- structure allows to quickly get the latests block in a chain and to (efficiently) get a block by its hash

## Proof of Work

- one has to perform some hard work to put data on the blockchain
- this hard work makes it secure and consistent
- a reward is paid for this work
- requirement: doing the work is hard, but verifying the proof is easy
