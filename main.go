/*
 * Author : Ismail Ash Shidiq (https://www.eulbyvan.com)
 * Created on : Fri Mar 03 2023 9:41:07 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
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