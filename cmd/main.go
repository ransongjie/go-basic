package main

import "com.xcrj.gobasic/internal/datatypeme"
import "com.xcrj.gobasic/internal/operatorme"
import "com.xcrj.gobasic/internal/codelogicme"

func main() {
	datatypeme.RunDataType()
	datatypeme.RunDataTypeConvert()
	datatypeme.RunDataTypeLen()
	datatypeme.RunDataTypeDefault()

	operatorme.RunOperator()

	codelogicme.RunCodeLogic()
}
