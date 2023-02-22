// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"gocloud.dev/secrets"
)

func main() {
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyType(&secrets.Keeper{}),
		),
		entc.FeatureNames("intercept"),
	}
	if err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by ent, DO NOT EDIT.
		`,
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}