package main

import (
    "errors"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type Customer struct {
    UID     string
    Name    string
    Address struct {
        StreetNo string
        Country  string
    }
}

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
    fmt.Printf("initialization done!!!")
    fmt.Printf("initialization done!!!")

    return nil, nil
}


func (t *SimpleChaincode) setDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    if len(args) < 3 {
        return nil, errors.New("insert Into Table failed. Must include 3 column values")
    }

    customer := &Customer{}
    customer.UID = args[0]
    customer.Name = args[1]
    customer.Address.Country = args[2]

    raw, err := json.Marshal(customer)
    if err != nil {
        return nil, err
    }

    err := stub.PutState(customer.UID, raw)
    if err != nil {
        return nil, err
    }

    return nil, nil
}

func (t *SimpleChaincode) getDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    return stub.GetState(args[0])
}





func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
    function, args := stub.GetFunctionAndParameters()
    fmt.Printf("Inside Invoke %s", function)
    if function == "setDetails" {
        return t.setDetails(stub, args)

    } else if function == "getDetails" {
        return t.getDetails(stub, args)
    }

    return nil, errors.New("Invalid invoke function name. Expecting  \"query\"")
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}