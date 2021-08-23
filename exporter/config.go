package exporter

type Config struct {
	InstanceId      string `yaml:"instance_id"`
	InstanceAlias   string `yaml:"instance_alias"`
	AccessKeyId     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}
