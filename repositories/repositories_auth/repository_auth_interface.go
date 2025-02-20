package repositories_auth

type RepositoryAuthInterface interface {
	Issue(payload string) (string, error)
	Validate(token string) (int, bool)
}
