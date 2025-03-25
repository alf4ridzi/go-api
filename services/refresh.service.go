package services

type RefreshService struct{}

func NewRefreshService() *RefreshService {
	return &RefreshService{}
}

func (r *RefreshService) RefreshToken(token string) (string, error) {
	return "", nil
}
