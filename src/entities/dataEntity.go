package entities

type DataEntity struct {
	Header []string
	Data   [][]string
}

func NewDataEntity(header []string, data [][]string) *DataEntity {
	return &DataEntity{
		Header: header,
		Data:   data,
	}
}

func (d *DataEntity) GetHeader() []string {
	return d.Header
}

func (d *DataEntity) GetData() [][]string {
	return d.Data
}

func (d *DataEntity) SetHeader(header []string) {
	d.Header = header
}

func (d *DataEntity) SetData(data [][]string) {
	d.Data = data
}

func (d *DataEntity) GetRow(row int) []string {
	return d.Data[row]
}
