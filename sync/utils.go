package sync

import (
	"io/ioutil"

	"encoding/json"

	"os"
	"path/filepath"

	"sync-mysql/errors"
	"net/url"
	"fmt"
	"bytes"
)

var (
	UTF8JSON = "# -*- coding:utf-8 -*-\n\n"
)


func AssetExists(relative string) (bool, error) {

	abs, err := filepath.Abs(os.Args[0])

	if err != nil {
		return false, errors.ToFormatError(
			err,
			"read asset(%s) fail.", relative,
		)
	}

	abs, _ = filepath.Split(abs)

	filename := filepath.Join(abs, "asset", relative)

	fmt.Println(filename)


	_, err = os.Stat(filename)


	if err != nil {

		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.ToFormatError(
			err,
			"check asset(%s) exist fail.", relative,
		)
	}

	return true, nil

}

func ReadAsString(filename string) (string,error){

	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return "",errors.ToFormatError(
			err,
			"read bytes from json file %s fail.", filename,
		)
	}

	return string(b),nil
}

func ReadAssetAsJSON(relative string, v interface{}) error {

	abs, err := filepath.Abs(os.Args[0])

	if err != nil {
		return errors.ToFormatError(
			err,
			"read asset(%s) fail.", relative,
		)
	}

	abs, _ = filepath.Split(abs)

	return ReadAsJSON(
		filepath.Join(abs, "asset", relative),
		v,
	)
}

func ReadFileList(relative string) ([]os.FileInfo,string, error){

	abs, err := filepath.Abs(os.Args[0])

	if err != nil {
		return nil,abs,errors.ToFormatError(
			err,
			"read asset(%s) fail.", relative,
		)
	}

	abs, _ = filepath.Split(abs)

	parent := filepath.Join(abs,"asset",relative)
	dirList, e := ioutil.ReadDir(parent)
	if e != nil {
		return nil,abs,errors.ToFormatError(
			err,
			"read asset(%s) fail.", relative,
		)
	}
	return dirList,parent,nil

}

func SaveAssetAsJSON(relative string, v interface{}) error {

	abs, err := filepath.Abs(os.Args[0])

	if err != nil {
		return errors.ToFormatError(
			err,
			"save asset(%s) fail.", relative,
		)
	}

	abs, _ = filepath.Split(abs)

	return SaveAsJson(
		filepath.Join(abs, "asset", relative),
		v,
	)
}

func SaveAssetFile(relative string, bytes []byte) error {

	abs, err := filepath.Abs(os.Args[0])

	if err != nil {
		return errors.ToFormatError(
			err,
			"save asset(%s) fail.", relative,
		)
	}

	abs, _ = filepath.Split(abs)

	return SaveFile(
		filepath.Join(abs, "asset", relative),
		bytes,
	)
}


func SaveFile(filename string, bytes []byte) error {

	dir, _ := filepath.Split(filename)

	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}


	if err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}


	if err := ioutil.WriteFile(filename, bytes, 0755); err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}

	return nil

}

func BytesCombine(seq string,pBytes ...[]byte) []byte {

	l := len(pBytes)
	s := make([][]byte, l)
	for index := 0; index < l; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte(seq)
	return bytes.Join(s, sep)
}


func SaveAsJson(filename string, v interface{}) error {

	dir, _ := filepath.Split(filename)

	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}

	jsonBytes, err := json.MarshalIndent(
		v,
		"",
		"    ")

	if err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}


	if err := ioutil.WriteFile(filename, jsonBytes, 0755); err != nil {
		return errors.ToFormatError(
			err,
			"marshal json to file %s fail.", filename,
		)
	}

	return nil

}

func ReadAsJSON(filename string, v interface{}) error {

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.ToFormatError(
			err,
			"read bytes from json file %s fail.", filename,
		)
	}

	err = json.Unmarshal(bytes, v)

	if err != nil {
		return errors.ToFormatError(
			err,
			"unmarshal from json file %s fail.", filename,
		)
	}

	return nil
}

type ConnectScheme struct {
	Language string
	Scheme string
	Username string
	Password string
	Host string
	Path string
	Fragment string
}

func ParseScheme(language string) (*ConnectScheme,error){
	fmt.Println(language)

	u, err := url.Parse(language)

	if err != nil{
		return nil,errors.ToFormatError(
			err,
			"parse language url error ",
		)
	}
	pwd,_ := u.User.Password()
	var result = &ConnectScheme{
		Language: language,
		Scheme :  u.Scheme,
		Username: u.User.Username(),
		Password: pwd,
		Host:     u.Host,
		Path:     u.Path,
		Fragment: u.Fragment,
	}
	return result,nil

}

func (scheme *ConnectScheme) ToGoMysql() string{

	return fmt.Sprintf("%s:%s@tcp(%s)/",
	scheme.Username,
	scheme.Password,
	scheme.Host,
	)
}

func (scheme *ConnectScheme) ToGoMysqlAndDB() string{

	return fmt.Sprintf("%s:%s@tcp(%s)%s",
		scheme.Username,
		scheme.Password,
		scheme.Host,
		scheme.Path,
	)
}

func (scheme *ConnectScheme) ToDataXMysql(db string) string{
	path := scheme.Path
	if db != ""{
		path = "/"+db
	}
	return fmt.Sprintf("jdbc:mysql://%s%s",
	scheme.Host,
		path)
}