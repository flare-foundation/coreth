// (c) 2021, Flare Networks Limited. All rights reserved.
// Please see the file LICENSE for licensing terms.

package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	testingChainID               = new(big.Int).SetUint64(16)
	stateConnectorActivationTime = new(big.Int).SetUint64(1636070400)
	tr                           = &http.Transport{
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     60 * time.Second,
		DisableCompression:  true,
	}
	client = &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	apiRetries    = 3
	apiRetryDelay = 1 * time.Second
)

func GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {
	// Return true if chainID is 16 or if block.timestamp is greater than the state connector activation time on any chain
	return chainID.Cmp(testingChainID) == 0 || blockTime.Cmp(stateConnectorActivationTime) > 0
}

func GetStateConnectorGasDivisor(blockTime *big.Int) uint64 {
	switch {
	default:
		return 3
	}
}

func GetMaxAllowedChains(blockTime *big.Int) uint32 {
	switch {
	default:
		return 5
	}
}

func GetStateConnectorContractAddr(blockTime *big.Int) string {
	switch {
	default:
		return "0x1000000000000000000000000000000000000001"
	}
}

func GetProveDataAvailabilityPeriodFinalitySelector(blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0xc5, 0xd6, 0x4c, 0xd1}
	}
}

func GetProvePaymentFinalitySelector(blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x38, 0x84, 0x92, 0xdd}
	}
}

func GetDisprovePaymentFinalitySelector(blockTime *big.Int) []byte {
	switch {
	default:
		return []byte{0x7f, 0x58, 0x24, 0x32}
	}
}

// =======================================================
// Proof of Work Common
// =======================================================

type GetPoWRequestPayload struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}
type GetPoWBlockCountResp struct {
	Result uint64      `json:"result"`
	Error  interface{} `json:"error"`
}

func GetPoWBlockCount(chainURL string, username string, password string) (uint64, bool) {
	data := GetPoWRequestPayload{
		Method: "getblockcount",
		Params: []string{},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return 0, true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return 0, true
	}
	req.Header.Set("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, true
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, true
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, true
	}
	var jsonResp GetPoWBlockCountResp
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return 0, true
	}
	if jsonResp.Error != nil {
		return 0, true
	}
	return jsonResp.Result, false
}

type GetPoWBlockHeaderResult struct {
	Hash          string `json:"hash"`
	Confirmations uint64 `json:"confirmations"`
	Height        uint64 `json:"height"`
}
type GetPoWBlockHeaderResp struct {
	Result GetPoWBlockHeaderResult `json:"result"`
	Error  interface{}             `json:"error"`
}

func GetPoWBlockHeader(ledgerHash string, requiredConfirmations uint64, chainURL string, username string, password string) (uint64, bool) {
	data := GetPoWRequestPayload{
		Method: "getblockheader",
		Params: []string{
			ledgerHash,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return 0, true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return 0, true
	}
	req.Header.Set("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, true
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, true
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, true
	}
	var jsonResp GetPoWBlockHeaderResp
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return 0, true
	}
	if jsonResp.Error != nil {
		return 0, false
	} else if jsonResp.Result.Confirmations < requiredConfirmations {
		return 0, false
	}
	return jsonResp.Result.Height, false
}

func ProveDataAvailabilityPeriodFinalityPoW(checkRet []byte, chainURL string, username string, password string) (bool, bool) {
	blockCount, err := GetPoWBlockCount(chainURL, username, password)
	if err {
		return false, true
	}
	ledger := binary.BigEndian.Uint64(checkRet[56:64])
	requiredConfirmations := binary.BigEndian.Uint64(checkRet[88:96])
	if blockCount < ledger+requiredConfirmations {
		return false, true
	}
	ledgerResp, err := GetPoWBlockHeader(hex.EncodeToString(checkRet[96:128]), requiredConfirmations, chainURL, username, password)
	if err {
		return false, true
	} else if ledgerResp > 0 && ledgerResp == ledger {
		return true, false
	} else {
		return false, false
	}
}

type GetPoWTxRequestParams struct {
	TxID    string `json:"txid"`
	Verbose bool   `json:"verbose"`
}
type GetPoWTxRequestPayload struct {
	Method string                `json:"method"`
	Params GetPoWTxRequestParams `json:"params"`
}
type GetPoWTxResult struct {
	TxID          string `json:"txid"`
	BlockHash     string `json:"blockhash"`
	Confirmations uint64 `json:"confirmations"`
	Vout          []struct {
		Value        float64 `json:"value"`
		N            uint64  `json:"n"`
		ScriptPubKey struct {
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}
type GetPoWTxResp struct {
	Result GetPoWTxResult `json:"result"`
	Error  interface{}    `json:"error"`
}

func GetPoWTx(txHash string, voutN uint64, latestAvailableBlock uint64, currencyCode string, chainURL string, username string, password string) ([]byte, uint64, bool) {
	data := GetPoWTxRequestPayload{
		Method: "getrawtransaction",
		Params: GetPoWTxRequestParams{
			TxID:    txHash[1:],
			Verbose: true,
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}, 0, true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return []byte{}, 0, true
	}
	req.Header.Set("Content-Type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, true
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, 0, true
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, 0, true
	}
	var jsonResp GetPoWTxResp
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return []byte{}, 0, true
	}
	if jsonResp.Error != nil {
		return []byte{}, 0, true
	}
	if uint64(len(jsonResp.Result.Vout)) <= voutN {
		return []byte{}, 0, false
	}
	if jsonResp.Result.Vout[voutN].ScriptPubKey.Type != "pubkeyhash" || len(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses) != 1 {
		return []byte{}, 0, false
	}
	inBlock, getBlockErr := GetPoWBlockHeader(jsonResp.Result.BlockHash, jsonResp.Result.Confirmations, chainURL, username, password)
	if getBlockErr {
		return []byte{}, 0, true
	}
	if inBlock == 0 || inBlock >= latestAvailableBlock {
		return []byte{}, 0, false
	}
	txIdHash := crypto.Keccak256([]byte(txHash))
	destinationHash := crypto.Keccak256([]byte(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses[0]))
	amountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp.Result.Vout[voutN].Value*math.Pow(10, 8)))), 32))
	currencyHash := crypto.Keccak256([]byte(currencyCode))
	return crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inBlock, false
}

func ProvePaymentFinalityPoW(checkRet []byte, isDisprove bool, currencyCode string, chainURL string, username string, password string) (bool, bool) {
	if len(checkRet) < 257 {
		return false, false
	}
	voutN, err := strconv.ParseUint(string(checkRet[192:193]), 16, 64)
	if err != nil {
		return false, false
	}
	paymentHash, inBlock, getPoWTxErr := GetPoWTx(string(checkRet[192:257]), voutN, binary.BigEndian.Uint64(checkRet[88:96]), currencyCode, chainURL, username, password)
	if getPoWTxErr {
		return false, true
	}
	if !isDisprove {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock == binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		}
	} else {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock > binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		} else if len(paymentHash) == 0 {
			return true, false
		}
	}
	return false, false
}

func ProvePoW(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, currencyCode string, chainURL string) (bool, bool) {
	var username, password string
	chainURLhash := sha256.Sum256([]byte(chainURL))
	chainURLchecksum := hex.EncodeToString(chainURLhash[0:4])
	switch currencyCode {
	case "btc":
		username = os.Getenv("BTC_U_" + chainURLchecksum)
		password = os.Getenv("BTC_P_" + chainURLchecksum)
	case "ltc":
		username = os.Getenv("LTC_U_" + chainURLchecksum)
		password = os.Getenv("LTC_P_" + chainURLchecksum)
	case "dog":
		username = os.Getenv("DOGE_U_" + chainURLchecksum)
		password = os.Getenv("DOGE_P_" + chainURLchecksum)
	}
	if bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {
		return ProveDataAvailabilityPeriodFinalityPoW(checkRet, chainURL, username, password)
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {
		return ProvePaymentFinalityPoW(checkRet, false, currencyCode, chainURL, username, password)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {
		return ProvePaymentFinalityPoW(checkRet, true, currencyCode, chainURL, username, password)
	}
	return false, false
}

// =======================================================
// XRP
// =======================================================

type GetXRPBlockRequestParams struct {
	LedgerIndex  uint64 `json:"ledger_index"`
	Full         bool   `json:"full"`
	Accounts     bool   `json:"accounts"`
	Transactions bool   `json:"transactions"`
	Expand       bool   `json:"expand"`
	OwnerFunds   bool   `json:"owner_funds"`
}
type GetXRPBlockRequestPayload struct {
	Method string                     `json:"method"`
	Params []GetXRPBlockRequestParams `json:"params"`
}
type CheckXRPErrorResponse struct {
	Error string `json:"error"`
}
type GetXRPBlockResponse struct {
	LedgerHash  string `json:"ledger_hash"`
	LedgerIndex int    `json:"ledger_index"`
	Validated   bool   `json:"validated"`
}

func GetXRPBlock(ledger uint64, chainURL string) (string, bool) {
	data := GetXRPBlockRequestPayload{
		Method: "ledger",
		Params: []GetXRPBlockRequestParams{
			GetXRPBlockRequestParams{
				LedgerIndex:  ledger,
				Full:         false,
				Accounts:     false,
				Transactions: false,
				Expand:       false,
				OwnerFunds:   false,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return "", true
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", true
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", true
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", true
	}
	var checkErrorResp map[string]CheckXRPErrorResponse
	err = json.Unmarshal(respBody, &checkErrorResp)
	if err != nil {
		return "", true
	}
	if checkErrorResp["result"].Error != "" {
		return "", true
	}
	var jsonResp map[string]GetXRPBlockResponse
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return "", true
	}
	if !jsonResp["result"].Validated {
		return "", true
	}
	return jsonResp["result"].LedgerHash, false
}

func ProveDataAvailabilityPeriodFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {
	ledger := binary.BigEndian.Uint64(checkRet[56:64])
	ledgerHashString, err := GetXRPBlock(ledger, chainURL)
	if err {
		return false, true
	}
	if ledgerHashString != "" && bytes.Equal(crypto.Keccak256([]byte(ledgerHashString)), checkRet[96:128]) {
		return true, false
	}
	return false, false
}

type GetXRPTxRequestParams struct {
	Transaction string `json:"transaction"`
	Binary      bool   `json:"binary"`
}
type GetXRPTxRequestPayload struct {
	Method string                  `json:"method"`
	Params []GetXRPTxRequestParams `json:"params"`
}
type GetXRPTxResponse struct {
	Destination     string `json:"Destination"`
	DestinationTag  int    `json:"DestinationTag"`
	TransactionType string `json:"TransactionType"`
	Hash            string `json:"hash"`
	InLedger        int    `json:"inLedger"`
	Validated       bool   `json:"validated"`
	Meta            struct {
		TransactionResult string      `json:"TransactionResult"`
		Amount            interface{} `json:"delivered_amount"`
	} `json:"meta"`
}

type GetXRPTxIssuedCurrency struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
	Value    string `json:"value"`
}

func GetXRPTx(txHash string, latestAvailableLedger uint64, chainURL string) ([]byte, uint64, bool) {
	data := GetXRPTxRequestPayload{
		Method: "tx",
		Params: []GetXRPTxRequestParams{
			GetXRPTxRequestParams{
				Transaction: txHash,
				Binary:      false,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}, 0, true
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", chainURL, body)
	if err != nil {
		return []byte{}, 0, true
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, 0, true
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return []byte{}, 0, true
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, 0, true
	}
	var checkErrorResp map[string]CheckXRPErrorResponse
	err = json.Unmarshal(respBody, &checkErrorResp)
	if err != nil {
		return []byte{}, 0, true
	}
	respErrString := checkErrorResp["result"].Error
	if respErrString != "" {
		if respErrString == "amendmentBlocked" ||
			respErrString == "failedToForward" ||
			respErrString == "invalid_API_version" ||
			respErrString == "noClosed" ||
			respErrString == "noCurrent" ||
			respErrString == "noNetwork" ||
			respErrString == "tooBusy" {
			return []byte{}, 0, true
		} else {
			return []byte{}, 0, false
		}
	}
	var jsonResp map[string]GetXRPTxResponse
	err = json.Unmarshal(respBody, &jsonResp)
	if err != nil {
		return []byte{}, 0, false
	}
	if jsonResp["result"].TransactionType != "Payment" || !jsonResp["result"].Validated || jsonResp["result"].Meta.TransactionResult != "tesSUCCESS" {
		return []byte{}, 0, false
	}
	inLedger := uint64(jsonResp["result"].InLedger)
	if inLedger == 0 || inLedger >= latestAvailableLedger || !jsonResp["result"].Validated {
		return []byte{}, 0, false
	}
	var currency string
	var amount uint64
	if stringAmount, ok := jsonResp["result"].Meta.Amount.(string); ok {
		amount, err = strconv.ParseUint(stringAmount, 10, 64)
		if err != nil {
			return []byte{}, 0, false
		}
		currency = "xrp"
	} else {
		amountStruct, err := json.Marshal(jsonResp["result"].Meta.Amount)
		if err != nil {
			return []byte{}, 0, false
		}
		var issuedCurrencyResp GetXRPTxIssuedCurrency
		err = json.Unmarshal(amountStruct, &issuedCurrencyResp)
		if err != nil {
			return []byte{}, 0, false
		}
		floatAmount, err := strconv.ParseFloat(issuedCurrencyResp.Value, 64)
		if err != nil {
			return []byte{}, 0, false
		}
		amount = uint64(floatAmount * math.Pow(10, 15))
		currency = issuedCurrencyResp.Currency + issuedCurrencyResp.Issuer
	}
	txIdHash := crypto.Keccak256([]byte(jsonResp["result"].Hash))
	destinationHash := crypto.Keccak256([]byte(jsonResp["result"].Destination))
	destinationTagHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp["result"].DestinationTag))), 32))
	destinationHash = crypto.Keccak256(destinationHash, destinationTagHash)
	amountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(amount))), 32))
	currencyHash := crypto.Keccak256([]byte(currency))
	return crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inLedger, false
}

func ProvePaymentFinalityXRP(checkRet []byte, isDisprove bool, chainURL string) (bool, bool) {
	paymentHash, inLedger, err := GetXRPTx(string(checkRet[192:]), binary.BigEndian.Uint64(checkRet[88:96]), chainURL)
	if err {
		return false, true
	}
	if !isDisprove {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger == binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		}
	} else {
		if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger > binary.BigEndian.Uint64(checkRet[56:64]) {
			return true, false
		} else if len(paymentHash) == 0 {
			return true, false
		}
	}
	return false, false
}

func ProveXRP(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {
	if bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {
		return ProveDataAvailabilityPeriodFinalityXRP(checkRet, chainURL)
	} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {
		return ProvePaymentFinalityXRP(checkRet, false, chainURL)
	} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {
		return ProvePaymentFinalityXRP(checkRet, true, chainURL)
	}
	return false, false
}

// =======================================================
// ALGO
// =======================================================

func ProveALGO(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {
	return false, false
}

// =======================================================
// Common
// =======================================================

func ProveChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainId uint32, chainURL string) (bool, bool) {
	switch chainId {
	case 0:
		return ProvePoW(sender, blockTime, functionSelector, checkRet, "btc", chainURL)
	case 1:
		return ProvePoW(sender, blockTime, functionSelector, checkRet, "ltc", chainURL)
	case 2:
		return ProvePoW(sender, blockTime, functionSelector, checkRet, "dog", chainURL)
	case 3:
		return ProveXRP(sender, blockTime, functionSelector, checkRet, chainURL)
	case 4:
		return ProveALGO(sender, blockTime, functionSelector, checkRet, chainURL)
	default:
		return false, true
	}
}

func ReadChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {
	chainId := binary.BigEndian.Uint32(checkRet[28:32])
	var chainURLs string
	switch chainId {
	case 0:
		chainURLs = os.Getenv("BTC_APIs")
	case 1:
		chainURLs = os.Getenv("LTC_APIs")
	case 2:
		chainURLs = os.Getenv("DOGE_APIs")
	case 3:
		chainURLs = os.Getenv("XRP_APIs")
	case 4:
		chainURLs = os.Getenv("ALGO_APIs")
	}
	if chainURLs == "" {
		return false
	}
	for i := 0; i < apiRetries; i++ {
		for _, chainURL := range strings.Split(chainURLs, ",") {
			if chainURL == "" {
				continue
			}
			verified, err := ProveChain(sender, blockTime, functionSelector, checkRet, chainId, chainURL)
			if !verified && err {
				continue
			}
			return verified
		}
		time.Sleep(apiRetryDelay)
	}
	return false
}

func GetVerificationPaths(functionSelector []byte, checkRet []byte) (string, string) {
	prefix := "cache/"
	acceptedPrefix := "ACCEPTED"
	rejectedPrefix := "REJECTED"
	functionHash := hex.EncodeToString(functionSelector[:])
	verificationHash := hex.EncodeToString(crypto.Keccak256(checkRet[0:64], checkRet[96:128]))
	suffix := "_" + functionHash + "_" + verificationHash
	return prefix + acceptedPrefix + suffix, prefix + rejectedPrefix + suffix
}

// Verify proof against underlying chain
func StateConnectorCall(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {
	if binary.BigEndian.Uint64(checkRet[88:96]) > 0 {
		go func() {
			acceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)
			_, errACCEPTED := os.Stat(acceptedPath)
			_, errREJECTED := os.Stat(rejectedPath)
			if errACCEPTED != nil && errREJECTED != nil {
				if ReadChain(sender, blockTime, functionSelector, checkRet) {
					verificationHashStore, err := os.Create(acceptedPath)
					verificationHashStore.Close()
					if err != nil {
						// Permissions problem
						panic(err)
					}
				} else {
					verificationHashStore, err := os.Create(rejectedPath)
					verificationHashStore.Close()
					if err != nil {
						// Permissions problem
						panic(err)
					}
				}
			}
		}()
		return true
	} else {
		acceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)
		_, errACCEPTED := os.Stat(acceptedPath)
		_, errREJECTED := os.Stat(rejectedPath)
		if errACCEPTED != nil && errREJECTED != nil {
			for i := 0; i < 2*apiRetries; i++ {
				_, errACCEPTED = os.Stat(acceptedPath)
				_, errREJECTED = os.Stat(rejectedPath)
				if errACCEPTED == nil || errREJECTED == nil {
					break
				}
				time.Sleep(apiRetryDelay)
			}
		}
		go func() {
			removeFulfilledAPIRequests := os.Getenv("REMOVE_FULFILLED_API_REQUESTS")
			if removeFulfilledAPIRequests == "1" {
				errDeleteACCEPTED := os.Remove(acceptedPath)
				errDeleteREJECTED := os.Remove(rejectedPath)
				if errDeleteACCEPTED != nil && errDeleteREJECTED != nil {
					// Permissions problem
					panic(errDeleteACCEPTED)
					panic(errDeleteREJECTED)
				}
			}
		}()
		return errACCEPTED == nil
	}
}
