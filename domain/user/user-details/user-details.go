package user_details

type UserDetails struct {
	Id                      string `json:"id"`
	UserEmail               string `json:"user_email"`
	BankId                  string `json:"bank"`
	CompanyRegisteredNumber string `json:"company_registered_number"`
	TaxNumber               string `json:"tax_number"`
}

type UserBank struct {
	Id         string `json:"id"`
	UserEmail  string `json:"user_email"`
	BankType   string `json:"bank_type"`
	BankName   string `json:"bank_name"`
	BranchCode string `json:"branch_code"`
	BankNumber string `json:"bank_number"`
	CvcCode    string `json:"cvc_code"`
}
