package utils

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"github.com/cxnky/goupdate/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**

 * Created by cxnky on 24/08/2018 at 18:13
 * utils
 * https://github.com/cxnky/

**/

// Unzip will unzip the downloaded update to the PWD.
func Unzip(src, dest string) error {

	r, err := zip.OpenReader(src)

	if err != nil {

		return errors.NewError(err.Error())

	}

	defer func() {

		if err := r.Close(); err != nil {

			panic(err)

		}

	}()

	os.MkdirAll(dest, 0755)

	extractAndWriteFile := func(f *zip.File) error {

		rc, err := f.Open()

		if err != nil {

			return errors.NewError(err.Error())

		}

		defer func() {

			if err := rc.Close(); err != nil {

				panic(err)

			}

		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return errors.NewError(err.Error())
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return errors.NewError(err.Error())
			}
		}
		return nil

	}

	for _, f := range r.File {

		err := extractAndWriteFile(f)

		if err != nil {

			return errors.NewError(err.Error())

		}

	}

	return nil

}

// ValidateChecksum will validate the checksum with the one that has been supplied at update.json to ensure it has been downloaded properly and not modified by a third party.
func ValidateChecksum(filePath, expectedChecksum string) bool {

	hasher := sha256.New()
	s, err := ioutil.ReadFile(filePath)
	hasher.Write(s)

	if err != nil {

		errors.NewError("unable to validate checksum of file")
		return false

	}

	actualChecksum := hex.EncodeToString(hasher.Sum(nil))

	return actualChecksum == expectedChecksum

}
