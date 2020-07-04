/*
* This file is Part of the Picnic-Lisp research project
*
* Copyright (c) 2020 Timo Sarkar
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to
* deal in the Software without restriction, including without limitation the
* rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
* sell copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in
* all copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
* FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
* IN THE SOFTWARE.
*/

package picnic

import (
	"path"

	"github.com/rakyll/statik/fs"
)

//go:generate statik -src=lib

func LoadLib(env *Env) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	dir, err := statikFS.Open("/")
	if err != nil {
		return err
	}
	defer dir.Close()

	fis, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		f, err := statikFS.Open(path.Join("/", fi.Name()))
		if err != nil {
			return err
		}
		node, err := NewParser(f).Parse()
		if err != nil {
			f.Close()
			return err
		}
		_, err = env.Eval(node)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	return nil
}
