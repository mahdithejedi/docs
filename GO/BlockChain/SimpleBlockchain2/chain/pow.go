package chain

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"sync"
)

type nonceResultType uint8

const (
	newNonce   nonceResultType = 0
	nonceFound                 = 1
)

const nonceRange uint64 = 100
const maxNonceWorker uint8 = 3

var wg sync.WaitGroup

var _lock sync.Mutex

type POW struct {
	Block  *block
	Target *big.Int
}

type nonceChannel struct {
	from  uint64
	to    uint64
	nonce int
	hash  []byte
}

const targetBits = 24
const maxNonce = math.MaxUint64

func NewPOW(b *block) *POW {
	t := big.NewInt(1)
	t.Lsh(t, uint(265-targetBits))
	return &POW{
		Block:  b,
		Target: t,
	}
}

func intToByte(i int64) []byte {
	if i > 0 {
		return append(big.NewInt(i).Bytes(), byte(1))
	}
	return append(big.NewInt(i).Bytes(), byte(0))
}

func (p *POW) prepareData(nounce int) []byte {
	return bytes.Join(
		[][]byte{
			p.Block.preBlock,
			p.Block.GetByte(),
			intToByte(p.Block.timestamp),
			intToByte(int64(targetBits)),
			intToByte(int64(nounce)),
		},
		[]byte{},
	)
}

func (p *POW) run(nonceRange chan nonceChannel, nonceResult chan nonceResultType) {
	defer wg.Done()
	for {
		_lock.Lock()
		newRange := <-nonceRange
		_lock.Unlock()
		nonce, result, error := p.proof(newRange.from, newRange.to)
		_lock.Lock()
		if error != nil {
			nonceResult <- newNonce
			_lock.Unlock()
		} else {
			nonceResult <- nonceFound
			nonceRange <- nonceChannel{
				nonce: nonce,
				hash:  result,
			}
			_lock.Unlock()
			break
		}
	}
}

func (p *POW) proof(from uint64, to uint64) (int, []byte, error) {
	var hashInt big.Int
	var hash [32]byte
	for from < to {
		data := p.prepareData(int(from))
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.Target) == -1 {
			return int(from), hash[:], nil
		} else {
			from++
		}
	}
	return 0, nil, errors.New("nonce now found")

}

func getNonce(nonce *uint64) nonceChannel {
	defer func(nonce *uint64) {
		*nonce += nonceRange
	}(nonce)
	if *nonce == 0 {
		*nonce = uint64(1)
	}
	return nonceChannel{
		from: *nonce,
		to:   *nonce + nonceRange,
	}
}

func (p *POW) Run() (int, []byte, error) {
	resultNonce := make(chan nonceResultType, maxNonceWorker)
	nonceChannelChan := make(chan nonceChannel, maxNonceWorker)
	var nonce uint64 = 0
	defer func() {
		wg.Wait()
		close(nonceChannelChan)
		close(resultNonce)
	}()
	for i := 0; i < int(maxNonceWorker); i++ {
		wg.Add(1)
		nonceChannelChan <- getNonce(&nonce)
		go p.run(nonceChannelChan, resultNonce)
	}
	for nonce < maxNonce {
		select {
		case res := <-resultNonce:
			if res == newNonce {
				nonceChannelChan <- getNonce(&nonce)
			} else if res == nonceFound {
				result := <-nonceChannelChan
				return result.nonce, result.hash, nil
			}
		}
	}
	return 0, nil, errors.New("oh shit no nonce")
}
func (p *POW) Validate() bool {
	var bigInt big.Int

	data := p.prepareData(p.Block.Nonce)
	hash := sha256.Sum256(data)
	bigInt.SetBytes(hash[:])
	fmt.Println("nonce is ", p.Block.Nonce, reflect.TypeOf(p.Block.Nonce))
	return bigInt.Cmp(p.Target) == -1
}
