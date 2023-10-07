## WSEI-BSI

desc

#### How to run the app

```console
git clone https://github.com/pstano1/wsei-bsi.git
cd wsei-bsi
go mod tidy
go run ./cmd/
```

#### Example Usage 

```console
go run ./cmd/ --action cipher --cipher caesar --text "AĄBCĆDEĘFGHIJKLŁMNŃOÓPQRSŚTUVWXYZŹŻ" --key 3
Result: cćdeęfghijklłmnńoópqrsśtuvwxyzźżaąb
```

##### Flags

| Flag | Description |
|----------|---------|
| action | Defines what action's gonna be performed |
| cipher | Name for cipher you wanna use |
| text | Value |
| key | - |

#### License
