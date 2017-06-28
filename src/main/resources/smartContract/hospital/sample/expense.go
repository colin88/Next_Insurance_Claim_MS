package main

import (
"fmt"
_ "strconv"
_ "encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	_"encoding/binary"
	_"bytes"
)

type MedicineDetail struct {
	Id string
	Name string
	Price int
	Number int
}

type ExpenseDetail struct {
	Uid string
	//yyyyMMddHHmmss
	ExpenseTime string
	Claimed bool
	Medicines []MedicineDetail
}


type HospitalChainCode struct {
}

func (t *HospitalChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// {"uid":"3702821982","expenseTime":"20001010010203","claimed":false,"medicines":[{"id":"1000","name":"med1000","price":10,
// "number":10},{"id":"2000","name":"med2000","price":20,"number":10},{"id":"3000","name":"med3000","price":30,"number":10}]}
func (t *HospitalChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}
	return shim.Error(`invalid invoke function name: "invoke" "query"`)
}

func (t *HospitalChainCode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	jsonVal := args[0]
	var jsonObj ExpenseDetail
	err := json.Unmarshal([]byte(jsonVal), &jsonObj)
	if err != nil {
		return shim.Error("Fail to unmarshal json data!")
	}

	 usrMapdata := map[string]ExpenseDetail{}

	// usrMapdataBytes is []byte
	usrMapdataBytes, err := stub.GetState(jsonObj.Uid)
	// map is not found
	if err == nil {
		jsonErr := json.Unmarshal(usrMapdataBytes, usrMapdata)
		if jsonErr != nil {
			return shim.Error("Failed to Unmarshal!")
		}
}
	usrMapdata[jsonObj.ExpenseTime] = jsonObj

	userMapJson, err := json.Marshal(usrMapdata)
	if  err != nil {
		return shim.Error("Failed to Marshal!")
	}

	//buf := new(bytes.Buffer)
	//binary.Write(buf, binary.BigEndian, usrMapdata)

	stub.PutState(jsonObj.Uid, userMapJson)

	return shim.Success([]byte("success!"))
}

func (t *HospitalChainCode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 -> user id")
	}

	uid := args[0]
	usrMapdataBytes, err := stub.GetState(uid)
	if err != nil {
		return shim.Success([]byte("Data is null!"))
	}

	usrMapdata := map[string]ExpenseDetail{}
	jsonData,err := json.Marshal(usrMapdata)
	if err != nil {
		return shim.Error("Fail to Marshal!")
	}

	return shim.Success(jsonData)
}

func main() {
	err := shim.Start(new(HospitalChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
