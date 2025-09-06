package models

type Debt struct {
    ID            string   `json:"id"`
    FirstName     string   `json:"first_name"`
    LastName      string   `json:"last_name"`
    PhoneNumber   string   `json:"phone_number"`
    Jshshir       string   `json:"jshshir"`
    Address       string   `json:"address"`
    BagID         string   `json:"bag_id"`
    Price         string   `json:"price"`
    PricePaid     string   `json:"price_paid"`
    Acquaintance  string   `json:"acquaintance"`
    Collateral    string   `json:"collateral"`
    Deadline      string   `json:"deadline"`
    DebtCUD       DebtCUD  `json:"debt_cud"`
}

type DebtCUD struct {
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt int64  `json:"deleted_at"`
}

type CreateDebtReq struct {
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    PhoneNumber  string `json:"phone_number"`
    Jshshir      string `json:"jshshir"`
    Address      string `json:"address"`
    BagID        string `json:"bag_id"`
    Price        string `json:"price"`
    PricePaid    string `json:"price_paid"`
    Acquaintance string `json:"acquaintance"`
    Collateral   string `json:"collateral"`
    Deadline     string `json:"deadline"`
}

type UpdateDebtReq struct {
    ID           string `json:"id"`
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    PhoneNumber  string `json:"phone_number"`
    Jshshir      string `json:"jshshir"`
    Address      string `json:"address"`
    BagID        string `json:"bag_id"`
    Price        string `json:"price"`
    PricePaid    string `json:"price_paid"`
    Acquaintance string `json:"acquaintance"`
    Collateral   string `json:"collateral"`
    Deadline     string `json:"deadline"`
}

type DeleteDebtReq struct {
    ID string `json:"id"`
}

type GetDebtByIdReq struct {
    ID string `json:"id"`
}

type GetDebtByFilterReq struct {
    Search string `json:"search"`
}

type GetDebtByFilterResp struct {
    Status      bool   `json:"status"`
    Message     string `json:"message"`
    GetCountResp int32  `json:"get_count_resp"`
    Debt        []Debt `json:"debt"`
}

type DebtResp struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Debt    Debt   `json:"debt"`
}
