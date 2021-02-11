package main

import (
	"encoding/json"
	"fmt"

	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//SmartContract ... The SmartContract
type SmartContract struct {
}

type Product struct {
	ProductName string `json:"product_name"`
	ProductId   string `json:"product_id"`
	ProductDescription  string `json:"product_description"`
	ProductPrice string `json:"product_price"`
	SellerID    string `json:"seller_id"`
	Status string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Category    string `json:"category"`
}

type ProductByIdResponse struct {
	ID      string  `json:"id"`
	Request Products `json:"product"`
}

type Response struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    DoctorByIdResponse `json:"data"`
}

//var logger = shim.NewLogger("example_cc0")

//Init Function
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	_, args := stub.GetFunctionAndParameters()
	var products = Products{
		ProductName: args[0],
		ProductId:   args[1],
		ProductDescription:      args[2],
		ProductPrice: args[3],
		SellerID:    args[4]
		Status:    args[5]
		CreatedAt:    args[6]
		UpdatedAt:    args[7]
		Category:    args[8]}

	productAsBytes, _ := json.Marshal(products)

	var uniqueID = args[1]

	err := stub.PutState(uniqueID, productAsBytes)

	if err != nil {
		fmt.Println("Error in Init")
	}

	return shim.Success([]byte("Chaincode Successfully initialized"))
}

//Createproduct ... this function is used to create product
func CreateProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var products = Products{
		ProductName: args[0],
		ProductId:   args[1],
		ProductDescription:      args[2],
		ProductPrice: args[3],
		SellerID:    args[4]
		Status:    args[5]
		CreatedAt:    args[6]
		UpdatedAt:    args[7]
		Category:    args[8]}

	productAsBytes, _ := json.Marshal(products)

	var uniqueID = args[1]

	err := stub.PutState(uniqueID, productAsBytes)

	if err != nil {
		fmt.Println("Erro in create product")
	}

	return shim.Success(nil)
}

//GetProductByID ... This function will return a particular product
func GetProductByID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("Before Len")

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.Expected 1 argument")
	}

	fmt.Println("After Len")

	query := `{
				"selector": {
				   "product_id": {
					  "$eq": "` + args[0] + `"
				   }
				}
			 }`

	fmt.Println("Queeryy =>>>> \n" + query)

	//resultsIterator, err := stub.GetQueryResult("{\"selector\":{\"prod_type\":\"products\",\"_id\":{\"$eq\": \"1\"}}}")

	resultsIterator, err := stub.GetQueryResult(query)

	if err != nil {
		fmt.Println("Error fetching reuslts")
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	fmt.Println("After results")
	// counter := 0
	//var resultKV
	for resultsIterator.HasNext() {
		fmt.Println("Inside hasnext")
		// Get the next element
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			fmt.Println("Err=" + err.Error())
			return shim.Success([]byte("Error in parse: " + err.Error()))
		}

		// Increment the counter
		// counter++
		key := queryResponse.GetKey()
		value := string(queryResponse.GetValue())

		// Print the receieved result on the console
		fmt.Printf("Result#   %s   %s \n\n", key, value)
		b, je := json.Marshal(value)
		if je != nil {
			return shim.Error(je.Error())
		}

		return shim.Success(b)
	}

	// Close the iterator
	resultsIterator.Close()
	return shim.Success(nil)

	//	return shim.Error("Could not find any doctors.")

}

//Invoke function
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fun, args := stub.GetFunctionAndParameters()
	if fun == "CreateProduct" {
		fmt.Println("Error occured ==> ")
		//logger.Info("########### create prods ###########")
		return CreateProduct(stub, args)
	} else if fun == "GetProductByID" {
		fmt.Println("Calling get  ==> ")
		return GetProductByID(stub, args)
	}

	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", fun))

}

func print {
fmt.println("Hello")
}

}

func main() {
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Print(err)
	}
}
