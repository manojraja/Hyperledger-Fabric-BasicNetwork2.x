package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Car :  Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type NotarizerData1 struct {
	DocType string `json:"docType"`
	BatchID string `json:"BatchID"`
	//payload
	BlockchainID string `json:"BlockchainID"`
	MetaDataHash string `json:"MetaDataHash"`
	DocHash      string `json:"DocHash"`
	//SRFH
	SROID                string `json:"SROID"`
	RegYear              string `json:"RegYear"`
	BookNumber           string `json:"BookNumber"`
	DocumentType         string `json:"DocumentType"`
	CreatedBy            string `json:"CreatedBy"`
	UpdatedBy            string `json:"UpdatedBy"`
	LastUpdatedTimestamp string `json:"LastUpdatedTimestamp"`
	Doc_Index_Id         string `json:"Doc_Index_Id"`
	DocSeqNo             string `json:"Doc_SeqNo"`
}


var logger = flogging.MustGetLogger("fabcar_cc")

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	NotarizerData1 := []NotarizerData1{
		{DocType: "fdfdsd", BatchID: "SA123", BlockchainID: "Block21", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
		{DocType: "fdfdsd", BatchID: "SA1123", BlockchainID: "Blockw1", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
		{DocType: "fdfdsd", BatchID: "SA1223", BlockchainID: "Blockr1", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
		{DocType: "fdfdsd", BatchID: "SA1423", BlockchainID: "Bloeck1", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
		{DocType: "fdfdsd", BatchID: "SA1423", BlockchainID: "Blockf1", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
		{DocType: "fdfdsd", BatchID: "SA12345", BlockchainID: "Blowck1", MetaDataHash: "Tomoko", DocHash: "dfdfdf", SROID: "dsd23", RegYear: "sdsd", BookNumber: "dd2323", DocumentType: "dsdsd", CreatedBy: "Harini", UpdatedBy: "Harini", LastUpdatedTimestamp: "d232323", Doc_Index_Id: "Harini123", DocSeqNo: "jhbds"},
	}

	for _, NotarizerData1 := range NotarizerData1 {
		NotarizerData1JSON, err := json.Marshal(NotarizerData1)
		if err != nil {
			return err
		}

		err = APIstub.PutState(NotarizerData1.BlockchainID, NotarizerData1JSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
}
}
func (s *SmartContract) invokeNotarizer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 14  {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}
	NotarizerData1 := NotarizerData1{
		DocType: args[0],
		BatchID: args[1],
		//payload
		BlockchainID: args[2],
		MetaDataHash: args[3],
		DocHash:      args[4],
		//SRFH
		SROID:                args[5],
		RegYear:              args[6],
		BookNumber:           args[7],
		DocumentType:         args[8],
		CreatedBy:            args[9],
		UpdatedBy:            args[10],
		LastUpdatedTimestamp: args[11],
		Doc_Index_Id:         args[12],
		DocSeqNo:             args[13],
	}
	NotarizerData1JSON, err := json.Marshal(NotarizerData1)
	if err != nil {
		return err
	}

	return APIstub.PutState(NotarizerData1JSON.BlockchainID, NotarizerData1JSON)
	
}
func (s *SmartContract) queryNotarizer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

// Invoke :  Method for INVOKING smart contract



// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
