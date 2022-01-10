package main

import (
	`github.com/storezhang/gox/file`
)

func copy(from string, to string) error {
	return file.Copy(from, to)
}
