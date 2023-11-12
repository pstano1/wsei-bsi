## WSEI-BSI

This repository cotains a CLI implementing some od most known cipher used in the past.

Supported ciphers:
- Caesar cipher
- Polybius square
- Beale cipher
- Trithemius cipher (tabula recta)
- Vigenère cipher

TO-DO:
- [ ] make a JSON config file for maps & character sets
- [ ] add comments abotu each cipher and it's modification if one's been done 

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

- `caesar`: Caesar cipher

- `polybius`: Polybius square

- `beale`: Beale cipher

##### Substitution Ciphers

- `trithemius`: Trithemius cipher (tabula recta)

- `vigenere`: Vigenère cipher

### License

This project is licensed under the MIT License.
