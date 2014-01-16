package app

import (
	"git.oschina.net/fuxiaohei/GoBlog.git/app/utils"
	"io/ioutil"
	"os"
	"path"
)

func LogError(bytes []byte) {
	dir := App.Config().StringOr("app.log_dir", "tmp/log")
	file := path.Join(dir, utils.DateInt64(utils.Now(), "MMDDHHmmss.log"))
	ioutil.WriteFile(file, bytes, os.ModePerm)
}
