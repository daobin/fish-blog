package util

// CheckState 校验状态是否有效
func CheckState(state int) bool {
	return state == StateEnable || state == StateDisable || state == StateDelete
}

// CheckObjType 校验对象类型是否有效
func CheckObjType(objType string) bool {
	return objType == ObjTypeCate || objType == ObjTypeArticle || objType == ObjTypeUser
}
