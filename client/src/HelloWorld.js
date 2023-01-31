import React from "react";
import { useEffect, useState } from "react";
import {helloWorldContract, connectWallet, updateMessage, loadCurrentMessage, getCurrentWalletConnected} from "./util/interact.js";

import alchemylogo from "./alchemylogo.svg";
import BookList from "./components/BookList.js";

const HelloWorld = () => {
  //state variables
  const [walletAddress, setWallet] = useState("");
  const [status, setStatus] = useState("");
  const [message, setMessage] = useState("No connection to the network."); //default message
  const [newBookName, setNewBookName] = useState("");
  const [newAuthorName, setNewAuthorName] = useState("");

  //called only once
  useEffect(() => {
    async function fetchMessage() {
      await loadCurrentMessage();
    }
    fetchMessage();
    addSmartContractListener();

    async function fetchWallet() {
      const {address, status} = await getCurrentWalletConnected();
      setWallet(address);
      setStatus(status);
    }
    fetchWallet();
    addWalletListener();
  }, []);

  function addSmartContractListener() {
    helloWorldContract.events.NewBook({}, (error, data) => {
      if (error) {
        setStatus("😥 " + error.message);
      } else {
        setMessage(data.returnValues[1]);
        setNewBookName("");
        setStatus("🎉 Your message has been updated!");
      }
    })
  }

  function addWalletListener() {
    if (window.ethereum) {
      window.ethereum.on("accountsChanged", (accounts) => {
        if (accounts.length > 0) {
          setWallet(accounts[0]);
          setStatus("👆🏽 Write a message in the text-field above.");
        } else {
          setWallet("");
          setStatus("🦊 Connect to Metamask using the top right button.");
        }
      });
    } else {
      setStatus(
        <p>
          {" "}
          🦊{" "}
          <a target="_blank" href={`https://metamask.io/download`}>
            You must install Metamask, a virtual Ethereum wallet, in your
            browser.
          </a>
        </p>
      );
    }
  }

  const connectWalletPressed = async () => { 
    const walletResponse = await connectWallet();
    setStatus(walletResponse.status);
    setWallet(walletResponse.address);
  };

  const onUpdatePressed = async () => {
    const { status } = await updateMessage(walletAddress, newBookName, newAuthorName);
    setStatus(status);
  };

  
  //the UI of our component
  return (
    <>
      <div id="container">
      <img id="logo" src={alchemylogo}></img>
      <button id="walletButton" onClick={connectWalletPressed}>
        {walletAddress.length > 0 ? (
          "Connected: " +
          String(walletAddress).substring(0, 6) +
          "..." +
          String(walletAddress).substring(38)
        ) : (
          <span>Connect Wallet</span>
        )}
      </button>

      <h2 style={{ paddingTop: "50px" }}>Current Message:</h2>
      <p>{message}</p>

      <div>
        <h2 style={{ paddingTop: "18px" }}>Name of Book:</h2>
        <input type="text" placeholder="Name of Book" onChange={(e) => setNewBookName(e.target.value)} value={newBookName}/>

        <h2 style={{ paddingTop: "18px" }}>Name of Author:</h2>
        <input type="text" placeholder="Name of Author" onChange={(e) => setNewAuthorName(e.target.value)} value={newAuthorName}/>

        <p id="status">{status}</p>

        <button id="publish" onClick={onUpdatePressed}>Register</button>
      </div>
    </div>

    <BookList />
    </>
  );
};

export default HelloWorld;
