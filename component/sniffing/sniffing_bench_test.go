/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2022-2023, daeuniverse Organization <dae@v2raya.org>
 */

package sniffing

import (
	"fmt"
	"github.com/mzz2017/softwind/pkg/fastrand"
	"testing"
)

var (
	httpMethodSet map[string]struct{}
)

func init() {
	httpMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "COPY", "HEAD", "OPTIONS", "LINK", "UNLINK", "PURGE", "LOCK", "UNLOCK", "PROPFIND"}
	httpMethodSet = make(map[string]struct{})
	for _, method := range httpMethods {
		httpMethodSet[method] = struct{}{}
	}
}

func BenchmarkStringSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var test [5]byte
		fastrand.Read(test[:])
		_, ok := httpMethodSet[string(test[:])]
		if !ok {
			fmt.Sprintf("%v", string(test[:]))
		}
	}
}

func BenchmarkStringSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var test [5]byte
		fastrand.Read(test[:])
		switch string(test[:]) {
		case "GET", "POST", "PUT", "PATCH", "DELETE", "COPY", "HEAD", "OPTIONS", "LINK", "UNLINK", "PURGE", "LOCK", "UNLOCK", "PROPFIND":
		default:
			fmt.Sprintf("%v", string(test[:]))
		}
	}
}
