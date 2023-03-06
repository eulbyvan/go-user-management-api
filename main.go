/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 9:41:07 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

/*
	Install package testify
	go get github.com/stretchr/testify
    go get github.com/DATA-DOG/go-sqlmock

	Rule of Thumbs
	1. Tidak boleh berkomunikasi dengan remote database
	2. Tidak boleh berkomunikasi secara lintas network
    3. Tidak boleh melakukan setup konfigurasi di environment


	go test ./...
*/

package main

import (
	"github.com/eulbyvan/go-user-management/delivery"
	_ "github.com/lib/pq"
)

func main() {
	// Run the server
	delivery.Server().Run()
}
