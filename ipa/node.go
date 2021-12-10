package ipa

import "errors"

func (node Node) GetConfigDrive() (ConfigDrive, error) {
	result := ConfigDrive{}

	cd, ok := node.InstanceInfo["configdrive"]
	if !ok {
		return result, errors.New("Node does not have a configdrive")
	}

	cdMap, ok := cd.(map[string]interface{})
	if !ok {
		return result, errors.New("Configdrive is not a mapping; pre-built configdrives are not supported")
	}

	userData, ok := cdMap["user_data"]
	if ok {
		userDataString, ok := userData.(string)
		if ok {
			result.UserData = userDataString
		} else {
			return result, errors.New("User data is not a string")
		}
	}

	return result, nil
}
