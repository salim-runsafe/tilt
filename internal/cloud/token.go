package cloud

import (
	"os"

	"github.com/google/uuid"
	"github.com/windmilleng/wmclient/pkg/dirs"

	"github.com/windmilleng/tilt/pkg/model"
)

const tokenFileName = "token"

func GetOrCreateToken(dir *dirs.WindmillDir) (model.CloudToken, error) {
	token, err := getExistingToken(dir)
	if os.IsNotExist(err) {
		u := uuid.New()
		newtoken := model.CloudToken(u.String())
		err := writeToken(dir, newtoken)
		if err != nil {
			return "", err
		}
		return newtoken, nil
	} else if err != nil {
		return "", err
	}

	return token, nil
}

func getExistingToken(dir *dirs.WindmillDir) (model.CloudToken, error) {
	token, err := dir.ReadFile(tokenFileName)
	if err != nil {
		return "", err
	}
	return model.CloudToken(token), nil
}

func writeToken(dir *dirs.WindmillDir, t model.CloudToken) error {
	return dir.WriteFile(tokenFileName, string(t))
}
