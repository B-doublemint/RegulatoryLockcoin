package mixcoin

import (
	"crypto/rsa"
	"crypto/rand"
	"math/big"
)
type PrivateKey struct {
	PublicKey            // public part.
	D         *big.Int   // private exponent
	Primes    []*big.Int // prime factors of N, has >= 2 elements.

	// Precomputed contains precomputed values that speed up private
	// operations, if available.
	Precomputed PrecomputedValues
}
type PrecomputedValues struct {
	Dp, Dq *big.Int // D mod (P-1) (or mod Q-1)
	Qinv   *big.Int // Q^-1 mod P

	// CRTValues is used for the 3rd and subsequent primes. Due to a
	// historical accident, the CRT for the first two primes is handled
	// differently in PKCS#1 and interoperability is sufficiently
	// important that we mirror this.
	CRTValues []CRTValue
}
// CRTValue contains the precomputed Chinese remainder theorem values.
type CRTValue struct {
	Exp   *big.Int // D mod (prime-1).
	Coeff *big.Int // R·Coeff ≡ 1 mod Prime.
	R     *big.Int // product of primes prior to this (inc p and q).
}
// A PublicKey represents the public part of an RSA key.
type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}
var privateKey PrivateKey
// Generate a key
func GenerateKey()(*rsa.PrivateKey){
	keysize := 2048
	key, _ := rsa.GenerateKey(rand.Reader, keysize)
	//println("公钥",key.PublicKey.N)
	//privateKey.PublicKey.E = key.PublicKey.E
	//privateKey.PublicKey.N = key.PublicKey.N
	//privateKey.D = key.D
	//privateKey.Primes = key.Primes
	//privateKey.Precomputed.Dp = key.Precomputed.Dp
	//privateKey.Precomputed.Dq = key.Precomputed.Dp
	//privateKey.Precomputed.Qinv = key.Precomputed.Qinv
	return key
}
func main (){
	GenerateKey()
}
