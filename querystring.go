package maestropanel

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
)

func convertObj2Form(writer *multipart.Writer, obj interface{}) (err error) {
	objVal := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)

	fieldCount := objVal.NumField()

	for i := 0; i < fieldCount; i++ {
		field := objVal.Field(i)

		tagName := objType.Field(i).Tag.Get("json")

		if tagName == "" {
			tagName = objType.Field(i).Name
		}

		fieldType := objType.Field(i).Tag.Get("type")
		v := toString(field)

		if fieldType == "filepath" {
			err := addFile2Writer(writer, tagName, v[0])
			if err != nil {
				return err
			}
		} else {
			if len(v) > 0 {
				for _, item := range v {
					if item != "" {
						writer.WriteField(tagName, item)
					}
				}
			}
		}
	}

	return err
}

func addFile2Writer(w *multipart.Writer, fieldName string, filePath string) (err error) {
	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(filePath))

	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)

	return err
}
