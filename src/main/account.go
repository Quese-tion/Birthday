package main

type accountingDAO interface {
  //CREATE USER
  insertUser(string) []User //Create
  selectAccountbyid(int) []Account
  selectAllAccounts() []Account
  selectAccountuid(int) []Account
  selectAccountPending() []Account
  //iterateDB(ResultSet resultSet) []Account
  approveAccount(int, bool) []Account
  depositAmount(int, float32) bool
  withdrawAmount(int, float32) bool
  transferAmount(int, float32, int) bool
  addUser(int, int) []Account
  cancelAccount(int) []Account
}
type Account struct { //ACCOUNT Data Model
  Name string `json:"account_name"`
  ID int `json:"id"`
  Status string `json:"status"`
  Users []int `json:"users"`
  Balance float32 `json:"balance"`
  //lastLoginAt string
}
