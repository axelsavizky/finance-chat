package configuration

const (
	Port         = ":8000"
	CertFilePath = ".localhost-ssl/localhost.crt"
	KeyFilePath  = ".localhost-ssl/localhost.key"

	DatabaseDriverName = "mysql"

	DatabaseUsername   = "root"
	DatabasePassword   = "root" // This must be stored with security in a vault
	DatabaseSchemaName = "finance-chat"

	ChatBotBasePath = "http://localhost:8001/%s"
)
