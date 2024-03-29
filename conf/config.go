package conf

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Context contains details on AiT instance.
type Context struct {
	Auth   string
	Domain string
	Name   string
}

// Config contains details on all config settings.
type Config struct {
	CurrentContext string
	Contexts       []Context
}

// Validate loops contexts to confirm if named context is valid.
func (c *Config) Validate(name string) bool {
	for _, x := range c.Contexts {
		if x.Name == name {
			return true
		}
	}
	return false
}

// GetContext returns the named context.
func (c *Config) GetContext(name string) (*Context, error) {
	for i, x := range c.Contexts {
		if x.Name == name {
			return &c.Contexts[i], nil
		}
	}
	return nil, errors.New("Context not found")
}

// AddContext adds new context to config.
func (c *Config) AddContext(context Context) error {
	if c.Validate(context.Name) {
		return errors.New("Context already exists")
	}
	c.Contexts = append(c.Contexts, context)
	return nil
}

// GetCurrentContext returns the current context.
func (c *Config) GetCurrentContext() (*Context, error) {
	return c.GetContext(c.CurrentContext)
}

// UpdateContext apply updates to named context.
func (c *Config) UpdateContext(name string, auth Auth, domain string) {
	context, err := c.GetContext(name)
	if err != nil {
		fmt.Println("Error getting context:", err)
		return
	}

	context.Auth = auth.Encode()
	if domain != "" {
		context.Domain = domain
	}
}

// LoadConfig to load yaml config from file.
func (c *Config) LoadConfig(path string) error {
	if c.ExistsConfig(path) != true {
		fmt.Println("Warning config file not found:", path)
		err := c.WriteConfig(path)
		if err != nil {
			return err
		}
		fmt.Println("Created config file:", path)
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return err
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		fmt.Println("Error parsing yaml config file:", err)
		return err
	}
	return nil
}

// WriteConfig to write any changed config back to file.
func (c *Config) WriteConfig(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("Error parsing config to yaml:", err)
		return err
	}
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("Error writing config to file:", err)
		return err
	}
	return nil
}

// ExistsConfig reports whether the config file exists.
func (c *Config) ExistsConfig(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
