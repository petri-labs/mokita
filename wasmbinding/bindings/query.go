package bindings

// MokisisQuery contains mokita custom queries.
// See https://github.com/petri-labs/mokita-bindings/blob/main/packages/bindings/src/query.rs
type MokisisQuery struct {
	/// Given a subdenom minted by a contract via `MokisisMsg::MintTokens`,
	/// returns the full denom as used by `BankMsg::Send`.
	FullDenom *FullDenom `json:"full_denom,omitempty"`
	/// Returns the admin of a denom, if the denom is a Token Factory denom.
	DenomAdmin *DenomAdmin `json:"denom_admin,omitempty"`
}

type FullDenom struct {
	CreatorAddr string `json:"creator_addr"`
	Subdenom    string `json:"subdenom"`
}

type DenomAdmin struct {
	Subdenom string `json:"subdenom"`
}

type DenomAdminResponse struct {
	Admin string `json:"admin"`
}

type FullDenomResponse struct {
	Denom string `json:"denom"`
}
