package generateType

const (
	ServiceParameters = "service_parameters"
)

func IsValid(s string) bool {
	return ServiceParameters == s
}
