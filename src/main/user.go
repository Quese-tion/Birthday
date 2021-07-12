package main

//DAO MODELING
type adminDAO interface  {
  createAccount(User)
  selectAllAccounts() []Account
  selectUser(int) []User //Read 1
  selectAllUsers() []User //Read All
  depositAmount(int, float32) bool
  withdrawAmount(int, float32) bool
  transferAmount(int, float32, int) bool
  //<User> iterateDB(ResultSet resultSet) throws SQLException;
  updateUser(User) []User //Update 1
  deleteUser(int) []User //Delete 1
  insertUser(string) []User
  addUser(int, int) []Account
  cancelAccount(int) []Account
}
type userDAO interface {
  createAccount(User)
  selectAllAccounts() []Account
  depositAmount(int, float32) bool
  withdrawAmount(int, float32) bool
  transferAmount(int, float32, int) bool
  insertUser(string) []User
  addUser(int, int) []Account
  cancelAccount(int) []Account
  deleteUser(int) []User
}

type User struct { //USER Data Model
  Username string `json:"User"`
  Password string `json:"pwd"`
  Name string `json:"full_name"`
  ID int `json:"id"`
  Role string `json:"role"`
  //Age int `json:"age"`
  Accounts []int `json:"accounts"`
  //Active bool `json:"active"`
  //lastLoginAt string
}

//Object Method Implementations
func (User) insertUser(string) []User{ return nil}
func (User) selectAllAccounts() []Account{ return nil}
func (User) selectUser(int) []User {//Read 1
  return nil}
func (User) selectAllUsers() []User {//Read All
  return nil}
func (User) depositAmount(int, float32) bool { return true}
func (User) withdrawAmount(int, float32) bool { return true}
func (User) transferAmount(int, float32, int) bool { return true}
//<User> iterateDB(ResultSet resultSet) throws SQLException;
func (User) updateUser(User) []User { //Update 1
  return nil}
func (User) deleteUser(int) []User {//Delete 1
  return nil}
func (User) addUser(int, int) []Account {return nil}
func (User) cancelAccount(int) []Account {return nil}
