# GoLang-EVM

GoLang-EVM is a Go program that connects to an Ethereum node and retrieves information about the latest blocks, transactions, gas price, network ID, and more.

## Prerequisites

- Go 1.16 or higher
- Ethereum node URL (e.g., Infura, QuickNode)
- Ethereum address for balance retrieval

## Getting Started

1. Clone the repository:

   ```shell
   git clone <repository_url>

2. Set up the environment variables:

- Create a .env file in the project root directory.
- Add the following lines to the .env file:

INFURA_ID=<your_infura_id>
QUICKNODE_URL=<your_quicknode_url>

- Replace <your_infura_id> with your Infura Project ID.
- Replace <your_quicknode_url> with the URL of your QuickNode Ethereum node.

3. Build and run the program:

go build
./GoLang-EVM

4. View the output:

The program will connect to the Ethereum node and fetch information about the latest blocks, transactions, gas price, network ID, and balance of the specified Ethereum address.


