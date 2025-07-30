package cloud

type CloudDb struct {
	url string
}

func NewCloudDB(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (cloudDb *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (cloudDb *CloudDb) Write(content []byte) {}
