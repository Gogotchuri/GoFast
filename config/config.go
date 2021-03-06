package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

const configFilename = "config/.env"

var conf *Config = nil //Global config instance
var once sync.Once

/*Config parameters*/
type Config struct {
	Domain   string         `json:"domain"`
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Pepper   string         `json:"pepper"`
	ServMail string         `json:"service_email"`
	ServPass string         `json:"service_password"`
	HMACKey  string         `json:"hmac_key"`
	Database PostgresConfig `json:"database"`

	Redis    RedisConfig    `json:"redis"`
	JWT      JWTConfig      `json:"jwt"`

	Google 	 SocialProviderConfig   `json:"google"`
	Facebook SocialProviderConfig 	`json:"facebook"`
}

/*JWTConfig secrets*/
type JWTConfig struct {
	AccessSecret  string        `json:"access_secret"`
	RefreshSecret string        `json:"refresh_secret"`
	AccessExp     time.Duration `json:"access_exp"`
	RefreshExp    time.Duration `json:"refresh_exp"`
}

/*Google oauth config*/
type SocialProviderConfig struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

/*OtacConfig parameters*/
type OtacConfig struct {
	EntryPrefix string `json:"otac_prefix"` // Prefix of otac key(prefix+email) in redis
	Expires     int64  `json:"otac_expire"` // Time before expiration
}

/*RedisConfig parameters*/
type RedisConfig struct {
	Host string     `json:"host"`
	Port int        `json:"port"`
	OTAC OtacConfig `json:"otac"`
}

/*PostgresConfig parameters*/
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

/*Dialect returns config dialect*/
func (PostgresConfig) Dialect() string {
	return "postgres"
}

/*DBConnectionInfo returns database connection info*/
func (c *PostgresConfig) DBConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}

/*IsProd returns true if current environment is production*/
func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

/*LoadConfig returns config for the app
 *Takes config filename as a argument
 *If empty string is passed, uses default file
 */
func LoadConfig(file string) *Config {
	if file == "" {
		file = configFilename
	}
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	conf = &c
	return &c
}

/*GetInstance Returns config singleton object,
 * if not initialized: returns after initialization
 */
func GetInstance() *Config {
	once.Do(func() {
		conf = LoadConfig(configFilename)
	})
	return conf
}
