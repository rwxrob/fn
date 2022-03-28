// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package maps

import (
	"sort"
	"strings"

	"github.com/rwxrob/fn"
)

// Note to maintainers: This file contains maps that require additional
// arguments and are therefore not able to call simple map functions
// from the mapf package. Please keep simple mapf-able maps in
// simple.go instead.

// Prefix adds a prefix to the string.
func Prefix(in []string, pre string) []string {
	return fn.Map(in, func(i string) string { return pre + i })
}

// Keys returns the keys in lexicographically sorted order.
func Keys[T any](m map[string]T) []string {
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// KeysWithPrefix returns only keys from the map that match
// strings.HasPrefix.
func KeysWithPrefix[T any](m map[string]T, pre string) []string {
	keys := []string{}
	for k, _ := range m {
		if strings.HasPrefix(k, pre) {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys
}
