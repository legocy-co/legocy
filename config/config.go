package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var appConfigInstance *AppConfig // private singleton variable

type AppConfig struct {
	BaseURL string `json:"base_url"`

	DbConf    DatabaseConfig `yaml:"database" json:"database"`
	JwtConf   JWTConfig      `yaml:"jwt" json:"jwt"`
	KafkaConf KafkaConfig    `yaml:"kafka" json:"kafka"`

	S3Host     string `json:"s3_host"`
	S3Port     string `json:"s3_port"`
	CDNBaseURL string `json:"cdn_base_url"`
}

func GetAppConfig() *AppConfig {
	return appConfigInstance
}

func SetAppConfig(cfg *AppConfig) error {
	if appConfigInstance != nil {
		return ErrConfigAlreadyExists
	}

	appConfigInstance = cfg
	return nil
}

type DatabaseConfig struct {
	Hostname     string `yaml:"hostname" json:"hostname"`
	Port         int    `yaml:"port" json:"port"`
	DbName       string `yaml:"db_name" json:"db_name"`
	DbUser       string `yaml:"db_user" json:"db_user"`
	DbPassword   string `yaml:"db_password" json:"db_password"`
	LoadFixtures bool   `yaml:"load_fixtures" json:"load_fixtures"`
}

type JWTConfig struct {
	SecretKey            string `yaml:"secret_key" json:"secret_key"`
	AccessTokenLifeTime  int    `yaml:"access_token_lifetime_hours" json:"access_token_lifetime_hours"`
	RefreshTokenLifeTime int    `json:"refresh_token_life_time_hours"`
}

type KafkaConfig struct {
	URI             string `json:"uri"`
	ConsumerGroupId string `json:"consumer_group_id"`
}

func GetDBConfig() *DatabaseConfig {
	cfg := GetAppConfig()
	if cfg == nil {
		return nil
	}

	return &cfg.DbConf
}

func SetupFromEnv() error {

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "localhost"
	}

	dbHost := os.Getenv("DB_HOST")
	log.Printf("DB_HOST = %v", dbHost)
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	log.Printf("DB_PORT = %v", dbPort)
	dbUser := os.Getenv("DB_USER")
	log.Printf("DB_USER = %v", dbUser)
	dbPassword := os.Getenv("DB_PASSWORD")
	log.Printf("DB_PASSWORD = %v", dbPassword)
	dbDatabaseName := os.Getenv("DB_DATABASE_NAME")
	loadFixtures := os.Getenv("DB_LOAD_FIXTURES") == "true"

	dbConfig := DatabaseConfig{
		Hostname:     dbHost,
		Port:         dbPort,
		DbName:       dbDatabaseName,
		DbUser:       dbUser,
		DbPassword:   dbPassword,
		LoadFixtures: loadFixtures,
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	jwtAccessTokenLifetimeHours, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_LIFETIME"))
	jwtRefreshTokenLifetimeHours, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_LIFETIME"))
	log.Printf("REFRESH_TOKEN = %v", jwtRefreshTokenLifetimeHours)

	jwtConfig := JWTConfig{
		SecretKey:            jwtSecretKey,
		AccessTokenLifeTime:  jwtAccessTokenLifetimeHours,
		RefreshTokenLifeTime: jwtRefreshTokenLifetimeHours,
	}

	kafkaUri := os.Getenv("KAFKA_URI")
	log.Printf("KAFKA_URL = %v", kafkaUri)
	kafkaConsumerGroupId := os.Getenv("KAFKA_CONSUMER_GROUP_ID")
	log.Printf("KAFKA_CONSUMER_GROUP_ID = %v", kafkaConsumerGroupId)

	kafkaConfig := KafkaConfig{kafkaUri, kafkaConsumerGroupId}

	s3Host := os.Getenv("S3_HOST")
	s3Port := os.Getenv("S3_PORT")
	cdnBaseUrl := os.Getenv("CDN_BASE_URL")

	appConfig := AppConfig{
		BaseURL:    baseUrl,
		DbConf:     dbConfig,
		JwtConf:    jwtConfig,
		KafkaConf:  kafkaConfig,
		S3Host:     s3Host,
		S3Port:     s3Port,
		CDNBaseURL: cdnBaseUrl,
	}

	return SetAppConfig(&appConfig)
}
