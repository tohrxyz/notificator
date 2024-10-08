# Notificator
Service that observes DeFi positions on AAVE lending protocol for given eth account and returns useful data, such as health factor.

Then it uses Signal messenger, to notify person about it. 

## Features 🚧
- [x] Get information about DeFi position
- [ ] Send the message notification to person

## How to run
1. set your `ETH_ACCOUNT` (Arbitrum AAVE v3 currently) and `THE_GRAPH_API_KEY` environment variables into your `.env` file
   (cc: `.env.example`. Get the api key from [The Graph Studio](https://thegraph.com/studio/)
2. have Golang installed (ofc)
3. Run it
  ```sh
  $ go run main.go
  ```

## Contribute
All contributions are welcome, I'm building something for myself, but it might satisfy even somebody else's needs. Let's make good shit together
