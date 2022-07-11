package main

func isOutOfStock(bookId string) bool {
	var count int

	userSQL := "SELECT current_stock FROM books WHERE book_id=$1"

	row := DB.QueryRow(userSQL, bookId)
	err := row.Scan(&count)
	if err != nil {
		//log.Fatal(err)
	}

	if count > 0 {
		return false
	} else {
		return true
	}

}
