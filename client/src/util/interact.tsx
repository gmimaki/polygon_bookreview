require('dotenv').config();
const alchemyKey = process.env.REACT_APP_ALCHEMY_KEY;
const { createAlchemyWeb3 } = require("@alch/alchemy-web3");
const web3 = createAlchemyWeb3(alchemyKey); 
const contractABI = require("../contract-abi.json");
const contractAddress = process.env.REACT_APP_CONTRACT_ADDRESS;

export const helloWorldContract = new web3.eth.Contract(
	contractABI.abi,
	contractAddress
);

export const connectWallet = async () => {
	if (window.ethereum) {
		try {
			const addressArray = await window.ethereum.request({
			  method: "eth_requestAccounts",
			});
			const obj = {
			  status: "👆🏽 Write a message in the text-field above.",
			  address: addressArray[0],
			};
			return obj;
		} catch (err: any) {
			return {
			  address: "",
			  status: "😥 " + err.message,
			};
		}
	} else {
		return {
			address: "",
			status: (
			  <span>
			    <p>
			      {" "}
			      🦊{" "}
			      <a target="_blank" href={`https://metamask.io/download`}>
				You must install Metamask, a virtual Ethereum wallet, in your
				browser.
			      </a>
			    </p>
			  </span>
			),
		};
	}
};

export const getCurrentWalletConnected = async () => {
	if (window.ethereum) {
		try {
			const addressArray = await window.ethereum.request({
			  method: "eth_accounts",
			});
			if (addressArray.length > 0) {
			  return {
			    address: addressArray[0],
			    status: "👆🏽 Write a message in the text-field above.",
			  };
			} else {
			  return {
			    address: "",
			    status: "🦊 Connect to Metamask using the top right button.",
			  };
			}
		} catch (err: any) {
			return {
			  address: "",
			  status: "😥 " + err.message,
			};
		}
	} else {
		return {
			address: "",
			status: (
			  <span>
			    <p>
			      {" "}
			      🦊{" "}
			      <a target="_blank" href={`https://metamask.io/download`}>
				You must install Metamask, a virtual Ethereum wallet, in your
				browser.
			      </a>
			    </p>
			  </span>
			),
		};
	}
};

export const updateMessage = async (address: string, bookName: string, authorName: string) => {
	if (!window.ethereum || address === null) {
		return {
		  status:
		    "💡 Connect your Metamask wallet to update the message on the blockchain.",
		};
	}
	    
	if (bookName.trim() === "") {
		return {
			status: "❌ Name of Book cannot be an empty string.",
		};
	}
	if (authorName.trim() === "") {
		return {
			status: "❌ Name of Author cannot be an empty string.",
		};
	}

	const transactionParameters = {
		to: contractAddress, // Required except during contract publications.
		from: address, // must match user's active address.
		data: helloWorldContract.methods.registerBook(bookName, authorName).encodeABI(),
	};

	try {
		const txHash = await window.ethereum.request({
			method: "eth_sendTransaction",
			params: [transactionParameters],
		});
		return {
			status: (
				<span>
					✅{" "}
					<a target="_blank" href={`https://goerli.etherscan.io/tx/${txHash}`}>
						View the status of your transaction on Etherscan!
					</a>
					<br />
					ℹ️ Once the transaction is verified by the network, the message will
					be updated automatically.
				</span>
			),
		};
	} catch (error: any) {
		return {
			status: "😥 " + error.message,
		};
	}
}

export const getBooks = async() => {
	if (!window.ethereum) {
		return {
		  status:
		    "💡 Connect your Metamask wallet to update the message on the blockchain.",
		};
	}

	const res = await helloWorldContract.methods.findBook().call()
	return res.map((b: any) => {
		return {
			name: b.name,
			author: b.author,
		}
	})
}