### Node Draft

The goal of the Flashbots node software is to provide mining pools with additional revenue with minimal operational overhead and no latency impact. This is implemented by exposing an API endpoint which gets accessed by DeFi traders, who propose block templates to your pool for a fee. Below are some questions that will help us accomplish that. Thanks for helping us make Flashbots better!

  

Node operation
--------------

1.  Do you typically run a vanilla/standard Geth version or a modified one? If modified, what are the core modifications?
    
2.  Do you use any special/non-standard geth configuration settings?
    
3.  Do you use clients other than Geth?
    
4.  What is a typical number of block templates you send to miners per block? In other words, how frequently do you send new blocks to miners?
    

API
---

5.  The MEV-geth JSON-RPC API needs to be exposed publicly in order to intake transaction bundles from traders. Do you anticipate any issues associated with this?
    

Reporting
---------

6.  Miner reporting. Do you operate custom software that miners use to view their revenue/financial metrics? If so, would you require splitting MEV-generated revenue into a separate line item from the standard mining revenue?
    
7.  Internal reporting. Do you operate custom software that you use to view financial metrics internally? What are the data points that you require from the MEV-geth in order to generate MEV-related metrics for internal reporting?
    

Logging / monitoring
--------------------

8.  What are logging/monitoring solutions you use that you’d like MEV-geth to be integrated with?