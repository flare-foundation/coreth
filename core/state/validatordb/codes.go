package validatordb

const (
	codeEpoch      = 0 // store active validator epoch
	codeMapping    = 1 // store provider to node ID mapping
	codeCandidates = 2 // store validator candidates for next epoch
	codeValidators = 3 // store active validators for current epoch
)
