package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

type SmartContract struct {
}

type serverConfig struct {
	CCID    string
	Address string
}

//const index = "color~name"

var logger = flogging.MustGetLogger("notarizer_cc")

// SimpleChaincode example simple Chaincode implementation
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

type HistoryQueryResult struct {
	Record    *NotarizerData1 `json:"record"`
	TxId      string          `json:"txId"`
	Timestamp time.Time       `json:"timestamp"`
	IsDelete  bool            `json:"isDelete"`
}

func (t *SmartContract) Init(APIstub shim.ChaincodeStubInterface) error {

	fmt.Println("In init fmt")
	log.Printf("In init fmt")
	logger.Info("In intt logger")
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

		err = ctx.GetStub().PutState(NotarizerData1.BlockchainID, NotarizerData1JSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil

}



func (t *SmartContract) StoreData(APIstub shim.ChaincodeStubInterface, docType, batchID string, blockchainID string, metaDataHash string, docHash string, sROID string, regYear string, bookNumber string, documentType string, createdBy string, updatedBy string, lastUpdatedTimestamp string, doc_Index_Id string, docSeqNo string) error {
	exists, err := t.DataExists(APIstub, blockchainID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the NotarizerData1 %s already exists", blockchainID)
	}

	// fmt.Print("Invoke is running StoreData")
	// logger.Info("In storedata batchid", batchID)

	NotarizerData1 := NotarizerData1{
		DocType: docType,
		BatchID: batchID,
		//payload
		BlockchainID: blockchainID,
		MetaDataHash: metaDataHash,
		DocHash:      docHash,
		//SRFH
		SROID:                sROID,
		RegYear:              regYear,
		BookNumber:           bookNumber,
		DocumentType:         documentType,
		CreatedBy:            createdBy,
		UpdatedBy:            updatedBy,
		LastUpdatedTimestamp: lastUpdatedTimestamp,
		Doc_Index_Id:         doc_Index_Id,
		DocSeqNo:             docSeqNo,
	}
	NotarizerData1JSON, err := json.Marshal(NotarizerData1)
	if err != nil {
		return err
	}

	return APIstub.GetStub().PutState(blockchainID, NotarizerData1JSON)
}

func (s *SmartContract) UpdateData(APIstub shim.ChaincodeStubInterface, docType, batchID string, blockchainID string, metaDataHash string, docHash string, sROID string, regYear string, bookNumber string, documentType string, createdBy string, updatedBy string, lastUpdatedTimestamp string, doc_Index_Id string, docSeqNo string) error {
	exists, err := s.DataExists(APIstub, blockchainID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", blockchainID)
	}

	// overwritting original asset with new asset
	NotarizerData1 := NotarizerData1{
		DocType: docType,
		BatchID: batchID,
		//payload
		BlockchainID: blockchainID,
		MetaDataHash: metaDataHash,
		DocHash:      docHash,
		//SRFH
		SROID:                sROID,
		RegYear:              regYear,
		BookNumber:           bookNumber,
		DocumentType:         documentType,
		CreatedBy:            createdBy,
		UpdatedBy:            updatedBy,
		LastUpdatedTimestamp: lastUpdatedTimestamp,
		Doc_Index_Id:         doc_Index_Id,
		DocSeqNo:             docSeqNo,
	}
	NotarizerData1JSON, err := json.Marshal(NotarizerData1)
	if err != nil {
		return err
	}

	return APIstub.GetStub().PutState(blockchainID, NotarizerData1JSON)
}

func (s *SmartContract) DataExists(APIstub shim.ChaincodeStubInterface, blockchainID string) (bool, error) {
	NotarizerData1JSON, err := APIstub.GetStub().GetState(blockchainID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return NotarizerData1JSON != nil, nil
}

// func (t *SmartContract) VerifyData(ctx contractapi.TransactionContextInterface, blockchainID string) (*NotarizerData1, error) {

// 	NotarizerData1JSON, err := ctx.GetStub().GetState(blockchainID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read from world state: %v", err)
// 	}
// 	if NotarizerData1JSON == nil {
// 		return nil, fmt.Errorf("the asset %s does not exist", blockchainID)
// 	}

// 	var asset NotarizerData1
// 	err = json.Unmarshal(NotarizerData1JSON, &asset)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &asset, nil
// }

func (t *SmartContract) GetHistoryForTransaction(APIstub shim.ChaincodeStubInterface, blockchainID string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: batchID %v", blockchainID)

	resultsIterator, err := APIstub.GetStub().GetHistoryForKey(blockchainID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset NotarizerData1
		fmt.Println("Asset = ", asset)
		log.Println("asset = ", asset)
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &asset)
			if err != nil {
				return nil, err
			}
		} else {
			asset = NotarizerData1{
				BlockchainID: blockchainID,
			}
		}

		// timestamp, err := ptypes.Timestamp(response.Timestamp)
		// if err != nil {
		// 	return nil, err
		// }

		record := HistoryQueryResult{
			TxId:      response.TxId,
			// Timestamp: timestamp,
			Record:    &asset,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	fmt.Println("records = ", records)
	log.Println("Records = ", records)

	return records, nil

}

// func (t *SmartContract) GetHistoryByBatchID(ctx contractapi.TransactionContextInterface, batchID string) ([]NotarizerData1, error) {

// 	fmt.Println("BatchID = ", batchID)

// 	log.Println("batchID = ", batchID)

// 	index := "batchid_compositekey"

// 	batchIDResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{batchID})

// 	if err != nil {

// 		return nil, err

// 	}

// 	defer batchIDResultsIterator.Close()

// 	log.Println("batchIDResultsIterator = ", batchIDResultsIterator)

// 	var assets []NotarizerData1

// 	for batchIDResultsIterator.HasNext() {

// 		responseRange, err := batchIDResultsIterator.Next()

// 		if err != nil {

// 			return nil, err

// 		}

// 		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)

// 		if err != nil {

// 			return nil, err

// 		}

// 		if len(compositeKeyParts) > 1 {

// 			returnedAssetID := compositeKeyParts[1]

// 			log.Println("compositekey = ", compositeKeyParts)

// 			asset, err := ctx.GetStub().GetState(returnedAssetID)

// 			if err != nil {

// 				return nil, fmt.Errorf("failed to read from world state: %v", err)

// 			}

// 			if err != nil {

// 				return nil, err

// 			}

// 			log.Println("Asset = ", asset)
// 			var assetValue NotarizerData1
// 			if err := json.Unmarshal([]byte(asset), &assetValue); err != nil {
// 				panic(err)
// 			}

// 			log.Println("Asset value after unmarshaling = ", assetValue)
// 			assets = append(assets, assetValue)
// 		}

// 	}
// 	return assets, nil
// }

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
