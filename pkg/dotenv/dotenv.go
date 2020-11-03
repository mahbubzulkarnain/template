package dotenv

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

var rootDir = "template-go"

func Load(args ...string) string {
	var key, defaultValue string
	if len(args) < 1 {
		return ""
	}

	key = args[0]
	if len(args) > 1 {
		defaultValue = args[1]
	}

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	loadFile()
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func loadFile(dirEnv ...string) {
	if len(dirEnv) > 0 {
		rootDir = dirEnv[0]
	}
	cwd, _ := os.Getwd()

	re := regexp.MustCompile(`^(.*` + rootDir + `)`)
	rootPath := string(re.Find([]byte(cwd)))
	for _, v := range []string{"/", "\\"} {
		cwdSplit := strings.Split(cwd, v)
		rootPathLength := len(strings.Split(rootPath, v))

		filePathENV := strings.Join(cwdSplit[:rootPathLength], v) + v + `.env`

		matches, _ := filepath.Glob(filePathENV)
		if len(matches) != 0 {
			_ = godotenv.Load(filePathENV)
			break
		}
	}
}
