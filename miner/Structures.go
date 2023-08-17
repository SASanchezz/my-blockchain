package Miner

type Config struct {
	Server struct {
		SeedNodeHost string `envconfig:"SEED_NODE_HOST"`
		SeedNodePort string `envconfig:"SEED_NODE_PORT"`
	}
}
