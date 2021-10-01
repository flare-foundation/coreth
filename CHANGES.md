[1mdiff --git a/CHANGES.md b/CHANGES.md[m
[1mnew file mode 100644[m
[1mindex 0000000..e69de29[m
[1mdiff --git a/accounts/accounts.go b/accounts/accounts.go[m
[1mindex dd7df0b..00420ca 100644[m
[1m--- a/accounts/accounts.go[m
[1m+++ b/accounts/accounts.go[m
[36m@@ -31,10 +31,10 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
 	"golang.org/x/crypto/sha3"[m
 )[m
 [m
[1mdiff --git a/accounts/external/backend.go b/accounts/external/backend.go[m
[1mindex 558700f..8302c0c 100644[m
[1m--- a/accounts/external/backend.go[m
[1m+++ b/accounts/external/backend.go[m
[36m@@ -31,15 +31,15 @@[m [mimport ([m
 	"math/big"[m
 	"sync"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
[31m-	"github.com/ava-labs/coreth/signer/core/apitypes"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/signer/core/apitypes"[m
 )[m
 [m
 type ExternalBackend struct {[m
[1mdiff --git a/accounts/keystore/account_cache.go b/accounts/keystore/account_cache.go[m
[1mindex 4c35aa7..fbe0a89 100644[m
[1m--- a/accounts/keystore/account_cache.go[m
[1m+++ b/accounts/keystore/account_cache.go[m
[36m@@ -37,10 +37,10 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	mapset "github.com/deckarep/golang-set"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 )[m
 [m
 // Minimum amount of time between cache reloads. This limit applies if the platform does[m
[1mdiff --git a/accounts/keystore/account_cache_test.go b/accounts/keystore/account_cache_test.go[m
[1mindex 3d341d5..81f80d6 100644[m
[1m--- a/accounts/keystore/account_cache_test.go[m
[1m+++ b/accounts/keystore/account_cache_test.go[m
[36m@@ -37,10 +37,10 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/cespare/cp"[m
 	"github.com/davecgh/go-spew/spew"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 )[m
 [m
 var ([m
[1mdiff --git a/accounts/keystore/key.go b/accounts/keystore/key.go[m
[1mindex 71402d3..ad6cc6c 100644[m
[1m--- a/accounts/keystore/key.go[m
[1m+++ b/accounts/keystore/key.go[m
[36m@@ -39,10 +39,10 @@[m [mimport ([m
 	"strings"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/google/uuid"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 )[m
 [m
 const ([m
[1mdiff --git a/accounts/keystore/keystore.go b/accounts/keystore/keystore.go[m
[1mindex ff82ef8..e1c177d 100644[m
[1m--- a/accounts/keystore/keystore.go[m
[1m+++ b/accounts/keystore/keystore.go[m
[36m@@ -42,11 +42,11 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 var ([m
[1mdiff --git a/accounts/keystore/keystore_test.go b/accounts/keystore/keystore_test.go[m
[1mindex 651ab70..70fbeb7 100644[m
[1m--- a/accounts/keystore/keystore_test.go[m
[1m+++ b/accounts/keystore/keystore_test.go[m
[36m@@ -38,10 +38,10 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 )[m
 [m
 var testSigData = make([]byte, 32)[m
[1mdiff --git a/accounts/keystore/passphrase.go b/accounts/keystore/passphrase.go[m
[1mindex 5da41f6..73f26d2 100644[m
[1m--- a/accounts/keystore/passphrase.go[m
[1m+++ b/accounts/keystore/passphrase.go[m
[36m@@ -48,11 +48,11 @@[m [mimport ([m
 	"os"[m
 	"path/filepath"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/google/uuid"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 	"golang.org/x/crypto/pbkdf2"[m
 	"golang.org/x/crypto/scrypt"[m
 )[m
[1mdiff --git a/accounts/keystore/presale.go b/accounts/keystore/presale.go[m
[1mindex 1dfbd9c..64f020d 100644[m
[1m--- a/accounts/keystore/presale.go[m
[1m+++ b/accounts/keystore/presale.go[m
[36m@@ -35,9 +35,9 @@[m [mimport ([m
 	"errors"[m
 	"fmt"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/google/uuid"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 	"golang.org/x/crypto/pbkdf2"[m
 )[m
 [m
[1mdiff --git a/accounts/keystore/wallet.go b/accounts/keystore/wallet.go[m
[1mindex 78eac1d..e603ac5 100644[m
[1m--- a/accounts/keystore/wallet.go[m
[1m+++ b/accounts/keystore/wallet.go[m
[36m@@ -29,11 +29,11 @@[m [mpackage keystore[m
 import ([m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // keystoreWallet implements the accounts.Wallet interface for the original[m
[1mdiff --git a/accounts/scwallet/hub.go b/accounts/scwallet/hub.go[m
[1mindex 7a630fa..c05ba6f 100644[m
[1m--- a/accounts/scwallet/hub.go[m
[1m+++ b/accounts/scwallet/hub.go[m
[36m@@ -51,11 +51,11 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	pcsc "github.com/gballet/go-libpcsclite"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
 )[m
 [m
 // Scheme is the URI prefix for smartcard wallets.[m
[1mdiff --git a/accounts/scwallet/wallet.go b/accounts/scwallet/wallet.go[m
[1mindex fcecc10..b912e21 100644[m
[1m--- a/accounts/scwallet/wallet.go[m
[1m+++ b/accounts/scwallet/wallet.go[m
[36m@@ -43,14 +43,14 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	pcsc "github.com/gballet/go-libpcsclite"[m
 	"github.com/status-im/keycard-go/derivationpath"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
 )[m
 [m
 // ErrPairingPasswordNeeded is returned if opening the smart card requires pairing with a pairing[m
[1mdiff --git a/chain/chain_test.go b/chain/chain_test.go[m
[1mindex 2d8c7e9..9057d9d 100644[m
[1m--- a/chain/chain_test.go[m
[1m+++ b/chain/chain_test.go[m
[36m@@ -9,20 +9,20 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts/keystore"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/eth"[m
[31m-	"github.com/ava-labs/coreth/eth/ethconfig"[m
[31m-	"github.com/ava-labs/coreth/node"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/keystore"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/ethconfig"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type testChain struct {[m
[1mdiff --git a/chain/coreth.go b/chain/coreth.go[m
[1mindex 95018be..bc107de 100644[m
[1m--- a/chain/coreth.go[m
[1m+++ b/chain/coreth.go[m
[36m@@ -7,15 +7,15 @@[m [mimport ([m
 	"fmt"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/eth"[m
[31m-	"github.com/ava-labs/coreth/node"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 var ([m
[1mdiff --git a/chain/counter_test.go b/chain/counter_test.go[m
[1mindex 7209ac5..4fd88fe 100644[m
[1m--- a/chain/counter_test.go[m
[1m+++ b/chain/counter_test.go[m
[36m@@ -15,8 +15,8 @@[m [mimport ([m
 [m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 [m
 	"github.com/ethereum/go-ethereum/log"[m
 )[m
[1mdiff --git a/chain/multicoin_test.go b/chain/multicoin_test.go[m
[1mindex a2ce4cb..f40a678 100644[m
[1m--- a/chain/multicoin_test.go[m
[1m+++ b/chain/multicoin_test.go[m
[36m@@ -28,19 +28,19 @@[m [mimport ([m
 	"strings"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts/keystore"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/eth"[m
[31m-	"github.com/ava-labs/coreth/eth/ethconfig"[m
[31m-	"github.com/ava-labs/coreth/node"[m
 	"github.com/ethereum/go-ethereum/accounts/abi"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/keystore"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/ethconfig"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
 )[m
 [m
 // TestMulticoin tests multicoin low-level state management and regular[m
[1mdiff --git a/chain/payment_test.go b/chain/payment_test.go[m
[1mindex 5d9da4d..f17e260 100644[m
[1m--- a/chain/payment_test.go[m
[1m+++ b/chain/payment_test.go[m
[36m@@ -7,8 +7,8 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // TestPayment tests basic payment (balance, not multi-coin)[m
[1mdiff --git a/chain/subscribe_accepted_heads_test.go b/chain/subscribe_accepted_heads_test.go[m
[1mindex 0a94bfd..ab7315e 100644[m
[1m--- a/chain/subscribe_accepted_heads_test.go[m
[1m+++ b/chain/subscribe_accepted_heads_test.go[m
[36m@@ -4,10 +4,10 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 func TestAcceptedHeadSubscriptions(t *testing.T) {[m
[1mdiff --git a/chain/subscribe_block_logs_test.go b/chain/subscribe_block_logs_test.go[m
[1mindex 26661da..2b2d2ab 100644[m
[1m--- a/chain/subscribe_block_logs_test.go[m
[1m+++ b/chain/subscribe_block_logs_test.go[m
[36m@@ -6,10 +6,10 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/eth/filters"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/filters"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 func TestBlockLogsAllowUnfinalized(t *testing.T) {[m
[1mdiff --git a/chain/subscribe_transactions_test.go b/chain/subscribe_transactions_test.go[m
[1mindex aac6db4..9388dae 100644[m
[1m--- a/chain/subscribe_transactions_test.go[m
[1m+++ b/chain/subscribe_transactions_test.go[m
[36m@@ -4,10 +4,10 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/eth/filters"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/filters"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 func TestSubscribeTransactions(t *testing.T) {[m
[1mdiff --git a/chain/test_chain.go b/chain/test_chain.go[m
[1mindex e4420de..3d71d82 100644[m
[1m--- a/chain/test_chain.go[m
[1m+++ b/chain/test_chain.go[m
[36m@@ -8,17 +8,17 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts/keystore"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/eth"[m
[31m-	"github.com/ava-labs/coreth/eth/ethconfig"[m
[31m-	"github.com/ava-labs/coreth/node"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/keystore"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/ethconfig"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/consensus/consensus.go b/consensus/consensus.go[m
[1mindex 6bc13a8..af462ca 100644[m
[1m--- a/consensus/consensus.go[m
[1m+++ b/consensus/consensus.go[m
[36m@@ -30,11 +30,11 @@[m [mpackage consensus[m
 import ([m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // ChainHeaderReader defines a small collection of methods needed to access the local[m
[1mdiff --git a/consensus/dummy/consensus.go b/consensus/dummy/consensus.go[m
[1mindex 89af334..e93724d 100644[m
[1m--- a/consensus/dummy/consensus.go[m
[1m+++ b/consensus/dummy/consensus.go[m
[36m@@ -10,13 +10,13 @@[m [mimport ([m
 	"math/big"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 type OnFinalizeCallbackType = func(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, receipts []*types.Receipt, uncles []*types.Header) error[m
[1mdiff --git a/consensus/dummy/dynamic_fees.go b/consensus/dummy/dynamic_fees.go[m
[1mindex 8050364..9221cec 100644[m
[1m--- a/consensus/dummy/dynamic_fees.go[m
[1m+++ b/consensus/dummy/dynamic_fees.go[m
[36m@@ -9,10 +9,10 @@[m [mimport ([m
 	"math/big"[m
 [m
 	"github.com/ava-labs/avalanchego/utils/wrappers"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/consensus/dummy/dynamic_fees_test.go b/consensus/dummy/dynamic_fees_test.go[m
[1mindex fd45082..bc6813d 100644[m
[1m--- a/consensus/dummy/dynamic_fees_test.go[m
[1m+++ b/consensus/dummy/dynamic_fees_test.go[m
[36m@@ -8,10 +8,10 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func testRollup(t *testing.T, longs []uint64, roll int) {[m
[1mdiff --git a/consensus/misc/dao.go b/consensus/misc/dao.go[m
[1mindex a0ab402..486f97a 100644[m
[1m--- a/consensus/misc/dao.go[m
[1m+++ b/consensus/misc/dao.go[m
[36m@@ -31,9 +31,9 @@[m [mimport ([m
 	"errors"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/consensus/misc/forks.go b/consensus/misc/forks.go[m
[1mindex 6199e11..63ce4f0 100644[m
[1m--- a/consensus/misc/forks.go[m
[1m+++ b/consensus/misc/forks.go[m
[36m@@ -29,9 +29,9 @@[m [mpackage misc[m
 import ([m
 	"fmt"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // VerifyForkHashes verifies that blocks conforming to network hard-forks do have[m
[1mdiff --git a/core/block_validator.go b/core/block_validator.go[m
[1mindex f7092ee..6993bf1 100644[m
[1m--- a/core/block_validator.go[m
[1m+++ b/core/block_validator.go[m
[36m@@ -29,11 +29,11 @@[m [mpackage core[m
 import ([m
 	"fmt"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // BlockValidator is responsible for validating block headers, uncles and[m
[1mdiff --git a/core/blockchain.go b/core/blockchain.go[m
[1mindex 7e03462..5073b6a 100644[m
[1m--- a/core/blockchain.go[m
[1m+++ b/core/blockchain.go[m
[36m@@ -36,19 +36,19 @@[m [mimport ([m
 	"sync/atomic"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/state/snapshot"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/trie"[m
 	lru "github.com/hashicorp/golang-lru"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state/snapshot"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/core/blockchain_test.go b/core/blockchain_test.go[m
[1mindex 07cf8ec..74981e8 100644[m
[1m--- a/core/blockchain_test.go[m
[1m+++ b/core/blockchain_test.go[m
[36m@@ -7,14 +7,14 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestArchiveBlockChain(t *testing.T) {[m
[1mdiff --git a/core/bloom_indexer.go b/core/bloom_indexer.go[m
[1mindex 07e50de..c2fa061 100644[m
[1m--- a/core/bloom_indexer.go[m
[1m+++ b/core/bloom_indexer.go[m
[36m@@ -20,12 +20,12 @@[m [mimport ([m
 	"context"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/bloombits"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/bitutil"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/bloombits"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/bloombits/generator.go b/core/bloombits/generator.go[m
[1mindex c0422ca..23a5b54 100644[m
[1m--- a/core/bloombits/generator.go[m
[1m+++ b/core/bloombits/generator.go[m
[36m@@ -29,7 +29,7 @@[m [mpackage bloombits[m
 import ([m
 	"errors"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 var ([m
[1mdiff --git a/core/bloombits/generator_test.go b/core/bloombits/generator_test.go[m
[1mindex 27dba47..b233daf 100644[m
[1m--- a/core/bloombits/generator_test.go[m
[1m+++ b/core/bloombits/generator_test.go[m
[36m@@ -31,7 +31,7 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // Tests that batched bloom bits are correctly rotated from the input bloom[m
[1mdiff --git a/core/chain_indexer.go b/core/chain_indexer.go[m
[1mindex ee6557f..eaf7881 100644[m
[1m--- a/core/chain_indexer.go[m
[1m+++ b/core/chain_indexer.go[m
[36m@@ -34,12 +34,12 @@[m [mimport ([m
 	"sync/atomic"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // ChainIndexerBackend defines the methods needed to process chain segments in[m
[1mdiff --git a/core/chain_indexer_test.go b/core/chain_indexer_test.go[m
[1mindex 3edf175..d8b1c43 100644[m
[1m--- a/core/chain_indexer_test.go[m
[1m+++ b/core/chain_indexer_test.go[m
[36m@@ -35,9 +35,9 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // Runs multiple tests with randomized parameters.[m
[1mdiff --git a/core/chain_makers.go b/core/chain_makers.go[m
[1mindex 4aca357..58fd71b 100644[m
[1m--- a/core/chain_makers.go[m
[1m+++ b/core/chain_makers.go[m
[36m@@ -30,15 +30,15 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/consensus/misc"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/misc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // BlockGen creates blocks for testing.[m
[1mdiff --git a/core/chain_makers_test.go b/core/chain_makers_test.go[m
[1mindex b67cac2..34881d4 100644[m
[1m--- a/core/chain_makers_test.go[m
[1m+++ b/core/chain_makers_test.go[m
[36m@@ -30,13 +30,13 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func ExampleGenerateChain() {[m
[1mdiff --git a/core/error.go b/core/error.go[m
[1mindex 0265448..4039ed7 100644[m
[1m--- a/core/error.go[m
[1m+++ b/core/error.go[m
[36m@@ -29,7 +29,7 @@[m [mpackage core[m
 import ([m
 	"errors"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 var ([m
[1mdiff --git a/core/events.go b/core/events.go[m
[1mindex 4898dbc..d57d378 100644[m
[1m--- a/core/events.go[m
[1m+++ b/core/events.go[m
[36m@@ -27,8 +27,8 @@[m
 package core[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // NewTxsEvent is posted when a batch of transactions enter the transaction pool.[m
[1mdiff --git a/core/evm.go b/core/evm.go[m
[1mindex d45f241..2dad5a4 100644[m
[1m--- a/core/evm.go[m
[1m+++ b/core/evm.go[m
[36m@@ -29,10 +29,10 @@[m [mpackage core[m
 import ([m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 	//"github.com/ethereum/go-ethereum/log"[m
 )[m
 [m
[1mdiff --git a/core/gen_genesis.go b/core/gen_genesis.go[m
[1mindex a4ec8f5..b247996 100644[m
[1m--- a/core/gen_genesis.go[m
[1m+++ b/core/gen_genesis.go[m
[36m@@ -7,10 +7,10 @@[m [mimport ([m
 	"errors"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var _ = (*genesisSpecMarshaling)(nil)[m
[1mdiff --git a/core/genesis.go b/core/genesis.go[m
[1mindex 185887e..153f51f 100644[m
[1m--- a/core/genesis.go[m
[1m+++ b/core/genesis.go[m
[36m@@ -34,16 +34,16 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 //go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go[m
[1mdiff --git a/core/headerchain.go b/core/headerchain.go[m
[1mindex eafc747..b8d75c1 100644[m
[1m--- a/core/headerchain.go[m
[1m+++ b/core/headerchain.go[m
[36m@@ -33,13 +33,13 @@[m [mimport ([m
 	mrand "math/rand"[m
 	"sync/atomic"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	lru "github.com/hashicorp/golang-lru"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/keeper.go b/core/keeper.go[m
[1mnew file mode 100644[m
[1mindex 0000000..abe3acc[m
[1m--- /dev/null[m
[1m+++ b/core/keeper.go[m
[36m@@ -0,0 +1,149 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m
[32m+[m[32mpackage core[m
[32m+[m
[32m+[m[32mimport ([m
[32m+[m	[32m"fmt"[m
[32m+[m	[32m"math/big"[m
[32m+[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/log"[m
[32m+[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m[32m)[m
[32m+[m
[32m+[m[32m// Define errors[m
[32m+[m[32mtype ErrInvalidKeeperData struct{}[m
[32m+[m
[32m+[m[32mfunc (e *ErrInvalidKeeperData) Error() string { return "invalid return data from keeper trigger" }[m
[32m+[m
[32m+[m[32mtype ErrKeeperDataEmpty struct{}[m
[32m+[m
[32m+[m[32mfunc (e *ErrKeeperDataEmpty) Error() string { return "return data from keeper trigger empty" }[m
[32m+[m
[32m+[m[32mtype ErrMaxMintExceeded struct {[m
[32m+[m	[32mmintMax     *big.Int[m
[32m+[m	[32mmintRequest *big.Int[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *ErrMaxMintExceeded) Error() string {[m
[32m+[m	[32mreturn fmt.Sprintf("mint request of %s exceeded max of %s", e.mintRequest.Text(10), e.mintMax.Text(10))[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype ErrMintNegative struct{}[m
[32m+[m
[32m+[m[32mfunc (e *ErrMintNegative) Error() string { return "mint request cannot be negative" }[m
[32m+[m
[32m+[m[32m// Define interface for dependencies[m
[32m+[m[32mtype EVMCaller interface {[m
[32m+[m	[32mCall(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)[m
[32m+[m	[32mGetBlockNumber() *big.Int[m
[32m+[m	[32mGetGasLimit() uint64[m
[32m+[m	[32mAddBalance(addr common.Address, amount *big.Int)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define maximums that can change by block height[m
[32m+[m[32mfunc GetKeeperGasMultiplier(blockNumber *big.Int) uint64 {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn 100[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetSystemTriggerContractAddr(blockNumber *big.Int) string {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn "0x1000000000000000000000000000000000000002"[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetSystemTriggerSelector(blockNumber *big.Int) []byte {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn []byte{0x7f, 0xec, 0x8d, 0x38}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetPrioritisedFTSOContract(blockTime *big.Int) string {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn "0x1000000000000000000000000000000000000003"[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetMaximumMintRequest(blockNumber *big.Int) *big.Int {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mmaxRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)[m
[32m+[m		[32mreturn maxRequest[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc triggerKeeper(evm EVMCaller) (*big.Int, error) {[m
[32m+[m	[32mbigZero := big.NewInt(0)[m
[32m+[m	[32m// Get the contract to call[m
[32m+[m	[32msystemTriggerContract := common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber()))[m
[32m+[m	[32m// Call the method[m
[32m+[m	[32mtriggerRet, _, triggerErr := evm.Call([m
[32m+[m		[32mvm.AccountRef(systemTriggerContract),[m
[32m+[m		[32msystemTriggerContract,[m
[32m+[m		[32mGetSystemTriggerSelector(evm.GetBlockNumber()),[m
[32m+[m		[32mGetKeeperGasMultiplier(evm.GetBlockNumber())*evm.GetGasLimit(),[m
[32m+[m		[32mbigZero)[m
[32m+[m	[32m// If no error and a value came back...[m
[32m+[m	[32mif triggerErr == nil && triggerRet != nil {[m
[32m+[m		[32m// Did we get one big int?[m
[32m+[m		[32mif len(triggerRet) == 32 {[m
[32m+[m			[32m// Convert to big int[m
[32m+[m			[32m// Mint request cannot be less than 0 as SetBytes treats value as unsigned[m
[32m+[m			[32mmintRequest := new(big.Int).SetBytes(triggerRet)[m
[32m+[m			[32m// return the mint request[m
[32m+[m			[32mreturn mintRequest, nil[m
[32m+[m		[32m} else {[m
[32m+[m			[32m// Returned length was not 32 bytes[m
[32m+[m			[32mreturn bigZero, &ErrInvalidKeeperData{}[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mif triggerErr != nil {[m
[32m+[m			[32mreturn bigZero, triggerErr[m
[32m+[m		[32m} else {[m
[32m+[m			[32mreturn bigZero, &ErrKeeperDataEmpty{}[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc mint(evm EVMCaller, mintRequest *big.Int) error {[m
[32m+[m	[32m// If the mint request is greater than zero and less than max[m
[32m+[m	[32mmax := GetMaximumMintRequest(evm.GetBlockNumber())[m
[32m+[m	[32mif mintRequest.Cmp(big.NewInt(0)) > 0 &&[m
[32m+[m		[32mmintRequest.Cmp(max) <= 0 {[m
[32m+[m		[32m// Mint the amount asked for on to the keeper contract[m
[32m+[m		[32mevm.AddBalance(common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber())), mintRequest)[m
[32m+[m	[32m} else if mintRequest.Cmp(max) > 0 {[m
[32m+[m		[32m// Return error[m
[32m+[m		[32mreturn &ErrMaxMintExceeded{[m
[32m+[m			[32mmintRequest: mintRequest,[m
[32m+[m			[32mmintMax:     max,[m
[32m+[m		[32m}[m
[32m+[m	[32m} else if mintRequest.Cmp(big.NewInt(0)) < 0 {[m
[32m+[m		[32m// Cannot mint negatives[m
[32m+[m		[32mreturn &ErrMintNegative{}[m
[32m+[m	[32m}[m
[32m+[m	[32m// No error[m
[32m+[m	[32mreturn nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc triggerKeeperAndMint(evm EVMCaller, log log.Logger) {[m
[32m+[m	[32m// Call the keeper[m
[32m+[m	[32mmintRequest, triggerErr := triggerKeeper(evm)[m
[32m+[m	[32m// If no error...[m
[32m+[m	[32mif triggerErr == nil {[m
[32m+[m		[32m// time to mint[m
[32m+[m		[32mif mintError := mint(evm, mintRequest); mintError != nil {[m
[32m+[m			[32mlog.Warn("Error minting inflation request", "error", mintError)[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mlog.Warn("Keeper trigger in error", "error", triggerErr)[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[1mdiff --git a/core/keeper_test.go b/core/keeper_test.go[m
[1mnew file mode 100644[m
[1mindex 0000000..10fff7c[m
[1m--- /dev/null[m
[1m+++ b/core/keeper_test.go[m
[36m@@ -0,0 +1,451 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m
[32m+[m[32mpackage core[m
[32m+[m
[32m+[m[32mimport ([m
[32m+[m	[32m"errors"[m
[32m+[m	[32m"math/big"[m
[32m+[m	[32m"testing"[m
[32m+[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/log"[m
[32m+[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m[32m)[m
[32m+[m
[32m+[m[32m// Define a mock structure to spy and mock values for keeper calls[m
[32m+[m[32mtype MockEVMCallerData struct {[m
[32m+[m	[32mcallCalls            int[m
[32m+[m	[32maddBalanceCalls      int[m
[32m+[m	[32mblockNumber          big.Int[m
[32m+[m	[32mgasLimit             uint64[m
[32m+[m	[32mmintRequestReturn    big.Int[m
[32m+[m	[32mlastAddBalanceAddr   common.Address[m
[32m+[m	[32mlastAddBalanceAmount *big.Int[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define a mock structure to spy and mock values for logger calls[m
[32m+[m[32mtype MockLoggerData struct {[m
[32m+[m	[32mwarnCalls int[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Set up default mock method calls[m
[32m+[m[32mfunc defautCall(e *MockEVMCallerData, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32me.callCalls++[m
[32m+[m
[32m+[m	[32mbuffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}[m
[32m+[m	[32mreturn e.mintRequestReturn.FillBytes(buffer), 0, nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc defaultGetBlockNumber(e *MockEVMCallerData) *big.Int {[m
[32m+[m	[32mreturn &e.blockNumber[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc defaultGetGasLimit(e *MockEVMCallerData) uint64 {[m
[32m+[m	[32mreturn e.gasLimit[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc defaultAddBalance(e *MockEVMCallerData, addr common.Address, amount *big.Int) {[m
[32m+[m	[32me.addBalanceCalls++[m
[32m+[m	[32me.lastAddBalanceAddr = addr[m
[32m+[m	[32me.lastAddBalanceAmount = amount[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define the default EVM mock and define default mock receiver functions[m
[32m+[m[32mtype DefaultEVMMock struct {[m
[32m+[m	[32mmockEVMCallerData MockEVMCallerData[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *DefaultEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32mreturn defautCall(&e.mockEVMCallerData, caller, addr, input, gas, value)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *DefaultEVMMock) GetBlockNumber() *big.Int {[m
[32m+[m	[32mreturn defaultGetBlockNumber(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *DefaultEVMMock) GetGasLimit() uint64 {[m
[32m+[m	[32mreturn defaultGetGasLimit(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *DefaultEVMMock) AddBalance(addr common.Address, amount *big.Int) {[m
[32m+[m	[32mdefaultAddBalance(&e.mockEVMCallerData, addr, amount)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldReturnMintRequest(t *testing.T) {[m
[32m+[m	[32mmintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *mintRequestReturn,[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mmintRequest, _ := triggerKeeper(defaultEVMMock)[m
[32m+[m
[32m+[m	[32mif mintRequest.Cmp(mintRequestReturn) != 0 {[m
[32m+[m		[32mt.Errorf("got %s want %q", mintRequest.Text(10), "50000000000000000000000000")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldNotLetMintRequestOverflow(t *testing.T) {[m
[32m+[m	[32mvar mintRequestReturn big.Int[m
[32m+[m	[32m// TODO: Compact with exponent?[m
[32m+[m	[32mbuffer := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}[m
[32m+[m	[32mmintRequestReturn.SetBytes(buffer)[m
[32m+[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: mintRequestReturn,[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mmintRequest, mintRequestError := triggerKeeper(defaultEVMMock)[m
[32m+[m
[32m+[m	[32mif mintRequestError != nil {[m
[32m+[m		[32mt.Errorf("received unexpected error %s", mintRequestError)[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mif mintRequest.Sign() < 1 {[m
[32m+[m		[32mt.Errorf("unexpected negative")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define a bad mint request return size mock[m
[32m+[m[32mtype BadMintReturnSizeEVMMock struct {[m
[32m+[m	[32mmockEVMCallerData MockEVMCallerData[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadMintReturnSizeEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32me.mockEVMCallerData.callCalls++[m
[32m+[m	[32m// Should be size 32 bytes[m
[32m+[m	[32mbuffer := []byte{0}[m
[32m+[m	[32mreturn e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadMintReturnSizeEVMMock) GetBlockNumber() *big.Int {[m
[32m+[m	[32mreturn defaultGetBlockNumber(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadMintReturnSizeEVMMock) GetGasLimit() uint64 {[m
[32m+[m	[32mreturn defaultGetGasLimit(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadMintReturnSizeEVMMock) AddBalance(addr common.Address, amount *big.Int) {[m
[32m+[m	[32mdefaultAddBalance(&e.mockEVMCallerData, addr, amount)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerValidatesMintRequestReturnValueSize(t *testing.T) {[m
[32m+[m	[32mvar mintRequestReturn big.Int[m
[32m+[m	[32m// TODO: Compact with exponent?[m
[32m+[m	[32mbuffer := []byte{255}[m
[32m+[m	[32mmintRequestReturn.SetBytes(buffer)[m
[32m+[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: mintRequestReturn,[m
[32m+[m	[32m}[m
[32m+[m	[32mbadMintReturnSizeEVMMock := &BadMintReturnSizeEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m	[32m// Call to return less than 32 bytes[m
[32m+[m	[32m_, err := triggerKeeper(badMintReturnSizeEVMMock)[m
[32m+[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mif err, ok := err.(*ErrInvalidKeeperData); !ok {[m
[32m+[m			[32mwant := &ErrInvalidKeeperData{}[m
[32m+[m			[32mt.Errorf("got '%s' want '%s'", err.Error(), want.Error())[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("no error returned as expected")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define a mock to simulate keeper trigger returning an error from Call[m
[32m+[m[32mtype BadTriggerCallEVMMock struct {[m
[32m+[m	[32mmockEVMCallerData MockEVMCallerData[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadTriggerCallEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32me.mockEVMCallerData.callCalls++[m
[32m+[m
[32m+[m	[32mbuffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}[m
[32m+[m	[32mreturn e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, errors.New("Call error happened")[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadTriggerCallEVMMock) GetBlockNumber() *big.Int {[m
[32m+[m	[32mreturn defaultGetBlockNumber(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadTriggerCallEVMMock) GetGasLimit() uint64 {[m
[32m+[m	[32mreturn defaultGetGasLimit(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *BadTriggerCallEVMMock) AddBalance(addr common.Address, amount *big.Int) {[m
[32m+[m	[32mdefaultAddBalance(&e.mockEVMCallerData, addr, amount)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerReturnsCallError(t *testing.T) {[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{}[m
[32m+[m	[32mbadTriggerCallEVMMock := &BadTriggerCallEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m	[32m// Call to return less than 32 bytes[m
[32m+[m	[32m_, err := triggerKeeper(badTriggerCallEVMMock)[m
[32m+[m
[32m+[m	[32mif err == nil {[m
[32m+[m		[32mt.Errorf("no error received")[m
[32m+[m	[32m} else {[m
[32m+[m		[32mif err.Error() != "Call error happened" {[m
[32m+[m			[32mt.Errorf("did not get expected error")[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype LoggerMock struct {[m
[32m+[m	[32mmockLoggerData MockLoggerData[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (l *LoggerMock) New(ctx ...interface{}) log.Logger {[m
[32m+[m	[32mreturn nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (l *LoggerMock) GetHandler() log.Handler {[m
[32m+[m	[32mreturn nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (l *LoggerMock) SetHandler(h log.Handler) {[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (l *LoggerMock) Trace(msg string, ctx ...interface{}) {}[m
[32m+[m[32mfunc (l *LoggerMock) Debug(msg string, ctx ...interface{}) {}[m
[32m+[m[32mfunc (l *LoggerMock) Info(msg string, ctx ...interface{})  {}[m
[32m+[m[32mfunc (l *LoggerMock) Error(msg string, ctx ...interface{}) {}[m
[32m+[m[32mfunc (l *LoggerMock) Crit(msg string, ctx ...interface{})  {}[m
[32m+[m
[32m+[m[32mfunc (l *LoggerMock) Warn(msg string, ctx ...interface{}) {[m
[32m+[m	[32ml.mockLoggerData.warnCalls++[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerAndMintLogsError(t *testing.T) {[m
[32m+[m	[32m// Assemble[m
[32m+[m	[32m// Set up mock EVM call to return an error[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{}[m
[32m+[m	[32mbadTriggerCallEVMMock := &BadTriggerCallEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m	[32m// Set up a mock logger[m
[32m+[m	[32mmockLoggerData := &MockLoggerData{}[m
[32m+[m	[32mloggerMock := &LoggerMock{[m
[32m+[m		[32mmockLoggerData: *mockLoggerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32m// Act[m
[32m+[m	[32mtriggerKeeperAndMint(badTriggerCallEVMMock, loggerMock)[m
[32m+[m
[32m+[m	[32m// Assert[m
[32m+[m	[32mif loggerMock.mockLoggerData.warnCalls != 1 {[m
[32m+[m		[32mt.Errorf("Logger.Warn not called as expected")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Define a mock to simulate keeper trigger returning nil for mint request[m
[32m+[m[32mtype ReturnNilMintRequestEVMMock struct {[m
[32m+[m	[32mmockEVMCallerData MockEVMCallerData[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *ReturnNilMintRequestEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32me.mockEVMCallerData.callCalls++[m
[32m+[m
[32m+[m	[32mreturn nil, 0, nil[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *ReturnNilMintRequestEVMMock) GetBlockNumber() *big.Int {[m
[32m+[m	[32mreturn defaultGetBlockNumber(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *ReturnNilMintRequestEVMMock) GetGasLimit() uint64 {[m
[32m+[m	[32mreturn defaultGetGasLimit(&e.mockEVMCallerData)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (e *ReturnNilMintRequestEVMMock) AddBalance(addr common.Address, amount *big.Int) {[m
[32m+[m	[32mdefaultAddBalance(&e.mockEVMCallerData, addr, amount)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerHandlesNilMintRequest(t *testing.T) {[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{}[m
[32m+[m	[32mreturnNilMintRequestEVMMock := &ReturnNilMintRequestEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m	[32m// Call to return less than 32 bytes[m
[32m+[m	[32m_, err := triggerKeeper(returnNilMintRequestEVMMock)[m
[32m+[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mif err, ok := err.(*ErrKeeperDataEmpty); !ok {[m
[32m+[m			[32mwant := &ErrKeeperDataEmpty{}[m
[32m+[m			[32mt.Errorf("got '%s' want '%s'", err.Error(), want.Error())[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("no error returned as expected")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldNotMintMoreThanMax(t *testing.T) {[m
[32m+[m	[32mmintRequest, _ := new(big.Int).SetString("50000000000000000000000001", 10)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *big.NewInt(0),[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32merr := mint(defaultEVMMock, mintRequest)[m
[32m+[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mif err, ok := err.(*ErrMaxMintExceeded); !ok {[m
[32m+[m			[32mwant := &ErrMaxMintExceeded{[m
[32m+[m				[32mmintRequest: mintRequest,[m
[32m+[m				[32mmintMax:     GetMaximumMintRequest(big.NewInt(0)),[m
[32m+[m			[32m}[m
[32m+[m			[32mt.Errorf("got '%s' want '%s'", err.Error(), want.Error())[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("no error returned as expected")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldNotMintNegative(t *testing.T) {[m
[32m+[m	[32mmintRequest := big.NewInt(-1)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *big.NewInt(0),[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32merr := mint(defaultEVMMock, mintRequest)[m
[32m+[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mif err, ok := err.(*ErrMintNegative); !ok {[m
[32m+[m			[32mwant := &ErrMintNegative{}[m
[32m+[m			[32mt.Errorf("got '%s' want '%s'", err.Error(), want.Error())[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("no error returned as expected")[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldMint(t *testing.T) {[m
[32m+[m	[32m// Assemble[m
[32m+[m	[32mmintRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *big.NewInt(0),[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32m// Act[m
[32m+[m	[32merr := mint(defaultEVMMock, mintRequest)[m
[32m+[m
[32m+[m	[32m// Assert[m
[32m+[m	[32mif err == nil {[m
[32m+[m		[32mif defaultEVMMock.mockEVMCallerData.addBalanceCalls != 1 {[m
[32m+[m			[32mt.Errorf("AddBalance not called as expected")[m
[32m+[m		[32m}[m
[32m+[m		[32mif defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr.String() != GetSystemTriggerContractAddr(big.NewInt(0)) {[m
[32m+[m			[32mt.Errorf("wanted addr %s; got addr %s", GetSystemTriggerContractAddr(big.NewInt(0)), defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr)[m
[32m+[m		[32m}[m
[32m+[m		[32mif defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Cmp(mintRequest) != 0 {[m
[32m+[m			[32mt.Errorf("wanted amount %s; got amount %s", mintRequest.Text(10), defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Text(10))[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("unexpected error returned; was = %s", err.Error())[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldNotErrorMintingZero(t *testing.T) {[m
[32m+[m	[32m// Assemble[m
[32m+[m	[32mmintRequest := big.NewInt(0)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *big.NewInt(0),[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32m// Act[m
[32m+[m	[32merr := mint(defaultEVMMock, mintRequest)[m
[32m+[m
[32m+[m	[32m// Assert[m
[32m+[m	[32mif err == nil {[m
[32m+[m		[32mif defaultEVMMock.mockEVMCallerData.addBalanceCalls != 0 {[m
[32m+[m			[32mt.Errorf("AddBalance called unexpectedly")[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mt.Errorf("unexpected error returned; was %s", err.Error())[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerFiredAndMinted(t *testing.T) {[m
[32m+[m	[32mmintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *mintRequestReturn,[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mlog := log.New()[m
[32m+[m	[32mtriggerKeeperAndMint(defaultEVMMock, log)[m
[32m+[m
[32m+[m	[32m// EVM Call function calling the keeper should have been cqlled[m
[32m+[m	[32mif defaultEVMMock.mockEVMCallerData.callCalls != 1 {[m
[32m+[m		[32mt.Errorf("EVM Call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.callCalls)[m
[32m+[m	[32m}[m
[32m+[m	[32m// AddBalance should have been called on the state database, minting the request asked for[m
[32m+[m	[32mif defaultEVMMock.mockEVMCallerData.addBalanceCalls != 1 {[m
[32m+[m		[32mt.Errorf("Add balance call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.addBalanceCalls)[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc TestKeeperTriggerShouldNotMintMoreThanLimit(t *testing.T) {[m
[32m+[m	[32mmintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000001", 10)[m
[32m+[m	[32mmockEVMCallerData := &MockEVMCallerData{[m
[32m+[m		[32mblockNumber:       *big.NewInt(0),[m
[32m+[m		[32mgasLimit:          0,[m
[32m+[m		[32mmintRequestReturn: *mintRequestReturn,[m
[32m+[m	[32m}[m
[32m+[m	[32mdefaultEVMMock := &DefaultEVMMock{[m
[32m+[m		[32mmockEVMCallerData: *mockEVMCallerData,[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mlog := log.New()[m
[32m+[m	[32mtriggerKeeperAndMint(defaultEVMMock, log)[m
[32m+[m
[32m+[m	[32m// EVM Call function calling the keeper should have been called[m
[32m+[m	[32mif defaultEVMMock.mockEVMCallerData.callCalls != 1 {[m
[32m+[m		[32mt.Errorf("EVM Call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.callCalls)[m
[32m+[m	[32m}[m
[32m+[m	[32m// AddBalance should not have been called on the state database, as the mint request was over the limit[m
[32m+[m	[32mif defaultEVMMock.mockEVMCallerData.addBalanceCalls != 0 {[m
[32m+[m		[32mt.Errorf("Add balance call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.addBalanceCalls)[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[1mdiff --git a/core/mkalloc.go b/core/mkalloc.go[m
[1mindex 96f9c31..046924b 100644[m
[1m--- a/core/mkalloc.go[m
[1m+++ b/core/mkalloc.go[m
[36m@@ -24,6 +24,7 @@[m
 // You should have received a copy of the GNU Lesser General Public License[m
 // along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.[m
 [m
[32m+[m[32m//go:build none[m
 // +build none[m
 [m
 /*[m
[36m@@ -44,8 +45,8 @@[m [mimport ([m
 	"sort"[m
 	"strconv"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
 )[m
 [m
 type allocItem struct{ Addr, Balance *big.Int }[m
[1mdiff --git a/core/rawdb/accessors_chain.go b/core/rawdb/accessors_chain.go[m
[1mindex 6ed9e0a..13e0f1a 100644[m
[1m--- a/core/rawdb/accessors_chain.go[m
[1m+++ b/core/rawdb/accessors_chain.go[m
[36m@@ -31,13 +31,13 @@[m [mimport ([m
 	"encoding/binary"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // ReadCanonicalHash retrieves the hash assigned to a canonical block number.[m
[1mdiff --git a/core/rawdb/accessors_chain_test.go b/core/rawdb/accessors_chain_test.go[m
[1mindex 38372a1..55b3ab4 100644[m
[1m--- a/core/rawdb/accessors_chain_test.go[m
[1m+++ b/core/rawdb/accessors_chain_test.go[m
[36m@@ -24,10 +24,10 @@[m [mimport ([m
 	"reflect"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 	"golang.org/x/crypto/sha3"[m
 )[m
 [m
[1mdiff --git a/core/rawdb/accessors_indexes.go b/core/rawdb/accessors_indexes.go[m
[1mindex 3c8b467..4cec301 100644[m
[1m--- a/core/rawdb/accessors_indexes.go[m
[1m+++ b/core/rawdb/accessors_indexes.go[m
[36m@@ -30,12 +30,12 @@[m [mimport ([m
 	"bytes"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // ReadTxLookupEntry retrieves the positional metadata associated with a transaction[m
[1mdiff --git a/core/rawdb/accessors_indexes_test.go b/core/rawdb/accessors_indexes_test.go[m
[1mindex a80ff1e..8efdc25 100644[m
[1m--- a/core/rawdb/accessors_indexes_test.go[m
[1m+++ b/core/rawdb/accessors_indexes_test.go[m
[36m@@ -22,10 +22,10 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 	"golang.org/x/crypto/sha3"[m
 )[m
 [m
[1mdiff --git a/core/rawdb/accessors_metadata.go b/core/rawdb/accessors_metadata.go[m
[1mindex 5b39253..47b9e9d 100644[m
[1m--- a/core/rawdb/accessors_metadata.go[m
[1m+++ b/core/rawdb/accessors_metadata.go[m
[36m@@ -29,11 +29,11 @@[m [mpackage rawdb[m
 import ([m
 	"encoding/json"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // ReadDatabaseVersion retrieves the version number of the database.[m
[1mdiff --git a/core/state/database.go b/core/state/database.go[m
[1mindex f4af1a0..feba572 100644[m
[1m--- a/core/state/database.go[m
[1m+++ b/core/state/database.go[m
[36m@@ -31,11 +31,11 @@[m [mimport ([m
 	"fmt"[m
 [m
 	"github.com/VictoriaMetrics/fastcache"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/trie"[m
 	lru "github.com/hashicorp/golang-lru"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/state/snapshot/conversion.go b/core/state/snapshot/conversion.go[m
[1mindex 0fd4e9b..d2b519f 100644[m
[1m--- a/core/state/snapshot/conversion.go[m
[1m+++ b/core/state/snapshot/conversion.go[m
[36m@@ -36,12 +36,12 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // trieKV represents a trie key-value pair[m
[1mdiff --git a/core/state/snapshot/disklayer.go b/core/state/snapshot/disklayer.go[m
[1mindex 106c451..c33f4a6 100644[m
[1m--- a/core/state/snapshot/disklayer.go[m
[1m+++ b/core/state/snapshot/disklayer.go[m
[36m@@ -32,11 +32,11 @@[m [mimport ([m
 	"time"[m
 [m
 	"github.com/VictoriaMetrics/fastcache"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // diskLayer is a low level persistent snapshot built on top of a key-value store.[m
[1mdiff --git a/core/state/snapshot/disklayer_test.go b/core/state/snapshot/disklayer_test.go[m
[1mindex f423c9a..9d758be 100644[m
[1m--- a/core/state/snapshot/disklayer_test.go[m
[1m+++ b/core/state/snapshot/disklayer_test.go[m
[36m@@ -32,12 +32,12 @@[m [mimport ([m
 	"os"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/ethdb/leveldb"[m
 	"github.com/ethereum/go-ethereum/ethdb/memorydb"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // reverse reverses the contents of a byte slice. It's used to update random accs[m
[1mdiff --git a/core/state/snapshot/generate.go b/core/state/snapshot/generate.go[m
[1mindex 45f927f..3a48d1f 100644[m
[1m--- a/core/state/snapshot/generate.go[m
[1m+++ b/core/state/snapshot/generate.go[m
[36m@@ -34,7 +34,6 @@[m [mimport ([m
 	"time"[m
 [m
 	"github.com/VictoriaMetrics/fastcache"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[36m@@ -42,6 +41,7 @@[m [mimport ([m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 var ([m
[1mdiff --git a/core/state/snapshot/iterator.go b/core/state/snapshot/iterator.go[m
[1mindex 7a5da6a..2c34768 100644[m
[1m--- a/core/state/snapshot/iterator.go[m
[1m+++ b/core/state/snapshot/iterator.go[m
[36m@@ -31,9 +31,9 @@[m [mimport ([m
 	"fmt"[m
 	"sort"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // Iterator is an iterator to step over all the accounts or the specific[m
[1mdiff --git a/core/state/snapshot/iterator_test.go b/core/state/snapshot/iterator_test.go[m
[1mindex 9d1b744..1916c49 100644[m
[1m--- a/core/state/snapshot/iterator_test.go[m
[1m+++ b/core/state/snapshot/iterator_test.go[m
[36m@@ -33,8 +33,8 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // TestAccountIteratorBasics tests some simple single-layer(diff and disk) iteration[m
[1mdiff --git a/core/state/snapshot/journal.go b/core/state/snapshot/journal.go[m
[1mindex 145f467..96c7c2e 100644[m
[1m--- a/core/state/snapshot/journal.go[m
[1m+++ b/core/state/snapshot/journal.go[m
[36m@@ -33,12 +33,12 @@[m [mimport ([m
 	"time"[m
 [m
 	"github.com/VictoriaMetrics/fastcache"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // journalGenerator is a disk layer entry containing the generator progress marker.[m
[1mdiff --git a/core/state/snapshot/snapshot.go b/core/state/snapshot/snapshot.go[m
[1mindex bbdba6b..e311d69 100644[m
[1m--- a/core/state/snapshot/snapshot.go[m
[1m+++ b/core/state/snapshot/snapshot.go[m
[36m@@ -36,12 +36,12 @@[m [mimport ([m
 	"time"[m
 [m
 	"github.com/VictoriaMetrics/fastcache"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/metrics"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/state/snapshot/snapshot_test.go b/core/state/snapshot/snapshot_test.go[m
[1mindex c28dafa..da5cd63 100644[m
[1m--- a/core/state/snapshot/snapshot_test.go[m
[1m+++ b/core/state/snapshot/snapshot_test.go[m
[36m@@ -32,9 +32,9 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // randomHash generates a random blob of data and returns it as a hash.[m
[1mdiff --git a/core/state/snapshot/wipe.go b/core/state/snapshot/wipe.go[m
[1mindex ca1f7e0..3cbf754 100644[m
[1m--- a/core/state/snapshot/wipe.go[m
[1m+++ b/core/state/snapshot/wipe.go[m
[36m@@ -30,10 +30,10 @@[m [mimport ([m
 	"bytes"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // wipeSnapshot starts a goroutine to iterate over the entire key-value database[m
[1mdiff --git a/core/state/snapshot/wipe_test.go b/core/state/snapshot/wipe_test.go[m
[1mindex 952370d..ba4ce69 100644[m
[1m--- a/core/state/snapshot/wipe_test.go[m
[1m+++ b/core/state/snapshot/wipe_test.go[m
[36m@@ -30,9 +30,9 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb/memorydb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 // Tests that given a database with random data content, all parts of a snapshot[m
[1mdiff --git a/core/state/state_test.go b/core/state/state_test.go[m
[1mindex 849f661..ec66ccc 100644[m
[1m--- a/core/state/state_test.go[m
[1m+++ b/core/state/state_test.go[m
[36m@@ -27,9 +27,9 @@[m
 package state[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 type stateTest struct {[m
[1mdiff --git a/core/state/statedb.go b/core/state/statedb.go[m
[1mindex acc779c..3fe062c 100644[m
[1m--- a/core/state/statedb.go[m
[1m+++ b/core/state/statedb.go[m
[36m@@ -34,15 +34,15 @@[m [mimport ([m
 	"sort"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state/snapshot"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/metrics"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state/snapshot"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 type revision struct {[m
[1mdiff --git a/core/state/statedb_test.go b/core/state/statedb_test.go[m
[1mindex 9c295f9..4841b6b 100644[m
[1m--- a/core/state/statedb_test.go[m
[1m+++ b/core/state/statedb_test.go[m
[36m@@ -39,11 +39,11 @@[m [mimport ([m
 	"testing"[m
 	"testing/quick"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state/snapshot"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state/snapshot"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // Tests that updating a state trie does not leak any database writes prior to[m
[1mdiff --git a/core/state_connector.go b/core/state_connector.go[m
[1mnew file mode 100644[m
[1mindex 0000000..db6665e[m
[1m--- /dev/null[m
[1m+++ b/core/state_connector.go[m
[36m@@ -0,0 +1,731 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m
[32m+[m[32mpackage core[m
[32m+[m
[32m+[m[32mimport ([m
[32m+[m	[32m"bytes"[m
[32m+[m	[32m"crypto/sha256"[m
[32m+[m	[32m"encoding/binary"[m
[32m+[m	[32m"encoding/hex"[m
[32m+[m	[32m"encoding/json"[m
[32m+[m	[32m"fmt"[m
[32m+[m	[32m"io/ioutil"[m
[32m+[m	[32m"math"[m
[32m+[m	[32m"math/big"[m
[32m+[m	[32m"net/http"[m
[32m+[m	[32m"os"[m
[32m+[m	[32m"strconv"[m
[32m+[m	[32m"strings"[m
[32m+[m	[32m"time"[m
[32m+[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m[32m)[m
[32m+[m
[32m+[m[32mvar ([m
[32m+[m	[32mtestingChainID               = new(big.Int).SetUint64(16)[m
[32m+[m	[32mstateConnectorActivationTime = new(big.Int).SetUint64(1636070400)[m
[32m+[m	[32mtr                           = &http.Transport{[m
[32m+[m		[32mMaxIdleConns:        100,[m
[32m+[m		[32mMaxConnsPerHost:     100,[m
[32m+[m		[32mMaxIdleConnsPerHost: 100,[m
[32m+[m		[32mIdleConnTimeout:     60 * time.Second,[m
[32m+[m		[32mDisableCompression:  true,[m
[32m+[m	[32m}[m
[32m+[m	[32mclient = &http.Client{[m
[32m+[m		[32mTransport: tr,[m
[32m+[m		[32mTimeout:   5 * time.Second,[m
[32m+[m	[32m}[m
[32m+[m	[32mapiRetries    = 3[m
[32m+[m	[32mapiRetryDelay = 1 * time.Second[m
[32m+[m[32m)[m
[32m+[m
[32m+[m[32mfunc GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {[m
[32m+[m	[32m// Return true if chainID is 16 or if block.timestamp is greater than the state connector activation time on any chain[m
[32m+[m	[32mreturn chainID.Cmp(testingChainID) == 0 || blockTime.Cmp(stateConnectorActivationTime) > 0[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetStateConnectorGasDivisor(blockTime *big.Int) uint64 {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn 3[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetMaxAllowedChains(blockTime *big.Int) uint32 {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn 5[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetStateConnectorContractAddr(blockTime *big.Int) string {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn "0x1000000000000000000000000000000000000001"[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetProveDataAvailabilityPeriodFinalitySelector(blockTime *big.Int) []byte {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn []byte{0xc5, 0xd6, 0x4c, 0xd1}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetProvePaymentFinalitySelector(blockTime *big.Int) []byte {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn []byte{0x38, 0x84, 0x92, 0xdd}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetDisprovePaymentFinalitySelector(blockTime *big.Int) []byte {[m
[32m+[m	[32mswitch {[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn []byte{0x7f, 0x58, 0x24, 0x32}[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// =======================================================[m
[32m+[m[32m// Proof of Work Common[m
[32m+[m[32m// =======================================================[m
[32m+[m
[32m+[m[32mtype GetPoWRequestPayload struct {[m
[32m+[m	[32mMethod string   `json:"method"`[m
[32m+[m	[32mParams []string `json:"params"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetPoWBlockCountResp struct {[m
[32m+[m	[32mResult uint64      `json:"result"`[m
[32m+[m	[32mError  interface{} `json:"error"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetPoWBlockCount(chainURL string, username string, password string) (uint64, bool) {[m
[32m+[m	[32mdata := GetPoWRequestPayload{[m
[32m+[m		[32mMethod: "getblockcount",[m
[32m+[m		[32mParams: []string{},[m
[32m+[m	[32m}[m
[32m+[m	[32mpayloadBytes, err := json.Marshal(data)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mbody := bytes.NewReader(payloadBytes)[m
[32m+[m	[32mreq, err := http.NewRequest("POST", chainURL, body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mreq.Header.Set("Content-Type", "application/json")[m
[32m+[m	[32mif username != "" && password != "" {[m
[32m+[m		[32mreq.SetBasicAuth(username, password)[m
[32m+[m	[32m}[m
[32m+[m	[32mresp, err := client.Do(req)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mdefer resp.Body.Close()[m
[32m+[m	[32mif resp.StatusCode != 200 {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespBody, err := ioutil.ReadAll(resp.Body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar jsonResp GetPoWBlockCountResp[m
[32m+[m	[32merr = json.Unmarshal(respBody, &jsonResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif jsonResp.Error != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn jsonResp.Result, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype GetPoWBlockHeaderResult struct {[m
[32m+[m	[32mHash          string `json:"hash"`[m
[32m+[m	[32mConfirmations uint64 `json:"confirmations"`[m
[32m+[m	[32mHeight        uint64 `json:"height"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetPoWBlockHeaderResp struct {[m
[32m+[m	[32mResult GetPoWBlockHeaderResult `json:"result"`[m
[32m+[m	[32mError  interface{}             `json:"error"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetPoWBlockHeader(ledgerHash string, requiredConfirmations uint64, chainURL string, username string, password string) (uint64, bool) {[m
[32m+[m	[32mdata := GetPoWRequestPayload{[m
[32m+[m		[32mMethod: "getblockheader",[m
[32m+[m		[32mParams: []string{[m
[32m+[m			[32mledgerHash,[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
[32m+[m	[32mpayloadBytes, err := json.Marshal(data)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mbody := bytes.NewReader(payloadBytes)[m
[32m+[m	[32mreq, err := http.NewRequest("POST", chainURL, body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mreq.Header.Set("Content-Type", "application/json")[m
[32m+[m	[32mif username != "" && password != "" {[m
[32m+[m		[32mreq.SetBasicAuth(username, password)[m
[32m+[m	[32m}[m
[32m+[m	[32mresp, err := client.Do(req)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mdefer resp.Body.Close()[m
[32m+[m	[32mif resp.StatusCode != 200 {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespBody, err := ioutil.ReadAll(resp.Body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar jsonResp GetPoWBlockHeaderResp[m
[32m+[m	[32merr = json.Unmarshal(respBody, &jsonResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif jsonResp.Error != nil {[m
[32m+[m		[32mreturn 0, false[m
[32m+[m	[32m} else if jsonResp.Result.Confirmations < requiredConfirmations {[m
[32m+[m		[32mreturn 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn jsonResp.Result.Height, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProveDataAvailabilityPeriodFinalityPoW(checkRet []byte, chainURL string, username string, password string) (bool, bool) {[m
[32m+[m	[32mblockCount, err := GetPoWBlockCount(chainURL, username, password)[m
[32m+[m	[32mif err {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m	[32mledger := binary.BigEndian.Uint64(checkRet[56:64])[m
[32m+[m	[32mrequiredConfirmations := binary.BigEndian.Uint64(checkRet[88:96])[m
[32m+[m	[32mif blockCount < ledger+requiredConfirmations {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m	[32mledgerResp, err := GetPoWBlockHeader(hex.EncodeToString(checkRet[96:128]), requiredConfirmations, chainURL, username, password)[m
[32m+[m	[32mif err {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m} else if ledgerResp > 0 && ledgerResp == ledger {[m
[32m+[m		[32mreturn true, false[m
[32m+[m	[32m} else {[m
[32m+[m		[32mreturn false, false[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype GetPoWTxRequestParams struct {[m
[32m+[m	[32mTxID    string `json:"txid"`[m
[32m+[m	[32mVerbose bool   `json:"verbose"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetPoWTxRequestPayload struct {[m
[32m+[m	[32mMethod string                `json:"method"`[m
[32m+[m	[32mParams GetPoWTxRequestParams `json:"params"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetPoWTxResult struct {[m
[32m+[m	[32mTxID          string `json:"txid"`[m
[32m+[m	[32mBlockHash     string `json:"blockhash"`[m
[32m+[m	[32mConfirmations uint64 `json:"confirmations"`[m
[32m+[m	[32mVout          []struct {[m
[32m+[m		[32mValue        float64 `json:"value"`[m
[32m+[m		[32mN            uint64  `json:"n"`[m
[32m+[m		[32mScriptPubKey struct {[m
[32m+[m			[32mType      string   `json:"type"`[m
[32m+[m			[32mAddresses []string `json:"addresses"`[m
[32m+[m		[32m} `json:"scriptPubKey"`[m
[32m+[m	[32m} `json:"vout"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetPoWTxResp struct {[m
[32m+[m	[32mResult GetPoWTxResult `json:"result"`[m
[32m+[m	[32mError  interface{}    `json:"error"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetPoWTx(txHash string, voutN uint64, latestAvailableBlock uint64, currencyCode string, chainURL string, username string, password string) ([]byte, uint64, bool) {[m
[32m+[m	[32mdata := GetPoWTxRequestPayload{[m
[32m+[m		[32mMethod: "getrawtransaction",[m
[32m+[m		[32mParams: GetPoWTxRequestParams{[m
[32m+[m			[32mTxID:    txHash[1:],[m
[32m+[m			[32mVerbose: true,[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
[32m+[m	[32mpayloadBytes, err := json.Marshal(data)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mbody := bytes.NewReader(payloadBytes)[m
[32m+[m	[32mreq, err := http.NewRequest("POST", chainURL, body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mreq.Header.Set("Content-Type", "application/json")[m
[32m+[m	[32mif username != "" && password != "" {[m
[32m+[m		[32mreq.SetBasicAuth(username, password)[m
[32m+[m	[32m}[m
[32m+[m	[32mresp, err := client.Do(req)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mdefer resp.Body.Close()[m
[32m+[m	[32mif resp.StatusCode != 200 {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespBody, err := ioutil.ReadAll(resp.Body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar jsonResp GetPoWTxResp[m
[32m+[m	[32merr = json.Unmarshal(respBody, &jsonResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif jsonResp.Error != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif uint64(len(jsonResp.Result.Vout)) <= voutN {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32mif jsonResp.Result.Vout[voutN].ScriptPubKey.Type != "pubkeyhash" || len(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses) != 1 {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32minBlock, getBlockErr := GetPoWBlockHeader(jsonResp.Result.BlockHash, jsonResp.Result.Confirmations, chainURL, username, password)[m
[32m+[m	[32mif getBlockErr {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif inBlock == 0 || inBlock >= latestAvailableBlock {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32mtxIdHash := crypto.Keccak256([]byte(txHash))[m
[32m+[m	[32mdestinationHash := crypto.Keccak256([]byte(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses[0]))[m
[32m+[m	[32mamountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp.Result.Vout[voutN].Value*math.Pow(10, 8)))), 32))[m
[32m+[m	[32mcurrencyHash := crypto.Keccak256([]byte(currencyCode))[m
[32m+[m	[32mreturn crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inBlock, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProvePaymentFinalityPoW(checkRet []byte, isDisprove bool, currencyCode string, chainURL string, username string, password string) (bool, bool) {[m
[32m+[m	[32mif len(checkRet) < 257 {[m
[32m+[m		[32mreturn false, false[m
[32m+[m	[32m}[m
[32m+[m	[32mvoutN, err := strconv.ParseUint(string(checkRet[192:193]), 16, 64)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn false, false[m
[32m+[m	[32m}[m
[32m+[m	[32mpaymentHash, inBlock, getPoWTxErr := GetPoWTx(string(checkRet[192:257]), voutN, binary.BigEndian.Uint64(checkRet[88:96]), currencyCode, chainURL, username, password)[m
[32m+[m	[32mif getPoWTxErr {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif !isDisprove {[m
[32m+[m		[32mif len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock == binary.BigEndian.Uint64(checkRet[56:64]) {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mif len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock > binary.BigEndian.Uint64(checkRet[56:64]) {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m} else if len(paymentHash) == 0 {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProvePoW(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, currencyCode string, chainURL string) (bool, bool) {[m
[32m+[m	[32mvar username, password string[m
[32m+[m	[32mchainURLhash := sha256.Sum256([]byte(chainURL))[m
[32m+[m	[32mchainURLchecksum := hex.EncodeToString(chainURLhash[0:4])[m
[32m+[m	[32mswitch currencyCode {[m
[32m+[m	[32mcase "btc":[m
[32m+[m		[32musername = os.Getenv("BTC_U_" + chainURLchecksum)[m
[32m+[m		[32mpassword = os.Getenv("BTC_P_" + chainURLchecksum)[m
[32m+[m	[32mcase "ltc":[m
[32m+[m		[32musername = os.Getenv("LTC_U_" + chainURLchecksum)[m
[32m+[m		[32mpassword = os.Getenv("LTC_P_" + chainURLchecksum)[m
[32m+[m	[32mcase "dog":[m
[32m+[m		[32musername = os.Getenv("DOGE_U_" + chainURLchecksum)[m
[32m+[m		[32mpassword = os.Getenv("DOGE_P_" + chainURLchecksum)[m
[32m+[m	[32m}[m
[32m+[m	[32mif bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProveDataAvailabilityPeriodFinalityPoW(checkRet, chainURL, username, password)[m
[32m+[m	[32m} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProvePaymentFinalityPoW(checkRet, false, currencyCode, chainURL, username, password)[m
[32m+[m	[32m} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProvePaymentFinalityPoW(checkRet, true, currencyCode, chainURL, username, password)[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// =======================================================[m
[32m+[m[32m// XRP[m
[32m+[m[32m// =======================================================[m
[32m+[m
[32m+[m[32mtype GetXRPBlockRequestParams struct {[m
[32m+[m	[32mLedgerIndex  uint64 `json:"ledger_index"`[m
[32m+[m	[32mFull         bool   `json:"full"`[m
[32m+[m	[32mAccounts     bool   `json:"accounts"`[m
[32m+[m	[32mTransactions bool   `json:"transactions"`[m
[32m+[m	[32mExpand       bool   `json:"expand"`[m
[32m+[m	[32mOwnerFunds   bool   `json:"owner_funds"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetXRPBlockRequestPayload struct {[m
[32m+[m	[32mMethod string                     `json:"method"`[m
[32m+[m	[32mParams []GetXRPBlockRequestParams `json:"params"`[m
[32m+[m[32m}[m
[32m+[m[32mtype CheckXRPErrorResponse struct {[m
[32m+[m	[32mError string `json:"error"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetXRPBlockResponse struct {[m
[32m+[m	[32mLedgerHash  string `json:"ledger_hash"`[m
[32m+[m	[32mLedgerIndex int    `json:"ledger_index"`[m
[32m+[m	[32mValidated   bool   `json:"validated"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetXRPBlock(ledger uint64, chainURL string) (string, bool) {[m
[32m+[m	[32mdata := GetXRPBlockRequestPayload{[m
[32m+[m		[32mMethod: "ledger",[m
[32m+[m		[32mParams: []GetXRPBlockRequestParams{[m
[32m+[m			[32m{[m
[32m+[m				[32mLedgerIndex:  ledger,[m
[32m+[m				[32mFull:         false,[m
[32m+[m				[32mAccounts:     false,[m
[32m+[m				[32mTransactions: false,[m
[32m+[m				[32mExpand:       false,[m
[32m+[m				[32mOwnerFunds:   false,[m
[32m+[m			[32m},[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
[32m+[m	[32mpayloadBytes, err := json.Marshal(data)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mbody := bytes.NewReader(payloadBytes)[m
[32m+[m	[32mreq, err := http.NewRequest("POST", chainURL, body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mreq.Header.Set("Content-Type", "application/json")[m
[32m+[m	[32mresp, err := client.Do(req)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mdefer resp.Body.Close()[m
[32m+[m	[32mif resp.StatusCode != 200 {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespBody, err := ioutil.ReadAll(resp.Body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar checkErrorResp map[string]CheckXRPErrorResponse[m
[32m+[m	[32merr = json.Unmarshal(respBody, &checkErrorResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mif checkErrorResp["result"].Error != "" {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar jsonResp map[string]GetXRPBlockResponse[m
[32m+[m	[32merr = json.Unmarshal(respBody, &jsonResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mif !jsonResp["result"].Validated {[m
[32m+[m		[32mreturn "", true[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn jsonResp["result"].LedgerHash, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProveDataAvailabilityPeriodFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {[m
[32m+[m	[32mledger := binary.BigEndian.Uint64(checkRet[56:64])[m
[32m+[m	[32mledgerHashString, err := GetXRPBlock(ledger, chainURL)[m
[32m+[m	[32mif err {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif ledgerHashString != "" && bytes.Equal(crypto.Keccak256([]byte(ledgerHashString)), checkRet[96:128]) {[m
[32m+[m		[32mreturn true, false[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype GetXRPTxRequestParams struct {[m
[32m+[m	[32mTransaction string `json:"transaction"`[m
[32m+[m	[32mBinary      bool   `json:"binary"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetXRPTxRequestPayload struct {[m
[32m+[m	[32mMethod string                  `json:"method"`[m
[32m+[m	[32mParams []GetXRPTxRequestParams `json:"params"`[m
[32m+[m[32m}[m
[32m+[m[32mtype GetXRPTxResponse struct {[m
[32m+[m	[32mDestination     string `json:"Destination"`[m
[32m+[m	[32mDestinationTag  int    `json:"DestinationTag"`[m
[32m+[m	[32mTransactionType string `json:"TransactionType"`[m
[32m+[m	[32mHash            string `json:"hash"`[m
[32m+[m	[32mInLedger        int    `json:"inLedger"`[m
[32m+[m	[32mValidated       bool   `json:"validated"`[m
[32m+[m	[32mMeta            struct {[m
[32m+[m		[32mTransactionResult string      `json:"TransactionResult"`[m
[32m+[m		[32mAmount            interface{} `json:"delivered_amount"`[m
[32m+[m	[32m} `json:"meta"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mtype GetXRPTxIssuedCurrency struct {[m
[32m+[m	[32mCurrency string `json:"currency"`[m
[32m+[m	[32mIssuer   string `json:"issuer"`[m
[32m+[m	[32mValue    string `json:"value"`[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetXRPTx(txHash string, latestAvailableLedger uint64, chainURL string) ([]byte, uint64, bool) {[m
[32m+[m	[32mdata := GetXRPTxRequestPayload{[m
[32m+[m		[32mMethod: "tx",[m
[32m+[m		[32mParams: []GetXRPTxRequestParams{[m
[32m+[m			[32m{[m
[32m+[m				[32mTransaction: txHash,[m
[32m+[m				[32mBinary:      false,[m
[32m+[m			[32m},[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
[32m+[m	[32mpayloadBytes, err := json.Marshal(data)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mbody := bytes.NewReader(payloadBytes)[m
[32m+[m	[32mreq, err := http.NewRequest("POST", chainURL, body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mreq.Header.Set("Content-Type", "application/json")[m
[32m+[m	[32mresp, err := client.Do(req)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mdefer resp.Body.Close()[m
[32m+[m	[32mif resp.StatusCode != 200 {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespBody, err := ioutil.ReadAll(resp.Body)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mvar checkErrorResp map[string]CheckXRPErrorResponse[m
[32m+[m	[32merr = json.Unmarshal(respBody, &checkErrorResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, true[m
[32m+[m	[32m}[m
[32m+[m	[32mrespErrString := checkErrorResp["result"].Error[m
[32m+[m	[32mif respErrString != "" {[m
[32m+[m		[32mif respErrString == "amendmentBlocked" ||[m
[32m+[m			[32mrespErrString == "failedToForward" ||[m
[32m+[m			[32mrespErrString == "invalid_API_version" ||[m
[32m+[m			[32mrespErrString == "noClosed" ||[m
[32m+[m			[32mrespErrString == "noCurrent" ||[m
[32m+[m			[32mrespErrString == "noNetwork" ||[m
[32m+[m			[32mrespErrString == "tooBusy" {[m
[32m+[m			[32mreturn []byte{}, 0, true[m
[32m+[m		[32m} else {[m
[32m+[m			[32mreturn []byte{}, 0, false[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m	[32mvar jsonResp map[string]GetXRPTxResponse[m
[32m+[m	[32merr = json.Unmarshal(respBody, &jsonResp)[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32mif jsonResp["result"].TransactionType != "Payment" || !jsonResp["result"].Validated || jsonResp["result"].Meta.TransactionResult != "tesSUCCESS" {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32minLedger := uint64(jsonResp["result"].InLedger)[m
[32m+[m	[32mif inLedger == 0 || inLedger >= latestAvailableLedger || !jsonResp["result"].Validated {[m
[32m+[m		[32mreturn []byte{}, 0, false[m
[32m+[m	[32m}[m
[32m+[m	[32mvar currency string[m
[32m+[m	[32mvar amount uint64[m
[32m+[m	[32mif stringAmount, ok := jsonResp["result"].Meta.Amount.(string); ok {[m
[32m+[m		[32mamount, err = strconv.ParseUint(stringAmount, 10, 64)[m
[32m+[m		[32mif err != nil {[m
[32m+[m			[32mreturn []byte{}, 0, false[m
[32m+[m		[32m}[m
[32m+[m		[32mcurrency = "xrp"[m
[32m+[m	[32m} else {[m
[32m+[m		[32mamountStruct, err := json.Marshal(jsonResp["result"].Meta.Amount)[m
[32m+[m		[32mif err != nil {[m
[32m+[m			[32mreturn []byte{}, 0, false[m
[32m+[m		[32m}[m
[32m+[m		[32mvar issuedCurrencyResp GetXRPTxIssuedCurrency[m
[32m+[m		[32merr = json.Unmarshal(amountStruct, &issuedCurrencyResp)[m
[32m+[m		[32mif err != nil {[m
[32m+[m			[32mreturn []byte{}, 0, false[m
[32m+[m		[32m}[m
[32m+[m		[32mfloatAmount, err := strconv.ParseFloat(issuedCurrencyResp.Value, 64)[m
[32m+[m		[32mif err != nil {[m
[32m+[m			[32mreturn []byte{}, 0, false[m
[32m+[m		[32m}[m
[32m+[m		[32mamount = uint64(floatAmount * math.Pow(10, 15))[m
[32m+[m		[32mcurrency = issuedCurrencyResp.Currency + issuedCurrencyResp.Issuer[m
[32m+[m	[32m}[m
[32m+[m	[32mtxIdHash := crypto.Keccak256([]byte(jsonResp["result"].Hash))[m
[32m+[m	[32mdestinationHash := crypto.Keccak256([]byte(jsonResp["result"].Destination))[m
[32m+[m	[32mdestinationTagHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp["result"].DestinationTag))), 32))[m
[32m+[m	[32mdestinationHash = crypto.Keccak256(destinationHash, destinationTagHash)[m
[32m+[m	[32mamountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(amount)), 32))[m
[32m+[m	[32mcurrencyHash := crypto.Keccak256([]byte(currency))[m
[32m+[m	[32mreturn crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inLedger, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProvePaymentFinalityXRP(checkRet []byte, isDisprove bool, chainURL string) (bool, bool) {[m
[32m+[m	[32mpaymentHash, inLedger, err := GetXRPTx(string(checkRet[192:]), binary.BigEndian.Uint64(checkRet[88:96]), chainURL)[m
[32m+[m	[32mif err {[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m	[32mif !isDisprove {[m
[32m+[m		[32mif len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger == binary.BigEndian.Uint64(checkRet[56:64]) {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mif len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger > binary.BigEndian.Uint64(checkRet[56:64]) {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m} else if len(paymentHash) == 0 {[m
[32m+[m			[32mreturn true, false[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ProveXRP(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {[m
[32m+[m	[32mif bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProveDataAvailabilityPeriodFinalityXRP(checkRet, chainURL)[m
[32m+[m	[32m} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProvePaymentFinalityXRP(checkRet, false, chainURL)[m
[32m+[m	[32m} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {[m
[32m+[m		[32mreturn ProvePaymentFinalityXRP(checkRet, true, chainURL)[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// =======================================================[m
[32m+[m[32m// ALGO[m
[32m+[m[32m// =======================================================[m
[32m+[m
[32m+[m[32mfunc ProveALGO(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {[m
[32m+[m	[32mreturn false, false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// =======================================================[m
[32m+[m[32m// Common[m
[32m+[m[32m// =======================================================[m
[32m+[m
[32m+[m[32mfunc ProveChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainId uint32, chainURL string) (bool, bool) {[m
[32m+[m	[32mswitch chainId {[m
[32m+[m	[32mcase 0:[m
[32m+[m		[32mreturn ProvePoW(sender, blockTime, functionSelector, checkRet, "btc", chainURL)[m
[32m+[m	[32mcase 1:[m
[32m+[m		[32mreturn ProvePoW(sender, blockTime, functionSelector, checkRet, "ltc", chainURL)[m
[32m+[m	[32mcase 2:[m
[32m+[m		[32mreturn ProvePoW(sender, blockTime, functionSelector, checkRet, "dog", chainURL)[m
[32m+[m	[32mcase 3:[m
[32m+[m		[32mreturn ProveXRP(sender, blockTime, functionSelector, checkRet, chainURL)[m
[32m+[m	[32mcase 4:[m
[32m+[m		[32mreturn ProveALGO(sender, blockTime, functionSelector, checkRet, chainURL)[m
[32m+[m	[32mdefault:[m
[32m+[m		[32mreturn false, true[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc ReadChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {[m
[32m+[m	[32mchainId := binary.BigEndian.Uint32(checkRet[28:32])[m
[32m+[m	[32mvar chainURLs string[m
[32m+[m	[32mswitch chainId {[m
[32m+[m	[32mcase 0:[m
[32m+[m		[32mchainURLs = os.Getenv("BTC_APIs")[m
[32m+[m	[32mcase 1:[m
[32m+[m		[32mchainURLs = os.Getenv("LTC_APIs")[m
[32m+[m	[32mcase 2:[m
[32m+[m		[32mchainURLs = os.Getenv("DOGE_APIs")[m
[32m+[m	[32mcase 3:[m
[32m+[m		[32mchainURLs = os.Getenv("XRP_APIs")[m
[32m+[m	[32mcase 4:[m
[32m+[m		[32mchainURLs = os.Getenv("ALGO_APIs")[m
[32m+[m	[32m}[m
[32m+[m	[32mif chainURLs == "" {[m
[32m+[m		[32mreturn false[m
[32m+[m	[32m}[m
[32m+[m	[32mfor i := 0; i < apiRetries; i++ {[m
[32m+[m		[32mfor _, chainURL := range strings.Split(chainURLs, ",") {[m
[32m+[m			[32mif chainURL == "" {[m
[32m+[m				[32mcontinue[m
[32m+[m			[32m}[m
[32m+[m			[32mverified, err := ProveChain(sender, blockTime, functionSelector, checkRet, chainId, chainURL)[m
[32m+[m			[32mif !verified && err {[m
[32m+[m				[32mcontinue[m
[32m+[m			[32m}[m
[32m+[m			[32mreturn verified[m
[32m+[m		[32m}[m
[32m+[m		[32mtime.Sleep(apiRetryDelay)[m
[32m+[m	[32m}[m
[32m+[m	[32mreturn false[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc GetVerificationPaths(functionSelector []byte, checkRet []byte) (string, string) {[m
[32m+[m	[32mprefix := "cache/"[m
[32m+[m	[32macceptedPrefix := "ACCEPTED"[m
[32m+[m	[32mrejectedPrefix := "REJECTED"[m
[32m+[m	[32mfunctionHash := hex.EncodeToString(functionSelector[:])[m
[32m+[m	[32mverificationHash := hex.EncodeToString(crypto.Keccak256(checkRet[0:64], checkRet[96:128]))[m
[32m+[m	[32msuffix := "_" + functionHash + "_" + verificationHash[m
[32m+[m	[32mreturn prefix + acceptedPrefix + suffix, prefix + rejectedPrefix + suffix[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32m// Verify proof against underlying chain[m
[32m+[m[32mfunc StateConnectorCall(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {[m
[32m+[m	[32mif binary.BigEndian.Uint64(checkRet[88:96]) > 0 {[m
[32m+[m		[32mgo func() {[m
[32m+[m			[32macceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)[m
[32m+[m			[32m_, errACCEPTED := os.Stat(acceptedPath)[m
[32m+[m			[32m_, errREJECTED := os.Stat(rejectedPath)[m
[32m+[m			[32mif errACCEPTED != nil && errREJECTED != nil {[m
[32m+[m				[32mif ReadChain(sender, blockTime, functionSelector, checkRet) {[m
[32m+[m					[32mverificationHashStore, err := os.Create(acceptedPath)[m
[32m+[m					[32mverificationHashStore.Close()[m
[32m+[m					[32mif err != nil {[m
[32m+[m						[32m// Permissions problem[m
[32m+[m						[32mpanic(err)[m
[32m+[m					[32m}[m
[32m+[m				[32m} else {[m
[32m+[m					[32mverificationHashStore, err := os.Create(rejectedPath)[m
[32m+[m					[32mverificationHashStore.Close()[m
[32m+[m					[32mif err != nil {[m
[32m+[m						[32m// Permissions problem[m
[32m+[m						[32mpanic(err)[m
[32m+[m					[32m}[m
[32m+[m				[32m}[m
[32m+[m			[32m}[m
[32m+[m		[32m}()[m
[32m+[m		[32mreturn true[m
[32m+[m	[32m} else {[m
[32m+[m		[32macceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)[m
[32m+[m		[32m_, errACCEPTED := os.Stat(acceptedPath)[m
[32m+[m		[32m_, errREJECTED := os.Stat(rejectedPath)[m
[32m+[m		[32mif errACCEPTED != nil && errREJECTED != nil {[m
[32m+[m			[32mfor i := 0; i < 2*apiRetries; i++ {[m
[32m+[m				[32m_, errACCEPTED = os.Stat(acceptedPath)[m
[32m+[m				[32m_, errREJECTED = os.Stat(rejectedPath)[m
[32m+[m				[32mif errACCEPTED == nil || errREJECTED == nil {[m
[32m+[m					[32mbreak[m
[32m+[m				[32m}[m
[32m+[m				[32mtime.Sleep(apiRetryDelay)[m
[32m+[m			[32m}[m
[32m+[m		[32m}[m
[32m+[m		[32mgo func() {[m
[32m+[m			[32mremoveFulfilledAPIRequests := os.Getenv("REMOVE_FULFILLED_API_REQUESTS")[m
[32m+[m			[32mif removeFulfilledAPIRequests == "1" {[m
[32m+[m				[32merrDeleteACCEPTED := os.Remove(acceptedPath)[m
[32m+[m				[32merrDeleteREJECTED := os.Remove(rejectedPath)[m
[32m+[m				[32mif errDeleteACCEPTED != nil && errDeleteREJECTED != nil {[m
[32m+[m					[32m// Permissions problem[m
[32m+[m					[32mpanic(fmt.Sprintf("%s\n%s", errDeleteACCEPTED, errDeleteREJECTED))[m
[32m+[m				[32m}[m
[32m+[m			[32m}[m
[32m+[m		[32m}()[m
[32m+[m		[32mreturn errACCEPTED == nil[m
[32m+[m	[32m}[m
[32m+[m[32m}[m
[1mdiff --git a/core/state_manager.go b/core/state_manager.go[m
[1mindex 8b76069..c95e03f 100644[m
[1m--- a/core/state_manager.go[m
[1m+++ b/core/state_manager.go[m
[36m@@ -30,9 +30,9 @@[m [mimport ([m
 	"fmt"[m
 	"math/rand"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/state_manager_test.go b/core/state_manager_test.go[m
[1mindex 17e2b1d..f57d105 100644[m
[1m--- a/core/state_manager_test.go[m
[1m+++ b/core/state_manager_test.go[m
[36m@@ -7,7 +7,7 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 [m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/stretchr/testify/assert"[m
[1mdiff --git a/core/state_prefetcher.go b/core/state_prefetcher.go[m
[1mindex a1e0cde..f0fd973 100644[m
[1m--- a/core/state_prefetcher.go[m
[1m+++ b/core/state_prefetcher.go[m
[36m@@ -30,11 +30,11 @@[m [mimport ([m
 	"math/big"[m
 	"sync/atomic"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // statePrefetcher is a basic Prefetcher, which blindly executes a block on top[m
[1mdiff --git a/core/state_processor.go b/core/state_processor.go[m
[1mindex b9c243f..ec6d709 100644[m
[1m--- a/core/state_processor.go[m
[1m+++ b/core/state_processor.go[m
[36m@@ -30,14 +30,14 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/consensus/misc"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/misc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // StateProcessor is a basic Processor, which takes care of transitioning[m
[1mdiff --git a/core/state_transition.go b/core/state_transition.go[m
[1mindex 381327c..876d39d 100644[m
[1m--- a/core/state_transition.go[m
[1m+++ b/core/state_transition.go[m
[36m@@ -1,3 +1,11 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m//[m
[32m+[m[32m// This file is a derived work, based on the avalanchego library whose original[m
[32m+[m[32m// notices appear below. It is distributed under a license compatible with the[m
[32m+[m[32m// licensing terms of the original code from which it is derived.[m
[32m+[m[32m// Please see the file LICENSE_AVALABS for licensing terms of the original work.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m[32m//[m
 // (c) 2019-2020, Ava Labs, Inc.[m
 //[m
 // This file is a derived work, based on the go-ethereum library whose original[m
[36m@@ -27,16 +35,19 @@[m
 package core[m
 [m
 import ([m
[32m+[m	[32m"bytes"[m
[32m+[m	[32m"encoding/binary"[m
 	"fmt"[m
 	"math"[m
 	"math/big"[m
 [m
[32m+[m	[32m"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"github.com/ethereum/go-ethereum/log"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var emptyCodeHash = crypto.Keccak256Hash(nil)[m
[36m@@ -115,6 +126,23 @@[m [mfunc (result *ExecutionResult) Return() []byte {[m
 	return common.CopyBytes(result.ReturnData)[m
 }[m
 [m
[32m+[m[32m// Implement the EVMCaller interface on the state transition structure; simply delegate the calls[m
[32m+[m[32mfunc (st *StateTransition) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {[m
[32m+[m	[32mreturn st.evm.Call(caller, addr, input, gas, value)[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (st *StateTransition) GetBlockNumber() *big.Int {[m
[32m+[m	[32mreturn st.evm.Context.BlockNumber[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (st *StateTransition) GetGasLimit() uint64 {[m
[32m+[m	[32mreturn st.evm.Context.GasLimit[m
[32m+[m[32m}[m
[32m+[m
[32m+[m[32mfunc (st *StateTransition) AddBalance(addr common.Address, amount *big.Int) {[m
[32m+[m	[32mst.state.AddBalance(addr, amount)[m
[32m+[m[32m}[m
[32m+[m
 // Revert returns the concrete revert reason if the execution is aborted by `REVERT`[m
 // opcode. Note the reason can be nil if no data supplied with revert opcode.[m
 func (result *ExecutionResult) Revert() []byte {[m
[36m@@ -321,19 +349,91 @@[m [mfunc (st *StateTransition) TransitionDb() (*ExecutionResult, error) {[m
 	if rules := st.evm.ChainConfig().AvalancheRules(st.evm.Context.BlockNumber, st.evm.Context.Time); rules.IsApricotPhase2 {[m
 		st.state.PrepareAccessList(msg.From(), msg.To(), vm.ActivePrecompiles(rules), msg.AccessList())[m
 	}[m
[32m+[m
 	var ([m
[31m-		ret   []byte[m
[31m-		vmerr error // vm errors do not effect consensus and are therefore not assigned to err[m
[32m+[m		[32mret                                       []byte[m
[32m+[m		[32mvmerr                                     error // vm errors do not affect consensus and are therefore not assigned to err[m
[32m+[m		[32mselectProveDataAvailabilityPeriodFinality bool[m
[32m+[m		[32mselectProvePaymentFinality                bool[m
[32m+[m		[32mselectDisprovePaymentFinality             bool[m
[32m+[m		[32mprioritisedFTSOContract                   bool[m
 	)[m
[31m-	if contractCreation {[m
[31m-		ret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)[m
[31m-	} else {[m
[32m+[m
[32m+[m	[32mif st.evm.Context.Coinbase != common.HexToAddress("0x0100000000000000000000000000000000000000") {[m
[32m+[m		[32mreturn nil, fmt.Errorf("Invalid value for block.coinbase")[m
[32m+[m	[32m}[m
[32m+[m	[32mif st.msg.From() == common.HexToAddress("0x0100000000000000000000000000000000000000") ||[m
[32m+[m		[32mst.msg.From() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) ||[m
[32m+[m		[32mst.msg.From() == common.HexToAddress(GetSystemTriggerContractAddr(st.evm.Context.Time)) {[m
[32m+[m		[32mreturn nil, fmt.Errorf("Invalid sender")[m
[32m+[m	[32m}[m
[32m+[m	[32mburnAddress := st.evm.Context.Coinbase[m
[32m+[m	[32mif !contractCreation {[m
[32m+[m		[32mif *msg.To() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) && len(st.data) >= 4 {[m
[32m+[m			[32mselectProveDataAvailabilityPeriodFinality = bytes.Equal(st.data[0:4], GetProveDataAvailabilityPeriodFinalitySelector(st.evm.Context.Time))[m
[32m+[m			[32mselectProvePaymentFinality = bytes.Equal(st.data[0:4], GetProvePaymentFinalitySelector(st.evm.Context.Time))[m
[32m+[m			[32mselectDisprovePaymentFinality = bytes.Equal(st.data[0:4], GetDisprovePaymentFinalitySelector(st.evm.Context.Time))[m
[32m+[m		[32m} else {[m
[32m+[m			[32mprioritisedFTSOContract = *msg.To() == common.HexToAddress(GetPrioritisedFTSOContract(st.evm.Context.Time))[m
[32m+[m		[32m}[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mif selectProveDataAvailabilityPeriodFinality || selectProvePaymentFinality || selectDisprovePaymentFinality {[m
 		// Increment the nonce for the next transaction[m
 		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)[m
[31m-		ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)[m
[32m+[m		[32mstateConnectorGas := st.gas / GetStateConnectorGasDivisor(st.evm.Context.Time)[m
[32m+[m		[32mcheckRet, _, checkVmerr := st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)[m
[32m+[m		[32mif checkVmerr == nil {[m
[32m+[m			[32mchainConfig := st.evm.ChainConfig()[m
[32m+[m			[32mif GetStateConnectorActivated(chainConfig.ChainID, st.evm.Context.Time) && binary.BigEndian.Uint32(checkRet[28:32]) < GetMaxAllowedChains(st.evm.Context.Time) {[m
[32m+[m				[32mif StateConnectorCall(msg.From(), st.evm.Context.Time, st.data[0:4], checkRet) {[m
[32m+[m					[32moriginalCoinbase := st.evm.Context.Coinbase[m
[32m+[m					[32mdefer func() {[m
[32m+[m						[32mst.evm.Context.Coinbase = originalCoinbase[m
[32m+[m					[32m}()[m
[32m+[m					[32mst.evm.Context.Coinbase = st.msg.From()[m
[32m+[m				[32m}[m
[32m+[m			[32m}[m
[32m+[m		[32m}[m
[32m+[m		[32mret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)[m
[32m+[m	[32m} else {[m
[32m+[m		[32mif contractCreation {[m
[32m+[m			[32mret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)[m
[32m+[m		[32m} else {[m
[32m+[m			[32m// Increment the nonce for the next transaction[m
[32m+[m			[32mst.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)[m
[32m+[m			[32mret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)[m
[32m+[m		[32m}[m
 	}[m
 	st.refundGas(apricotPhase1)[m
[31m-	st.state.AddBalance(st.evm.Context.Coinbase, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))[m
[32m+[m	[32mif vmerr == nil && prioritisedFTSOContract {[m
[32m+[m		[32mnominalGasUsed := uint64(21000)[m
[32m+[m		[32mnominalGasPrice := uint64(225_000_000_000)[m
[32m+[m		[32mnominalFee := new(big.Int).Mul(new(big.Int).SetUint64(nominalGasUsed), new(big.Int).SetUint64(nominalGasPrice))[m
[32m+[m		[32mactualGasUsed := st.gasUsed()[m
[32m+[m		[32mactualGasPrice := st.gasPrice[m
[32m+[m		[32mactualFee := new(big.Int).Mul(new(big.Int).SetUint64(actualGasUsed), actualGasPrice)[m
[32m+[m		[32mif actualFee.Cmp(nominalFee) > 0 {[m
[32m+[m			[32mfeeRefund := new(big.Int).Sub(actualFee, nominalFee)[m
[32m+[m			[32mst.state.AddBalance(st.msg.From(), feeRefund)[m
[32m+[m			[32mst.state.AddBalance(burnAddress, nominalFee)[m
[32m+[m		[32m} else {[m
[32m+[m			[32mst.state.AddBalance(burnAddress, actualFee)[m
[32m+[m		[32m}[m
[32m+[m	[32m} else {[m
[32m+[m		[32mst.state.AddBalance(burnAddress, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32m// Call the keeper contract trigger method if there is no vm error[m
[32m+[m	[32mif vmerr == nil {[m
[32m+[m		[32m// Temporarily disable EVM debugging[m
[32m+[m		[32moldDebug := st.evm.Config.Debug[m
[32m+[m		[32mst.evm.Config.Debug = false[m
[32m+[m		[32m// Call the keeper contract trigger[m
[32m+[m		[32mlog := log.Root()[m
[32m+[m		[32mtriggerKeeperAndMint(st, log)[m
[32m+[m		[32mst.evm.Config.Debug = oldDebug[m
[32m+[m	[32m}[m
 [m
 	return &ExecutionResult{[m
 		UsedGas:    st.gasUsed(),[m
[1mdiff --git a/core/test_blockchain.go b/core/test_blockchain.go[m
[1mindex 604ed0f..cfdd2b9 100644[m
[1m--- a/core/test_blockchain.go[m
[1m+++ b/core/test_blockchain.go[m
[36m@@ -8,13 +8,13 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type ChainTest struct {[m
[1mdiff --git a/core/tx_cacher.go b/core/tx_cacher.go[m
[1mindex 86679ff..03099f2 100644[m
[1m--- a/core/tx_cacher.go[m
[1m+++ b/core/tx_cacher.go[m
[36m@@ -29,7 +29,7 @@[m [mpackage core[m
 import ([m
 	"runtime"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // senderCacher is a concurrent transaction sender recoverer and cacher.[m
[1mdiff --git a/core/tx_journal.go b/core/tx_journal.go[m
[1mindex b2bfa53..17b4419 100644[m
[1m--- a/core/tx_journal.go[m
[1m+++ b/core/tx_journal.go[m
[36m@@ -31,10 +31,10 @@[m [mimport ([m
 	"io"[m
 	"os"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // errNoActiveJournal is returned if a transaction is attempted to be inserted[m
[1mdiff --git a/core/tx_list.go b/core/tx_list.go[m
[1mindex 6788bad..2e28237 100644[m
[1m--- a/core/tx_list.go[m
[1m+++ b/core/tx_list.go[m
[36m@@ -33,8 +33,8 @@[m [mimport ([m
 	"sort"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // nonceHeap is a heap.Interface implementation over 64bit unsigned integers for[m
[1mdiff --git a/core/tx_list_test.go b/core/tx_list_test.go[m
[1mindex 86cf79a..91e4d6e 100644[m
[1m--- a/core/tx_list_test.go[m
[1m+++ b/core/tx_list_test.go[m
[36m@@ -31,8 +31,8 @@[m [mimport ([m
 	"math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // Tests that transactions can be added to strict lists and list contents and[m
[1mdiff --git a/core/tx_noncer.go b/core/tx_noncer.go[m
[1mindex bb5968b..439b9c7 100644[m
[1m--- a/core/tx_noncer.go[m
[1m+++ b/core/tx_noncer.go[m
[36m@@ -29,8 +29,8 @@[m [mpackage core[m
 import ([m
 	"sync"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
 )[m
 [m
 // txNoncer is a tiny virtual state database to manage the executable nonces of[m
[1mdiff --git a/core/tx_pool.go b/core/tx_pool.go[m
[1mindex 457be8f..9663b17 100644[m
[1m--- a/core/tx_pool.go[m
[1m+++ b/core/tx_pool.go[m
[36m@@ -35,15 +35,15 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/prque"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/metrics"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 const ([m
[1mdiff --git a/core/tx_pool_test.go b/core/tx_pool_test.go[m
[1mindex 84353e9..97e6c7e 100644[m
[1m--- a/core/tx_pool_test.go[m
[1m+++ b/core/tx_pool_test.go[m
[36m@@ -29,14 +29,14 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/core/types.go b/core/types.go[m
[1mindex 9607212..6a0d938 100644[m
[1m--- a/core/types.go[m
[1m+++ b/core/types.go[m
[36m@@ -27,9 +27,9 @@[m
 package core[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 )[m
 [m
 // Validator is an interface which defines the standard for block validation. It[m
[1mdiff --git a/core/types/block_test.go b/core/types/block_test.go[m
[1mindex e2f5623..85db341 100644[m
[1m--- a/core/types/block_test.go[m
[1m+++ b/core/types/block_test.go[m
[36m@@ -33,11 +33,11 @@[m [mimport ([m
 	"reflect"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 	"golang.org/x/crypto/sha3"[m
 )[m
 [m
[1mdiff --git a/core/types/hashing_test.go b/core/types/hashing_test.go[m
[1mindex 424798d..e537e47 100644[m
[1m--- a/core/types/hashing_test.go[m
[1m+++ b/core/types/hashing_test.go[m
[36m@@ -34,12 +34,12 @@[m [mimport ([m
 	mrand "math/rand"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 func TestDeriveSha(t *testing.T) {[m
[1mdiff --git a/core/types/receipt.go b/core/types/receipt.go[m
[1mindex 6da7c15..90366de 100644[m
[1m--- a/core/types/receipt.go[m
[1m+++ b/core/types/receipt.go[m
[36m@@ -34,11 +34,11 @@[m [mimport ([m
 	"math/big"[m
 	"unsafe"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 //go:generate gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go[m
[1mdiff --git a/core/types/receipt_test.go b/core/types/receipt_test.go[m
[1mindex 95fe68a..d2a2508 100644[m
[1m--- a/core/types/receipt_test.go[m
[1m+++ b/core/types/receipt_test.go[m
[36m@@ -33,10 +33,10 @@[m [mimport ([m
 	"reflect"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestDecodeEmptyTypedReceipt(t *testing.T) {[m
[1mdiff --git a/core/types/transaction_signing.go b/core/types/transaction_signing.go[m
[1mindex a717749..982d10d 100644[m
[1m--- a/core/types/transaction_signing.go[m
[1m+++ b/core/types/transaction_signing.go[m
[36m@@ -32,9 +32,9 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ErrInvalidChainId = errors.New("invalid chain id for signer")[m
[1mdiff --git a/core/vm/access_list_tracer.go b/core/vm/access_list_tracer.go[m
[1mindex c481c54..22c8f3e 100644[m
[1m--- a/core/vm/access_list_tracer.go[m
[1m+++ b/core/vm/access_list_tracer.go[m
[36m@@ -30,8 +30,8 @@[m [mimport ([m
 	"math/big"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // accessList is an accumulator for the set of accounts and storage slots an EVM[m
[1mdiff --git a/core/vm/contracts.go b/core/vm/contracts.go[m
[1mindex cfb2671..31b8d8f 100644[m
[1m--- a/core/vm/contracts.go[m
[1m+++ b/core/vm/contracts.go[m
[36m@@ -32,13 +32,13 @@[m [mimport ([m
 	"errors"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/crypto/blake2b"[m
 	"github.com/ethereum/go-ethereum/crypto/bls12381"[m
 	"github.com/ethereum/go-ethereum/crypto/bn256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
 	//lint:ignore SA1019 Needed for precompile[m
 	"golang.org/x/crypto/ripemd160"[m
[1mdiff --git a/core/vm/contracts_stateful.go b/core/vm/contracts_stateful.go[m
[1mindex 981ebd7..05b0361 100644[m
[1m--- a/core/vm/contracts_stateful.go[m
[1m+++ b/core/vm/contracts_stateful.go[m
[36m@@ -7,9 +7,9 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // PrecompiledContractsApricot contains the default set of pre-compiled Ethereum[m
[1mdiff --git a/core/vm/contracts_stateful_test.go b/core/vm/contracts_stateful_test.go[m
[1mindex 6126078..1278a67 100644[m
[1m--- a/core/vm/contracts_stateful_test.go[m
[1m+++ b/core/vm/contracts_stateful_test.go[m
[36m@@ -7,12 +7,12 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/stretchr/testify/assert"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestPrecompiledContractSpendsGas(t *testing.T) {[m
[1mdiff --git a/core/vm/eips.go b/core/vm/eips.go[m
[1mindex e525a73..399a874 100644[m
[1m--- a/core/vm/eips.go[m
[1m+++ b/core/vm/eips.go[m
[36m@@ -30,8 +30,8 @@[m [mimport ([m
 	"fmt"[m
 	"sort"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var activators = map[int]func(*JumpTable){[m
[1mdiff --git a/core/vm/evm.go b/core/vm/evm.go[m
[1mindex 9ad9db1..6f78cc6 100644[m
[1m--- a/core/vm/evm.go[m
[1m+++ b/core/vm/evm.go[m
[36m@@ -31,10 +31,10 @@[m [mimport ([m
 	"sync/atomic"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // emptyCodeHash is used by create to ensure deployment is disallowed to already[m
[1mdiff --git a/core/vm/gas_table.go b/core/vm/gas_table.go[m
[1mindex 383fb6d..d611fc5 100644[m
[1m--- a/core/vm/gas_table.go[m
[1m+++ b/core/vm/gas_table.go[m
[36m@@ -29,9 +29,9 @@[m [mpackage vm[m
 import ([m
 	"errors"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // memoryGasCost calculates the quadratic gas for memory expansion. It does so[m
[1mdiff --git a/core/vm/gas_table_test.go b/core/vm/gas_table_test.go[m
[1mindex 92d5d30..2150bb2 100644[m
[1m--- a/core/vm/gas_table_test.go[m
[1m+++ b/core/vm/gas_table_test.go[m
[36m@@ -31,11 +31,11 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestMemoryGasCost(t *testing.T) {[m
[1mdiff --git a/core/vm/instructions.go b/core/vm/instructions.go[m
[1mindex 81f3d3e..3fe1c4b 100644[m
[1m--- a/core/vm/instructions.go[m
[1m+++ b/core/vm/instructions.go[m
[36m@@ -29,10 +29,10 @@[m [mpackage vm[m
 import ([m
 	"errors"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 	"golang.org/x/crypto/sha3"[m
 )[m
 [m
[1mdiff --git a/core/vm/instructions_test.go b/core/vm/instructions_test.go[m
[1mindex 1570b36..50fb6db 100644[m
[1m--- a/core/vm/instructions_test.go[m
[1m+++ b/core/vm/instructions_test.go[m
[36m@@ -33,10 +33,10 @@[m [mimport ([m
 	"io/ioutil"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type TwoOperandTestcase struct {[m
[1mdiff --git a/core/vm/interface.go b/core/vm/interface.go[m
[1mindex bde4b08..6148970 100644[m
[1m--- a/core/vm/interface.go[m
[1m+++ b/core/vm/interface.go[m
[36m@@ -29,8 +29,8 @@[m [mpackage vm[m
 import ([m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // StateDB is an EVM database for full state querying.[m
[1mdiff --git a/core/vm/jump_table.go b/core/vm/jump_table.go[m
[1mindex d833e52..249736d 100644[m
[1m--- a/core/vm/jump_table.go[m
[1m+++ b/core/vm/jump_table.go[m
[36m@@ -27,7 +27,7 @@[m
 package vm[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type ([m
[1mdiff --git a/core/vm/logger.go b/core/vm/logger.go[m
[1mindex f8d088d..e49d4f7 100644[m
[1m--- a/core/vm/logger.go[m
[1m+++ b/core/vm/logger.go[m
[36m@@ -34,12 +34,12 @@[m [mimport ([m
 	"strings"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // Storage represents a contract's storage.[m
[1mdiff --git a/core/vm/logger_test.go b/core/vm/logger_test.go[m
[1mindex 1d87e46..a35c4e6 100644[m
[1m--- a/core/vm/logger_test.go[m
[1m+++ b/core/vm/logger_test.go[m
[36m@@ -30,10 +30,10 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/holiman/uint256"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type dummyContractRef struct {[m
[1mdiff --git a/core/vm/operations_acl.go b/core/vm/operations_acl.go[m
[1mindex adfe772..d9e2f06 100644[m
[1m--- a/core/vm/operations_acl.go[m
[1m+++ b/core/vm/operations_acl.go[m
[36m@@ -29,9 +29,9 @@[m [mpackage vm[m
 import ([m
 	"errors"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // gasSStoreEIP2929 implements gas cost for SSTORE according to EIP-2929[m
[1mdiff --git a/core/vm/runtime/env.go b/core/vm/runtime/env.go[m
[1mindex 5cae37d..6ebd9c6 100644[m
[1m--- a/core/vm/runtime/env.go[m
[1m+++ b/core/vm/runtime/env.go[m
[36m@@ -27,8 +27,8 @@[m
 package runtime[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 )[m
 [m
 func NewEnv(cfg *Config) *vm.EVM {[m
[1mdiff --git a/core/vm/runtime/runtime.go b/core/vm/runtime/runtime.go[m
[1mindex eb6e3f6..bd24dba 100644[m
[1m--- a/core/vm/runtime/runtime.go[m
[1m+++ b/core/vm/runtime/runtime.go[m
[36m@@ -31,12 +31,12 @@[m [mimport ([m
 	"math/big"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // Config is a basic type specifying certain configuration flags for running[m
[1mdiff --git a/core/vm/runtime/runtime_example_test.go b/core/vm/runtime/runtime_example_test.go[m
[1mindex 9850e28..c126459 100644[m
[1m--- a/core/vm/runtime/runtime_example_test.go[m
[1m+++ b/core/vm/runtime/runtime_example_test.go[m
[36m@@ -29,8 +29,8 @@[m [mpackage runtime_test[m
 import ([m
 	"fmt"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/vm/runtime"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm/runtime"[m
 )[m
 [m
 func ExampleExecute() {[m
[1mdiff --git a/core/vm/runtime/runtime_test.go b/core/vm/runtime/runtime_test.go[m
[1mindex 4b8d0a9..8561e31 100644[m
[1m--- a/core/vm/runtime/runtime_test.go[m
[1m+++ b/core/vm/runtime/runtime_test.go[m
[36m@@ -34,16 +34,16 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/accounts/abi"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/core/asm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestDefaults(t *testing.T) {[m
[1mdiff --git a/core/vm/stack_table.go b/core/vm/stack_table.go[m
[1mindex 487acae..e75f21a 100644[m
[1m--- a/core/vm/stack_table.go[m
[1m+++ b/core/vm/stack_table.go[m
[36m@@ -27,7 +27,7 @@[m
 package vm[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func minSwapStack(n int) int {[m
[1mdiff --git a/eth/api.go b/eth/api.go[m
[1mindex 62dd67b..f0d2bce 100644[m
[1m--- a/eth/api.go[m
[1m+++ b/eth/api.go[m
[36m@@ -35,16 +35,16 @@[m [mimport ([m
 	"os"[m
 	"strings"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/internal/ethapi"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/internal/ethapi"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // PublicEthereumAPI provides an API to access Ethereum full node-related[m
[1mdiff --git a/eth/api_backend.go b/eth/api_backend.go[m
[1mindex 5152b20..29e81d9 100644[m
[1m--- a/eth/api_backend.go[m
[1m+++ b/eth/api_backend.go[m
[36m@@ -32,21 +32,21 @@[m [mimport ([m
 	"math/big"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/bloombits"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/eth/gasprice"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/eth/downloader"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/bloombits"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/gasprice"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 var ([m
[1mdiff --git a/eth/backend.go b/eth/backend.go[m
[1mindex 380548e..8c32ac2 100644[m
[1m--- a/eth/backend.go[m
[1m+++ b/eth/backend.go[m
[36m@@ -33,28 +33,28 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/bloombits"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/eth/ethconfig"[m
[31m-	"github.com/ava-labs/coreth/eth/filters"[m
[31m-	"github.com/ava-labs/coreth/eth/gasprice"[m
[31m-	"github.com/ava-labs/coreth/eth/tracers"[m
[31m-	"github.com/ava-labs/coreth/internal/ethapi"[m
[31m-	"github.com/ava-labs/coreth/miner"[m
[31m-	"github.com/ava-labs/coreth/node"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/eth/downloader"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/bloombits"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/ethconfig"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/filters"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/gasprice"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/tracers"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/internal/ethapi"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/miner"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Config contains the configuration options of the ETH protocol.[m
[1mdiff --git a/eth/bloombits.go b/eth/bloombits.go[m
[1mindex ecc0aaf..b45ed5b 100644[m
[1m--- a/eth/bloombits.go[m
[1m+++ b/eth/bloombits.go[m
[36m@@ -29,8 +29,8 @@[m [mpackage eth[m
 import ([m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
 	"github.com/ethereum/go-ethereum/common/bitutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
 )[m
 [m
 const ([m
[1mdiff --git a/eth/ethconfig/config.go b/eth/ethconfig/config.go[m
[1mindex 99ea015..bbfa10b 100644[m
[1m--- a/eth/ethconfig/config.go[m
[1m+++ b/eth/ethconfig/config.go[m
[36m@@ -29,10 +29,10 @@[m [mpackage ethconfig[m
 import ([m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/eth/gasprice"[m
[31m-	"github.com/ava-labs/coreth/miner"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/gasprice"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/miner"[m
 )[m
 [m
 // DefaultFullGPOConfig contains default gasprice oracle settings for full node.[m
[1mdiff --git a/eth/filters/api.go b/eth/filters/api.go[m
[1mindex f543d6e..50339df 100644[m
[1m--- a/eth/filters/api.go[m
[1m+++ b/eth/filters/api.go[m
[36m@@ -35,13 +35,13 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // filter is a helper struct that holds meta information over the filter type[m
[1mdiff --git a/eth/filters/api_test.go b/eth/filters/api_test.go[m
[1mindex 6f35639..2903c0b 100644[m
[1m--- a/eth/filters/api_test.go[m
[1m+++ b/eth/filters/api_test.go[m
[36m@@ -21,8 +21,8 @@[m [mimport ([m
 	"fmt"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 func TestUnmarshalJSONNewFilterArgs(t *testing.T) {[m
[1mdiff --git a/eth/filters/filter.go b/eth/filters/filter.go[m
[1mindex 94769a1..e8f9ee7 100644[m
[1m--- a/eth/filters/filter.go[m
[1m+++ b/eth/filters/filter.go[m
[36m@@ -32,15 +32,15 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/bloombits"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/bloombits"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 type Backend interface {[m
[1mdiff --git a/eth/filters/filter_system.go b/eth/filters/filter_system.go[m
[1mindex 86f5a0b..db72584 100644[m
[1m--- a/eth/filters/filter_system.go[m
[1m+++ b/eth/filters/filter_system.go[m
[36m@@ -34,14 +34,14 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Type determines the kind of filter and is used to put the filter in to[m
[1mdiff --git a/eth/gasprice/gasprice.go b/eth/gasprice/gasprice.go[m
[1mindex b1d415b..954b124 100644[m
[1m--- a/eth/gasprice/gasprice.go[m
[1m+++ b/eth/gasprice/gasprice.go[m
[36m@@ -33,12 +33,12 @@[m [mimport ([m
 	"sync"[m
 [m
 	"github.com/ava-labs/avalanchego/utils/timer"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 const sampleNumber = 3 // Number of transactions sampled in a block[m
[1mdiff --git a/eth/gasprice/gasprice_test.go b/eth/gasprice/gasprice_test.go[m
[1mindex cfea2d4..97733eb 100644[m
[1m--- a/eth/gasprice/gasprice_test.go[m
[1m+++ b/eth/gasprice/gasprice_test.go[m
[36m@@ -32,15 +32,15 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 type testBackend struct {[m
[1mdiff --git a/eth/state_accessor.go b/eth/state_accessor.go[m
[1mindex e94b2d4..adad3d1 100644[m
[1m--- a/eth/state_accessor.go[m
[1m+++ b/eth/state_accessor.go[m
[36m@@ -32,13 +32,13 @@[m [mimport ([m
 	"math/big"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 )[m
 [m
 // stateAtBlock retrieves the state database associated with a certain block.[m
[1mdiff --git a/eth/tracers/api.go b/eth/tracers/api.go[m
[1mindex c54cd93..46e7505 100644[m
[1m--- a/eth/tracers/api.go[m
[1m+++ b/eth/tracers/api.go[m
[36m@@ -36,19 +36,19 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/internal/ethapi"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/internal/ethapi"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 const ([m
[1mdiff --git a/eth/tracers/api_test.go b/eth/tracers/api_test.go[m
[1mindex 454f35c..4122c90 100644[m
[1m--- a/eth/tracers/api_test.go[m
[1m+++ b/eth/tracers/api_test.go[m
[36m@@ -38,20 +38,20 @@[m [mimport ([m
 	"sort"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/internal/ethapi"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/internal/ethapi"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 var ([m
[1mdiff --git a/eth/tracers/tracer.go b/eth/tracers/tracer.go[m
[1mindex 3f4325a..4670281 100644[m
[1m--- a/eth/tracers/tracer.go[m
[1m+++ b/eth/tracers/tracer.go[m
[36m@@ -35,12 +35,12 @@[m [mimport ([m
 	"time"[m
 	"unsafe"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
 	"gopkg.in/olebedev/go-duktape.v3"[m
 )[m
 [m
[1mdiff --git a/eth/tracers/tracer_test.go b/eth/tracers/tracer_test.go[m
[1mindex 3d4e94c..e8ce389 100644[m
[1m--- a/eth/tracers/tracer_test.go[m
[1m+++ b/eth/tracers/tracer_test.go[m
[36m@@ -33,10 +33,10 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type account struct{}[m
[1mdiff --git a/eth/tracers/tracers.go b/eth/tracers/tracers.go[m
[1mindex dbe53f3..049cb18 100644[m
[1m--- a/eth/tracers/tracers.go[m
[1m+++ b/eth/tracers/tracers.go[m
[36m@@ -31,7 +31,7 @@[m [mimport ([m
 	"strings"[m
 	"unicode"[m
 [m
[31m-	"github.com/ava-labs/coreth/eth/tracers/internal/tracers"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/tracers/internal/tracers"[m
 )[m
 [m
 // all contains all the built in JavaScript tracers by name.[m
[1mdiff --git a/eth/tracers/tracers_test.go b/eth/tracers/tracers_test.go[m
[1mindex f9e55e2..ded3c86 100644[m
[1m--- a/eth/tracers/tracers_test.go[m
[1m+++ b/eth/tracers/tracers_test.go[m
[36m@@ -37,17 +37,17 @@[m [mimport ([m
 	"strings"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/rawdb"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/tests"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/rawdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/tests"[m
 )[m
 [m
 // To generate a new callTracer test, copy paste the makeTest method below into[m
[1mdiff --git a/ethclient/corethclient/corethclient.go b/ethclient/corethclient/corethclient.go[m
[1mindex 8d7654f..da9ca7f 100644[m
[1m--- a/ethclient/corethclient/corethclient.go[m
[1m+++ b/ethclient/corethclient/corethclient.go[m
[36m@@ -33,12 +33,12 @@[m [mimport ([m
 	"runtime"[m
 	"runtime/debug"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/ethclient"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/ethclient"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Client is a wrapper around rpc.Client that implements geth-specific functionality.[m
[1mdiff --git a/ethclient/ethclient.go b/ethclient/ethclient.go[m
[1mindex 0959a0a..8d30ca2 100644[m
[1m--- a/ethclient/ethclient.go[m
[1m+++ b/ethclient/ethclient.go[m
[36m@@ -35,11 +35,11 @@[m [mimport ([m
 	"math/big"[m
 [m
 	"github.com/ava-labs/avalanchego/ids"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/interfaces"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/interfaces"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Client defines typed wrappers for the Ethereum RPC API.[m
[1mdiff --git a/ethclient/signer.go b/ethclient/signer.go[m
[1mindex dafa943..45ef52d 100644[m
[1m--- a/ethclient/signer.go[m
[1m+++ b/ethclient/signer.go[m
[36m@@ -30,8 +30,8 @@[m [mimport ([m
 	"errors"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // senderFromServer is a types.Signer that remembers the sender address returned by the RPC[m
[1mdiff --git a/go.mod b/go.mod[m
[1mindex 981a551..fac1e12 100644[m
[1m--- a/go.mod[m
[1m+++ b/go.mod[m
[36m@@ -1,4 +1,4 @@[m
[31m-module github.com/ava-labs/coreth[m
[32m+[m[32mmodule gitlab.com/flarenetwork/coreth[m
 [m
 go 1.15[m
 [m
[1mdiff --git a/interfaces/interfaces.go b/interfaces/interfaces.go[m
[1mindex efcde08..be8870f 100644[m
[1m--- a/interfaces/interfaces.go[m
[1m+++ b/interfaces/interfaces.go[m
[36m@@ -32,8 +32,8 @@[m [mimport ([m
 	"errors"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 // NotFound is returned by API methods if the requested item does not exist.[m
[1mdiff --git a/internal/ethapi/api.go b/internal/ethapi/api.go[m
[1mindex b9ad558..e110a19 100644[m
[1m--- a/internal/ethapi/api.go[m
[1m+++ b/internal/ethapi/api.go[m
[36m@@ -35,15 +35,6 @@[m [mimport ([m
 	"time"[m
 [m
 	"github.com/ava-labs/avalanchego/ids"[m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/accounts/keystore"[m
[31m-	"github.com/ava-labs/coreth/accounts/scwallet"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/davecgh/go-spew/spew"[m
 	"github.com/ethereum/go-ethereum/accounts/abi"[m
 	"github.com/ethereum/go-ethereum/common"[m
[36m@@ -53,6 +44,15 @@[m [mimport ([m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 	"github.com/tyler-smith/go-bip39"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/keystore"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/scwallet"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // PublicEthereumAPI provides an API to access Ethereum related information.[m
[1mdiff --git a/internal/ethapi/backend.go b/internal/ethapi/backend.go[m
[1mindex 3d03734..f3b1643 100644[m
[1m--- a/internal/ethapi/backend.go[m
[1m+++ b/internal/ethapi/backend.go[m
[36m@@ -31,19 +31,19 @@[m [mimport ([m
 	"context"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/bloombits"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/eth/downloader"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/bloombits"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Backend interface provides the common API services (that are provided by[m
[1mdiff --git a/internal/ethapi/transaction_args.go b/internal/ethapi/transaction_args.go[m
[1mindex 6b06742..dcadfd7 100644[m
[1m--- a/internal/ethapi/transaction_args.go[m
[1m+++ b/internal/ethapi/transaction_args.go[m
[36m@@ -33,12 +33,12 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // TransactionArgs represents the arguments to construct a new transaction[m
[1mdiff --git a/miner/miner.go b/miner/miner.go[m
[1mindex 263439c..24ddc22 100644[m
[1m--- a/miner/miner.go[m
[1m+++ b/miner/miner.go[m
[36m@@ -28,12 +28,12 @@[m
 package miner[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/event"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // Backend wraps all methods required for mining.[m
[1mdiff --git a/miner/worker.go b/miner/worker.go[m
[1mindex c8c7268..69fa1f6 100644[m
[1m--- a/miner/worker.go[m
[1m+++ b/miner/worker.go[m
[36m@@ -36,16 +36,16 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/consensus"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/consensus/misc"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/event"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/misc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // environment is the worker's current environment and holds all of the current state information.[m
[1mdiff --git a/node/api.go b/node/api.go[m
[1mindex c38b2df..0c36fc7 100644[m
[1m--- a/node/api.go[m
[1m+++ b/node/api.go[m
[36m@@ -27,10 +27,10 @@[m
 package node[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/internal/debug"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/internal/debug"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // apis returns the collection of built-in RPC APIs.[m
[1mdiff --git a/node/config.go b/node/config.go[m
[1mindex 4d2b83a..b51e67a 100644[m
[1m--- a/node/config.go[m
[1m+++ b/node/config.go[m
[36m@@ -32,12 +32,12 @@[m [mimport ([m
 	"os"[m
 	"path/filepath"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/accounts/external"[m
[31m-	"github.com/ava-labs/coreth/accounts/keystore"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/external"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts/keystore"[m
 [m
[31m-	"github.com/ava-labs/coreth/rpc"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Config represents a small collection of configuration values to fine tune the[m
[1mdiff --git a/node/defaults.go b/node/defaults.go[m
[1mindex e4c826b..e3e2257 100644[m
[1m--- a/node/defaults.go[m
[1m+++ b/node/defaults.go[m
[36m@@ -27,7 +27,7 @@[m
 package node[m
 [m
 import ([m
[31m-	"github.com/ava-labs/coreth/rpc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 const ([m
[1mdiff --git a/node/node.go b/node/node.go[m
[1mindex 265e00d..1272a9f 100644[m
[1m--- a/node/node.go[m
[1m+++ b/node/node.go[m
[36m@@ -31,8 +31,8 @@[m [mimport ([m
 [m
 	"github.com/ethereum/go-ethereum/event"[m
 [m
[31m-	"github.com/ava-labs/coreth/accounts"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/accounts"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 )[m
 [m
 // Node is a container on which services can be registered.[m
[1mdiff --git a/plugin/evm/block.go b/plugin/evm/block.go[m
[1mindex cae83af..d749229 100644[m
[1m--- a/plugin/evm/block.go[m
[1m+++ b/plugin/evm/block.go[m
[36m@@ -12,8 +12,8 @@[m [mimport ([m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
 	"github.com/ava-labs/avalanchego/ids"[m
 	"github.com/ava-labs/avalanchego/snow/choices"[m
[1mdiff --git a/plugin/evm/block_verification.go b/plugin/evm/block_verification.go[m
[1mindex 03910af..1e52707 100644[m
[1m--- a/plugin/evm/block_verification.go[m
[1m+++ b/plugin/evm/block_verification.go[m
[36m@@ -7,11 +7,11 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	coreth "github.com/ava-labs/coreth/chain"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/trie"[m
[32m+[m	[32mcoreth "gitlab.com/flarenetwork/coreth/chain"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 var ([m
[1mdiff --git a/plugin/evm/config.go b/plugin/evm/config.go[m
[1mindex f184a09..721bf75 100644[m
[1m--- a/plugin/evm/config.go[m
[1m+++ b/plugin/evm/config.go[m
[36m@@ -7,8 +7,8 @@[m [mimport ([m
 	"encoding/json"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/eth"[m
 	"github.com/spf13/cast"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
 )[m
 [m
 const ([m
[1mdiff --git a/plugin/evm/export_tx.go b/plugin/evm/export_tx.go[m
[1mindex 3b50915..be2275f 100644[m
[1m--- a/plugin/evm/export_tx.go[m
[1m+++ b/plugin/evm/export_tx.go[m
[36m@@ -1,5 +1,12 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m//[m
[32m+[m[32m// This file is a derived work, based on the avalanchego library whose original[m
[32m+[m[32m// notice appears below. It is distributed under a license compatible with the[m
[32m+[m[32m// licensing terms of the original code from which it is derived.[m
[32m+[m[32m// Please see the file LICENSE_AVALABS for licensing terms of the original work.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m[32m//[m
 // (c) 2019-2020, Ava Labs, Inc. All rights reserved.[m
[31m-// See the file LICENSE for licensing terms.[m
 [m
 package evm[m
 [m
[36m@@ -7,19 +14,14 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
[31m-	"github.com/ava-labs/avalanchego/chains/atomic"[m
 	"github.com/ava-labs/avalanchego/database"[m
 	"github.com/ava-labs/avalanchego/ids"[m
 	"github.com/ava-labs/avalanchego/snow"[m
 	"github.com/ava-labs/avalanchego/utils/crypto"[m
[31m-	"github.com/ava-labs/avalanchego/utils/math"[m
 	"github.com/ava-labs/avalanchego/vms/components/avax"[m
[31m-	"github.com/ava-labs/avalanchego/vms/secp256k1fx"[m
[31m-	"github.com/ethereum/go-ethereum/common"[m
[31m-	"github.com/ethereum/go-ethereum/log"[m
 )[m
 [m
 // UnsignedExportTx is an unsigned ExportTx[m
[36m@@ -46,75 +48,16 @@[m [mfunc (tx *UnsignedExportTx) Verify([m
 	ctx *snow.Context,[m
 	rules params.Rules,[m
 ) error {[m
[31m-	switch {[m
[31m-	case tx == nil:[m
[31m-		return errNilTx[m
[31m-	case tx.DestinationChain != avmID:[m
[31m-		return errWrongChainID[m
[31m-	case len(tx.ExportedOutputs) == 0:[m
[31m-		return errNoExportOutputs[m
[31m-	case tx.NetworkID != ctx.NetworkID:[m
[31m-		return errWrongNetworkID[m
[31m-	case ctx.ChainID != tx.BlockchainID:[m
[31m-		return errWrongBlockchainID[m
[31m-	}[m
[31m-[m
[31m-	for _, in := range tx.Ins {[m
[31m-		if err := in.Verify(); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	for _, out := range tx.ExportedOutputs {[m
[31m-		if err := out.Verify(); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-	}[m
[31m-	if !avax.IsSortedTransferableOutputs(tx.ExportedOutputs, Codec) {[m
[31m-		return errOutputsNotSorted[m
[31m-	}[m
[31m-	if rules.IsApricotPhase1 && !IsSortedAndUniqueEVMInputs(tx.Ins) {[m
[31m-		return errInputsNotSortedUnique[m
[31m-	}[m
[31m-[m
[31m-	return nil[m
[32m+[m	[32mreturn errWrongChainID[m
 }[m
 [m
 func (tx *UnsignedExportTx) Cost() (uint64, error) {[m
[31m-	byteCost := calcBytesCost(len(tx.UnsignedBytes()))[m
[31m-	numSigs := uint64(len(tx.Ins))[m
[31m-	sigCost, err := math.Mul64(numSigs, secp256k1fx.CostPerSignature)[m
[31m-	if err != nil {[m
[31m-		return 0, err[m
[31m-	}[m
[31m-	return math.Add64(byteCost, sigCost)[m
[32m+[m	[32mreturn 0, fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // Amount of [assetID] burned by this transaction[m
 func (tx *UnsignedExportTx) Burned(assetID ids.ID) (uint64, error) {[m
[31m-	var ([m
[31m-		spent uint64[m
[31m-		input uint64[m
[31m-		err   error[m
[31m-	)[m
[31m-	for _, out := range tx.ExportedOutputs {[m
[31m-		if out.AssetID() == assetID {[m
[31m-			spent, err = math.Add64(spent, out.Output().Amount())[m
[31m-			if err != nil {[m
[31m-				return 0, err[m
[31m-			}[m
[31m-		}[m
[31m-	}[m
[31m-	for _, in := range tx.Ins {[m
[31m-		if in.AssetID == assetID {[m
[31m-			input, err = math.Add64(input, in.Amount)[m
[31m-			if err != nil {[m
[31m-				return 0, err[m
[31m-			}[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	return math.Sub64(input, spent)[m
[32m+[m	[32mreturn 0, fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // SemanticVerify this transaction is valid.[m
[36m@@ -125,105 +68,12 @@[m [mfunc (tx *UnsignedExportTx) SemanticVerify([m
 	baseFee *big.Int,[m
 	rules params.Rules,[m
 ) error {[m
[31m-	if err := tx.Verify(vm.ctx.XChainID, vm.ctx, rules); err != nil {[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	// Check the transaction consumes and produces the right amounts[m
[31m-	fc := avax.NewFlowChecker()[m
[31m-	switch {[m
[31m-	// Apply dynamic fees to export transactions as of Apricot Phase 3[m
[31m-	case rules.IsApricotPhase3:[m
[31m-		cost, err := stx.Cost()[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		txFee, err := calculateDynamicFee(cost, baseFee)[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		fc.Produce(vm.ctx.AVAXAssetID, txFee)[m
[31m-[m
[31m-	// Apply fees to export transactions before Apricot Phase 3[m
[31m-	default:[m
[31m-		fc.Produce(vm.ctx.AVAXAssetID, params.AvalancheAtomicTxFee)[m
[31m-	}[m
[31m-	for _, out := range tx.ExportedOutputs {[m
[31m-		fc.Produce(out.AssetID(), out.Output().Amount())[m
[31m-	}[m
[31m-	for _, in := range tx.Ins {[m
[31m-		fc.Consume(in.AssetID, in.Amount)[m
[31m-	}[m
[31m-[m
[31m-	if err := fc.Verify(); err != nil {[m
[31m-		return fmt.Errorf("export tx flow check failed due to: %w", err)[m
[31m-	}[m
[31m-[m
[31m-	if len(tx.Ins) != len(stx.Creds) {[m
[31m-		return fmt.Errorf("export tx contained mismatched number of inputs/credentials (%d vs. %d)", len(tx.Ins), len(stx.Creds))[m
[31m-	}[m
[31m-[m
[31m-	for i, input := range tx.Ins {[m
[31m-		cred, ok := stx.Creds[i].(*secp256k1fx.Credential)[m
[31m-		if !ok {[m
[31m-			return fmt.Errorf("expected *secp256k1fx.Credential but got %T", cred)[m
[31m-		}[m
[31m-		if err := cred.Verify(); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-[m
[31m-		if len(cred.Sigs) != 1 {[m
[31m-			return fmt.Errorf("expected one signature for EVM Input Credential, but found: %d", len(cred.Sigs))[m
[31m-		}[m
[31m-		pubKeyIntf, err := vm.secpFactory.RecoverPublicKey(tx.UnsignedBytes(), cred.Sigs[0][:])[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		pubKey, ok := pubKeyIntf.(*crypto.PublicKeySECP256K1R)[m
[31m-		if !ok {[m
[31m-			// This should never happen[m
[31m-			return fmt.Errorf("expected *crypto.PublicKeySECP256K1R but got %T", pubKeyIntf)[m
[31m-		}[m
[31m-		if input.Address != PublicKeyToEthAddress(pubKey) {[m
[31m-			return errPublicKeySignatureMismatch[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	return nil[m
[32m+[m	[32mreturn fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // Accept this transaction.[m
 func (tx *UnsignedExportTx) Accept(ctx *snow.Context, batch database.Batch) error {[m
[31m-	txID := tx.ID()[m
[31m-[m
[31m-	elems := make([]*atomic.Element, len(tx.ExportedOutputs))[m
[31m-	for i, out := range tx.ExportedOutputs {[m
[31m-		utxo := &avax.UTXO{[m
[31m-			UTXOID: avax.UTXOID{[m
[31m-				TxID:        txID,[m
[31m-				OutputIndex: uint32(i),[m
[31m-			},[m
[31m-			Asset: avax.Asset{ID: out.AssetID()},[m
[31m-			Out:   out.Out,[m
[31m-		}[m
[31m-[m
[31m-		utxoBytes, err := Codec.Marshal(codecVersion, utxo)[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		utxoID := utxo.InputID()[m
[31m-		elem := &atomic.Element{[m
[31m-			Key:   utxoID[:],[m
[31m-			Value: utxoBytes,[m
[31m-		}[m
[31m-		if out, ok := utxo.Out.(avax.Addressable); ok {[m
[31m-			elem.Traits = out.Addresses()[m
[31m-		}[m
[31m-[m
[31m-		elems[i] = elem[m
[31m-	}[m
[31m-[m
[31m-	return ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{tx.DestinationChain: {PutRequests: elems}}, batch)[m
[32m+[m	[32mreturn fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // newExportTx returns a new ExportTx[m
[36m@@ -235,121 +85,10 @@[m [mfunc (vm *VM) newExportTx([m
 	baseFee *big.Int, // fee to use post-AP3[m
 	keys []*crypto.PrivateKeySECP256K1R, // Pay the fee and provide the tokens[m
 ) (*Tx, error) {[m
[31m-	if vm.ctx.XChainID != chainID {[m
[31m-		return nil, errWrongChainID[m
[31m-	}[m
[31m-[m
[31m-	outs := []*avax.TransferableOutput{{ // Exported to X-Chain[m
[31m-		Asset: avax.Asset{ID: assetID},[m
[31m-		Out: &secp256k1fx.TransferOutput{[m
[31m-			Amt: amount,[m
[31m-			OutputOwners: secp256k1fx.OutputOwners{[m
[31m-				Locktime:  0,[m
[31m-				Threshold: 1,[m
[31m-				Addrs:     []ids.ShortID{to},[m
[31m-			},[m
[31m-		},[m
[31m-	}}[m
[31m-[m
[31m-	var ([m
[31m-		avaxNeeded           uint64 = 0[m
[31m-		ins, avaxIns         []EVMInput[m
[31m-		signers, avaxSigners [][]*crypto.PrivateKeySECP256K1R[m
[31m-		err                  error[m
[31m-	)[m
[31m-[m
[31m-	// consume non-AVAX[m
[31m-	if assetID != vm.ctx.AVAXAssetID {[m
[31m-		ins, signers, err = vm.GetSpendableFunds(keys, assetID, amount)[m
[31m-		if err != nil {[m
[31m-			return nil, fmt.Errorf("couldn't generate tx inputs/signers: %w", err)[m
[31m-		}[m
[31m-	} else {[m
[31m-		avaxNeeded = amount[m
[31m-	}[m
[31m-[m
[31m-	rules := vm.currentRules()[m
[31m-	switch {[m
[31m-	case rules.IsApricotPhase3:[m
[31m-		utx := &UnsignedExportTx{[m
[31m-			NetworkID:        vm.ctx.NetworkID,[m
[31m-			BlockchainID:     vm.ctx.ChainID,[m
[31m-			DestinationChain: chainID,[m
[31m-			Ins:              ins,[m
[31m-			ExportedOutputs:  outs,[m
[31m-		}[m
[31m-		tx := &Tx{UnsignedAtomicTx: utx}[m
[31m-		if err := tx.Sign(vm.codec, nil); err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-[m
[31m-		var cost uint64[m
[31m-		cost, err = tx.Cost()[m
[31m-		if err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-[m
[31m-		avaxIns, avaxSigners, err = vm.GetSpendableAVAXWithFee(keys, avaxNeeded, cost, baseFee)[m
[31m-	default:[m
[31m-		var newAvaxNeeded uint64[m
[31m-		newAvaxNeeded, err = math.Add64(avaxNeeded, params.AvalancheAtomicTxFee)[m
[31m-		if err != nil {[m
[31m-			return nil, errOverflowExport[m
[31m-		}[m
[31m-		avaxIns, avaxSigners, err = vm.GetSpendableFunds(keys, vm.ctx.AVAXAssetID, newAvaxNeeded)[m
[31m-	}[m
[31m-	if err != nil {[m
[31m-		return nil, fmt.Errorf("couldn't generate tx inputs/signers: %w", err)[m
[31m-	}[m
[31m-	ins = append(ins, avaxIns...)[m
[31m-	signers = append(signers, avaxSigners...)[m
[31m-[m
[31m-	SortEVMInputsAndSigners(ins, signers)[m
[31m-[m
[31m-	// Create the transaction[m
[31m-	utx := &UnsignedExportTx{[m
[31m-		NetworkID:        vm.ctx.NetworkID,[m
[31m-		BlockchainID:     vm.ctx.ChainID,[m
[31m-		DestinationChain: chainID,[m
[31m-		Ins:              ins,[m
[31m-		ExportedOutputs:  outs,[m
[31m-	}[m
[31m-	tx := &Tx{UnsignedAtomicTx: utx}[m
[31m-	if err := tx.Sign(vm.codec, signers); err != nil {[m
[31m-		return nil, err[m
[31m-	}[m
[31m-	return tx, utx.Verify(vm.ctx.XChainID, vm.ctx, vm.currentRules())[m
[32m+[m	[32mreturn nil, errWrongChainID[m
 }[m
 [m
 // EVMStateTransfer executes the state update from the atomic export transaction[m
 func (tx *UnsignedExportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {[m
[31m-	addrs := map[[20]byte]uint64{}[m
[31m-	for _, from := range tx.Ins {[m
[31m-		if from.AssetID == ctx.AVAXAssetID {[m
[31m-			log.Debug("crosschain C->X", "addr", from.Address, "amount", from.Amount, "assetID", "AVAX")[m
[31m-			// We multiply the input amount by x2cRate to convert AVAX back to the appropriate[m
[31m-			// denomination before export.[m
[31m-			amount := new(big.Int).Mul([m
[31m-				new(big.Int).SetUint64(from.Amount), x2cRate)[m
[31m-			if state.GetBalance(from.Address).Cmp(amount) < 0 {[m
[31m-				return errInsufficientFunds[m
[31m-			}[m
[31m-			state.SubBalance(from.Address, amount)[m
[31m-		} else {[m
[31m-			log.Debug("crosschain C->X", "addr", from.Address, "amount", from.Amount, "assetID", from.AssetID)[m
[31m-			amount := new(big.Int).SetUint64(from.Amount)[m
[31m-			if state.GetBalanceMultiCoin(from.Address, common.Hash(from.AssetID)).Cmp(amount) < 0 {[m
[31m-				return errInsufficientFunds[m
[31m-			}[m
[31m-			state.SubBalanceMultiCoin(from.Address, common.Hash(from.AssetID), amount)[m
[31m-		}[m
[31m-		if state.GetNonce(from.Address) != from.Nonce {[m
[31m-			return errInvalidNonce[m
[31m-		}[m
[31m-		addrs[from.Address] = from.Nonce[m
[31m-	}[m
[31m-	for addr, nonce := range addrs {[m
[31m-		state.SetNonce(addr, nonce+1)[m
[31m-	}[m
[31m-	return nil[m
[32m+[m	[32mreturn errInsufficientFunds[m
 }[m
[1mdiff --git a/plugin/evm/export_tx_test.go b/plugin/evm/export_tx_test.go[m
[1mindex 44d6b6a..03af06a 100644[m
[1m--- a/plugin/evm/export_tx_test.go[m
[1m+++ b/plugin/evm/export_tx_test.go[m
[36m@@ -7,7 +7,7 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
 	"github.com/ava-labs/avalanchego/chains/atomic"[m
 	"github.com/ava-labs/avalanchego/ids"[m
[1mdiff --git a/plugin/evm/gasprice_update.go b/plugin/evm/gasprice_update.go[m
[1mindex b96fa69..2c98fb2 100644[m
[1m--- a/plugin/evm/gasprice_update.go[m
[1m+++ b/plugin/evm/gasprice_update.go[m
[36m@@ -8,7 +8,7 @@[m [mimport ([m
 	"sync"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type gasPriceUpdater struct {[m
[1mdiff --git a/plugin/evm/gasprice_update_test.go b/plugin/evm/gasprice_update_test.go[m
[1mindex 179dbb8..4ab77ad 100644[m
[1m--- a/plugin/evm/gasprice_update_test.go[m
[1m+++ b/plugin/evm/gasprice_update_test.go[m
[36m@@ -9,7 +9,7 @@[m [mimport ([m
 	"testing"[m
 	"time"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 type mockGasPriceSetter struct {[m
[1mdiff --git a/plugin/evm/import_tx.go b/plugin/evm/import_tx.go[m
[1mindex 0349820..8013ece 100644[m
[1m--- a/plugin/evm/import_tx.go[m
[1m+++ b/plugin/evm/import_tx.go[m
[36m@@ -1,5 +1,12 @@[m
[32m+[m[32m// (c) 2021, Flare Networks Limited. All rights reserved.[m
[32m+[m[32m//[m
[32m+[m[32m// This file is a derived work, based on the avalanchego library whose original[m
[32m+[m[32m// notice appears below. It is distributed under a license compatible with the[m
[32m+[m[32m// licensing terms of the original code from which it is derived.[m
[32m+[m[32m// Please see the file LICENSE_AVALABS for licensing terms of the original work.[m
[32m+[m[32m// Please see the file LICENSE for licensing terms.[m
[32m+[m[32m//[m
 // (c) 2019-2020, Ava Labs, Inc. All rights reserved.[m
[31m-// See the file LICENSE for licensing terms.[m
 [m
 package evm[m
 [m
[36m@@ -7,19 +14,15 @@[m [mimport ([m
 	"fmt"[m
 	"math/big"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
[31m-	"github.com/ava-labs/avalanchego/chains/atomic"[m
 	"github.com/ava-labs/avalanchego/database"[m
 	"github.com/ava-labs/avalanchego/ids"[m
 	"github.com/ava-labs/avalanchego/snow"[m
 	"github.com/ava-labs/avalanchego/utils/crypto"[m
[31m-	"github.com/ava-labs/avalanchego/utils/math"[m
 	"github.com/ava-labs/avalanchego/vms/components/avax"[m
[31m-	"github.com/ava-labs/avalanchego/vms/secp256k1fx"[m
 	"github.com/ethereum/go-ethereum/common"[m
[31m-	"github.com/ethereum/go-ethereum/log"[m
 )[m
 [m
 // UnsignedImportTx is an unsigned ImportTx[m
[36m@@ -52,89 +55,16 @@[m [mfunc (tx *UnsignedImportTx) Verify([m
 	ctx *snow.Context,[m
 	rules params.Rules,[m
 ) error {[m
[31m-	switch {[m
[31m-	case tx == nil:[m
[31m-		return errNilTx[m
[31m-	case tx.SourceChain != avmID:[m
[31m-		return errWrongChainID[m
[31m-	case len(tx.ImportedInputs) == 0:[m
[31m-		return errNoImportInputs[m
[31m-	case tx.NetworkID != ctx.NetworkID:[m
[31m-		return errWrongNetworkID[m
[31m-	case ctx.ChainID != tx.BlockchainID:[m
[31m-		return errWrongBlockchainID[m
[31m-	case rules.IsApricotPhase3 && len(tx.Outs) == 0:[m
[31m-		return errNoEVMOutputs[m
[31m-	}[m
[31m-[m
[31m-	for _, out := range tx.Outs {[m
[31m-		if err := out.Verify(); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	for _, in := range tx.ImportedInputs {[m
[31m-		if err := in.Verify(); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-	}[m
[31m-	if !avax.IsSortedAndUniqueTransferableInputs(tx.ImportedInputs) {[m
[31m-		return errInputsNotSortedUnique[m
[31m-	}[m
[31m-[m
[31m-	if rules.IsApricotPhase2 {[m
[31m-		if !IsSortedAndUniqueEVMOutputs(tx.Outs) {[m
[31m-			return errOutputsNotSortedUnique[m
[31m-		}[m
[31m-	} else if rules.IsApricotPhase1 {[m
[31m-		if !IsSortedEVMOutputs(tx.Outs) {[m
[31m-			return errOutputsNotSorted[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	return nil[m
[32m+[m	[32mreturn errWrongChainID[m
 }[m
 [m
 func (tx *UnsignedImportTx) Cost() (uint64, error) {[m
[31m-	cost := calcBytesCost(len(tx.UnsignedBytes()))[m
[31m-	for _, in := range tx.ImportedInputs {[m
[31m-		inCost, err := in.In.Cost()[m
[31m-		if err != nil {[m
[31m-			return 0, err[m
[31m-		}[m
[31m-		cost, err = math.Add64(cost, inCost)[m
[31m-		if err != nil {[m
[31m-			return 0, err[m
[31m-		}[m
[31m-	}[m
[31m-	return cost, nil[m
[32m+[m	[32mreturn 0, fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // Amount of [assetID] burned by this transaction[m
 func (tx *UnsignedImportTx) Burned(assetID ids.ID) (uint64, error) {[m
[31m-	var ([m
[31m-		spent uint64[m
[31m-		input uint64[m
[31m-		err   error[m
[31m-	)[m
[31m-	for _, out := range tx.Outs {[m
[31m-		if out.AssetID == assetID {[m
[31m-			spent, err = math.Add64(spent, out.Amount)[m
[31m-			if err != nil {[m
[31m-				return 0, err[m
[31m-			}[m
[31m-		}[m
[31m-	}[m
[31m-	for _, in := range tx.ImportedInputs {[m
[31m-		if in.AssetID() == assetID {[m
[31m-			input, err = math.Add64(input, in.Input().Amount())[m
[31m-			if err != nil {[m
[31m-				return 0, err[m
[31m-			}[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	return math.Sub64(input, spent)[m
[32m+[m	[32mreturn 0, fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // SemanticVerify this transaction is valid.[m
[36m@@ -145,82 +75,7 @@[m [mfunc (tx *UnsignedImportTx) SemanticVerify([m
 	baseFee *big.Int,[m
 	rules params.Rules,[m
 ) error {[m
[31m-	if err := tx.Verify(vm.ctx.XChainID, vm.ctx, rules); err != nil {[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	// Check the transaction consumes and produces the right amounts[m
[31m-	fc := avax.NewFlowChecker()[m
[31m-	switch {[m
[31m-	// Apply dynamic fees to import transactions as of Apricot Phase 3[m
[31m-	case rules.IsApricotPhase3:[m
[31m-		cost, err := stx.Cost()[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		txFee, err := calculateDynamicFee(cost, baseFee)[m
[31m-		if err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-		fc.Produce(vm.ctx.AVAXAssetID, txFee)[m
[31m-[m
[31m-	// Apply fees to import transactions as of Apricot Phase 2[m
[31m-	case rules.IsApricotPhase2:[m
[31m-		fc.Produce(vm.ctx.AVAXAssetID, params.AvalancheAtomicTxFee)[m
[31m-	}[m
[31m-	for _, out := range tx.Outs {[m
[31m-		fc.Produce(out.AssetID, out.Amount)[m
[31m-	}[m
[31m-	for _, in := range tx.ImportedInputs {[m
[31m-		fc.Consume(in.AssetID(), in.Input().Amount())[m
[31m-	}[m
[31m-[m
[31m-	if err := fc.Verify(); err != nil {[m
[31m-		return fmt.Errorf("import tx flow check failed due to: %w", err)[m
[31m-	}[m
[31m-[m
[31m-	if len(stx.Creds) != len(tx.ImportedInputs) {[m
[31m-		return fmt.Errorf("export tx contained mismatched number of inputs/credentials (%d vs. %d)", len(tx.ImportedInputs), len(stx.Creds))[m
[31m-	}[m
[31m-[m
[31m-	if !vm.ctx.IsBootstrapped() {[m
[31m-		// Allow for force committing during bootstrapping[m
[31m-		return nil[m
[31m-	}[m
[31m-[m
[31m-	utxoIDs := make([][]byte, len(tx.ImportedInputs))[m
[31m-	for i, in := range tx.ImportedInputs {[m
[31m-		inputID := in.UTXOID.InputID()[m
[31m-		utxoIDs[i] = inputID[:][m
[31m-	}[m
[31m-	// allUTXOBytes is guaranteed to be the same length as utxoIDs[m
[31m-	allUTXOBytes, err := vm.ctx.SharedMemory.Get(tx.SourceChain, utxoIDs)[m
[31m-	if err != nil {[m
[31m-		return fmt.Errorf("failed to fetch import UTXOs from %s with %w", tx.SourceChain, err)[m
[31m-	}[m
[31m-[m
[31m-	for i, in := range tx.ImportedInputs {[m
[31m-		utxoBytes := allUTXOBytes[i][m
[31m-[m
[31m-		utxo := &avax.UTXO{}[m
[31m-		if _, err := vm.codec.Unmarshal(utxoBytes, utxo); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-[m
[31m-		cred := stx.Creds[i][m
[31m-[m
[31m-		utxoAssetID := utxo.AssetID()[m
[31m-		inAssetID := in.AssetID()[m
[31m-		if utxoAssetID != inAssetID {[m
[31m-			return errAssetIDMismatch[m
[31m-		}[m
[31m-[m
[31m-		if err := vm.fx.VerifyTransfer(tx, in.In, cred, utxo.Out); err != nil {[m
[31m-			return err[m
[31m-		}[m
[31m-	}[m
[31m-[m
[31m-	return vm.conflicts(tx.InputUTXOs(), parent)[m
[32m+[m	[32mreturn fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // Accept this transaction and spend imported inputs[m
[36m@@ -229,12 +84,7 @@[m [mfunc (tx *UnsignedImportTx) SemanticVerify([m
 // only to have the transaction not be Accepted. This would be inconsistent.[m
 // Recall that imported UTXOs are not kept in a versionDB.[m
 func (tx *UnsignedImportTx) Accept(ctx *snow.Context, batch database.Batch) error {[m
[31m-	utxoIDs := make([][]byte, len(tx.ImportedInputs))[m
[31m-	for i, in := range tx.ImportedInputs {[m
[31m-		inputID := in.InputID()[m
[31m-		utxoIDs[i] = inputID[:][m
[31m-	}[m
[31m-	return ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{tx.SourceChain: {RemoveRequests: utxoIDs}}, batch)[m
[32m+[m	[32mreturn fmt.Errorf("exportTx transactions disabled")[m
 }[m
 [m
 // newImportTx returns a new ImportTx[m
[36m@@ -244,160 +94,11 @@[m [mfunc (vm *VM) newImportTx([m
 	baseFee *big.Int, // fee to use post-AP3[m
 	keys []*crypto.PrivateKeySECP256K1R, // Keys to import the funds[m
 ) (*Tx, error) {[m
[31m-	if vm.ctx.XChainID != chainID {[m
[31m-		return nil, errWrongChainID[m
[31m-	}[m
[31m-[m
[31m-	kc := secp256k1fx.NewKeychain()[m
[31m-	for _, key := range keys {[m
[31m-		kc.Add(key)[m
[31m-	}[m
[31m-[m
[31m-	atomicUTXOs, _, _, err := vm.GetAtomicUTXOs(chainID, kc.Addresses(), ids.ShortEmpty, ids.Empty, -1)[m
[31m-	if err != nil {[m
[31m-		return nil, fmt.Errorf("problem retrieving atomic UTXOs: %w", err)[m
[31m-	}[m
[31m-[m
[31m-	importedInputs := []*avax.TransferableInput{}[m
[31m-	signers := [][]*crypto.PrivateKeySECP256K1R{}[m
[31m-[m
[31m-	importedAmount := make(map[ids.ID]uint64)[m
[31m-	now := vm.clock.Unix()[m
[31m-	for _, utxo := range atomicUTXOs {[m
[31m-		inputIntf, utxoSigners, err := kc.Spend(utxo.Out, now)[m
[31m-		if err != nil {[m
[31m-			continue[m
[31m-		}[m
[31m-		input, ok := inputIntf.(avax.TransferableIn)[m
[31m-		if !ok {[m
[31m-			continue[m
[31m-		}[m
[31m-		aid := utxo.AssetID()[m
[31m-		importedAmount[aid], err = math.Add64(importedAmount[aid], input.Amount())[m
[31m-		if err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-		importedInputs = append(importedInputs, &avax.TransferableInput{[m
[31m-			UTXOID: utxo.UTXOID,[m
[31m-			Asset:  utxo.Asset,[m
[31m-			In:     input,[m
[31m-		})[m
[31m-		signers = append(signers, utxoSigners)[m
[31m-	}[m
[31m-	avax.SortTransferableInputsWithSigners(importedInputs, signers)[m
[31m-	importedAVAXAmount := importedAmount[vm.ctx.AVAXAssetID][m
[31m-[m
[31m-	outs := make([]EVMOutput, 0, len(importedAmount))[m
[31m-	// This will create unique outputs (in the context of sorting)[m
[31m-	// since each output will have a unique assetID[m
[31m-	for assetID, amount := range importedAmount {[m
[31m-		// Skip the AVAX amount since it is included separately to account for[m
[31m-		// the fee[m
[31m-		if assetID == vm.ctx.AVAXAssetID || amount == 0 {[m
[31m-			continue[m
[31m-		}[m
[31m-		outs = append(outs, EVMOutput{[m
[31m-			Address: to,[m
[31m-			Amount:  amount,[m
[31m-			AssetID: assetID,[m
[31m-		})[m
[31m-	}[m
[31m-[m
[31m-	rules := vm.currentRules()[m
[31m-[m
[31m-	var ([m
[31m-		txFeeWithoutChange uint64[m
[31m-		txFeeWithChange    uint64[m
[31m-	)[m
[31m-	switch {[m
[31m-	case rules.IsApricotPhase3:[m
[31m-		if baseFee == nil {[m
[31m-			return nil, errNilBaseFeeApricotPhase3[m
[31m-		}[m
[31m-		utx := &UnsignedImportTx{[m
[31m-			NetworkID:      vm.ctx.NetworkID,[m
[31m-			BlockchainID:   vm.ctx.ChainID,[m
[31m-			Outs:           outs,[m
[31m-			ImportedInputs: importedInputs,[m
[31m-			SourceChain:    chainID,[m
[31m-		}[m
[31m-		tx := &Tx{UnsignedAtomicTx: utx}[m
[31m-		if err := tx.Sign(vm.codec, nil); err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-[m
[31m-		costWithoutChange, err := tx.Cost()[m
[31m-		if err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-		costWithChange := costWithoutChange + EVMOutputGas[m
[31m-[m
[31m-		txFeeWithoutChange, err = calculateDynamicFee(costWithoutChange, baseFee)[m
[31m-		if err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-		txFeeWithChange, err = calculateDynamicFee(costWithChange, baseFee)[m
[31m-		if err != nil {[m
[31m-			return nil, err[m
[31m-		}[m
[31m-	case rules.IsApricotPhase2:[m
[31m-		txFeeWithoutChange = params.AvalancheAtomicTxFee[m
[31m-		txFeeWithChange = params.AvalancheAtomicTxFee[m
[31m-	}[m
[31m-[m
[31m-	// AVAX output[m
[31m-	if importedAVAXAmount < txFeeWithoutChange { // imported amount goes toward paying tx fee[m
[31m-		return nil, errInsufficientFundsForFee[m
[31m-	}[m
[31m-[m
[31m-	if importedAVAXAmount > txFeeWithChange {[m
[31m-		outs = append(outs, EVMOutput{[m
[31m-			Address: to,[m
[31m-			Amount:  importedAVAXAmount - txFeeWithChange,[m
[31m-			AssetID: vm.ctx.AVAXAssetID,[m
[31m-		})[m
[31m-	}[m
[31m-[m
[31m-	// If no outputs are produced, return an error.[m
[31m-	// Note: this can happen if there is exactly enough AVAX to pay the[m
[31m-	// transaction fee, but no other funds to be imported.[m
[31m-	if len(outs) == 0 {[m
[31m-		return nil, errNoEVMOutputs[m
[31m-	}[m
[31m-[m
[31m-	SortEVMOutputs(outs)[m
[31m-[m
[31m-	// Create the transaction[m
[31m-	utx := &UnsignedImportTx{[m
[31m-		NetworkID:      vm.ctx.NetworkID,[m
[31m-		BlockchainID:   vm.ctx.ChainID,[m
[31m-		Outs:           outs,[m
[31m-		ImportedInputs: importedInputs,[m
[31m-		SourceChain:    chainID,[m
[31m-	}[m
[31m-	tx := &Tx{UnsignedAtomicTx: utx}[m
[31m-	if err := tx.Sign(vm.codec, signers); err != nil {[m
[31m-		return nil, err[m
[31m-	}[m
[31m-	return tx, utx.Verify(vm.ctx.XChainID, vm.ctx, vm.currentRules())[m
[32m+[m	[32mreturn nil, errWrongChainID[m
 }[m
 [m
 // EVMStateTransfer performs the state transfer to increase the balances of[m
 // accounts accordingly with the imported EVMOutputs[m
 func (tx *UnsignedImportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {[m
[31m-	for _, to := range tx.Outs {[m
[31m-		if to.AssetID == ctx.AVAXAssetID {[m
[31m-			log.Debug("crosschain X->C", "addr", to.Address, "amount", to.Amount, "assetID", "AVAX")[m
[31m-			// If the asset is AVAX, convert the input amount in nAVAX to gWei by[m
[31m-			// multiplying by the x2c rate.[m
[31m-			amount := new(big.Int).Mul([m
[31m-				new(big.Int).SetUint64(to.Amount), x2cRate)[m
[31m-			state.AddBalance(to.Address, amount)[m
[31m-		} else {[m
[31m-			log.Debug("crosschain X->C", "addr", to.Address, "amount", to.Amount, "assetID", to.AssetID)[m
[31m-			amount := new(big.Int).SetUint64(to.Amount)[m
[31m-			state.AddBalanceMultiCoin(to.Address, common.Hash(to.AssetID), amount)[m
[31m-		}[m
[31m-	}[m
[31m-	return nil[m
[32m+[m	[32mreturn errInsufficientFunds[m
 }[m
[1mdiff --git a/plugin/evm/import_tx_test.go b/plugin/evm/import_tx_test.go[m
[1mindex c4c7df8..6551bb3 100644[m
[1m--- a/plugin/evm/import_tx_test.go[m
[1m+++ b/plugin/evm/import_tx_test.go[m
[36m@@ -7,7 +7,7 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
 	"github.com/ava-labs/avalanchego/chains/atomic"[m
 	"github.com/ava-labs/avalanchego/ids"[m
[1mdiff --git a/plugin/evm/service.go b/plugin/evm/service.go[m
[1mindex a6f56f2..ee079bb 100644[m
[1m--- a/plugin/evm/service.go[m
[1m+++ b/plugin/evm/service.go[m
[36m@@ -17,11 +17,11 @@[m [mimport ([m
 	"github.com/ava-labs/avalanchego/utils/crypto"[m
 	"github.com/ava-labs/avalanchego/utils/formatting"[m
 	"github.com/ava-labs/avalanchego/utils/json"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	ethcrypto "github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/log"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // test constants[m
[1mdiff --git a/plugin/evm/static_service.go b/plugin/evm/static_service.go[m
[1mindex 9c59225..f19fe78 100644[m
[1m--- a/plugin/evm/static_service.go[m
[1m+++ b/plugin/evm/static_service.go[m
[36m@@ -8,7 +8,7 @@[m [mimport ([m
 	"encoding/json"[m
 [m
 	"github.com/ava-labs/avalanchego/utils/formatting"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
 )[m
 [m
 // StaticService defines the static API services exposed by the evm[m
[1mdiff --git a/plugin/evm/tx.go b/plugin/evm/tx.go[m
[1mindex ac4f798..b15365e 100644[m
[1m--- a/plugin/evm/tx.go[m
[1m+++ b/plugin/evm/tx.go[m
[36m@@ -12,8 +12,8 @@[m [mimport ([m
 [m
 	"github.com/ethereum/go-ethereum/common"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 [m
 	"github.com/ava-labs/avalanchego/codec"[m
 	"github.com/ava-labs/avalanchego/database"[m
[1mdiff --git a/plugin/evm/tx_test.go b/plugin/evm/tx_test.go[m
[1mindex d1d4947..ca73402 100644[m
[1m--- a/plugin/evm/tx_test.go[m
[1m+++ b/plugin/evm/tx_test.go[m
[36m@@ -7,7 +7,7 @@[m [mimport ([m
 	"math/big"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func TestCalculateDynamicFee(t *testing.T) {[m
[1mdiff --git a/plugin/evm/vm.go b/plugin/evm/vm.go[m
[1mindex 5850762..198f899 100644[m
[1m--- a/plugin/evm/vm.go[m
[1m+++ b/plugin/evm/vm.go[m
[36m@@ -10,25 +10,26 @@[m [mimport ([m
 	"errors"[m
 	"fmt"[m
 	"math/big"[m
[32m+[m	[32m"os"[m
 	"path/filepath"[m
 	"strings"[m
 	"sync"[m
 	"time"[m
 [m
 	"github.com/ava-labs/avalanchego/database/versiondb"[m
[31m-	coreth "github.com/ava-labs/coreth/chain"[m
[31m-	"github.com/ava-labs/coreth/consensus/dummy"[m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/eth/ethconfig"[m
[31m-	"github.com/ava-labs/coreth/node"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
[32m+[m	[32mcoreth "gitlab.com/flarenetwork/coreth/chain"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/consensus/dummy"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth/ethconfig"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/node"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/log"[m
 	"github.com/ethereum/go-ethereum/rlp"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 [m
 	avalancheRPC "github.com/gorilla/rpc/v2"[m
 [m
[36m@@ -279,6 +280,26 @@[m [mfunc (vm *VM) Initialize([m
 			return fmt.Errorf("failed to unmarshal config %s: %w", string(configBytes), err)[m
 		}[m
 	}[m
[32m+[m	[32mvm.config.EthAPIEnabled = false[m
[32m+[m	[32mvm.config.NetAPIEnabled = false[m
[32m+[m	[32mvm.config.Web3APIEnabled = false[m
[32m+[m	[32mvm.config.DebugAPIEnabled = false[m
[32m+[m	[32mvm.config.MaxBlocksPerRequest = 1[m
[32m+[m	[32mweb3API := os.Getenv("WEB3_API")[m
[32m+[m	[32mif web3API == "enabled" {[m
[32m+[m		[32mvm.config.EthAPIEnabled = true[m
[32m+[m		[32mvm.config.NetAPIEnabled = true[m
[32m+[m		[32mvm.config.Web3APIEnabled = true[m
[32m+[m	[32m} else if web3API == "debug" {[m
[32m+[m		[32mvm.config.EthAPIEnabled = true[m
[32m+[m		[32mvm.config.NetAPIEnabled = true[m
[32m+[m		[32mvm.config.Web3APIEnabled = true[m
[32m+[m		[32mvm.config.DebugAPIEnabled = true[m
[32m+[m		[32mvm.config.TxPoolAPIEnabled = true[m
[32m+[m		[32mvm.config.Pruning = false[m
[32m+[m		[32mvm.config.MaxBlocksPerRequest = 0[m
[32m+[m	[32m}[m
[32m+[m
 	if b, err := json.Marshal(vm.config); err == nil {[m
 		log.Info("Initializing Coreth VM", "Version", Version, "Config", string(b))[m
 	} else {[m
[36m@@ -337,7 +358,7 @@[m [mfunc (vm *VM) Initialize([m
 	ethConfig.SnapshotVerify = vm.config.SnapshotVerify[m
 [m
 	vm.chainConfig = g.Config[m
[31m-	vm.networkID = ethConfig.NetworkId[m
[32m+[m	[32mvm.networkID = g.Config.ChainID.Uint64()[m
 	vm.secpFactory = crypto.FactorySECP256K1R{Cache: cache.LRU{Size: secpFactoryCacheSize}}[m
 [m
 	nodecfg := node.Config{[m
[1mdiff --git a/plugin/evm/vm_test.go b/plugin/evm/vm_test.go[m
[1mindex 80a6953..f4ec7e0 100644[m
[1m--- a/plugin/evm/vm_test.go[m
[1m+++ b/plugin/evm/vm_test.go[m
[36m@@ -40,13 +40,13 @@[m [mimport ([m
 [m
 	engCommon "github.com/ava-labs/avalanchego/snow/engine/common"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/eth"[m
[31m-	"github.com/ava-labs/coreth/params"[m
[31m-	"github.com/ava-labs/coreth/rpc"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/eth"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/rpc"[m
 [m
[31m-	accountKeystore "github.com/ava-labs/coreth/accounts/keystore"[m
[32m+[m	[32maccountKeystore "gitlab.com/flarenetwork/coreth/accounts/keystore"[m
 )[m
 [m
 var ([m
[1mdiff --git a/plugin/main.go b/plugin/main.go[m
[1mindex 9637bea..af1f80a 100644[m
[1m--- a/plugin/main.go[m
[1m+++ b/plugin/main.go[m
[36m@@ -12,7 +12,7 @@[m [mimport ([m
 [m
 	"github.com/ava-labs/avalanchego/vms/rpcchainvm"[m
 [m
[31m-	"github.com/ava-labs/coreth/plugin/evm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/plugin/evm"[m
 )[m
 [m
 func main() {[m
[1mdiff --git a/signer/core/apitypes/types.go b/signer/core/apitypes/types.go[m
[1mindex cab7f9c..32a9ec4 100644[m
[1m--- a/signer/core/apitypes/types.go[m
[1m+++ b/signer/core/apitypes/types.go[m
[36m@@ -32,9 +32,9 @@[m [mimport ([m
 	"math/big"[m
 	"strings"[m
 [m
[31m-	"github.com/ava-labs/coreth/core/types"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
 )[m
 [m
 type ValidationInfo struct {[m
[1mdiff --git a/tests/init.go b/tests/init.go[m
[1mindex c924236..1fe1359 100644[m
[1m--- a/tests/init.go[m
[1m+++ b/tests/init.go[m
[36m@@ -31,7 +31,7 @@[m [mimport ([m
 	"math/big"[m
 	"sort"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // Forks table defines supported forks and their chain config.[m
[1mdiff --git a/tests/init_test.go b/tests/init_test.go[m
[1mindex 7f30fde..b72be17 100644[m
[1m--- a/tests/init_test.go[m
[1m+++ b/tests/init_test.go[m
[36m@@ -40,7 +40,7 @@[m [mimport ([m
 	"strings"[m
 	"testing"[m
 [m
[31m-	"github.com/ava-labs/coreth/params"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 func readJSON(reader io.Reader, value interface{}) error {[m
[1mdiff --git a/tests/state_test_util.go b/tests/state_test_util.go[m
[1mindex 7df6737..43e3990 100644[m
[1m--- a/tests/state_test_util.go[m
[1m+++ b/tests/state_test_util.go[m
[36m@@ -34,17 +34,17 @@[m [mimport ([m
 	"strconv"[m
 	"strings"[m
 [m
[31m-	"github.com/ava-labs/coreth/core"[m
[31m-	"github.com/ava-labs/coreth/core/state"[m
[31m-	"github.com/ava-labs/coreth/core/state/snapshot"[m
[31m-	"github.com/ava-labs/coreth/core/types"[m
[31m-	"github.com/ava-labs/coreth/core/vm"[m
[31m-	"github.com/ava-labs/coreth/params"[m
 	"github.com/ethereum/go-ethereum/common"[m
 	"github.com/ethereum/go-ethereum/common/hexutil"[m
 	"github.com/ethereum/go-ethereum/common/math"[m
 	"github.com/ethereum/go-ethereum/crypto"[m
 	"github.com/ethereum/go-ethereum/ethdb"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/state/snapshot"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/types"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/core/vm"[m
[32m+[m	[32m"gitlab.com/flarenetwork/coreth/params"[m
 )[m
 [m
 // StateTest checks transaction processing without block context.[m
