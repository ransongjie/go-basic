package main

import "com.xcrj.gobasic/internal/data_type"
import "com.xcrj.gobasic/internal/operator"
import "com.xcrj.gobasic/internal/control"
import "com.xcrj.gobasic/internal/generic"

func main() {
	data_type.RunDataType()
	data_type.RunDataTypeConvert()
	data_type.RunDataTypeLen()
	data_type.RunDataTypeDefault()

	operator.RunOperator()

	control.RunCodeLogic()

	generic.RunGeneric();
}
