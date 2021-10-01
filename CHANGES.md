<b>diff --git a/CHANGES.md b/CHANGES.md</b>
<b>new file mode 100644</b>
<b>index 0000000..e69de29</b>
<b>diff --git a/accounts/accounts.go b/accounts/accounts.go</b>
<b>index dd7df0b..00420ca 100644</b>
<b>--- a/accounts/accounts.go</b>
<b>+++ b/accounts/accounts.go</b>
<span style="color:#0AA">@@ -31,10 +31,10 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
 	"golang.org/x/crypto/sha3"
 )
 
<b>diff --git a/accounts/external/backend.go b/accounts/external/backend.go</b>
<b>index 558700f..8302c0c 100644</b>
<b>--- a/accounts/external/backend.go</b>
<b>+++ b/accounts/external/backend.go</b>
<span style="color:#0AA">@@ -31,15 +31,15 @@</span> import (
 	"math/big"
 	"sync"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/signer/core/apitypes"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/signer/core/apitypes"</span>
 )
 
 type ExternalBackend struct {
<b>diff --git a/accounts/keystore/account_cache.go b/accounts/keystore/account_cache.go</b>
<b>index 4c35aa7..fbe0a89 100644</b>
<b>--- a/accounts/keystore/account_cache.go</b>
<b>+++ b/accounts/keystore/account_cache.go</b>
<span style="color:#0AA">@@ -37,10 +37,10 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	mapset "github.com/deckarep/golang-set"
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 )
 
 // Minimum amount of time between cache reloads. This limit applies if the platform does
<b>diff --git a/accounts/keystore/account_cache_test.go b/accounts/keystore/account_cache_test.go</b>
<b>index 3d341d5..81f80d6 100644</b>
<b>--- a/accounts/keystore/account_cache_test.go</b>
<b>+++ b/accounts/keystore/account_cache_test.go</b>
<span style="color:#0AA">@@ -37,10 +37,10 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/cespare/cp"
 	"github.com/davecgh/go-spew/spew"
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 )
 
 var (
<b>diff --git a/accounts/keystore/key.go b/accounts/keystore/key.go</b>
<b>index 71402d3..ad6cc6c 100644</b>
<b>--- a/accounts/keystore/key.go</b>
<b>+++ b/accounts/keystore/key.go</b>
<span style="color:#0AA">@@ -39,10 +39,10 @@</span> import (
 	"strings"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/google/uuid"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 )
 
 const (
<b>diff --git a/accounts/keystore/keystore.go b/accounts/keystore/keystore.go</b>
<b>index ff82ef8..e1c177d 100644</b>
<b>--- a/accounts/keystore/keystore.go</b>
<b>+++ b/accounts/keystore/keystore.go</b>
<span style="color:#0AA">@@ -42,11 +42,11 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 var (
<b>diff --git a/accounts/keystore/keystore_test.go b/accounts/keystore/keystore_test.go</b>
<b>index 651ab70..70fbeb7 100644</b>
<b>--- a/accounts/keystore/keystore_test.go</b>
<b>+++ b/accounts/keystore/keystore_test.go</b>
<span style="color:#0AA">@@ -38,10 +38,10 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 )
 
 var testSigData = make([]byte, 32)
<b>diff --git a/accounts/keystore/passphrase.go b/accounts/keystore/passphrase.go</b>
<b>index 5da41f6..73f26d2 100644</b>
<b>--- a/accounts/keystore/passphrase.go</b>
<b>+++ b/accounts/keystore/passphrase.go</b>
<span style="color:#0AA">@@ -48,11 +48,11 @@</span> import (
 	"os"
 	"path/filepath"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/google/uuid"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 	"golang.org/x/crypto/pbkdf2"
 	"golang.org/x/crypto/scrypt"
 )
<b>diff --git a/accounts/keystore/presale.go b/accounts/keystore/presale.go</b>
<b>index 1dfbd9c..64f020d 100644</b>
<b>--- a/accounts/keystore/presale.go</b>
<b>+++ b/accounts/keystore/presale.go</b>
<span style="color:#0AA">@@ -35,9 +35,9 @@</span> import (
 	"errors"
 	"fmt"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/google/uuid"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 	"golang.org/x/crypto/pbkdf2"
 )
 
<b>diff --git a/accounts/keystore/wallet.go b/accounts/keystore/wallet.go</b>
<b>index 78eac1d..e603ac5 100644</b>
<b>--- a/accounts/keystore/wallet.go</b>
<b>+++ b/accounts/keystore/wallet.go</b>
<span style="color:#0AA">@@ -29,11 +29,11 @@</span> package keystore
 import (
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // keystoreWallet implements the accounts.Wallet interface for the original
<b>diff --git a/accounts/scwallet/hub.go b/accounts/scwallet/hub.go</b>
<b>index 7a630fa..c05ba6f 100644</b>
<b>--- a/accounts/scwallet/hub.go</b>
<b>+++ b/accounts/scwallet/hub.go</b>
<span style="color:#0AA">@@ -51,11 +51,11 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
 	pcsc "github.com/gballet/go-libpcsclite"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
 )
 
 // Scheme is the URI prefix for smartcard wallets.
<b>diff --git a/accounts/scwallet/wallet.go b/accounts/scwallet/wallet.go</b>
<b>index fcecc10..b912e21 100644</b>
<b>--- a/accounts/scwallet/wallet.go</b>
<b>+++ b/accounts/scwallet/wallet.go</b>
<span style="color:#0AA">@@ -43,14 +43,14 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/log"
 	pcsc "github.com/gballet/go-libpcsclite"
 	"github.com/status-im/keycard-go/derivationpath"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
 )
 
 // ErrPairingPasswordNeeded is returned if opening the smart card requires pairing with a pairing
<b>diff --git a/chain/chain_test.go b/chain/chain_test.go</b>
<b>index 2d8c7e9..9057d9d 100644</b>
<b>--- a/chain/chain_test.go</b>
<b>+++ b/chain/chain_test.go</b>
<span style="color:#0AA">@@ -9,20 +9,20 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/ethconfig"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/ethconfig"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type testChain struct {
<b>diff --git a/chain/coreth.go b/chain/coreth.go</b>
<b>index 95018be..bc107de 100644</b>
<b>--- a/chain/coreth.go</b>
<b>+++ b/chain/coreth.go</b>
<span style="color:#0AA">@@ -7,15 +7,15 @@</span> import (
 	"fmt"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 var (
<b>diff --git a/chain/counter_test.go b/chain/counter_test.go</b>
<b>index 7209ac5..4fd88fe 100644</b>
<b>--- a/chain/counter_test.go</b>
<b>+++ b/chain/counter_test.go</b>
<span style="color:#0AA">@@ -15,8 +15,8 @@</span> import (
 
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 
 	"github.com/ethereum/go-ethereum/log"
 )
<b>diff --git a/chain/multicoin_test.go b/chain/multicoin_test.go</b>
<b>index a2ce4cb..f40a678 100644</b>
<b>--- a/chain/multicoin_test.go</b>
<b>+++ b/chain/multicoin_test.go</b>
<span style="color:#0AA">@@ -28,19 +28,19 @@</span> import (
 	"strings"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/ethconfig"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
 	"github.com/ethereum/go-ethereum/accounts/abi"
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/ethconfig"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
 )
 
 // TestMulticoin tests multicoin low-level state management and regular
<b>diff --git a/chain/payment_test.go b/chain/payment_test.go</b>
<b>index 5d9da4d..f17e260 100644</b>
<b>--- a/chain/payment_test.go</b>
<b>+++ b/chain/payment_test.go</b>
<span style="color:#0AA">@@ -7,8 +7,8 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // TestPayment tests basic payment (balance, not multi-coin)
<b>diff --git a/chain/subscribe_accepted_heads_test.go b/chain/subscribe_accepted_heads_test.go</b>
<b>index 0a94bfd..ab7315e 100644</b>
<b>--- a/chain/subscribe_accepted_heads_test.go</b>
<b>+++ b/chain/subscribe_accepted_heads_test.go</b>
<span style="color:#0AA">@@ -4,10 +4,10 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 func TestAcceptedHeadSubscriptions(t *testing.T) {
<b>diff --git a/chain/subscribe_block_logs_test.go b/chain/subscribe_block_logs_test.go</b>
<b>index 26661da..2b2d2ab 100644</b>
<b>--- a/chain/subscribe_block_logs_test.go</b>
<b>+++ b/chain/subscribe_block_logs_test.go</b>
<span style="color:#0AA">@@ -6,10 +6,10 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/filters"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/filters"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 func TestBlockLogsAllowUnfinalized(t *testing.T) {
<b>diff --git a/chain/subscribe_transactions_test.go b/chain/subscribe_transactions_test.go</b>
<b>index aac6db4..9388dae 100644</b>
<b>--- a/chain/subscribe_transactions_test.go</b>
<b>+++ b/chain/subscribe_transactions_test.go</b>
<span style="color:#0AA">@@ -4,10 +4,10 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/filters"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/filters"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 func TestSubscribeTransactions(t *testing.T) {
<b>diff --git a/chain/test_chain.go b/chain/test_chain.go</b>
<b>index e4420de..3d71d82 100644</b>
<b>--- a/chain/test_chain.go</b>
<b>+++ b/chain/test_chain.go</b>
<span style="color:#0AA">@@ -8,17 +8,17 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/ethconfig"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/ethconfig"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/consensus/consensus.go b/consensus/consensus.go</b>
<b>index 6bc13a8..af462ca 100644</b>
<b>--- a/consensus/consensus.go</b>
<b>+++ b/consensus/consensus.go</b>
<span style="color:#0AA">@@ -30,11 +30,11 @@</span> package consensus
 import (
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // ChainHeaderReader defines a small collection of methods needed to access the local
<b>diff --git a/consensus/dummy/consensus.go b/consensus/dummy/consensus.go</b>
<b>index 89af334..e93724d 100644</b>
<b>--- a/consensus/dummy/consensus.go</b>
<b>+++ b/consensus/dummy/consensus.go</b>
<span style="color:#0AA">@@ -10,13 +10,13 @@</span> import (
 	"math/big"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 type OnFinalizeCallbackType = func(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, receipts []*types.Receipt, uncles []*types.Header) error
<b>diff --git a/consensus/dummy/dynamic_fees.go b/consensus/dummy/dynamic_fees.go</b>
<b>index 8050364..9221cec 100644</b>
<b>--- a/consensus/dummy/dynamic_fees.go</b>
<b>+++ b/consensus/dummy/dynamic_fees.go</b>
<span style="color:#0AA">@@ -9,10 +9,10 @@</span> import (
 	"math/big"
 
 	"github.com/ava-labs/avalanchego/utils/wrappers"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/consensus/dummy/dynamic_fees_test.go b/consensus/dummy/dynamic_fees_test.go</b>
<b>index fd45082..bc6813d 100644</b>
<b>--- a/consensus/dummy/dynamic_fees_test.go</b>
<b>+++ b/consensus/dummy/dynamic_fees_test.go</b>
<span style="color:#0AA">@@ -8,10 +8,10 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func testRollup(t *testing.T, longs []uint64, roll int) {
<b>diff --git a/consensus/misc/dao.go b/consensus/misc/dao.go</b>
<b>index a0ab402..486f97a 100644</b>
<b>--- a/consensus/misc/dao.go</b>
<b>+++ b/consensus/misc/dao.go</b>
<span style="color:#0AA">@@ -31,9 +31,9 @@</span> import (
 	"errors"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/consensus/misc/forks.go b/consensus/misc/forks.go</b>
<b>index 6199e11..63ce4f0 100644</b>
<b>--- a/consensus/misc/forks.go</b>
<b>+++ b/consensus/misc/forks.go</b>
<span style="color:#0AA">@@ -29,9 +29,9 @@</span> package misc
 import (
 	"fmt"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // VerifyForkHashes verifies that blocks conforming to network hard-forks do have
<b>diff --git a/core/block_validator.go b/core/block_validator.go</b>
<b>index f7092ee..6993bf1 100644</b>
<b>--- a/core/block_validator.go</b>
<b>+++ b/core/block_validator.go</b>
<span style="color:#0AA">@@ -29,11 +29,11 @@</span> package core
 import (
 	"fmt"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // BlockValidator is responsible for validating block headers, uncles and
<b>diff --git a/core/blockchain.go b/core/blockchain.go</b>
<b>index 7e03462..5073b6a 100644</b>
<b>--- a/core/blockchain.go</b>
<b>+++ b/core/blockchain.go</b>
<span style="color:#0AA">@@ -36,19 +36,19 @@</span> import (
 	"sync/atomic"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state/snapshot"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/trie"
 	lru "github.com/hashicorp/golang-lru"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state/snapshot"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/core/blockchain_test.go b/core/blockchain_test.go</b>
<b>index 07cf8ec..74981e8 100644</b>
<b>--- a/core/blockchain_test.go</b>
<b>+++ b/core/blockchain_test.go</b>
<span style="color:#0AA">@@ -7,14 +7,14 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestArchiveBlockChain(t *testing.T) {
<b>diff --git a/core/bloom_indexer.go b/core/bloom_indexer.go</b>
<b>index 07e50de..c2fa061 100644</b>
<b>--- a/core/bloom_indexer.go</b>
<b>+++ b/core/bloom_indexer.go</b>
<span style="color:#0AA">@@ -20,12 +20,12 @@</span> import (
 	"context"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/bloombits"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/bitutil"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/bloombits"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 const (
<b>diff --git a/core/bloombits/generator.go b/core/bloombits/generator.go</b>
<b>index c0422ca..23a5b54 100644</b>
<b>--- a/core/bloombits/generator.go</b>
<b>+++ b/core/bloombits/generator.go</b>
<span style="color:#0AA">@@ -29,7 +29,7 @@</span> package bloombits
 import (
 	"errors"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 var (
<b>diff --git a/core/bloombits/generator_test.go b/core/bloombits/generator_test.go</b>
<b>index 27dba47..b233daf 100644</b>
<b>--- a/core/bloombits/generator_test.go</b>
<b>+++ b/core/bloombits/generator_test.go</b>
<span style="color:#0AA">@@ -31,7 +31,7 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // Tests that batched bloom bits are correctly rotated from the input bloom
<b>diff --git a/core/chain_indexer.go b/core/chain_indexer.go</b>
<b>index ee6557f..eaf7881 100644</b>
<b>--- a/core/chain_indexer.go</b>
<b>+++ b/core/chain_indexer.go</b>
<span style="color:#0AA">@@ -34,12 +34,12 @@</span> import (
 	"sync/atomic"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // ChainIndexerBackend defines the methods needed to process chain segments in
<b>diff --git a/core/chain_indexer_test.go b/core/chain_indexer_test.go</b>
<b>index 3edf175..d8b1c43 100644</b>
<b>--- a/core/chain_indexer_test.go</b>
<b>+++ b/core/chain_indexer_test.go</b>
<span style="color:#0AA">@@ -35,9 +35,9 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // Runs multiple tests with randomized parameters.
<b>diff --git a/core/chain_makers.go b/core/chain_makers.go</b>
<b>index 4aca357..58fd71b 100644</b>
<b>--- a/core/chain_makers.go</b>
<b>+++ b/core/chain_makers.go</b>
<span style="color:#0AA">@@ -30,15 +30,15 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/misc"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/misc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // BlockGen creates blocks for testing.
<b>diff --git a/core/chain_makers_test.go b/core/chain_makers_test.go</b>
<b>index b67cac2..34881d4 100644</b>
<b>--- a/core/chain_makers_test.go</b>
<b>+++ b/core/chain_makers_test.go</b>
<span style="color:#0AA">@@ -30,13 +30,13 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func ExampleGenerateChain() {
<b>diff --git a/core/error.go b/core/error.go</b>
<b>index 0265448..4039ed7 100644</b>
<b>--- a/core/error.go</b>
<b>+++ b/core/error.go</b>
<span style="color:#0AA">@@ -29,7 +29,7 @@</span> package core
 import (
 	"errors"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 var (
<b>diff --git a/core/events.go b/core/events.go</b>
<b>index 4898dbc..d57d378 100644</b>
<b>--- a/core/events.go</b>
<b>+++ b/core/events.go</b>
<span style="color:#0AA">@@ -27,8 +27,8 @@</span>
 package core
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // NewTxsEvent is posted when a batch of transactions enter the transaction pool.
<b>diff --git a/core/evm.go b/core/evm.go</b>
<b>index d45f241..2dad5a4 100644</b>
<b>--- a/core/evm.go</b>
<b>+++ b/core/evm.go</b>
<span style="color:#0AA">@@ -29,10 +29,10 @@</span> package core
 import (
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 	//"github.com/ethereum/go-ethereum/log"
 )
 
<b>diff --git a/core/gen_genesis.go b/core/gen_genesis.go</b>
<b>index a4ec8f5..b247996 100644</b>
<b>--- a/core/gen_genesis.go</b>
<b>+++ b/core/gen_genesis.go</b>
<span style="color:#0AA">@@ -7,10 +7,10 @@</span> import (
 	"errors"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var _ = (*genesisSpecMarshaling)(nil)
<b>diff --git a/core/genesis.go b/core/genesis.go</b>
<b>index 185887e..153f51f 100644</b>
<b>--- a/core/genesis.go</b>
<b>+++ b/core/genesis.go</b>
<span style="color:#0AA">@@ -34,16 +34,16 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 //go:generate gencodec -type Genesis -field-override genesisSpecMarshaling -out gen_genesis.go
<b>diff --git a/core/headerchain.go b/core/headerchain.go</b>
<b>index eafc747..b8d75c1 100644</b>
<b>--- a/core/headerchain.go</b>
<b>+++ b/core/headerchain.go</b>
<span style="color:#0AA">@@ -33,13 +33,13 @@</span> import (
 	mrand "math/rand"
 	"sync/atomic"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	lru "github.com/hashicorp/golang-lru"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 const (
<b>diff --git a/core/keeper.go b/core/keeper.go</b>
<b>new file mode 100644</b>
<b>index 0000000..abe3acc</b>
<b>--- /dev/null</b>
<b>+++ b/core/keeper.go</b>
<span style="color:#0AA">@@ -0,0 +1,149 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">package core</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">import (</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"fmt"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"math/big"</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/common"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/log"</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span><span style="color:#0A0">)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define errors</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type ErrInvalidKeeperData struct{}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ErrInvalidKeeperData) Error() string { return "invalid return data from keeper trigger" }</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type ErrKeeperDataEmpty struct{}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ErrKeeperDataEmpty) Error() string { return "return data from keeper trigger empty" }</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type ErrMaxMintExceeded struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintMax     *big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest *big.Int</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ErrMaxMintExceeded) Error() string {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return fmt.Sprintf("mint request of %s exceeded max of %s", e.mintRequest.Text(10), e.mintMax.Text(10))</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type ErrMintNegative struct{}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ErrMintNegative) Error() string { return "mint request cannot be negative" }</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define interface for dependencies</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type EVMCaller interface {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">GetBlockNumber() *big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">GetGasLimit() uint64</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">AddBalance(addr common.Address, amount *big.Int)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define maximums that can change by block height</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetKeeperGasMultiplier(blockNumber *big.Int) uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 100</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetSystemTriggerContractAddr(blockNumber *big.Int) string {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "0x1000000000000000000000000000000000000002"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetSystemTriggerSelector(blockNumber *big.Int) []byte {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{0x7f, 0xec, 0x8d, 0x38}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetPrioritisedFTSOContract(blockTime *big.Int) string {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "0x1000000000000000000000000000000000000003"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetMaximumMintRequest(blockNumber *big.Int) *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">maxRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return maxRequest</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func triggerKeeper(evm EVMCaller) (*big.Int, error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">bigZero := big.NewInt(0)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Get the contract to call</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">systemTriggerContract := common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber()))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call the method</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">triggerRet, _, triggerErr := evm.Call(</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.AccountRef(systemTriggerContract),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">systemTriggerContract,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">GetSystemTriggerSelector(evm.GetBlockNumber()),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">GetKeeperGasMultiplier(evm.GetBlockNumber())*evm.GetGasLimit(),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">bigZero)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// If no error and a value came back...</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if triggerErr == nil && triggerRet != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Did we get one big int?</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if len(triggerRet) == 32 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">// Convert to big int</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">// Mint request cannot be less than 0 as SetBytes treats value as unsigned</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">mintRequest := new(big.Int).SetBytes(triggerRet)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">// return the mint request</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return mintRequest, nil</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">// Returned length was not 32 bytes</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return bigZero, &ErrInvalidKeeperData{}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if triggerErr != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return bigZero, triggerErr</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return bigZero, &ErrKeeperDataEmpty{}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func mint(evm EVMCaller, mintRequest *big.Int) error {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// If the mint request is greater than zero and less than max</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">max := GetMaximumMintRequest(evm.GetBlockNumber())</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if mintRequest.Cmp(big.NewInt(0)) > 0 &&</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequest.Cmp(max) <= 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Mint the amount asked for on to the keeper contract</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">evm.AddBalance(common.HexToAddress(GetSystemTriggerContractAddr(evm.GetBlockNumber())), mintRequest)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if mintRequest.Cmp(max) > 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Return error</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return &ErrMaxMintExceeded{</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">mintRequest: mintRequest,</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">mintMax:     max,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if mintRequest.Cmp(big.NewInt(0)) < 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Cannot mint negatives</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return &ErrMintNegative{}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// No error</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func triggerKeeperAndMint(evm EVMCaller, log log.Logger) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call the keeper</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest, triggerErr := triggerKeeper(evm)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// If no error...</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if triggerErr == nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// time to mint</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if mintError := mint(evm, mintRequest); mintError != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">log.Warn("Error minting inflation request", "error", mintError)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">log.Warn("Keeper trigger in error", "error", triggerErr)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<b>diff --git a/core/keeper_test.go b/core/keeper_test.go</b>
<b>new file mode 100644</b>
<b>index 0000000..10fff7c</b>
<b>--- /dev/null</b>
<b>+++ b/core/keeper_test.go</b>
<span style="color:#0AA">@@ -0,0 +1,451 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">package core</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">import (</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"errors"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"math/big"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"testing"</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/common"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/log"</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span><span style="color:#0A0">)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define a mock structure to spy and mock values for keeper calls</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type MockEVMCallerData struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">callCalls            int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">addBalanceCalls      int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">blockNumber          big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">gasLimit             uint64</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn    big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">lastAddBalanceAddr   common.Address</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">lastAddBalanceAmount *big.Int</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define a mock structure to spy and mock values for logger calls</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type MockLoggerData struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">warnCalls int</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Set up default mock method calls</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func defautCall(e *MockEVMCallerData, caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.callCalls++</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">buffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return e.mintRequestReturn.FillBytes(buffer), 0, nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func defaultGetBlockNumber(e *MockEVMCallerData) *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return &e.blockNumber</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func defaultGetGasLimit(e *MockEVMCallerData) uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return e.gasLimit</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func defaultAddBalance(e *MockEVMCallerData, addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.addBalanceCalls++</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.lastAddBalanceAddr = addr</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.lastAddBalanceAmount = amount</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define the default EVM mock and define default mock receiver functions</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type DefaultEVMMock struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData MockEVMCallerData</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *DefaultEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defautCall(&e.mockEVMCallerData, caller, addr, input, gas, value)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *DefaultEVMMock) GetBlockNumber() *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetBlockNumber(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *DefaultEVMMock) GetGasLimit() uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetGasLimit(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *DefaultEVMMock) AddBalance(addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultAddBalance(&e.mockEVMCallerData, addr, amount)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldReturnMintRequest(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *mintRequestReturn,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest, _ := triggerKeeper(defaultEVMMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if mintRequest.Cmp(mintRequestReturn) != 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("got %s want %q", mintRequest.Text(10), "50000000000000000000000000")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldNotLetMintRequestOverflow(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var mintRequestReturn big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// TODO: Compact with exponent?</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">buffer := []byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn.SetBytes(buffer)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: mintRequestReturn,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest, mintRequestError := triggerKeeper(defaultEVMMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if mintRequestError != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("received unexpected error %s", mintRequestError)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if mintRequest.Sign() < 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("unexpected negative")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define a bad mint request return size mock</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type BadMintReturnSizeEVMMock struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData MockEVMCallerData</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadMintReturnSizeEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.mockEVMCallerData.callCalls++</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Should be size 32 bytes</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">buffer := []byte{0}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadMintReturnSizeEVMMock) GetBlockNumber() *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetBlockNumber(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadMintReturnSizeEVMMock) GetGasLimit() uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetGasLimit(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadMintReturnSizeEVMMock) AddBalance(addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultAddBalance(&e.mockEVMCallerData, addr, amount)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerValidatesMintRequestReturnValueSize(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var mintRequestReturn big.Int</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// TODO: Compact with exponent?</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">buffer := []byte{255}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn.SetBytes(buffer)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: mintRequestReturn,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">badMintReturnSizeEVMMock := &BadMintReturnSizeEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call to return less than 32 bytes</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">_, err := triggerKeeper(badMintReturnSizeEVMMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err, ok := err.(*ErrInvalidKeeperData); !ok {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">want := &ErrInvalidKeeperData{}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("got '%s' want '%s'", err.Error(), want.Error())</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("no error returned as expected")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define a mock to simulate keeper trigger returning an error from Call</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type BadTriggerCallEVMMock struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData MockEVMCallerData</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadTriggerCallEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.mockEVMCallerData.callCalls++</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">buffer := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return e.mockEVMCallerData.mintRequestReturn.FillBytes(buffer), 0, errors.New("Call error happened")</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadTriggerCallEVMMock) GetBlockNumber() *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetBlockNumber(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadTriggerCallEVMMock) GetGasLimit() uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetGasLimit(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *BadTriggerCallEVMMock) AddBalance(addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultAddBalance(&e.mockEVMCallerData, addr, amount)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerReturnsCallError(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">badTriggerCallEVMMock := &BadTriggerCallEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call to return less than 32 bytes</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">_, err := triggerKeeper(badTriggerCallEVMMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err == nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("no error received")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err.Error() != "Call error happened" {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("did not get expected error")</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type LoggerMock struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockLoggerData MockLoggerData</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) New(ctx ...interface{}) log.Logger {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) GetHandler() log.Handler {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) SetHandler(h log.Handler) {</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Trace(msg string, ctx ...interface{}) {}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Debug(msg string, ctx ...interface{}) {}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Info(msg string, ctx ...interface{})  {}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Error(msg string, ctx ...interface{}) {}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Crit(msg string, ctx ...interface{})  {}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (l *LoggerMock) Warn(msg string, ctx ...interface{}) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">l.mockLoggerData.warnCalls++</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerAndMintLogsError(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assemble</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Set up mock EVM call to return an error</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">badTriggerCallEVMMock := &BadTriggerCallEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Set up a mock logger</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockLoggerData := &MockLoggerData{}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">loggerMock := &LoggerMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockLoggerData: *mockLoggerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Act</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">triggerKeeperAndMint(badTriggerCallEVMMock, loggerMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assert</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if loggerMock.mockLoggerData.warnCalls != 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("Logger.Warn not called as expected")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Define a mock to simulate keeper trigger returning nil for mint request</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type ReturnNilMintRequestEVMMock struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData MockEVMCallerData</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ReturnNilMintRequestEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">e.mockEVMCallerData.callCalls++</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil, 0, nil</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ReturnNilMintRequestEVMMock) GetBlockNumber() *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetBlockNumber(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ReturnNilMintRequestEVMMock) GetGasLimit() uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return defaultGetGasLimit(&e.mockEVMCallerData)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (e *ReturnNilMintRequestEVMMock) AddBalance(addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultAddBalance(&e.mockEVMCallerData, addr, amount)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerHandlesNilMintRequest(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">returnNilMintRequestEVMMock := &ReturnNilMintRequestEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call to return less than 32 bytes</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">_, err := triggerKeeper(returnNilMintRequestEVMMock)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err, ok := err.(*ErrKeeperDataEmpty); !ok {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">want := &ErrKeeperDataEmpty{}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("got '%s' want '%s'", err.Error(), want.Error())</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("no error returned as expected")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldNotMintMoreThanMax(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest, _ := new(big.Int).SetString("50000000000000000000000001", 10)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *big.NewInt(0),</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err := mint(defaultEVMMock, mintRequest)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err, ok := err.(*ErrMaxMintExceeded); !ok {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">want := &ErrMaxMintExceeded{</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">mintRequest: mintRequest,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">mintMax:     GetMaximumMintRequest(big.NewInt(0)),</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("got '%s' want '%s'", err.Error(), want.Error())</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("no error returned as expected")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldNotMintNegative(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest := big.NewInt(-1)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *big.NewInt(0),</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err := mint(defaultEVMMock, mintRequest)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err, ok := err.(*ErrMintNegative); !ok {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">want := &ErrMintNegative{}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("got '%s' want '%s'", err.Error(), want.Error())</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("no error returned as expected")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldMint(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assemble</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest, _ := new(big.Int).SetString("50000000000000000000000000", 10)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *big.NewInt(0),</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Act</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err := mint(defaultEVMMock, mintRequest)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assert</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err == nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 1 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("AddBalance not called as expected")</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr.String() != GetSystemTriggerContractAddr(big.NewInt(0)) {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("wanted addr %s; got addr %s", GetSystemTriggerContractAddr(big.NewInt(0)), defaultEVMMock.mockEVMCallerData.lastAddBalanceAddr)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Cmp(mintRequest) != 0 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("wanted amount %s; got amount %s", mintRequest.Text(10), defaultEVMMock.mockEVMCallerData.lastAddBalanceAmount.Text(10))</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("unexpected error returned; was = %s", err.Error())</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldNotErrorMintingZero(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assemble</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequest := big.NewInt(0)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *big.NewInt(0),</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Act</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err := mint(defaultEVMMock, mintRequest)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Assert</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err == nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 0 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">t.Errorf("AddBalance called unexpectedly")</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("unexpected error returned; was %s", err.Error())</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerFiredAndMinted(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *mintRequestReturn,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">log := log.New()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">triggerKeeperAndMint(defaultEVMMock, log)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// EVM Call function calling the keeper should have been cqlled</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.callCalls != 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("EVM Call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.callCalls)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// AddBalance should have been called on the state database, minting the request asked for</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("Add balance call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.addBalanceCalls)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func TestKeeperTriggerShouldNotMintMoreThanLimit(t *testing.T) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000001", 10)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">mockEVMCallerData := &MockEVMCallerData{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">blockNumber:       *big.NewInt(0),</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">gasLimit:          0,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mintRequestReturn: *mintRequestReturn,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defaultEVMMock := &DefaultEVMMock{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">mockEVMCallerData: *mockEVMCallerData,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">log := log.New()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">triggerKeeperAndMint(defaultEVMMock, log)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// EVM Call function calling the keeper should have been called</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.callCalls != 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("EVM Call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.callCalls)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// AddBalance should not have been called on the state database, as the mint request was over the limit</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if defaultEVMMock.mockEVMCallerData.addBalanceCalls != 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">t.Errorf("Add balance call count not as expected. got %d want 1", defaultEVMMock.mockEVMCallerData.addBalanceCalls)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<b>diff --git a/core/mkalloc.go b/core/mkalloc.go</b>
<b>index 96f9c31..046924b 100644</b>
<b>--- a/core/mkalloc.go</b>
<b>+++ b/core/mkalloc.go</b>
<span style="color:#0AA">@@ -24,6 +24,7 @@</span>
 // You should have received a copy of the GNU Lesser General Public License
 // along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
 
<span style="color:#0A0">+</span><span style="color:#0A0">//go:build none</span>
 // +build none
 
 /*
<span style="color:#0AA">@@ -44,8 +45,8 @@</span> import (
 	"sort"
 	"strconv"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
 )
 
 type allocItem struct{ Addr, Balance *big.Int }
<b>diff --git a/core/rawdb/accessors_chain.go b/core/rawdb/accessors_chain.go</b>
<b>index 6ed9e0a..13e0f1a 100644</b>
<b>--- a/core/rawdb/accessors_chain.go</b>
<b>+++ b/core/rawdb/accessors_chain.go</b>
<span style="color:#0AA">@@ -31,13 +31,13 @@</span> import (
 	"encoding/binary"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // ReadCanonicalHash retrieves the hash assigned to a canonical block number.
<b>diff --git a/core/rawdb/accessors_chain_test.go b/core/rawdb/accessors_chain_test.go</b>
<b>index 38372a1..55b3ab4 100644</b>
<b>--- a/core/rawdb/accessors_chain_test.go</b>
<b>+++ b/core/rawdb/accessors_chain_test.go</b>
<span style="color:#0AA">@@ -24,10 +24,10 @@</span> import (
 	"reflect"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 	"golang.org/x/crypto/sha3"
 )
 
<b>diff --git a/core/rawdb/accessors_indexes.go b/core/rawdb/accessors_indexes.go</b>
<b>index 3c8b467..4cec301 100644</b>
<b>--- a/core/rawdb/accessors_indexes.go</b>
<b>+++ b/core/rawdb/accessors_indexes.go</b>
<span style="color:#0AA">@@ -30,12 +30,12 @@</span> import (
 	"bytes"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // ReadTxLookupEntry retrieves the positional metadata associated with a transaction
<b>diff --git a/core/rawdb/accessors_indexes_test.go b/core/rawdb/accessors_indexes_test.go</b>
<b>index a80ff1e..8efdc25 100644</b>
<b>--- a/core/rawdb/accessors_indexes_test.go</b>
<b>+++ b/core/rawdb/accessors_indexes_test.go</b>
<span style="color:#0AA">@@ -22,10 +22,10 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 	"golang.org/x/crypto/sha3"
 )
 
<b>diff --git a/core/rawdb/accessors_metadata.go b/core/rawdb/accessors_metadata.go</b>
<b>index 5b39253..47b9e9d 100644</b>
<b>--- a/core/rawdb/accessors_metadata.go</b>
<b>+++ b/core/rawdb/accessors_metadata.go</b>
<span style="color:#0AA">@@ -29,11 +29,11 @@</span> package rawdb
 import (
 	"encoding/json"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // ReadDatabaseVersion retrieves the version number of the database.
<b>diff --git a/core/state/database.go b/core/state/database.go</b>
<b>index f4af1a0..feba572 100644</b>
<b>--- a</b><b>/core/state/database.go</b>
<b>+++ b/core/state/database.go</b>
<span style="color:#0AA">@@ -31,11 +31,11 @@</span> import (
 	"fmt"
 
 	"github.com/VictoriaMetrics/fastcache"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/trie"
 	lru "github.com/hashicorp/golang-lru"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 const (
<b>diff --git a/core/state/snapshot/conversion.go b/core/state/snapshot/conversion.go</b>
<b>index 0fd4e9b..d2b519f 100644</b>
<b>--- a/core/state/snapshot/conversion.go</b>
<b>+++ b/core/state/snapshot/conversion.go</b>
<span style="color:#0AA">@@ -36,12 +36,12 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // trieKV represents a trie key-value pair
<b>diff --git a/core/state/snapshot/disklayer.go b/core/state/snapshot/disklayer.go</b>
<b>index 106c451..c33f4a6 100644</b>
<b>--- a/core/state/snapshot/disklayer.go</b>
<b>+++ b/core/state/snapshot/disklayer.go</b>
<span style="color:#0AA">@@ -32,11 +32,11 @@</span> import (
 	"time"
 
 	"github.com/VictoriaMetrics/fastcache"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // diskLayer is a low level persistent snapshot built on top of a key-value store.
<b>diff --git a/core/state/snapshot/disklayer_test.go b/core/state/snapshot/disklayer_test.go</b>
<b>index f423c9a..9d758be 100644</b>
<b>--- a/core/state/snapshot/disklayer_test.go</b>
<b>+++ b/core/state/snapshot/disklayer_test.go</b>
<span style="color:#0AA">@@ -32,12 +32,12 @@</span> import (
 	"os"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/ethdb/leveldb"
 	"github.com/ethereum/go-ethereum/ethdb/memorydb"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // reverse reverses the contents of a byte slice. It's used to update random accs
<b>diff --git a/core/state/snapshot/generate.go b/core/state/snapshot/generate.go</b>
<b>index 45f927f..3a48d1f 100644</b>
<b>--- a/core/state/snapshot/generate.go</b>
<b>+++ b/core/state/snapshot/generate.go</b>
<span style="color:#0AA">@@ -34,7 +34,6 @@</span> import (
 	"time"
 
 	"github.com/VictoriaMetrics/fastcache"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0AA">@@ -42,6 +41,7 @@</span> import (
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 var (
<b>diff --git a/core/state/snapshot/iterator.go b/core/state/snapshot/iterator.go</b>
<b>index 7a5da6a..2c34768 100644</b>
<b>--- a/core/state/snapshot/iterator.go</b>
<b>+++ b/core/state/snapshot/iterator.go</b>
<span style="color:#0AA">@@ -31,9 +31,9 @@</span> import (
 	"fmt"
 	"sort"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // Iterator is an iterator to step over all the accounts or the specific
<b>diff --git a/core/state/snapshot/iterator_test.go b/core/state/snapshot/iterator_test.go</b>
<b>index 9d1b744..1916c49 100644</b>
<b>--- a/core/state/snapshot/iterator_test.go</b>
<b>+++ b/core/state/snapshot/iterator_test.go</b>
<span style="color:#0AA">@@ -33,8 +33,8 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // TestAccountIteratorBasics tests some simple single-layer(diff and disk) iteration
<b>diff --git a/core/state/snapshot/journal.go b/core/state/snapshot/journal.go</b>
<b>index 145f467..96c7c2e 100644</b>
<b>--- a/core/state/snapshot/journal.go</b>
<b>+++ b/core/state/snapshot/journal.go</b>
<span style="color:#0AA">@@ -33,12 +33,12 @@</span> import (
 	"time"
 
 	"github.com/VictoriaMetrics/fastcache"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // journalGenerator is a disk layer entry containing the generator progress marker.
<b>diff --git a/core/state/snapshot/snapshot.go b/core/state/snapshot/snapshot.go</b>
<b>index bbdba6b..e311d69 100644</b>
<b>--- a/core/state/snapshot/snapshot.go</b>
<b>+++ b/core/state/snapshot/snapshot.go</b>
<span style="color:#0AA">@@ -36,12 +36,12 @@</span> import (
 	"time"
 
 	"github.com/VictoriaMetrics/fastcache"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/metrics"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 const (
<b>diff --git a/core/state/snapshot/snapshot_test.go b/core/state/snapshot/snapshot_test.go</b>
<b>index c28dafa..da5cd63 100644</b>
<b>--- a/core/state/snapshot/snapshot_test.go</b>
<b>+++ b/core/state/snapshot/snapshot_test.go</b>
<span style="color:#0AA">@@ -32,9 +32,9 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // randomHash generates a random blob of data and returns it as a hash.
<b>diff --git a/core/state/snapshot/wipe.go b/core/state/snapshot/wipe.go</b>
<b>index ca1f7e0..3cbf754 100644</b>
<b>--- a/core/state/snapshot/wipe.go</b>
<b>+++ b/core/state/snapshot/wipe.go</b>
<span style="color:#0AA">@@ -30,10 +30,10 @@</span> import (
 	"bytes"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // wipeSnapshot starts a goroutine to iterate over the entire key-value database
<b>diff --git a/core/state/snapshot/wipe_test.go b/core/state/snapshot/wipe_test.go</b>
<b>index 952370d..ba4ce69 100644</b>
<b>--- a/core/state/snapshot/wipe_test.go</b>
<b>+++ b/core/state/snapshot/wipe_test.go</b>
<span style="color:#0AA">@@ -30,9 +30,9 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb/memorydb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 // Tests that given a database with random data content, all parts of a snapshot
<b>diff --git a/core/state/state_test.go b/core/state/state_test.go</b>
<b>index 849f661..ec66ccc 100644</b>
<b>--- a/core/state/state_test.go</b>
<b>+++ b/core/state/state_test.go</b>
<span style="color:#0AA">@@ -27,9 +27,9 @@</span>
 package state
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 type stateTest struct {
<b>diff --git a/core/state/statedb.go b/core/state/statedb.go</b>
<b>index acc779c..3fe062c 100644</b>
<b>--- a/core/state/statedb.go</b>
<b>+++ b/core/state/statedb.go</b>
<span style="color:#0AA">@@ -34,15 +34,15 @@</span> import (
 	"sort"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state/snapshot"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/metrics"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state/snapshot"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 type revision struct {
<b>diff --git a/core/state/statedb_test.go b/core/state/statedb_test.go</b>
<b>index 9c295f9..4841b6b 100644</b>
<b>--- a/core/state/statedb_test.go</b>
<b>+++ b/core/state/statedb_test.go</b>
<span style="color:#0AA">@@ -39,11 +39,11 @@</span> import (
 	"testing"
 	"testing/quick"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state/snapshot"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state/snapshot"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // Tests that updating a state trie does not leak any database writes prior to
<b>diff --git a/core/state_connector.go b/core/state_connector.go</b>
<b>new file mode 100644</b>
<b>index 0000000..db6665e</b>
<b>--- /dev/null</b>
<b>+++ b/core/state_connector.go</b>
<span style="color:#0AA">@@ -0,0 +1,731 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">package core</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">import (</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"bytes"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"crypto/sha256"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"encoding/binary"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"encoding/hex"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"encoding/json"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"fmt"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"io/ioutil"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"math"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"math/big"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"net/http"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"os"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"strconv"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"strings"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"time"</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/common"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/common/hexutil"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/crypto"</span>
<span style="color:#0A0">+</span><span style="color:#0A0">)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">var (</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">testingChainID               = new(big.Int).SetUint64(16)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">stateConnectorActivationTime = new(big.Int).SetUint64(1636070400)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">tr                           = &http.Transport{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">MaxIdleConns:        100,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">MaxConnsPerHost:     100,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">MaxIdleConnsPerHost: 100,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">IdleConnTimeout:     60 * time.Second,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">DisableCompression:  true,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">client = &http.Client{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Transport: tr,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Timeout:   5 * time.Second,</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">apiRetries    = 3</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">apiRetryDelay = 1 * time.Second</span>
<span style="color:#0A0">+</span><span style="color:#0A0">)</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetStateConnectorActivated(chainID *big.Int, blockTime *big.Int) bool {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Return true if chainID is 16 or if block.timestamp is greater than the state connector activation time on any chain</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return chainID.Cmp(testingChainID) == 0 || blockTime.Cmp(stateConnectorActivationTime) > 0</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetStateConnectorGasDivisor(blockTime *big.Int) uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 3</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetMaxAllowedChains(blockTime *big.Int) uint32 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 5</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetStateConnectorContractAddr(blockTime *big.Int) string {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "0x1000000000000000000000000000000000000001"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetProveDataAvailabilityPeriodFinalitySelector(blockTime *big.Int) []byte {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{0xc5, 0xd6, 0x4c, 0xd1}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetProvePaymentFinalitySelector(blockTime *big.Int) []byte {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{0x38, 0x84, 0x92, 0xdd}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetDisprovePaymentFinalitySelector(blockTime *big.Int) []byte {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{0x7f, 0x58, 0x24, 0x32}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Proof of Work Common</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWRequestPayload struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Method string   `json:"method"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Params []string `json:"params"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWBlockCountResp struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Result uint64      `json:"result"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Error  interface{} `json:"error"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetPoWBlockCount(chainURL string, username string, password string) (uint64, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">data := GetPoWRequestPayload{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Method: "getblockcount",</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Params: []string{},</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">payloadBytes, err := json.Marshal(data)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">body := bytes.NewReader(payloadBytes)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req, err := http.NewRequest("POST", chainURL, body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req.Header.Set("Content-Type", "application/json")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if username != "" && password != "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">req.SetBasicAuth(username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">resp, err := client.Do(req)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defer resp.Body.Close()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if resp.StatusCode != 200 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respBody, err := ioutil.ReadAll(resp.Body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var jsonResp GetPoWBlockCountResp</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &jsonResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if jsonResp.Error != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return jsonResp.Result, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWBlockHeaderResult struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Hash          string `json:"hash"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Confirmations uint64 `json:"confirmations"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Height        uint64 `json:"height"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWBlockHeaderResp struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Result GetPoWBlockHeaderResult `json:"result"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Error  interface{}             `json:"error"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetPoWBlockHeader(ledgerHash string, requiredConfirmations uint64, chainURL string, username string, password string) (uint64, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">data := GetPoWRequestPayload{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Method: "getblockheader",</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Params: []string{</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">ledgerHash,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">payloadBytes, err := json.Marshal(data)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">body := bytes.NewReader(payloadBytes)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req, err := http.NewRequest("POST", chainURL, body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req.Header.Set("Content-Type", "application/json")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if username != "" && password != "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">req.SetBasicAuth(username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">resp, err := client.Do(req)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defer resp.Body.Close()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if resp.StatusCode != 200 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respBody, err := ioutil.ReadAll(resp.Body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var jsonResp GetPoWBlockHeaderResp</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &jsonResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if jsonResp.Error != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if jsonResp.Result.Confirmations < requiredConfirmations {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return jsonResp.Result.Height, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProveDataAvailabilityPeriodFinalityPoW(checkRet []byte, chainURL string, username string, password string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">blockCount, err := GetPoWBlockCount(chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">ledger := binary.BigEndian.Uint64(checkRet[56:64])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">requiredConfirmations := binary.BigEndian.Uint64(checkRet[88:96])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if blockCount < ledger+requiredConfirmations {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">ledgerResp, err := GetPoWBlockHeader(hex.EncodeToString(checkRet[96:128]), requiredConfirmations, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if ledgerResp > 0 && ledgerResp == ledger {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWTxRequestParams struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">TxID    string `json:"txid"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Verbose bool   `json:"verbose"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWTxRequestPayload struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Method string                `json:"method"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Params GetPoWTxRequestParams `json:"params"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWTxResult struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">TxID          string `json:"txid"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">BlockHash     string `json:"blockhash"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Confirmations uint64 `json:"confirmations"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Vout          []struct {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Value        float64 `json:"value"`</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">N            uint64  `json:"n"`</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">ScriptPubKey struct {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">Type      string   `json:"type"`</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">Addresses []string `json:"addresses"`</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} `json:"scriptPubKey"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} `json:"vout"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetPoWTxResp struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Result GetPoWTxResult `json:"result"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Error  interface{}    `json:"error"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetPoWTx(txHash string, voutN uint64, latestAvailableBlock uint64, currencyCode string, chainURL string, username string, password string) ([]byte, uint64, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">data := GetPoWTxRequestPayload{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Method: "getrawtransaction",</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Params: GetPoWTxRequestParams{</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">TxID:    txHash[1:],</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">Verbose: true,</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">payloadBytes, err := json.Marshal(data)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">body := bytes.NewReader(payloadBytes)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req, err := http.NewRequest("POST", chainURL, body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req.Header.Set("Content-Type", "application/json")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if username != "" && password != "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">req.SetBasicAuth(username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">resp, err := client.Do(req)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defer resp.Body.Close()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if resp.StatusCode != 200 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respBody, err := ioutil.ReadAll(resp.Body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var jsonResp GetPoWTxResp</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &jsonResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if jsonResp.Error != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if uint64(len(jsonResp.Result.Vout)) <= voutN {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if jsonResp.Result.Vout[voutN].ScriptPubKey.Type != "pubkeyhash" || len(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses) != 1 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">inBlock, getBlockErr := GetPoWBlockHeader(jsonResp.Result.BlockHash, jsonResp.Result.Confirmations, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if getBlockErr {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if inBlock == 0 || inBlock >= latestAvailableBlock {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">txIdHash := crypto.Keccak256([]byte(txHash))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">destinationHash := crypto.Keccak256([]byte(jsonResp.Result.Vout[voutN].ScriptPubKey.Addresses[0]))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">amountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp.Result.Vout[voutN].Value*math.Pow(10, 8)))), 32))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">currencyHash := crypto.Keccak256([]byte(currencyCode))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inBlock, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProvePaymentFinalityPoW(checkRet []byte, isDisprove bool, currencyCode string, chainURL string, username string, password string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if len(checkRet) < 257 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">voutN, err := strconv.ParseUint(string(checkRet[192:193]), 16, 64)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">paymentHash, inBlock, getPoWTxErr := GetPoWTx(string(checkRet[192:257]), voutN, binary.BigEndian.Uint64(checkRet[88:96]), currencyCode, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if getPoWTxErr {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if !isDisprove {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock == binary.BigEndian.Uint64(checkRet[56:64]) {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inBlock > binary.BigEndian.Uint64(checkRet[56:64]) {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else if len(paymentHash) == 0 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProvePoW(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, currencyCode string, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var username, password string</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">chainURLhash := sha256.Sum256([]byte(chainURL))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">chainURLchecksum := hex.EncodeToString(chainURLhash[0:4])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch currencyCode {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case "btc":</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">username = os.Getenv("BTC_U_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">password = os.Getenv("BTC_P_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case "ltc":</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">username = os.Getenv("LTC_U_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">password = os.Getenv("LTC_P_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case "dog":</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">username = os.Getenv("DOGE_U_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">password = os.Getenv("DOGE_P_" + chainURLchecksum)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProveDataAvailabilityPeriodFinalityPoW(checkRet, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePaymentFinalityPoW(checkRet, false, currencyCode, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePaymentFinalityPoW(checkRet, true, currencyCode, chainURL, username, password)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// XRP</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPBlockRequestParams struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">LedgerIndex  uint64 `json:"ledger_index"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Full         bool   `json:"full"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Accounts     bool   `json:"accounts"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Transactions bool   `json:"transactions"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Expand       bool   `json:"expand"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">OwnerFunds   bool   `json:"owner_funds"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPBlockRequestPayload struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Method string                     `json:"method"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Params []GetXRPBlockRequestParams `json:"params"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type CheckXRPErrorResponse struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Error string `json:"error"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPBlockResponse struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">LedgerHash  string `json:"ledger_hash"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">LedgerIndex int    `json:"ledger_index"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Validated   bool   `json:"validated"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetXRPBlock(ledger uint64, chainURL string) (string, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">data := GetXRPBlockRequestPayload{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Method: "ledger",</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Params: []GetXRPBlockRequestParams{</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">{</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">LedgerIndex:  ledger,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Full:         false,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Accounts:     false,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Transactions: false,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Expand:       false,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">OwnerFunds:   false,</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">payloadBytes, err := json.Marshal(data)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">body := bytes.NewReader(payloadBytes)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req, err := http.NewRequest("POST", chainURL, body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req.Header.Set("Content-Type", "application/json")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">resp, err := client.Do(req)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defer resp.Body.Close()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if resp.StatusCode != 200 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respBody, err := ioutil.ReadAll(resp.Body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var checkErrorResp map[string]CheckXRPErrorResponse</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &checkErrorResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if checkErrorResp["result"].Error != "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var jsonResp map[string]GetXRPBlockResponse</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &jsonResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if !jsonResp["result"].Validated {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return "", true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return jsonResp["result"].LedgerHash, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProveDataAvailabilityPeriodFinalityXRP(checkRet []byte, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">ledger := binary.BigEndian.Uint64(checkRet[56:64])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">ledgerHashString, err := GetXRPBlock(ledger, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if ledgerHashString != "" && bytes.Equal(crypto.Keccak256([]byte(ledgerHashString)), checkRet[96:128]) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPTxRequestParams struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Transaction string `json:"transaction"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Binary      bool   `json:"binary"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPTxRequestPayload struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Method string                  `json:"method"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Params []GetXRPTxRequestParams `json:"params"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPTxResponse struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Destination     string `json:"Destination"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">DestinationTag  int    `json:"DestinationTag"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">TransactionType string `json:"TransactionType"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Hash            string `json:"hash"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">InLedger        int    `json:"inLedger"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Validated       bool   `json:"validated"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Meta            struct {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">TransactionResult string      `json:"TransactionResult"`</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Amount            interface{} `json:"delivered_amount"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} `json:"meta"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">type GetXRPTxIssuedCurrency struct {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Currency string `json:"currency"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Issuer   string `json:"issuer"`</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">Value    string `json:"value"`</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetXRPTx(txHash string, latestAvailableLedger uint64, chainURL string) ([]byte, uint64, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">data := GetXRPTxRequestPayload{</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Method: "tx",</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">Params: []GetXRPTxRequestParams{</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">{</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Transaction: txHash,</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">Binary:      false,</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">},</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">payloadBytes, err := json.Marshal(data)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">body := bytes.NewReader(payloadBytes)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req, err := http.NewRequest("POST", chainURL, body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">req.Header.Set("Content-Type", "application/json")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">resp, err := client.Do(req)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">defer resp.Body.Close()</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if resp.StatusCode != 200 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respBody, err := ioutil.ReadAll(resp.Body)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var checkErrorResp map[string]CheckXRPErrorResponse</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &checkErrorResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">respErrString := checkErrorResp["result"].Error</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if respErrString != "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if respErrString == "amendmentBlocked" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "failedToForward" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "invalid_API_version" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "noClosed" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "noCurrent" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "noNetwork" ||</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">respErrString == "tooBusy" {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var jsonResp map[string]GetXRPTxResponse</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">err = json.Unmarshal(respBody, &jsonResp)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if jsonResp["result"].TransactionType != "Payment" || !jsonResp["result"].Validated || jsonResp["result"].Meta.TransactionResult != "tesSUCCESS" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">inLedger := uint64(jsonResp["result"].InLedger)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if inLedger == 0 || inLedger >= latestAvailableLedger || !jsonResp["result"].Validated {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var currency string</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var amount uint64</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if stringAmount, ok := jsonResp["result"].Meta.Amount.(string); ok {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">amount, err = strconv.ParseUint(stringAmount, 10, 64)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">currency = "xrp"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">amountStruct, err := json.Marshal(jsonResp["result"].Meta.Amount)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">var issuedCurrencyResp GetXRPTxIssuedCurrency</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">err = json.Unmarshal(amountStruct, &issuedCurrencyResp)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">floatAmount, err := strconv.ParseFloat(issuedCurrencyResp.Value, 64)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return []byte{}, 0, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">amount = uint64(floatAmount * math.Pow(10, 15))</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">currency = issuedCurrencyResp.Currency + issuedCurrencyResp.Issuer</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">txIdHash := crypto.Keccak256([]byte(jsonResp["result"].Hash))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">destinationHash := crypto.Keccak256([]byte(jsonResp["result"].Destination))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">destinationTagHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(uint64(jsonResp["result"].DestinationTag))), 32))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">destinationHash = crypto.Keccak256(destinationHash, destinationTagHash)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">amountHash := crypto.Keccak256(common.LeftPadBytes(common.FromHex(hexutil.EncodeUint64(amount)), 32))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">currencyHash := crypto.Keccak256([]byte(currency))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return crypto.Keccak256(txIdHash, destinationHash, amountHash, currencyHash), inLedger, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProvePaymentFinalityXRP(checkRet []byte, isDisprove bool, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">paymentHash, inLedger, err := GetXRPTx(string(checkRet[192:]), binary.BigEndian.Uint64(checkRet[88:96]), chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if err {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if !isDisprove {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger == binary.BigEndian.Uint64(checkRet[56:64]) {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if len(paymentHash) > 0 && bytes.Equal(paymentHash, checkRet[96:128]) && inLedger > binary.BigEndian.Uint64(checkRet[56:64]) {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else if len(paymentHash) == 0 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return true, false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProveXRP(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if bytes.Equal(functionSelector, GetProveDataAvailabilityPeriodFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProveDataAvailabilityPeriodFinalityXRP(checkRet, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if bytes.Equal(functionSelector, GetProvePaymentFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePaymentFinalityXRP(checkRet, false, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if bytes.Equal(functionSelector, GetDisprovePaymentFinalitySelector(blockTime)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePaymentFinalityXRP(checkRet, true, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// ALGO</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProveALGO(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false, false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Common</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// =======================================================</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ProveChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte, chainId uint32, chainURL string) (bool, bool) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch chainId {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 0:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePoW(sender, blockTime, functionSelector, checkRet, "btc", chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 1:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePoW(sender, blockTime, functionSelector, checkRet, "ltc", chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 2:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProvePoW(sender, blockTime, functionSelector, checkRet, "dog", chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 3:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProveXRP(sender, blockTime, functionSelector, checkRet, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 4:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return ProveALGO(sender, blockTime, functionSelector, checkRet, chainURL)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">default:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false, true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func ReadChain(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">chainId := binary.BigEndian.Uint32(checkRet[28:32])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">var chainURLs string</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">switch chainId {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 0:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">chainURLs = os.Getenv("BTC_APIs")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 1:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">chainURLs = os.Getenv("LTC_APIs")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 2:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">chainURLs = os.Getenv("DOGE_APIs")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 3:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">chainURLs = os.Getenv("XRP_APIs")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">case 4:</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">chainURLs = os.Getenv("ALGO_APIs")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if chainURLs == "" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">for i := 0; i < apiRetries; i++ {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">for _, chainURL := range strings.Split(chainURLs, ",") {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">if chainURL == "" {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">continue</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">verified, err := ProveChain(sender, blockTime, functionSelector, checkRet, chainId, chainURL)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">if !verified && err {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">continue</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">return verified</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">time.Sleep(apiRetryDelay)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return false</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func GetVerificationPaths(functionSelector []byte, checkRet []byte) (string, string) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">prefix := "cache/"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">acceptedPrefix := "ACCEPTED"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">rejectedPrefix := "REJECTED"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">functionHash := hex.EncodeToString(functionSelector[:])</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">verificationHash := hex.EncodeToString(crypto.Keccak256(checkRet[0:64], checkRet[96:128]))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">suffix := "_" + functionHash + "_" + verificationHash</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return prefix + acceptedPrefix + suffix, prefix + rejectedPrefix + suffix</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Verify proof against underlying chain</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func StateConnectorCall(sender common.Address, blockTime *big.Int, functionSelector []byte, checkRet []byte) bool {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if binary.BigEndian.Uint64(checkRet[88:96]) > 0 {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">go func() {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">acceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">_, errACCEPTED := os.Stat(acceptedPath)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">_, errREJECTED := os.Stat(rejectedPath)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">if errACCEPTED != nil && errREJECTED != nil {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">if ReadChain(sender, blockTime, functionSelector, checkRet) {</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">verificationHashStore, err := os.Create(acceptedPath)</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">verificationHashStore.Close()</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>						<span style="color:#0A0">// Permissions problem</span>
<span style="color:#0A0">+</span>						<span style="color:#0A0">panic(err)</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">verificationHashStore, err := os.Create(rejectedPath)</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">verificationHashStore.Close()</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">if err != nil {</span>
<span style="color:#0A0">+</span>						<span style="color:#0A0">// Permissions problem</span>
<span style="color:#0A0">+</span>						<span style="color:#0A0">panic(err)</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}()</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">acceptedPath, rejectedPath := GetVerificationPaths(functionSelector, checkRet)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">_, errACCEPTED := os.Stat(acceptedPath)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">_, errREJECTED := os.Stat(rejectedPath)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if errACCEPTED != nil && errREJECTED != nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">for i := 0; i < 2*apiRetries; i++ {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">_, errACCEPTED = os.Stat(acceptedPath)</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">_, errREJECTED = os.Stat(rejectedPath)</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">if errACCEPTED == nil || errREJECTED == nil {</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">break</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">time.Sleep(apiRetryDelay)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">go func() {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">removeFulfilledAPIRequests := os.Getenv("REMOVE_FULFILLED_API_REQUESTS")</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">if removeFulfilledAPIRequests == "1" {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">errDeleteACCEPTED := os.Remove(acceptedPath)</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">errDeleteREJECTED := os.Remove(rejectedPath)</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">if errDeleteACCEPTED != nil && errDeleteREJECTED != nil {</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">// Permissions problem</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">panic(fmt.Sprintf("%s\n%s", errDeleteACCEPTED, errDeleteREJECTED))</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}()</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return errACCEPTED == nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<b>diff --git a/core/state_manager.go b/core/state_manager.go</b>
<b>index 8b76069..c95e03f 100644</b>
<b>--- a/core/state_manager.go</b>
<b>+++ b/core/state_manager.go</b>
<span style="color:#0AA">@@ -30,9 +30,9 @@</span> import (
 	"fmt"
 	"math/rand"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 const (
<b>diff --git a/core/state_manager_test.go b/core/state_manager_test.go</b>
<b>index 17e2b1d..f57d105 100644</b>
<b>--- a/core/state_manager_test.go</b>
<b>+++ b/core/state_manager_test.go</b>
<span style="color:#0AA">@@ -7,7 +7,7 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/stretchr/testify/assert"
<b>diff --git a/core/state_prefetcher.go b/core/state_prefetcher.go</b>
<b>index a1e0cde..f0fd973 100644</b>
<b>--- a/core/state_prefetcher.go</b>
<b>+++ b/core/state_prefetcher.go</b>
<span style="color:#0AA">@@ -30,11 +30,11 @@</span> import (
 	"math/big"
 	"sync/atomic"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // statePrefetcher is a basic Prefetcher, which blindly executes a block on top
<b>diff --git a/core/state_processor.go b/core/state_processor.go</b>
<b>index b9c243f..ec6d709 100644</b>
<b>--- a/core/state_processor.go</b>
<b>+++ b/core/state_processor.go</b>
<span style="color:#0AA">@@ -30,14 +30,14 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/misc"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/misc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // StateProcessor is a basic Processor, which takes care of transitioning
<b>diff --git a/core/state_transition.go b/core/state_transition.go</b>
<b>index 381327c..876d39d 100644</b>
<b>--- a/core/state_transition.go</b>
<b>+++ b/core/state_transition.go</b>
<span style="color:#0AA">@@ -1,3 +1,11 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// This file is a derived work, based on the avalanchego library whose original</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// notices appear below. It is distributed under a license compatible with the</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// licensing terms of the original code from which it is derived.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE_AVALABS for licensing terms of the original work.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
 // (c) 2019-2020, Ava Labs, Inc.
 //
 // This file is a derived work, based on the go-ethereum library whose original
<span style="color:#0AA">@@ -27,16 +35,19 @@</span>
 package core
 
 import (
<span style="color:#0A0">+</span>	<span style="color:#0A0">"bytes"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"encoding/binary"</span>
 	"fmt"
 	"math"
 	"math/big"
 
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/common"</span>
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"github.com/ethereum/go-ethereum/log"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ethereum/go-ethereum/common"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var emptyCodeHash = crypto.Keccak256Hash(nil)
<span style="color:#0AA">@@ -115,6 +126,23 @@</span> func (result *ExecutionResult) Return() []byte {
 	return common.CopyBytes(result.ReturnData)
 }
 
<span style="color:#0A0">+</span><span style="color:#0A0">// Implement the EVMCaller interface on the state transition structure; simply delegate the calls</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (st *StateTransition) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return st.evm.Call(caller, addr, input, gas, value)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (st *StateTransition) GetBlockNumber() *big.Int {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return st.evm.Context.BlockNumber</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (st *StateTransition) GetGasLimit() uint64 {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return st.evm.Context.GasLimit</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span><span style="color:#0A0">func (st *StateTransition) AddBalance(addr common.Address, amount *big.Int) {</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">st.state.AddBalance(addr, amount)</span>
<span style="color:#0A0">+</span><span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
 // Revert returns the concrete revert reason if the execution is aborted by `REVERT`
 // opcode. Note the reason can be nil if no data supplied with revert opcode.
 func (result *ExecutionResult) Revert() []byte {
<span style="color:#0AA">@@ -321,19 +349,91 @@</span> func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {
 	if rules := st.evm.ChainConfig().AvalancheRules(st.evm.Context.BlockNumber, st.evm.Context.Time); rules.IsApricotPhase2 {
 		st.state.PrepareAccessList(msg.From(), msg.To(), vm.ActivePrecompiles(rules), msg.AccessList())
 	}
<span style="color:#0A0">+</span>
 	var (
<span style="color:#A00">-		ret   []byte</span>
<span style="color:#A00">-		vmerr error // vm errors do not effect consensus and are therefore not assigned to err</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">ret                                       []byte</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vmerr                                     error // vm errors do not affect consensus and are therefore not assigned to err</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">selectProveDataAvailabilityPeriodFinality bool</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">selectProvePaymentFinality                bool</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">selectDisprovePaymentFinality             bool</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">prioritisedFTSOContract                   bool</span>
 	)
<span style="color:#A00">-	if contractCreation {</span>
<span style="color:#A00">-		ret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)</span>
<span style="color:#A00">-	} else {</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if st.evm.Context.Coinbase != common.HexToAddress("0x0100000000000000000000000000000000000000") {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return nil, fmt.Errorf("Invalid value for block.coinbase")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if st.msg.From() == common.HexToAddress("0x0100000000000000000000000000000000000000") ||</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">st.msg.From() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) ||</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">st.msg.From() == common.HexToAddress(GetSystemTriggerContractAddr(st.evm.Context.Time)) {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">return nil, fmt.Errorf("Invalid sender")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">burnAddress := st.evm.Context.Coinbase</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if !contractCreation {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if *msg.To() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) && len(st.data) >= 4 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">selectProveDataAvailabilityPeriodFinality = bytes.Equal(st.data[0:4], GetProveDataAvailabilityPeriodFinalitySelector(st.evm.Context.Time))</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">selectProvePaymentFinality = bytes.Equal(st.data[0:4], GetProvePaymentFinalitySelector(st.evm.Context.Time))</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">selectDisprovePaymentFinality = bytes.Equal(st.data[0:4], GetDisprovePaymentFinalitySelector(st.evm.Context.Time))</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">prioritisedFTSOContract = *msg.To() == common.HexToAddress(GetPrioritisedFTSOContract(st.evm.Context.Time))</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if selectProveDataAvailabilityPeriodFinality || selectProvePaymentFinality || selectDisprovePaymentFinality {</span>
 		// Increment the nonce for the next transaction
 		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)
<span style="color:#A00">-		ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">stateConnectorGas := st.gas / GetStateConnectorGasDivisor(st.evm.Context.Time)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">checkRet, _, checkVmerr := st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if checkVmerr == nil {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">chainConfig := st.evm.ChainConfig()</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">if GetStateConnectorActivated(chainConfig.ChainID, st.evm.Context.Time) && binary.BigEndian.Uint32(checkRet[28:32]) < GetMaxAllowedChains(st.evm.Context.Time) {</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">if StateConnectorCall(msg.From(), st.evm.Context.Time, st.data[0:4], checkRet) {</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">originalCoinbase := st.evm.Context.Coinbase</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">defer func() {</span>
<span style="color:#0A0">+</span>						<span style="color:#0A0">st.evm.Context.Coinbase = originalCoinbase</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">}()</span>
<span style="color:#0A0">+</span>					<span style="color:#0A0">st.evm.Context.Coinbase = st.msg.From()</span>
<span style="color:#0A0">+</span>				<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if contractCreation {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">ret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">// Increment the nonce for the next transaction</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
 	}
 	st.refundGas(apricotPhase1)
<span style="color:#A00">-	st.state.AddBalance(st.evm.Context.Coinbase, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if vmerr == nil && prioritisedFTSOContract {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">nominalGasUsed := uint64(21000)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">nominalGasPrice := uint64(225_000_000_000)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">nominalFee := new(big.Int).Mul(new(big.Int).SetUint64(nominalGasUsed), new(big.Int).SetUint64(nominalGasPrice))</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">actualGasUsed := st.gasUsed()</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">actualGasPrice := st.gasPrice</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">actualFee := new(big.Int).Mul(new(big.Int).SetUint64(actualGasUsed), actualGasPrice)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">if actualFee.Cmp(nominalFee) > 0 {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">feeRefund := new(big.Int).Sub(actualFee, nominalFee)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">st.state.AddBalance(st.msg.From(), feeRefund)</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">st.state.AddBalance(burnAddress, nominalFee)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>			<span style="color:#0A0">st.state.AddBalance(burnAddress, actualFee)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">st.state.AddBalance(burnAddress, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">// Call the keeper contract trigger method if there is no vm error</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if vmerr == nil {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Temporarily disable EVM debugging</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">oldDebug := st.evm.Config.Debug</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">st.evm.Config.Debug = false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">// Call the keeper contract trigger</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">log := log.Root()</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">triggerKeeperAndMint(st, log)</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">st.evm.Config.Debug = oldDebug</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
 
 	return &ExecutionResult{
 		UsedGas:    st.gasUsed(),
<b>diff --git a/core/test_blockchain.go b/core/test_blockchain.go</b>
<b>index 604ed0f..cfdd2b9 100644</b>
<b>--- a/core/test_blockchain.go</b>
<b>+++ b/core/test_blockchain.go</b>
<span style="color:#0AA">@@ -8,13 +8,13 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type ChainTest struct {
<b>diff --git a/core/tx_cacher.go b/core/tx_cacher.go</b>
<b>index 86679ff..03099f2 100644</b>
<b>--- a/core/tx_cacher.go</b>
<b>+++ b/core/tx_cacher.go</b>
<span style="color:#0AA">@@ -29,7 +29,7 @@</span> package core
 import (
 	"runtime"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // senderCacher is a concurrent transaction sender recoverer and cacher.
<b>diff --git a/core/tx_journal.go b/core/tx_journal.go</b>
<b>index b2bfa53..17b4419 100644</b>
<b>--- a/core/tx_journal.go</b>
<b>+++ b/core/tx_journal.go</b>
<span style="color:#0AA">@@ -31,10 +31,10 @@</span> import (
 	"io"
 	"os"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // errNoActiveJournal is returned if a transaction is attempted to be inserted
<b>diff --git a/core/tx_list.go b/core/tx_list.go</b>
<b>index 6788bad..2e28237 100644</b>
<b>--- a/core/tx_list.go</b>
<b>+++ b/core/tx_list.go</b>
<span style="color:#0AA">@@ -33,8 +33,8 @@</span> import (
 	"sort"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // nonceHeap is a heap.Interface implementation over 64bit unsigned integers for
<b>diff --git a/core/tx_list_test.go b/core/tx_list_test.go</b>
<b>index 86cf79a..91e4d6e 100644</b>
<b>--- a/core/tx_list_test.go</b>
<b>+++ b/core/tx_list_test.go</b>
<span style="color:#0AA">@@ -31,8 +31,8 @@</span> import (
 	"math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // Tests that transactions can be added to strict lists and list contents and
<b>diff --git a/core/tx_noncer.go b/core/tx_noncer.go</b>
<b>index bb5968b..439b9c7 100644</b>
<b>--- a/core/tx_noncer.go</b>
<b>+++ b/core/tx_noncer.go</b>
<span style="color:#0AA">@@ -29,8 +29,8 @@</span> package core
 import (
 	"sync"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
 )
 
 // txNoncer is a tiny virtual state database to manage the executable nonces of
<b>diff --git a/core/tx_pool.go b/core/tx_pool.go</b>
<b>index 457be8f..9663b17 100644</b>
<b>--- a/core/tx_pool.go</b>
<b>+++ b/core/tx_pool.go</b>
<span style="color:#0AA">@@ -35,15 +35,15 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/prque"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/metrics"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 const (
<b>diff --git a/core/tx_pool_test.go b/core/tx_pool_test.go</b>
<b>index 84353e9..97e6c7e 100644</b>
<b>--- a/core/tx_pool_test.go</b>
<b>+++ b/core/tx_pool_test.go</b>
<span style="color:#0AA">@@ -29,14 +29,14 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/core/types.go b/core/types.go</b>
<b>index 9607212..6a0d938 100644</b>
<b>--- a/core/types.go</b>
<b>+++ b/core/types.go</b>
<span style="color:#0AA">@@ -27,9 +27,9 @@</span>
 package core
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 )
 
 // Validator is an interface which defines the standard for block validation. It
<b>diff --git a/core/types/block_test.go b/core/types/block_test.go</b>
<b>index e2f5623..85db341 100644</b>
<b>--- a/core/types/block_test.go</b>
<b>+++ b/core/types/block_test.go</b>
<span style="color:#0AA">@@ -33,11 +33,11 @@</span> import (
 	"reflect"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 	"golang.org/x/crypto/sha3"
 )
 
<b>diff --git a/core/types/hashing_test.go b/core/types/hashing_test.go</b>
<b>index 424798d..e537e47 100644</b>
<b>--- a/core/types/hashing_test.go</b>
<b>+++ b/core/types/hashing_test.go</b>
<span style="color:#0AA">@@ -34,12 +34,12 @@</span> import (
 	mrand "math/rand"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 func TestDeriveSha(t *testing.T) {
<b>diff --git a/core/types/receipt.go b/core/types/receipt.go</b>
<b>index 6da7c15..90366de 100644</b>
<b>--- a/core/types/receipt.go</b>
<b>+++ b/core/types/receipt.go</b>
<span style="color:#0AA">@@ -34,11 +34,11 @@</span> import (
 	"math/big"
 	"unsafe"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 //go:generate gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go
<b>diff --git a/core/types/receipt_test.go b/core/types/receipt_test.go</b>
<b>index 95fe68a..d2a2508 100644</b>
<b>--- a/core/types/receipt_test.go</b>
<b>+++ b/core/types/receipt_test.go</b>
<span style="color:#0AA">@@ -33,10 +33,10 @@</span> import (
 	"reflect"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestDecodeEmptyTypedReceipt(t *testing.T) {
<b>diff --git a/core/types/transaction_signing.go b/core/types/transaction_signing.go</b>
<b>index a717749..982d10d 100644</b>
<b>--- a/core/types/transaction_signing.go</b>
<b>+++ b/core/types/transaction_signing.go</b>
<span style="color:#0AA">@@ -32,9 +32,9 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var ErrInvalidChainId = errors.New("invalid chain id for signer")
<b>diff --git a/core/vm/access_list_tracer.go b/core/vm/access_list_tracer.go</b>
<b>index c481c54..22c8f3e 100644</b>
<b>--- a/core/vm/access_list_tracer.go</b>
[1m+++ b/core/vm/access_list_tracer.go
<span style="color:#0AA">@@ -30,8 +30,8 @@</span> import (
 	"math/big"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // accessList is an accumulator for the set of accounts and storage slots an EVM
<b>diff --git a/core/vm/contracts.go b/core/vm/contracts.go</b>
<b>index cfb2671..31b8d8f 100644</b>
<b>--- a/core/vm/contracts.go</b>
<b>+++ b/core/vm/contracts.go</b>
<span style="color:#0AA">@@ -32,13 +32,13 @@</span> import (
 	"errors"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/crypto/blake2b"
 	"github.com/ethereum/go-ethereum/crypto/bls12381"
 	"github.com/ethereum/go-ethereum/crypto/bn256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
 	//lint:ignore SA1019 Needed for precompile
 	"golang.org/x/crypto/ripemd160"
<b>diff --git a/core/vm/contracts_stateful.go b/core/vm/contracts_stateful.go</b>
<b>index 981ebd7..05b0361 100644</b>
<b>--- a/core/vm/contracts_stateful.go</b>
<b>+++ b/core/vm/contracts_stateful.go</b>
<span style="color:#0AA">@@ -7,9 +7,9 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // PrecompiledContractsApricot contains the default set of pre-compiled Ethereum
<b>diff --git a/core/vm/contracts_stateful_test.go b/core/vm/contracts_stateful_test.go</b>
<b>index 6126078..1278a67 100644</b>
<b>--- a/core/vm/contracts_stateful_test.go</b>
<b>+++ b/core/vm/contracts_stateful_test.go</b>
<span style="color:#0AA">@@ -7,12 +7,12 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/stretchr/testify/assert"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestPrecompiledContractSpendsGas(t *testing.T) {
<b>diff --git a/core/vm/eips.go b/core/vm/eips.go</b>
<b>index e525a73..399a874 100644</b>
<b>--- a/core/vm/eips.go</b>
<b>+++ b/core/vm/eips.go</b>
<span style="color:#0AA">@@ -30,8 +30,8 @@</span> import (
 	"fmt"
 	"sort"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var activators = map[int]func(*JumpTable){
<b>diff --git a/core/vm/evm.go b/core/vm/evm.go</b>
<b>index 9ad9db1..6f78cc6 100644</b>
<b>--- a/core/vm/evm.go</b>
<b>+++ b/core/vm/evm.go</b>
<span style="color:#0AA">@@ -31,10 +31,10 @@</span> import (
 	"sync/atomic"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // emptyCodeHash is used by create to ensure deployment is disallowed to already
<b>diff --git a/core/vm/gas_table.go b/core/vm/gas_table.go</b>
<b>index 383fb6d..d611fc5 100644</b>
<b>--- a/core/vm/gas_table.go</b>
<b>+++ b/core/vm/gas_table.go</b>
<span style="color:#0AA">@@ -29,9 +29,9 @@</span> package vm
 import (
 	"errors"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // memoryGasCost calculates the quadratic gas for memory expansion. It does so
<b>diff --git a/core/vm/gas_table_test.go b/core/vm/gas_table_test.go</b>
<b>index 92d5d30..2150bb2 100644</b>
<b>--- a/core/vm/gas_table_test.go</b>
<b>+++ b/core/vm/gas_table_test.go</b>
<span style="color:#0AA">@@ -31,11 +31,11 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestMemoryGasCost(t *testing.T) {
<b>diff --git a/core/vm/instructions.go b/core/vm/instructions.go</b>
<b>index 81f3d3e..3fe1c4b 100644</b>
<b>--- a/core/vm/instructions.go</b>
<b>+++ b/core/vm/instructions.go</b>
<span style="color:#0AA">@@ -29,10 +29,10 @@</span> package vm
 import (
 	"errors"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 	"golang.org/x/crypto/sha3"
 )
 
<b>diff --git a/core/vm/instructions_test.go b/core/vm/instructions_test.go</b>
<b>index 1570b36..50fb6db 100644</b>
<b>--- a/core/vm/instructions_test.go</b>
<b>+++ b/core/vm/instructions_test.go</b>
<span style="color:#0AA">@@ -33,10 +33,10 @@</span> import (
 	"io/ioutil"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type TwoOperandTestcase struct {
<b>diff --git a/core/vm/interface.go b/core/vm/interface.go</b>
<b>index bde4b08..6148970 100644</b>
<b>--- a/core/vm/interface.go</b>
<b>+++ b/core/vm/interface.go</b>
<span style="color:#0AA">@@ -29,8 +29,8 @@</span> package vm
 import (
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // StateDB is an EVM database for full state querying.
<b>diff --git a/core/vm/jump_table.go b/core/vm/jump_table.go</b>
<b>index d833e52..249736d 100644</b>
<b>--- a/core/vm/jump_table.go</b>
<b>+++ b/core/vm/jump_table.go</b>
<span style="color:#0AA">@@ -27,7 +27,7 @@</span>
 package vm
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type (
<b>diff --git a/core/vm/logger.go b/core/vm/logger.go</b>
<b>index f8d088d..e49d4f7 100644</b>
<b>--- a/core/vm/logger.go</b>
<b>+++ b/core/vm/logger.go</b>
<span style="color:#0AA">@@ -34,12 +34,12 @@</span> import (
 	"strings"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // Storage represents a contract's storage.
<b>diff --git a/core/vm/logger_test.go b/core/vm/logger_test.go</b>
<b>index 1d87e46..a35c4e6 100644</b>
<b>--- a/core/vm/logger_test.go</b>
<b>+++ b/core/vm/logger_test.go</b>
<span style="color:#0AA">@@ -30,10 +30,10 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/holiman/uint256"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type dummyContractRef struct {
<b>diff --git a/core/vm/operations_acl.go b/core/vm/operations_acl.go</b>
<b>index adfe772..d9e2f06 100644</b>
<b>--- a/core/vm/operations_acl.go</b>
<b>+++ b/core/vm/operations_acl.go</b>
<span style="color:#0AA">@@ -29,9 +29,9 @@</span> package vm
 import (
 	"errors"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/math"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // gasSStoreEIP2929 implements gas cost for SSTORE according to EIP-2929
<b>diff --git a/core/vm/runtime/env.go b/core/vm/runtime/env.go</b>
<b>index 5cae37d..6ebd9c6 100644</b>
<b>--- a/core/vm/runtime/env.go</b>
<b>+++ b/core/vm/runtime/env.go</b>
<span style="color:#0AA">@@ -27,8 +27,8 @@</span>
 package runtime
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 )
 
 func NewEnv(cfg *Config) *vm.EVM {
<b>diff --git a/core/vm/runtime/runtime.go b/core/vm/runtime/runtime.go</b>
<b>index eb6e3f6..bd24dba 100644</b>
<b>--- a/core/vm/runtime/runtime.go</b>
<b>+++ b/core/vm/runtime/runtime.go</b>
<span style="color:#0AA">@@ -31,12 +31,12 @@</span> import (
 	"math/big"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // Config is a basic type specifying certain configuration flags for running
<b>diff --git a/core/vm/runtime/runtime_example_test.go b/core/vm/runtime/runtime_example_test.go</b>
<b>index 9850e28..c126459 100644</b>
<b>--- a/core/vm/runtime/runtime_example_test.go</b>
<b>+++ b/core/vm/runtime/runtime_example_test.go</b>
<span style="color:#0AA">@@ -29,8 +29,8 @@</span> package runtime_test
 import (
 	"fmt"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm/runtime"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm/runtime"</span>
 )
 
 func ExampleExecute() {
<b>diff --git a/core/vm/runtime/runtime_test.go b/core/vm/runtime/runtime_test.go</b>
<b>index 4b8d0a9..8561e31 100644</b>
<b>--- a/core/vm/runtime/runtime_test.go</b>
<b>+++ b/core/vm/runtime/runtime_test.go</b>
<span style="color:#0AA">@@ -34,16 +34,16 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/accounts/abi"
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/core/asm"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestDefaults(t *testing.T) {
<b>diff --git a/core/vm/stack_table.go b/core/vm/stack_table.go</b>
<b>index 487acae..e75f21a 100644</b>
<b>--- a/core/vm/stack_table.go</b>
<b>+++ b/core/vm/stack_table.go</b>
<span style="color:#0AA">@@ -27,7 +27,7 @@</span>
 package vm
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func minSwapStack(n int) int {
<b>diff --git a/eth/api.go b/eth/api.go</b>
<b>index 62dd67b..f0d2bce 100644</b>
<b>--- a/eth/api.go</b>
<b>+++ b/eth/api.go</b>
<span style="color:#0AA">@@ -35,16 +35,16 @@</span> import (
 	"os"
 	"strings"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/internal/ethapi"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/internal/ethapi"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // PublicEthereumAPI provides an API to access Ethereum full node-related
<b>diff --git a/eth/api_backend.go b/eth/api_backend.go</b>
<b>index 5152b20..29e81d9 100644</b>
<b>--- a/eth/api_backend.go</b>
<b>+++ b/eth/api_backend.go</b>
<span style="color:#0AA">@@ -32,21 +32,21 @@</span> import (
 	"math/big"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/bloombits"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/gasprice"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/eth/downloader"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/bloombits"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/gasprice"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 var (
<b>diff --git a/eth/backend.go b/eth/backend.go</b>
<b>index 380548e..8c32ac2 100644</b>
<b>--- a/eth/backend.go</b>
<b>+++ b/eth/backend.go</b>
<span style="color:#0AA">@@ -33,28 +33,28 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/bloombits"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/ethconfig"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/filters"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/gasprice"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/tracers"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/internal/ethapi"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/miner"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/eth/downloader"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/bloombits"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/ethconfig"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/filters"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/gasprice"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/tracers"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/internal/ethapi"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/miner"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Config contains the configuration options of the ETH protocol.
<b>diff --git a/eth/bloombits.go b/eth/bloombits.go</b>
<b>index ecc0aaf..b45ed5b 100644</b>
<b>--- a/eth/bloombits.go</b>
<b>+++ b/eth/bloombits.go</b>
<span style="color:#0AA">@@ -29,8 +29,8 @@</span> package eth
 import (
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
 	"github.com/ethereum/go-ethereum/common/bitutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
 )
 
 const (
<b>diff --git a/eth/ethconfig/config.go b/eth/ethconfig/config.go</b>
<b>index 99ea015..bbfa10b 100644</b>
<b>--- a/eth/ethconfig/config.go</b>
<b>+++ b/eth/ethconfig/config.go</b>
<span style="color:#0AA">@@ -29,10 +29,10 @@</span> package ethconfig
 import (
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/gasprice"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/miner"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/gasprice"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/miner"</span>
 )
 
 // DefaultFullGPOConfig contains default gasprice oracle settings for full node.
<b>diff --git a/eth/filters/api.go b/eth/filters/api.go</b>
<b>index f543d6e..50339df 100644</b>
<b>--- a/eth/filters/api.go</b>
<b>+++ b/eth/filters/api.go</b>
<span style="color:#0AA">@@ -35,13 +35,13 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // filter is a helper struct that holds meta information over the filter type
<b>diff --git a/eth/filters/api_test.go b/eth/filters/api_test.go</b>
<b>index 6f35639..2903c0b 100644</b>
<b>--- a/eth/filters/api_test.go</b>
<b>+++ b/eth/filters/api_test.go</b>
<span style="color:#0AA">@@ -21,8 +21,8 @@</span> import (
 	"fmt"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 func TestUnmarshalJSONNewFilterArgs(t *testing.T) {
<b>diff --git a/eth/filters/filter.go b/eth/filters/filter.go</b>
<b>index 94769a1..e8f9ee7 100644</b>
<b>--- a/eth/filters/filter.go</b>
<b>+++ b/eth/filters/filter.go</b>
<span style="color:#0AA">@@ -32,15 +32,15 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/bloombits"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/bloombits"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 type Backend interface {
<b>diff --git a/eth/filters/filter_system.go b/eth/filters/filter_system.go</b>
<b>index 86f5a0b..db72584 100644</b>
<b>--- a/eth/filters/filter_system.go</b>
<b>+++ b/eth/filters/filter_system.go</b>
<span style="color:#0AA">@@ -34,14 +34,14 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Type determines the kind of filter and is used to put the filter in to
<b>diff --git a/eth/gasprice/gasprice.go b/eth/gasprice/gasprice.go</b>
<b>index b1d415b..954b124 100644</b>
<b>--- a/eth/gasprice/gasprice.go</b>
<b>+++ b/eth/gasprice/gasprice.go</b>
<span style="color:#0AA">@@ -33,12 +33,12 @@</span> import (
 	"sync"
 
 	"github.com/ava-labs/avalanchego/utils/timer"
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 const sampleNumber = 3 // Number of transactions sampled in a block
<b>diff --git a/eth/gasprice/gasprice_test.go b/eth/gasprice/gasprice_test.go</b>
<b>index cfea2d4..97733eb 100644</b>
<b>--- a/eth/gasprice/gasprice_test.go</b>
<b>+++ b/eth/gasprice/gasprice_test.go</b>
<span style="color:#0AA">@@ -32,15 +32,15 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 type testBackend struct {
<b>diff --git a/eth/state_accessor.go b/eth/state_accessor.go</b>
<b>index e94b2d4..adad3d1 100644</b>
<b>--- a/eth/state_accessor.go</b>
<b>+++ b/eth/state_accessor.go</b>
<span style="color:#0AA">@@ -32,13 +32,13 @@</span> import (
 	"math/big"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 )
 
 // stateAtBlock retrieves the state database associated with a certain block.
<b>diff --git a/eth/tracers/api.go b/eth/tracers/api.go</b>
<b>index c54cd93..46e7505 100644</b>
<b>--- a/eth/tracers/api.go</b>
<b>+++ b/eth/tracers/api.go</b>
<span style="color:#0AA">@@ -36,19 +36,19 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/internal/ethapi"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/internal/ethapi"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 const (
<b>diff --git a/eth/tracers/api_test.go b/eth/tracers/api_test.go</b>
<b>index 454f35c..4122c90 100644</b>
<b>--- a/eth/tracers/api_test.go</b>
<b>+++ b/eth/tracers/api_test.go</b>
<span style="color:#0AA">@@ -38,20 +38,20 @@</span> import (
 	"sort"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/internal/ethapi"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/internal/ethapi"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 var (
<b>diff --git a/eth/tracers/tracer.go b/eth/tracers/tracer.go</b>
<b>index 3f4325a..4670281 100644</b>
<b>--- a/eth/tracers/tracer.go</b>
<b>+++ b/eth/tracers/tracer.go</b>
<span style="color:#0AA">@@ -35,12 +35,12 @@</span> import (
 	"time"
 	"unsafe"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
 	"gopkg.in/olebedev/go-duktape.v3"
 )
 
<b>diff --git a/eth/tracers/tracer_test.go b/eth/tracers/tracer_test.go</b>
<b>index 3d4e94c..e8ce389 100644</b>
<b>--- a/eth/tracers/tracer_test.go</b>
<b>+++ b/eth/tracers/tracer_test.go</b>
<span style="color:#0AA">@@ -33,10 +33,10 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type account struct{}
<b>diff --git a/eth/tracers/tracers.go b/eth/tracers/tracers.go</b>
<b>index dbe53f3..049cb18 100644</b>
<b>--- a/eth/tracers/tracers.go</b>
<b>+++ b/eth/tracers/tracers.go</b>
<span style="color:#0AA">@@ -31,7 +31,7 @@</span> import (
 	"strings"
 	"unicode"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/tracers/internal/tracers"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/tracers/internal/tracers"</span>
 )
 
 // all contains all the built in JavaScript tracers by name.
<b>diff --git a/eth/tracers/tracers_test.go b/eth/tracers/tracers_test.go</b>
<b>index f9e55e2..ded3c86 100644</b>
<b>--- a/eth/tracers/tracers_test.go</b>
<b>+++ b/eth/tracers/tracers_test.go</b>
<span style="color:#0AA">@@ -37,17 +37,17 @@</span> import (
 	"strings"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/rawdb"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/tests"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/rawdb"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/tests"</span>
 )
 
 // To generate a new callTracer test, copy paste the makeTest method below into
<b>diff --git a/ethclient/corethclient/corethclient.go b/ethclient/corethclient/corethclient.go</b>
<b>index 8d7654f..da9ca7f 100644</b>
<b>--- a/ethclient/corethclient/corethclient.go</b>
<b>+++ b/ethclient/corethclient/corethclient.go</b>
<span style="color:#0AA">@@ -33,12 +33,12 @@</span> import (
 	"runtime"
 	"runtime/debug"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/ethclient"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/ethclient"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Client is a wrapper around rpc.Client that implements geth-specific functionality.
<b>diff --git a/ethclient/ethclient.go b/ethclient/ethclient.go</b>
<b>index 0959a0a..8d30ca2 100644</b>
<b>--- a/ethclient/ethclient.go</b>
<b>+++ b/ethclient/ethclient.go</b>
<span style="color:#0AA">@@ -35,11 +35,11 @@</span> import (
 	"math/big"
 
 	"github.com/ava-labs/avalanchego/ids"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/interfaces"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/interfaces"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Client defines typed wrappers for the Ethereum RPC API.
<b>diff --git a/ethclient/signer.go b/ethclient/signer.go</b>
<b>index dafa943..45ef52d 100644</b>
<b>--- a/ethclient/signer.go</b>
<b>+++ b/ethclient/signer.go</b>
<span style="color:#0AA">@@ -30,8 +30,8 @@</span> import (
 	"errors"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // senderFromServer is a types.Signer that remembers the sender address returned by the RPC
<b>diff --git a/go.mod b/go.mod</b>
<b>index 981a551..fac1e12 100644</b>
<b>--- a/go.mod</b>
<b>+++ b/go.mod</b>
<span style="color:#0AA">@@ -1,4 +1,4 @@</span>
<span style="color:#A00">-module github.com/ava-labs/coreth</span>
<span style="color:#0A0">+</span><span style="color:#0A0">module gitlab.com/flarenetwork/coreth</span>
 
 go 1.15
 
<b>diff --git a/interfaces/interfaces.go b/interfaces/interfaces.go</b>
<b>index efcde08..be8870f 100644</b>
<b>--- a/interfaces/interfaces.go</b>
<b>+++ b/interfaces/interfaces.go</b>
<span style="color:#0AA">@@ -32,8 +32,8 @@</span> import (
 	"errors"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 // NotFound is returned by API methods if the requested item does not exist.
<b>diff --git a/internal/ethapi/api.go b/internal/ethapi/api.go</b>
<b>index b9ad558..e110a19 100644</b>
<b>--- a/internal/ethapi/api.go</b>
<b>+++ b/internal/ethapi/api.go</b>
<span style="color:#0AA">@@ -35,15 +35,6 @@</span> import (
 	"time"
 
 	"github.com/ava-labs/avalanchego/ids"
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/scwallet"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/davecgh/go-spew/spew"
 	"github.com/ethereum/go-ethereum/accounts/abi"
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#0AA">@@ -53,6 +44,15 @@</span> import (
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
 	"github.com/tyler-smith/go-bip39"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/scwallet"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // PublicEthereumAPI provides an API to access Ethereum related information.
<b>diff --git a/internal/ethapi/backend.go b/internal/ethapi/backend.go</b>
<b>index 3d03734..f3b1643 100644</b>
<b>--- a/internal/ethapi/backend.go</b>
<b>+++ b/internal/ethapi/backend.go</b>
<span style="color:#0AA">@@ -31,19 +31,19 @@</span> import (
 	"context"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/bloombits"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/eth/downloader"
 	"github.com/ethereum/go-ethereum/ethdb"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/bloombits"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Backend interface provides the common API services (that are provided by
<b>diff --git a/internal/ethapi/transaction_args.go b/internal/ethapi/transaction_args.go</b>
<b>index 6b06742..dcadfd7 100644</b>
<b>--- a/internal/ethapi/transaction_args.go</b>
<b>+++ b/internal/ethapi/transaction_args.go</b>
<span style="color:#0AA">@@ -33,12 +33,12 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // TransactionArgs represents the arguments to construct a new transaction
<b>diff --git a/miner/miner.go b/miner/miner.go</b>
<b>index 263439c..24ddc22 100644</b>
<b>--- a/miner/miner.go</b>
<b>+++ b/miner/miner.go</b>
<span style="color:#0AA">@@ -28,12 +28,12 @@</span>
 package miner
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/event"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // Backend wraps all methods required for mining.
<b>diff --git a/miner/worker.go b/miner/worker.go</b>
<b>index c8c7268..69fa1f6 100644</b>
<b>--- a/miner/worker.go</b>
<b>+++ b/miner/worker.go</b>
<span style="color:#0AA">@@ -36,16 +36,16 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/misc"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/event"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/misc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // environment is the worker's current environment and holds all of the current state information.
<b>diff --git a/node/api.go b/node/api.go</b>
<b>index c38b2df..0c36fc7 100644</b>
<b>--- a/node/api.go</b>
<b>+++ b/node/api.go</b>
<span style="color:#0AA">@@ -27,10 +27,10 @@</span>
 package node
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/internal/debug"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/crypto"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/internal/debug"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // apis returns the collection of built-in RPC APIs.
<b>diff --git a/node/config.go b/node/config.go</b>
<b>index 4d2b83a..b51e67a 100644</b>
<b>--- a/node/config.go</b>
<b>+++ b/node/config.go</b>
<span style="color:#0AA">@@ -32,12 +32,12 @@</span> import (
 	"os"
 	"path/filepath"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/external"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/external"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Config represents a small collection of configuration values to fine tune the
<b>diff --git a/node/defaults.go b/node/defaults.go</b>
<b>index e4c826b..e3e2257 100644</b>
<b>--- a/node/defaults.go</b>
<b>+++ b/node/defaults.go</b>
<span style="color:#0AA">@@ -27,7 +27,7 @@</span>
 package node
 
 import (
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 const (
<b>diff --git a/node/node.go b/node/node.go</b>
<b>index 265e00d..1272a9f 100644</b>
<b>--- a/node/node.go</b>
<b>+++ b/node/node.go</b>
<span style="color:#0AA">@@ -31,8 +31,8 @@</span> import (
 
 	"github.com/ethereum/go-ethereum/event"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/accounts"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/accounts"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 )
 
 // Node is a container on which services can be registered.
<b>diff --git a/plugin/evm/block.go b/plugin/evm/block.go</b>
<b>index cae83af..d749229 100644</b>
<b>--- a/plugin/evm/block.go</b>
<b>+++ b/plugin/evm/block.go</b>
<span style="color:#0AA">@@ -12,8 +12,8 @@</span> import (
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
 	"github.com/ava-labs/avalanchego/ids"
 	"github.com/ava-labs/avalanchego/snow/choices"
<b>diff --git a/plugin/evm/block_verification.go b/plugin/evm/block_verification.go</b>
<b>index 03910af..1e52707 100644</b>
<b>--- a/plugin/evm/block_verification.go</b>
<b>+++ b/plugin/evm/block_verification.go</b>
<span style="color:#0AA">@@ -7,11 +7,11 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	coreth "github.com/ava-labs/coreth/chain"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/trie"
<span style="color:#0A0">+</span>	<span style="color:#0A0">coreth "gitlab.com/flarenetwork/coreth/chain"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 var (
<b>diff --git a/plugin/evm/config.go b/plugin/evm/config.go</b>
<b>index f184a09..721bf75 100644</b>
<b>--- a/plugin/evm/config.go</b>
<b>+++ b/plugin/evm/config.go</b>
<span style="color:#0AA">@@ -7,8 +7,8 @@</span> import (
 	"encoding/json"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
 	"github.com/spf13/cast"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
 )
 
 const (
<b>diff --git a/plugin/evm/export_tx.go b/plugin/evm/export_tx.go</b>
<b>index 3b50915..be2275f 100644</b>
<b>--- a/plugin/evm/export_tx.go</b>
<b>+++ b/plugin/evm/export_tx.go</b>
<span style="color:#0AA">@@ -1,5 +1,12 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// This file is a derived work, based on the avalanchego library whose original</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// notice appears below. It is distributed under a license compatible with the</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// licensing terms of the original code from which it is derived.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE_AVALABS for licensing terms of the original work.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
 // (c) 2019-2020, Ava Labs, Inc. All rights reserved.
<span style="color:#A00">-// See the file LICENSE for licensing terms.</span>
 
 package evm
 
<span style="color:#0AA">@@ -7,19 +14,14 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/chains/atomic"</span>
 	"github.com/ava-labs/avalanchego/database"
 	"github.com/ava-labs/avalanchego/ids"
 	"github.com/ava-labs/avalanchego/snow"
 	"github.com/ava-labs/avalanchego/utils/crypto"
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/utils/math"</span>
 	"github.com/ava-labs/avalanchego/vms/components/avax"
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/vms/secp256k1fx"</span>
<span style="color:#A00">-	"github.com/ethereum/go-ethereum/common"</span>
<span style="color:#A00">-	"github.com/ethereum/go-ethereum/log"</span>
 )
 
 // UnsignedExportTx is an unsigned ExportTx
<span style="color:#0AA">@@ -46,75 +48,16 @@</span> func (tx *UnsignedExportTx) Verify(
 	ctx *snow.Context,
 	rules params.Rules,
 ) error {
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	case tx == nil:</span>
<span style="color:#A00">-		return errNilTx</span>
<span style="color:#A00">-	case tx.DestinationChain != avmID:</span>
<span style="color:#A00">-		return errWrongChainID</span>
<span style="color:#A00">-	case len(tx.ExportedOutputs) == 0:</span>
<span style="color:#A00">-		return errNoExportOutputs</span>
<span style="color:#A00">-	case tx.NetworkID != ctx.NetworkID:</span>
<span style="color:#A00">-		return errWrongNetworkID</span>
<span style="color:#A00">-	case ctx.ChainID != tx.BlockchainID:</span>
<span style="color:#A00">-		return errWrongBlockchainID</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for _, in := range tx.Ins {</span>
<span style="color:#A00">-		if err := in.Verify(); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for _, out := range tx.ExportedOutputs {</span>
<span style="color:#A00">-		if err := out.Verify(); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	if !avax.IsSortedTransferableOutputs(tx.ExportedOutputs, Codec) {</span>
<span style="color:#A00">-		return errOutputsNotSorted</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	if rules.IsApricotPhase1 && !IsSortedAndUniqueEVMInputs(tx.Ins) {</span>
<span style="color:#A00">-		return errInputsNotSortedUnique</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return errWrongChainID</span>
 }
 
 func (tx *UnsignedExportTx) Cost() (uint64, error) {
<span style="color:#A00">-	byteCost := calcBytesCost(len(tx.UnsignedBytes()))</span>
<span style="color:#A00">-	numSigs := uint64(len(tx.Ins))</span>
<span style="color:#A00">-	sigCost, err := math.Mul64(numSigs, secp256k1fx.CostPerSignature)</span>
<span style="color:#A00">-	if err != nil {</span>
<span style="color:#A00">-		return 0, err</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return math.Add64(byteCost, sigCost)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return 0, fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // Amount of [assetID] burned by this transaction
 func (tx *UnsignedExportTx) Burned(assetID ids.ID) (uint64, error) {
<span style="color:#A00">-	var (</span>
<span style="color:#A00">-		spent uint64</span>
<span style="color:#A00">-		input uint64</span>
<span style="color:#A00">-		err   error</span>
<span style="color:#A00">-	)</span>
<span style="color:#A00">-	for _, out := range tx.ExportedOutputs {</span>
<span style="color:#A00">-		if out.AssetID() == assetID {</span>
<span style="color:#A00">-			spent, err = math.Add64(spent, out.Output().Amount())</span>
<span style="color:#A00">-			if err != nil {</span>
<span style="color:#A00">-				return 0, err</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, in := range tx.Ins {</span>
<span style="color:#A00">-		if in.AssetID == assetID {</span>
<span style="color:#A00">-			input, err = math.Add64(input, in.Amount)</span>
<span style="color:#A00">-			if err != nil {</span>
<span style="color:#A00">-				return 0, err</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return math.Sub64(input, spent)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return 0, fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // SemanticVerify this transaction is valid.
<span style="color:#0AA">@@ -125,105 +68,12 @@</span> func (tx *UnsignedExportTx) SemanticVerify(
 	baseFee *big.Int,
 	rules params.Rules,
 ) error {
<span style="color:#A00">-	if err := tx.Verify(vm.ctx.XChainID, vm.ctx, rules); err != nil {</span>
<span style="color:#A00">-		return err</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Check the transaction consumes and produces the right amounts</span>
<span style="color:#A00">-	fc := avax.NewFlowChecker()</span>
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	// Apply dynamic fees to export transactions as of Apricot Phase 3</span>
<span style="color:#A00">-	case rules.IsApricotPhase3:</span>
<span style="color:#A00">-		cost, err := stx.Cost()</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		txFee, err := calculateDynamicFee(cost, baseFee)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		fc.Produce(vm.ctx.AVAXAssetID, txFee)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Apply fees to export transactions before Apricot Phase 3</span>
<span style="color:#A00">-	default:</span>
<span style="color:#A00">-		fc.Produce(vm.ctx.AVAXAssetID, params.AvalancheAtomicTxFee)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, out := range tx.ExportedOutputs {</span>
<span style="color:#A00">-		fc.Produce(out.AssetID(), out.Output().Amount())</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, in := range tx.Ins {</span>
<span style="color:#A00">-		fc.Consume(in.AssetID, in.Amount)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if err := fc.Verify(); err != nil {</span>
<span style="color:#A00">-		return fmt.Errorf("export tx flow check failed due to: %w", err)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if len(tx.Ins) != len(stx.Creds) {</span>
<span style="color:#A00">-		return fmt.Errorf("export tx contained mismatched number of inputs/credentials (%d vs. %d)", len(tx.Ins), len(stx.Creds))</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for i, input := range tx.Ins {</span>
<span style="color:#A00">-		cred, ok := stx.Creds[i].(*secp256k1fx.Credential)</span>
<span style="color:#A00">-		if !ok {</span>
<span style="color:#A00">-			return fmt.Errorf("expected *secp256k1fx.Credential but got %T", cred)</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		if err := cred.Verify(); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		if len(cred.Sigs) != 1 {</span>
<span style="color:#A00">-			return fmt.Errorf("expected one signature for EVM Input Credential, but found: %d", len(cred.Sigs))</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		pubKeyIntf, err := vm.secpFactory.RecoverPublicKey(tx.UnsignedBytes(), cred.Sigs[0][:])</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		pubKey, ok := pubKeyIntf.(*crypto.PublicKeySECP256K1R)</span>
<span style="color:#A00">-		if !ok {</span>
<span style="color:#A00">-			// This should never happen</span>
<span style="color:#A00">-			return fmt.Errorf("expected *crypto.PublicKeySECP256K1R but got %T", pubKeyIntf)</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		if input.Address != PublicKeyToEthAddress(pubKey) {</span>
<span style="color:#A00">-			return errPublicKeySignatureMismatch</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // Accept this transaction.
 func (tx *UnsignedExportTx) Accept(ctx *snow.Context, batch database.Batch) error {
<span style="color:#A00">-	txID := tx.ID()</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	elems := make([]*atomic.Element, len(tx.ExportedOutputs))</span>
<span style="color:#A00">-	for i, out := range tx.ExportedOutputs {</span>
<span style="color:#A00">-		utxo := &avax.UTXO{</span>
<span style="color:#A00">-			UTXOID: avax.UTXOID{</span>
<span style="color:#A00">-				TxID:        txID,</span>
<span style="color:#A00">-				OutputIndex: uint32(i),</span>
<span style="color:#A00">-			},</span>
<span style="color:#A00">-			Asset: avax.Asset{ID: out.AssetID()},</span>
<span style="color:#A00">-			Out:   out.Out,</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		utxoBytes, err := Codec.Marshal(codecVersion, utxo)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		utxoID := utxo.InputID()</span>
<span style="color:#A00">-		elem := &atomic.Element{</span>
<span style="color:#A00">-			Key:   utxoID[:],</span>
<span style="color:#A00">-			Value: utxoBytes,</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		if out, ok := utxo.Out.(avax.Addressable); ok {</span>
<span style="color:#A00">-			elem.Traits = out.Addresses()</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		elems[i] = elem</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{tx.DestinationChain: {PutRequests: elems}}, batch)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // newExportTx returns a new ExportTx
<span style="color:#0AA">@@ -235,121 +85,10 @@</span> func (vm *VM) newExportTx(
 	baseFee *big.Int, // fee to use post-AP3
 	keys []*crypto.PrivateKeySECP256K1R, // Pay the fee and provide the tokens
 ) (*Tx, error) {
<span style="color:#A00">-	if vm.ctx.XChainID != chainID {</span>
<span style="color:#A00">-		return nil, errWrongChainID</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	outs := []*avax.TransferableOutput{{ // Exported to X-Chain</span>
<span style="color:#A00">-		Asset: avax.Asset{ID: assetID},</span>
<span style="color:#A00">-		Out: &secp256k1fx.TransferOutput{</span>
<span style="color:#A00">-			Amt: amount,</span>
<span style="color:#A00">-			OutputOwners: secp256k1fx.OutputOwners{</span>
<span style="color:#A00">-				Locktime:  0,</span>
<span style="color:#A00">-				Threshold: 1,</span>
<span style="color:#A00">-				Addrs:     []ids.ShortID{to},</span>
<span style="color:#A00">-			},</span>
<span style="color:#A00">-		},</span>
<span style="color:#A00">-	}}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	var (</span>
<span style="color:#A00">-		avaxNeeded           uint64 = 0</span>
<span style="color:#A00">-		ins, avaxIns         []EVMInput</span>
<span style="color:#A00">-		signers, avaxSigners [][]*crypto.PrivateKeySECP256K1R</span>
<span style="color:#A00">-		err                  error</span>
<span style="color:#A00">-	)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// consume non-AVAX</span>
<span style="color:#A00">-	if assetID != vm.ctx.AVAXAssetID {</span>
<span style="color:#A00">-		ins, signers, err = vm.GetSpendableFunds(keys, assetID, amount)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, fmt.Errorf("couldn't generate tx inputs/signers: %w", err)</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	} else {</span>
<span style="color:#A00">-		avaxNeeded = amount</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	rules := vm.currentRules()</span>
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	case rules.IsApricotPhase3:</span>
<span style="color:#A00">-		utx := &UnsignedExportTx{</span>
<span style="color:#A00">-			NetworkID:        vm.ctx.NetworkID,</span>
<span style="color:#A00">-			BlockchainID:     vm.ctx.ChainID,</span>
<span style="color:#A00">-			DestinationChain: chainID,</span>
<span style="color:#A00">-			Ins:              ins,</span>
<span style="color:#A00">-			ExportedOutputs:  outs,</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		tx := &Tx{UnsignedAtomicTx: utx}</span>
<span style="color:#A00">-		if err := tx.Sign(vm.codec, nil); err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		var cost uint64</span>
<span style="color:#A00">-		cost, err = tx.Cost()</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		avaxIns, avaxSigners, err = vm.GetSpendableAVAXWithFee(keys, avaxNeeded, cost, baseFee)</span>
<span style="color:#A00">-	default:</span>
<span style="color:#A00">-		var newAvaxNeeded uint64</span>
<span style="color:#A00">-		newAvaxNeeded, err = math.Add64(avaxNeeded, params.AvalancheAtomicTxFee)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, errOverflowExport</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		avaxIns, avaxSigners, err = vm.GetSpendableFunds(keys, vm.ctx.AVAXAssetID, newAvaxNeeded)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	if err != nil {</span>
<span style="color:#A00">-		return nil, fmt.Errorf("couldn't generate tx inputs/signers: %w", err)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	ins = append(ins, avaxIns...)</span>
<span style="color:#A00">-	signers = append(signers, avaxSigners...)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	SortEVMInputsAndSigners(ins, signers)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Create the transaction</span>
<span style="color:#A00">-	utx := &UnsignedExportTx{</span>
<span style="color:#A00">-		NetworkID:        vm.ctx.NetworkID,</span>
<span style="color:#A00">-		BlockchainID:     vm.ctx.ChainID,</span>
<span style="color:#A00">-		DestinationChain: chainID,</span>
<span style="color:#A00">-		Ins:              ins,</span>
<span style="color:#A00">-		ExportedOutputs:  outs,</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	tx := &Tx{UnsignedAtomicTx: utx}</span>
<span style="color:#A00">-	if err := tx.Sign(vm.codec, signers); err != nil {</span>
<span style="color:#A00">-		return nil, err</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return tx, utx.Verify(vm.ctx.XChainID, vm.ctx, vm.currentRules())</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil, errWrongChainID</span>
 }
 
 // EVMStateTransfer executes the state update from the atomic export transaction
 func (tx *UnsignedExportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
<span style="color:#A00">-	addrs := map[[20]byte]uint64{}</span>
<span style="color:#A00">-	for _, from := range tx.Ins {</span>
<span style="color:#A00">-		if from.AssetID == ctx.AVAXAssetID {</span>
<span style="color:#A00">-			log.Debug("crosschain C->X", "addr", from.Address, "amount", from.Amount, "assetID", "AVAX")</span>
<span style="color:#A00">-			// We multiply the input amount by x2cRate to convert AVAX back to the appropriate</span>
<span style="color:#A00">-			// denomination before export.</span>
<span style="color:#A00">-			amount := new(big.Int).Mul(</span>
<span style="color:#A00">-				new(big.Int).SetUint64(from.Amount), x2cRate)</span>
<span style="color:#A00">-			if state.GetBalance(from.Address).Cmp(amount) < 0 {</span>
<span style="color:#A00">-				return errInsufficientFunds</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-			state.SubBalance(from.Address, amount)</span>
<span style="color:#A00">-		} else {</span>
<span style="color:#A00">-			log.Debug("crosschain C->X", "addr", from.Address, "amount", from.Amount, "assetID", from.AssetID)</span>
<span style="color:#A00">-			amount := new(big.Int).SetUint64(from.Amount)</span>
<span style="color:#A00">-			if state.GetBalanceMultiCoin(from.Address, common.Hash(from.AssetID)).Cmp(amount) < 0 {</span>
<span style="color:#A00">-				return errInsufficientFunds</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-			state.SubBalanceMultiCoin(from.Address, common.Hash(from.AssetID), amount)</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		if state.GetNonce(from.Address) != from.Nonce {</span>
<span style="color:#A00">-			return errInvalidNonce</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		addrs[from.Address] = from.Nonce</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for addr, nonce := range addrs {</span>
<span style="color:#A00">-		state.SetNonce(addr, nonce+1)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return errInsufficientFunds</span>
 }
<b>diff --git a/plugin/evm/export_tx_test.go b/plugin/evm/export_tx_test.go</b>
<b>index 44d6b6a..03af06a 100644</b>
<b>--- a/plugin/evm/export_tx_test.go</b>
<b>+++ b/plugin/evm/export_tx_test.go</b>
<span style="color:#0AA">@@ -7,7 +7,7 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
 	"github.com/ava-labs/avalanchego/chains/atomic"
 	"github.com/ava-labs/avalanchego/ids"
<b>diff --git a/plugin/evm/gasprice_update.go b/plugin/evm/gasprice_update.go</b>
<b>index b96fa69..2c98fb2 100644</b>
<b>--- a/plugin/evm/gasprice_update.go</b>
<b>+++ b/plugin/evm/gasprice_update.go</b>
<span style="color:#0AA">@@ -8,7 +8,7 @@</span> import (
 	"sync"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type gasPriceUpdater struct {
<b>diff --git a/plugin/evm/gasprice_update_test.go b/plugin/evm/gasprice_update_test.go</b>
<b>index 179dbb8..4ab77ad 100644</b>
<b>--- a/plugin/evm/gasprice_update_test.go</b>
<b>+++ b/plugin/evm/gasprice_update_test.go</b>
<span style="color:#0AA">@@ -9,7 +9,7 @@</span> import (
 	"testing"
 	"time"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 type mockGasPriceSetter struct {
<b>diff --git a/plugin/evm/import_tx.go b/plugin/evm/import_tx.go</b>
<b>index 0349820..8013ece 100644</b>
<b>--- a/plugin/evm/import_tx.go</b>
<b>+++ b/plugin/evm/import_tx.go</b>
<span style="color:#0AA">@@ -1,5 +1,12 @@</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// (c) 2021, Flare Networks Limited. All rights reserved.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// This file is a derived work, based on the avalanchego library whose original</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// notice appears below. It is distributed under a license compatible with the</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// licensing terms of the original code from which it is derived.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE_AVALABS for licensing terms of the original work.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">// Please see the file LICENSE for licensing terms.</span>
<span style="color:#0A0">+</span><span style="color:#0A0">//</span>
 // (c) 2019-2020, Ava Labs, Inc. All rights reserved.
<span style="color:#A00">-// See the file LICENSE for licensing terms.</span>
 
 package evm
 
<span style="color:#0AA">@@ -7,19 +14,15 @@</span> import (
 	"fmt"
 	"math/big"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/chains/atomic"</span>
 	"github.com/ava-labs/avalanchego/database"
 	"github.com/ava-labs/avalanchego/ids"
 	"github.com/ava-labs/avalanchego/snow"
 	"github.com/ava-labs/avalanchego/utils/crypto"
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/utils/math"</span>
 	"github.com/ava-labs/avalanchego/vms/components/avax"
<span style="color:#A00">-	"github.com/ava-labs/avalanchego/vms/secp256k1fx"</span>
 	"github.com/ethereum/go-ethereum/common"
<span style="color:#A00">-	"github.com/ethereum/go-ethereum/log"</span>
 )
 
 // UnsignedImportTx is an unsigned ImportTx
<span style="color:#0AA">@@ -52,89 +55,16 @@</span> func (tx *UnsignedImportTx) Verify(
 	ctx *snow.Context,
 	rules params.Rules,
 ) error {
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	case tx == nil:</span>
<span style="color:#A00">-		return errNilTx</span>
<span style="color:#A00">-	case tx.SourceChain != avmID:</span>
<span style="color:#A00">-		return errWrongChainID</span>
<span style="color:#A00">-	case len(tx.ImportedInputs) == 0:</span>
<span style="color:#A00">-		return errNoImportInputs</span>
<span style="color:#A00">-	case tx.NetworkID != ctx.NetworkID:</span>
<span style="color:#A00">-		return errWrongNetworkID</span>
<span style="color:#A00">-	case ctx.ChainID != tx.BlockchainID:</span>
<span style="color:#A00">-		return errWrongBlockchainID</span>
<span style="color:#A00">-	case rules.IsApricotPhase3 && len(tx.Outs) == 0:</span>
<span style="color:#A00">-		return errNoEVMOutputs</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for _, out := range tx.Outs {</span>
<span style="color:#A00">-		if err := out.Verify(); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for _, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		if err := in.Verify(); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	if !avax.IsSortedAndUniqueTransferableInputs(tx.ImportedInputs) {</span>
<span style="color:#A00">-		return errInputsNotSortedUnique</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if rules.IsApricotPhase2 {</span>
<span style="color:#A00">-		if !IsSortedAndUniqueEVMOutputs(tx.Outs) {</span>
<span style="color:#A00">-			return errOutputsNotSortedUnique</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	} else if rules.IsApricotPhase1 {</span>
<span style="color:#A00">-		if !IsSortedEVMOutputs(tx.Outs) {</span>
<span style="color:#A00">-			return errOutputsNotSorted</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return errWrongChainID</span>
 }
 
 func (tx *UnsignedImportTx) Cost() (uint64, error) {
<span style="color:#A00">-	cost := calcBytesCost(len(tx.UnsignedBytes()))</span>
<span style="color:#A00">-	for _, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		inCost, err := in.In.Cost()</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return 0, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		cost, err = math.Add64(cost, inCost)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return 0, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return cost, nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return 0, fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // Amount of [assetID] burned by this transaction
 func (tx *UnsignedImportTx) Burned(assetID ids.ID) (uint64, error) {
<span style="color:#A00">-	var (</span>
<span style="color:#A00">-		spent uint64</span>
<span style="color:#A00">-		input uint64</span>
<span style="color:#A00">-		err   error</span>
<span style="color:#A00">-	)</span>
<span style="color:#A00">-	for _, out := range tx.Outs {</span>
<span style="color:#A00">-		if out.AssetID == assetID {</span>
<span style="color:#A00">-			spent, err = math.Add64(spent, out.Amount)</span>
<span style="color:#A00">-			if err != nil {</span>
<span style="color:#A00">-				return 0, err</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		if in.AssetID() == assetID {</span>
<span style="color:#A00">-			input, err = math.Add64(input, in.Input().Amount())</span>
<span style="color:#A00">-			if err != nil {</span>
<span style="color:#A00">-				return 0, err</span>
<span style="color:#A00">-			}</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return math.Sub64(input, spent)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return 0, fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // SemanticVerify this transaction is valid.
<span style="color:#0AA">@@ -145,82 +75,7 @@</span> func (tx *UnsignedImportTx) SemanticVerify(
 	baseFee *big.Int,
 	rules params.Rules,
 ) error {
<span style="color:#A00">-	if err := tx.Verify(vm.ctx.XChainID, vm.ctx, rules); err != nil {</span>
<span style="color:#A00">-		return err</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Check the transaction consumes and produces the right amounts</span>
<span style="color:#A00">-	fc := avax.NewFlowChecker()</span>
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	// Apply dynamic fees to import transactions as of Apricot Phase 3</span>
<span style="color:#A00">-	case rules.IsApricotPhase3:</span>
<span style="color:#A00">-		cost, err := stx.Cost()</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		txFee, err := calculateDynamicFee(cost, baseFee)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		fc.Produce(vm.ctx.AVAXAssetID, txFee)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Apply fees to import transactions as of Apricot Phase 2</span>
<span style="color:#A00">-	case rules.IsApricotPhase2:</span>
<span style="color:#A00">-		fc.Produce(vm.ctx.AVAXAssetID, params.AvalancheAtomicTxFee)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, out := range tx.Outs {</span>
<span style="color:#A00">-		fc.Produce(out.AssetID, out.Amount)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	for _, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		fc.Consume(in.AssetID(), in.Input().Amount())</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if err := fc.Verify(); err != nil {</span>
<span style="color:#A00">-		return fmt.Errorf("import tx flow check failed due to: %w", err)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if len(stx.Creds) != len(tx.ImportedInputs) {</span>
<span style="color:#A00">-		return fmt.Errorf("export tx contained mismatched number of inputs/credentials (%d vs. %d)", len(tx.ImportedInputs), len(stx.Creds))</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if !vm.ctx.IsBootstrapped() {</span>
<span style="color:#A00">-		// Allow for force committing during bootstrapping</span>
<span style="color:#A00">-		return nil</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	utxoIDs := make([][]byte, len(tx.ImportedInputs))</span>
<span style="color:#A00">-	for i, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		inputID := in.UTXOID.InputID()</span>
<span style="color:#A00">-		utxoIDs[i] = inputID[:]</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	// allUTXOBytes is guaranteed to be the same length as utxoIDs</span>
<span style="color:#A00">-	allUTXOBytes, err := vm.ctx.SharedMemory.Get(tx.SourceChain, utxoIDs)</span>
<span style="color:#A00">-	if err != nil {</span>
<span style="color:#A00">-		return fmt.Errorf("failed to fetch import UTXOs from %s with %w", tx.SourceChain, err)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	for i, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		utxoBytes := allUTXOBytes[i]</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		utxo := &avax.UTXO{}</span>
<span style="color:#A00">-		if _, err := vm.codec.Unmarshal(utxoBytes, utxo); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		cred := stx.Creds[i]</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		utxoAssetID := utxo.AssetID()</span>
<span style="color:#A00">-		inAssetID := in.AssetID()</span>
<span style="color:#A00">-		if utxoAssetID != inAssetID {</span>
<span style="color:#A00">-			return errAssetIDMismatch</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		if err := vm.fx.VerifyTransfer(tx, in.In, cred, utxo.Out); err != nil {</span>
<span style="color:#A00">-			return err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	return vm.conflicts(tx.InputUTXOs(), pa</span><span style="color:#A00">rent)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // Accept this transaction and spend imported inputs
<span style="color:#0AA">@@ -229,12 +84,7 @@</span> func (tx *UnsignedImportTx) SemanticVerify(
 // only to have the transaction not be Accepted. This would be inconsistent.
 // Recall that imported UTXOs are not kept in a versionDB.
 func (tx *UnsignedImportTx) Accept(ctx *snow.Context, batch database.Batch) error {
<span style="color:#A00">-	utxoIDs := make([][]byte, len(tx.ImportedInputs))</span>
<span style="color:#A00">-	for i, in := range tx.ImportedInputs {</span>
<span style="color:#A00">-		inputID := in.InputID()</span>
<span style="color:#A00">-		utxoIDs[i] = inputID[:]</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{tx.SourceChain: {RemoveRequests: utxoIDs}}, batch)</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return fmt.Errorf("exportTx transactions disabled")</span>
 }
 
 // newImportTx returns a new ImportTx
<span style="color:#0AA">@@ -244,160 +94,11 @@</span> func (vm *VM) newImportTx(
 	baseFee *big.Int, // fee to use post-AP3
 	keys []*crypto.PrivateKeySECP256K1R, // Keys to import the funds
 ) (*Tx, error) {
<span style="color:#A00">-	if vm.ctx.XChainID != chainID {</span>
<span style="color:#A00">-		return nil, errWrongChainID</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	kc := secp256k1fx.NewKeychain()</span>
<span style="color:#A00">-	for _, key := range keys {</span>
<span style="color:#A00">-		kc.Add(key)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	atomicUTXOs, _, _, err := vm.GetAtomicUTXOs(chainID, kc.Addresses(), ids.ShortEmpty, ids.Empty, -1)</span>
<span style="color:#A00">-	if err != nil {</span>
<span style="color:#A00">-		return nil, fmt.Errorf("problem retrieving atomic UTXOs: %w", err)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	importedInputs := []*avax.TransferableInput{}</span>
<span style="color:#A00">-	signers := [][]*crypto.PrivateKeySECP256K1R{}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	importedAmount := make(map[ids.ID]uint64)</span>
<span style="color:#A00">-	now := vm.clock.Unix()</span>
<span style="color:#A00">-	for _, utxo := range atomicUTXOs {</span>
<span style="color:#A00">-		inputIntf, utxoSigners, err := kc.Spend(utxo.Out, now)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			continue</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		input, ok := inputIntf.(avax.TransferableIn)</span>
<span style="color:#A00">-		if !ok {</span>
<span style="color:#A00">-			continue</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		aid := utxo.AssetID()</span>
<span style="color:#A00">-		importedAmount[aid], err = math.Add64(importedAmount[aid], input.Amount())</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		importedInputs = append(importedInputs, &avax.TransferableInput{</span>
<span style="color:#A00">-			UTXOID: utxo.UTXOID,</span>
<span style="color:#A00">-			Asset:  utxo.Asset,</span>
<span style="color:#A00">-			In:     input,</span>
<span style="color:#A00">-		})</span>
<span style="color:#A00">-		signers = append(signers, utxoSigners)</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	avax.SortTransferableInputsWithSigners(importedInputs, signers)</span>
<span style="color:#A00">-	importedAVAXAmount := importedAmount[vm.ctx.AVAXAssetID]</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	outs := make([]EVMOutput, 0, len(importedAmount))</span>
<span style="color:#A00">-	// This will create unique outputs (in the context of sorting)</span>
<span style="color:#A00">-	// since each output will have a unique assetID</span>
<span style="color:#A00">-	for assetID, amount := range importedAmount {</span>
<span style="color:#A00">-		// Skip the AVAX amount since it is included separately to account for</span>
<span style="color:#A00">-		// the fee</span>
<span style="color:#A00">-		if assetID == vm.ctx.AVAXAssetID || amount == 0 {</span>
<span style="color:#A00">-			continue</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		outs = append(outs, EVMOutput{</span>
<span style="color:#A00">-			Address: to,</span>
<span style="color:#A00">-			Amount:  amount,</span>
<span style="color:#A00">-			AssetID: assetID,</span>
<span style="color:#A00">-		})</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	rules := vm.currentRules()</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	var (</span>
<span style="color:#A00">-		txFeeWithoutChange uint64</span>
<span style="color:#A00">-		txFeeWithChange    uint64</span>
<span style="color:#A00">-	)</span>
<span style="color:#A00">-	switch {</span>
<span style="color:#A00">-	case rules.IsApricotPhase3:</span>
<span style="color:#A00">-		if baseFee == nil {</span>
<span style="color:#A00">-			return nil, errNilBaseFeeApricotPhase3</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		utx := &UnsignedImportTx{</span>
<span style="color:#A00">-			NetworkID:      vm.ctx.NetworkID,</span>
<span style="color:#A00">-			BlockchainID:   vm.ctx.ChainID,</span>
<span style="color:#A00">-			Outs:           outs,</span>
<span style="color:#A00">-			ImportedInputs: importedInputs,</span>
<span style="color:#A00">-			SourceChain:    chainID,</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		tx := &Tx{UnsignedAtomicTx: utx}</span>
<span style="color:#A00">-		if err := tx.Sign(vm.codec, nil); err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		costWithoutChange, err := tx.Cost()</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		costWithChange := costWithoutChange + EVMOutputGas</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-		txFeeWithoutChange, err = calculateDynamicFee(costWithoutChange, baseFee)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-		txFeeWithChange, err = calculateDynamicFee(costWithChange, baseFee)</span>
<span style="color:#A00">-		if err != nil {</span>
<span style="color:#A00">-			return nil, err</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	case rules.IsApricotPhase2:</span>
<span style="color:#A00">-		txFeeWithoutChange = params.AvalancheAtomicTxFee</span>
<span style="color:#A00">-		txFeeWithChange = params.AvalancheAtomicTxFee</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// AVAX output</span>
<span style="color:#A00">-	if importedAVAXAmount < txFeeWithoutChange { // imported amount goes toward paying tx fee</span>
<span style="color:#A00">-		return nil, errInsufficientFundsForFee</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	if importedAVAXAmount > txFeeWithChange {</span>
<span style="color:#A00">-		outs = append(outs, EVMOutput{</span>
<span style="color:#A00">-			Address: to,</span>
<span style="color:#A00">-			Amount:  importedAVAXAmount - txFeeWithChange,</span>
<span style="color:#A00">-			AssetID: vm.ctx.AVAXAssetID,</span>
<span style="color:#A00">-		})</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// If no outputs are produced, return an error.</span>
<span style="color:#A00">-	// Note: this can happen if there is exactly enough AVAX to pay the</span>
<span style="color:#A00">-	// transaction fee, but no other funds to be imported.</span>
<span style="color:#A00">-	if len(outs) == 0 {</span>
<span style="color:#A00">-		return nil, errNoEVMOutputs</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	SortEVMOutputs(outs)</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	// Create the transaction</span>
<span style="color:#A00">-	utx := &UnsignedImportTx{</span>
<span style="color:#A00">-		NetworkID:      vm.ctx.NetworkID,</span>
<span style="color:#A00">-		BlockchainID:   vm.ctx.ChainID,</span>
<span style="color:#A00">-		Outs:           outs,</span>
<span style="color:#A00">-		ImportedInputs: importedInputs,</span>
<span style="color:#A00">-		SourceChain:    chainID,</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	tx := &Tx{UnsignedAtomicTx: utx}</span>
<span style="color:#A00">-	if err := tx.Sign(vm.codec, signers); err != nil {</span>
<span style="color:#A00">-		return nil, err</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return tx, utx.Verify(vm.ctx.XChainID, vm.ctx, vm.currentRules())</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return nil, errWrongChainID</span>
 }
 
 // EVMStateTransfer performs the state transfer to increase the balances of
 // accounts accordingly with the imported EVMOutputs
 func (tx *UnsignedImportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
<span style="color:#A00">-	for _, to := range tx.Outs {</span>
<span style="color:#A00">-		if to.AssetID == ctx.AVAXAssetID {</span>
<span style="color:#A00">-			log.Debug("crosschain X->C", "addr", to.Address, "amount", to.Amount, "assetID", "AVAX")</span>
<span style="color:#A00">-			// If the asset is AVAX, convert the input amount in nAVAX to gWei by</span>
<span style="color:#A00">-			// multiplying by the x2c rate.</span>
<span style="color:#A00">-			amount := new(big.Int).Mul(</span>
<span style="color:#A00">-				new(big.Int).SetUint64(to.Amount), x2cRate)</span>
<span style="color:#A00">-			state.AddBalance(to.Address, amount)</span>
<span style="color:#A00">-		} else {</span>
<span style="color:#A00">-			log.Debug("crosschain X->C", "addr", to.Address, "amount", to.Amount, "assetID", to.AssetID)</span>
<span style="color:#A00">-			amount := new(big.Int).SetUint64(to.Amount)</span>
<span style="color:#A00">-			state.AddBalanceMultiCoin(to.Address, common.Hash(to.AssetID), amount)</span>
<span style="color:#A00">-		}</span>
<span style="color:#A00">-	}</span>
<span style="color:#A00">-	return nil</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">return errInsufficientFunds</span>
 }
<b>diff --git a/plugin/evm/import_tx_test.go b/plugin/evm/import_tx_test.go</b>
<b>index c4c7df8..6551bb3 100644</b>
<b>--- a/plugin/evm/import_tx_test.go</b>
<b>+++ b/plugin/evm/import_tx_test.go</b>
<span style="color:#0AA">@@ -7,7 +7,7 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
 	"github.com/ava-labs/avalanchego/chains/atomic"
 	"github.com/ava-labs/avalanchego/ids"
<b>diff --git a/plugin/evm/service.go b/plugin/evm/service.go</b>
<b>index a6f56f2..ee079bb 100644</b>
<b>--- a/plugin/evm/service.go</b>
<b>+++ b/plugin/evm/service.go</b>
<span style="color:#0AA">@@ -17,11 +17,11 @@</span> import (
 	"github.com/ava-labs/avalanchego/utils/crypto"
 	"github.com/ava-labs/avalanchego/utils/formatting"
 	"github.com/ava-labs/avalanchego/utils/json"
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	ethcrypto "github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/log"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // test constants
<b>diff --git a/plugin/evm/static_service.go b/plugin/evm/static_service.go</b>
<b>index 9c59225..f19fe78 100644</b>
<b>--- a/plugin/evm/static_service.go</b>
<b>+++ b/plugin/evm/static_service.go</b>
<span style="color:#0AA">@@ -8,7 +8,7 @@</span> import (
 	"encoding/json"
 
 	"github.com/ava-labs/avalanchego/utils/formatting"
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
 )
 
 // StaticService defines the static API services exposed by the evm
<b>diff --git a/plugin/evm/tx.go b/plugin/evm/tx.go</b>
<b>index ac4f798..b15365e 100644</b>
<b>--- a/plugin/evm/tx.go</b>
<b>+++ b/plugin/evm/tx.go</b>
<span style="color:#0AA">@@ -12,8 +12,8 @@</span> import (
 
 	"github.com/ethereum/go-ethereum/common"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 
 	"github.com/ava-labs/avalanchego/codec"
 	"github.com/ava-labs/avalanchego/database"
<b>diff --git a/plugin/evm/tx_test.go b/plugin/evm/tx_test.go</b>
<b>index d1d4947..ca73402 100644</b>
<b>--- a/plugin/evm/tx_test.go</b>
<b>+++ b/plugin/evm/tx_test.go</b>
<span style="color:#0AA">@@ -7,7 +7,7 @@</span> import (
 	"math/big"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func TestCalculateDynamicFee(t *testing.T) {
<b>diff --git a/plugin/evm/vm.go b/plugin/evm/vm.go</b>
<b>index 5850762..198f899 100644</b>
<b>--- a/plugin/evm/vm.go</b>
<b>+++ b/plugin/evm/vm.go</b>
<span style="color:#0AA">@@ -10,25 +10,26 @@</span> import (
 	"errors"
 	"fmt"
 	"math/big"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"os"</span>
 	"path/filepath"
 	"strings"
 	"sync"
 	"time"
 
 	"github.com/ava-labs/avalanchego/database/versiondb"
<span style="color:#A00">-	coreth "github.com/ava-labs/coreth/chain"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/consensus/dummy"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth/ethconfig"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/node"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">coreth "gitlab.com/flarenetwork/coreth/chain"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/consensus/dummy"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth/ethconfig"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/node"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/log"
 	"github.com/ethereum/go-ethereum/rlp"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 
 	avalancheRPC "github.com/gorilla/rpc/v2"
 
<span style="color:#0AA">@@ -279,6 +280,26 @@</span> func (vm *VM) Initialize(
 			return fmt.Errorf("failed to unmarshal config %s: %w", string(configBytes), err)
 		}
 	}
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.config.EthAPIEnabled = false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.config.NetAPIEnabled = false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.config.Web3APIEnabled = false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.config.DebugAPIEnabled = false</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.config.MaxBlocksPerRequest = 1</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">web3API := os.Getenv("WEB3_API")</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">if web3API == "enabled" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.EthAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.NetAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.Web3APIEnabled = true</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">} else if web3API == "debug" {</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.EthAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.NetAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.Web3APIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.DebugAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.TxPoolAPIEnabled = true</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.Pruning = false</span>
<span style="color:#0A0">+</span>		<span style="color:#0A0">vm.config.MaxBlocksPerRequest = 0</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">}</span>
<span style="color:#0A0">+</span>
 	if b, err := json.Marshal(vm.config); err == nil {
 		log.Info("Initializing Coreth VM", "Version", Version, "Config", string(b))
 	} else {
<span style="color:#0AA">@@ -337,7 +358,7 @@</span> func (vm *VM) Initialize(
 	ethConfig.SnapshotVerify = vm.config.SnapshotVerify
 
 	vm.chainConfig = g.Config
<span style="color:#A00">-	vm.networkID = ethConfig.NetworkId</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">vm.networkID = g.Config.ChainID.Uint64()</span>
 	vm.secpFactory = crypto.FactorySECP256K1R{Cache: cache.LRU{Size: secpFactoryCacheSize}}
 
 	nodecfg := node.Config{
<b>diff --git a/plugin/evm/vm_test.go b/plugin/evm/vm_test.go</b>
<b>index 80a6953..f4ec7e0 100644</b>
<b>--- a/plugin/evm/vm_test.go</b>
<b>+++ b/plugin/evm/vm_test.go</b>
<span style="color:#0AA">@@ -40,13 +40,13 @@</span> import (
 
 	engCommon "github.com/ava-labs/avalanchego/snow/engine/common"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/eth"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/rpc"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/eth"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/rpc"</span>
 
<span style="color:#A00">-	accountKeystore "github.com/ava-labs/coreth/accounts/keystore"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">accountKeystore "gitlab.com/flarenetwork/coreth/accounts/keystore"</span>
 )
 
 var (
<b>diff --git a/plugin/main.go b/plugin/main.go</b>
<b>index 9637bea..af1f80a 100644</b>
<b>--- a/plugin/main.go</b>
<b>+++ b/plugin/main.go</b>
<span style="color:#0AA">@@ -12,7 +12,7 @@</span> import (
 
 	"github.com/ava-labs/avalanchego/vms/rpcchainvm"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/plugin/evm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/plugin/evm"</span>
 )
 
 func main() {
<b>diff --git a/signer/core/apitypes/types.go b/signer/core/apitypes/types.go</b>
<b>index cab7f9c..32a9ec4 100644</b>
<b>--- a/signer/core/apitypes/types.go</b>
<b>+++ b/signer/core/apitypes/types.go</b>
<span style="color:#0AA">@@ -32,9 +32,9 @@</span> import (
 	"math/big"
 	"strings"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
 )
 
 type ValidationInfo struct {
<b>diff --git a/tests/init.go b/tests/init.go</b>
<b>index c924236..1fe1359 100644</b>
<b>--- a/tests/init.go</b>
<b>+++ b/tests/init.go</b>
<span style="color:#0AA">@@ -31,7 +31,7 @@</span> import (
 	"math/big"
 	"sort"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // Forks table defines supported forks and their chain config.
<b>diff --git a/tests/init_test.go b/tests/init_test.go</b>
<b>index 7f30fde..b72be17 100644</b>
<b>--- a/tests/init_test.go</b>
<b>+++ b/tests/init_test.go</b>
<span style="color:#0AA">@@ -40,7 +40,7 @@</span> import (
 	"strings"
 	"testing"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 func readJSON(reader io.Reader, value interface{}) error {
<b>diff --git a/tests/state_test_util.go b/tests/state_test_util.go</b>
<b>index 7df6737..43e3990 100644</b>
<b>--- a/tests/state_test_util.go</b>
<b>+++ b/tests/state_test_util.go</b>
<span style="color:#0AA">@@ -34,17 +34,17 @@</span> import (
 	"strconv"
 	"strings"
 
<span style="color:#A00">-	"github.com/ava-labs/coreth/core"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/state/snapshot"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/types"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/core/vm"</span>
<span style="color:#A00">-	"github.com/ava-labs/coreth/params"</span>
 	"github.com/ethereum/go-ethereum/common"
 	"github.com/ethereum/go-ethereum/common/hexutil"
 	"github.com/ethereum/go-ethereum/common/math"
 	"github.com/ethereum/go-ethereum/crypto"
 	"github.com/ethereum/go-ethereum/ethdb"
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/state/snapshot"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/types"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/core/vm"</span>
<span style="color:#0A0">+</span>	<span style="color:#0A0">"gitlab.com/flarenetwork/coreth/params"</span>
 )
 
 // StateTest checks transaction processing without block context.
