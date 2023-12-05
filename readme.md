# Kagi by doddy-s

## Description
Kagi is a simple package to encrypt and decrypt string using AES-256 algorithm.

## How to use

### Get the package and create kagi instance

```
package main

import (
	"fmt"

	"github.com/doddy-s/kagi"
)

func main() {
	kagi := kagi.New("1234567890123456789012ThisIsAKey")
    // Make sure your key is 32 bytes long
}
```

### Encrypt
```
    someString := "ThisIsTheStringToBeEncrypted"
    
	encrypted := kagi.Encrypt(someString)
	fmt.Println("Encrypted string = ", encrypted)
```

### Decrypt
```
    decrypted := kagi.Decrypt(encrypted)
	fmt.Println("Decrypted string = ", decrypted)
```