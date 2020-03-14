package jwt


type Secret []byte

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

const JSON="json"
const EXP="exp"
