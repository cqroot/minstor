package drive

import (
	"os"

	"github.com/cqroot/minstor/internal/logger"
	"golang.org/x/exp/slog"
)

type Drive struct {
	Path string
}

func (d Drive) Write(name string, content []byte) error {
	name = d.fileName(name)
	logger.Debug("Start writing content to the file.", slog.String("name", name))

	err := os.WriteFile(name, content, 0o666)
	if err != nil {
		logger.ErrorE("Failed to write content to the file.", err,
			slog.String("name", name))
		return err
	}

	return nil
}

func (d Drive) Read(name string) ([]byte, error) {
	name = d.fileName(name)
	logger.Debug("Start reading content from the file.", slog.String("name", name))

	content, err := os.ReadFile(name)
	if err != nil {
		logger.ErrorE("Failed to read content from the file.", err,
			slog.String("name", name))
		return nil, err
	}

	return content, nil
}

func (d Drive) Delete(name string) error {
	name = d.fileName(name)
	logger.Debug("Start deleting the file.", slog.String("name", name))

	err := os.Remove(name)
	if err != nil {
		logger.ErrorE("Failed to remove the file.", err,
			slog.String("name", name))
		return err
	}

	return nil
}
