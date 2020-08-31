// Copyright (c) Alex Ellis 2017. All rights reserved.
// Copyright (c) OpenFaaS Author(s) 2018. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
	"net/http"

	"function"
)

import "time"

func main() {
	// input, err := ioutil.ReadAll(os.Stdin)
	// if err != nil {
	// 	log.Fatalf("Unable to read standard input: %s", err.Error())
	// }

	// //fmt.Println(function.Handle(input))
	// http.HandleFunc("/Handle", function.Handle).Methods("POST")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8002),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	http.HandleFunc("/", function.Handle) //.Methods("POST")
	log.Fatal(s.ListenAndServe())
	http.ListenAndServe(":8002", nil)

}
