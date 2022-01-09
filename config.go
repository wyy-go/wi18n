package wi18n

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	defaultFormatBundleFile = "yaml"
	defaultRootPath         = "./_example/localize"
)

var (
	defaultLanguage       = language.English
	defaultUnmarshalFunc  = yaml.Unmarshal
	defaultAcceptLanguage = []language.Tag{
		defaultLanguage,
		language.German,
		language.French,
	}

	defaultBundleConfig = &Config{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
	}
)

type (
	// GetLngHandler ...
	GetLngHandler = func(context *gin.Context, defaultLng string) string

	// Option ...
	Option func(GinI18n)
)

// Config ...
type Config struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
}

// WithBundle ...
func WithBundle(cfg *Config) Option {
	return func(g GinI18n) {
		g.setBundleConfig(cfg)
	}
}

// WithGetLngHandle ...
func WithGetLngHandle(handler GetLngHandler) Option {
	return func(g GinI18n) {
		g.setGetLngHandler(handler)
	}
}
