package model

type User struct {
	ID       int    `json:"id"`
	NamaUser string `json:"namaUser"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Categories struct {
	ID   int    `json:"id" binding:"required"`
	Nama_kategori string `json:"namaKategori" binding:"required"`
}

type Menus struct {
	ID   int    `json:"id" binding:"required"`
	Kategori_id      int    `json:"kategoriid" binding:"required"`
	Nama_menu string `json:"namaMenu" binding:"required"`
	Harga       int `json:"harga" binding:"required"`
	Stok         int `json:"stok" binding:"required"`
	Image         string `json:"image" binding:"required"`
}

type Transactiondetails struct {
	ID  int `json:"id" binding:"required"`
	Transaksi_id int `json:"transaksiid" binding:"required"`
	Menu_id  int `json:"menuId" binding:"required"`
	Jumlah     int `json:"jumlah" binding:"required"`
	Subtotal   int `json:"subtotal" binding:"required"`
}

type Transactions struct {
	ID   int `json:"id" binding:"required"`
	User_id  int `json:"userId" binding:"required"`
	Nama_cust  int `json:"namaCust" binding:"required"`
	Tanggal   int `json:"tanggal" binding:"required"`
	Total_harga   int `json:"totalHarga" binding:"required"`
	Metode_bayar   int `json:"metodeBayar" binding:"required"`
}