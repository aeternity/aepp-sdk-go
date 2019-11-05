/*
Package account contains code for managing account keypairs, generating a HD
wallet as well as saving them securely to disk.

HD wallet functionality per AEX-10
(https://github.com/aeternity/AEXs/blob/master/AEXS/aex-10.md) is implemented in
this fashion:

            +----------------------------+
            |   BIP39 Mnemonic           |
            +-------------+--------------+
                          |
            +-------------+--------------+
            |   Binary Seed              |
            +-------------+--------------+
                          |
            +-------------+--------------+
            |   BIP32 Master Key         |
            +----X-----------------------+
                XXX
                XXX
        +-----X---+      +---------+     +----------+
        |         |      |         |     |          |
        |         |      |         |     |          |
        |         |      |         |     |          |
        +---------XX     +---------+     +----------+
                    XXX
    +---------+     +--XX-----+
    |         |     |         |
    |         |     |         |
    |         |     |         |
    +---------+     +-------XX+
                            X                           +---------------------+
                        +----XX+----+                   |                     |
                        |           |                   |                     |
                        | Child Key +------------------->  aeternity Account  |
                        |           |                   |                     |
                        +-----------+                   |                     |
                                                        +---------------------+



*/
package account
