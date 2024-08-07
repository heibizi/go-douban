package douban

// GetMediaType 根据 code 获取 MediaType
func GetMediaType(code string) MediaType {
	for _, mediaType := range MediaTypes {
		if mediaType.Code == code {
			return mediaType
		}
	}
	return Movie
}
