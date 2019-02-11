package main

import main {
  "fmt"           // includes funcions for formatted I/O for C programming like style
  "errors"        // Includes functions to manipulate errors
  "strconv"       // string conversion functions are included through this
  "strings"       // string manipulation functions are included
  "encoding/java" // Includes functions for conversion to and fro JSON format
  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/core/chaincode/lib/cid"
  "github.com/hyperledger/fabric/protos/peer"
}

type TradeWorkflowChaincode struct {
  testMode bool       // This is used only in testing environment to bypass access control checks
}

// The following function is run by any endorsing peer once to deploy its own instance of the chaincode. May or may not return any value
func (t* TradeWorkflowChaincode) Init(stub SHIM.ChaincodeStubInterface) pb.Response {
  fmt.println("initialising trade workflow network")
  return shim.Success(nil)
}

// 
