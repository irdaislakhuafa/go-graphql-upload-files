package files

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
)

func SaveFile(ctx context.Context, targetPath, fileName string, file graphql.Upload, mode fs.FileMode) (string, error) {
	targetFile := fmt.Sprintf("%s/%s", targetPath, fileName)
	_, err := os.Stat(targetPath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(targetPath, mode); err != nil {
			log.Println(err)
			return "", err
		}
	}

	bytesStream, err := io.ReadAll(file.File)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := os.WriteFile(targetFile, bytesStream, mode); err != nil {
		log.Println(err)
		return "", err
	}

	return fileName, nil
}
