package logrus

import (
	"errors"
	"github.com/alecthomas/units"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/xhit/go-str2duration/v2"
	"io"
	"path/filepath"
)

var ErrNilOption = errors.New("option can not be nil")

type logBundle struct {
	logger *logrus.Logger
}

type Option struct {
	Path, Pattern, MaxAge, RotationTime, RotationSize string
}

//New is constructor of the logrus package
func New(opt *Option) (log.Logger, error) {

	if opt == nil {
		return nil, ErrNilOption
	}
	l := &logBundle{logger: logrus.New()}
	writer, err := getLoggerWriter(opt)
	if err != nil {
		return nil, err
	}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l, nil
}

func getLoggerWriter(opt *Option) (io.Writer, error) {
	maxAge, err := str2duration.ParseDuration(opt.MaxAge)
	if err != nil {
		return nil, err
	}
	rotationTime, err := str2duration.ParseDuration(opt.RotationTime)
	if err != nil {
		return nil, err
	}
	rotationSize, err := units.ParseBase2Bytes(opt.RotationSize)
	if err != nil {
		return nil, err
	}
	return rotators.New(
		filepath.Join(opt.Path, opt.Pattern),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
		rotators.WithRotationSize(int64(rotationSize)),
	)
}
func (l *logBundle) Info(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Info(field.Message)
}

func (l *logBundle) Warning(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Warning(field.Message)
}

func (l *logBundle) Error(field *log.Field) {
	l.logger.WithFields(logrus.Fields{
		"section":  field.Section,
		"function": field.Function,
		"params":   field.Params,
	}).Error(field.Message)
}
