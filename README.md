# Notificator
Service that observes DeFi positions on AAVE lending protocol for given eth account and returns useful data, such as health factor.

Then it uses Matrix protocol, to notify person about it (thanks [matrix-commander](https://github.com/8go/matrix-commander))

### Current functionality
The program runs indifinitely (until it crashes lol) and periodically (1 hour) fetches data from TheGraph DeFi protocol that queries AAVE v3 Arbitrum position
for specified ETH account (Arb ETH) and then it sends health factor to specified matrix room ID as a message.

`Health Factor: 1.59 @ 2024-10-09 19:27:16`

<img width="464" alt="image" src="https://github.com/user-attachments/assets/78305642-5517-4cf7-82d0-a0fb0f4da6e9">


## How to run
1. have [matrix-commander](https://github.com/8go/matrix-commander) installed and logged-in
2. have [The Graph](https://thegraph.com/studio/) api key
3. have Golang installed (ofc)
4. set env vars to `.env` (see `.env.example`)
5. Run it
  ```sh
  $ go run main.go
  ```

### or you can build it
```sh
$ go build -o notificator
```
```sh
$ ./notificator
```

or in the background
```sh
$ nohup ./notificator nohup.out
```
## Contribute
All contributions are welcome, I'm building something for myself, but it might satisfy even somebody else's needs. Let's make good shit together
