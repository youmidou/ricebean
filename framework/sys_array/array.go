package sys_array

func Delete(slice []*interface{}, index int) {
	slice = append(slice[:index], slice[index+1:]...)

}
