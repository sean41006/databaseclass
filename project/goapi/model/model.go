package model

type Test struct {
	Id    int    `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Phone string `form:"phone" json:"phone"`
}

type Category struct {
	C_id  int    `form:"c_id" json:"c_id"`
	Name  string `form:"name" json:"name"`
	Floor string `form:"floor" json:"floor"`
}

type Brand struct {
	B_id     int    `form:"b_id" json:"b_id"`
	Name     string `form:"name" json:"name"`
	Floor    string `form:"floor" json:"floor"`
	Position string `form:"position" json:"position"`
	C_id     int    `form:"c_id" json:"c_id"`
}

type Employee struct {
	E_id        int    `form:"e_id" json:"e_id"`
	Name        string `form:"name" json:"name"`
	Phone       string `form:"phone" json:"phone"`
	Account     string `form:"account" json:"account"`
	Packcard_id int    `form:"parkcard_id" json:"parkcard_id"`
	B_id        int    `form:"b_id" json:"b_id"`
}

type Parkrecord struct {
	P_id        int    `form:"p_id" json:"p_id"`
	Carnum      string `form:"carnumber" json:"carnumber"`
	Parktime    string `form:"parktime" json:"parktime"`
	Packcard_id int    `form:"parkcard_id" json:"parkcard_id"`
}

type Member struct {
	M_id  int    `form:"m_id" json:"m_id"`
	Level string `form:"level" json:"level"`
}

type Memberdetail struct {
	M_id    int    `form:"m_id" json:"m_id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type Memberall struct {
	M_id    int    `form:"m_id" json:"m_id"`
	Level   string `form:"level" json:"level"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

type Transactionrecord struct {
	T_id    int    `form:"t_id" json:"t_id"`
	Time    string `form:"time" json:"time"`
	Product string `form:"product" json:"product"`
	Amount  string `form:"amount" json:"amount"`
	M_id    int    `form:"m_id" json:"m_id"`
	B_id    int    `form:"b_id" json:"b_id"`
}

type Transname struct {
	T_id    int    `form:"t_id" json:"t_id"`
	Time    string `form:"time" json:"time"`
	Product string `form:"product" json:"product"`
	Amount  string `form:"amount" json:"amount"`
	M_name  string `form:"m_name" json:"m_name"`
	B_name  string `form:"b_name" json:"b_name"`
}

type Account struct {
	Id       int    `form:"id" json:"id"`
	Account  string `form:"account" json:"account"`
	Password string `form:"password" json:"password"`
}

// type Response struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// 	Data    []Test
// }
