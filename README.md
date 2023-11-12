## WSEI-BSI

This repository houses a Command-Line Interface (CLI) that implements various historical ciphers. These ciphers were used in the past for secure communication and encryption.

Supported ciphers:
- Caesar cipher
- Polybius square
- Homophonic cipher
- Trithemius cipher (tabula recta)
- Vigenère cipher

### config.json

Please note that `config.json` may contain sensitive information, and for production scenarios, it's recommended to secure this file and avoid public exposure. However, in the context of this educational project, I've included it for instructional purposes and as a reference example.

### How to run the app

```console
git clone https://github.com/pstano1/wsei-bsi.git
cd wsei-bsi
go mod tidy
go run ./cmd/ <with flags>
```

### Example Usage 

```console
go run ./cmd/ --action code --cipher caesar --text "AĄBCĆDEĘFGHIJKLŁMNŃOÓPQRSŚTUVWXYZŹŻ" --key 3
Result: cćdeęfghijklłmnńoópqrsśtuvwxyzźżaąb
```

### Flags

| Flag | Description |
|----------|---------|
| action | Defines what action's gonna be performed |
| cipher | Name for cipher you wanna use |
| text | Value |
| key | - |


#### Available Actions

- `code`: Encode a plaintext.
- `decode`: Decode a ciphertext.

#### Available Ciphers

##### Symmetric Ciphers

- `caesar`

- `polybius`

- `homophonic`

##### Substitution Ciphers

- `trithemius`

- `vigenere`

### License

This project is licensed under the MIT License.
