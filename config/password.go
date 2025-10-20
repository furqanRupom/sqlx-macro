package config

import (
	"github.com/furqanrupom/sqlx-macro/util"
	"log"
	"os"
)

type PasswordConfig struct {
	Memory  uint32
	Time    uint32
	Threads uint8
	KeyLen  uint32
}

func PasswordConf() *PasswordConfig {
	memory := os.Getenv("MEMORY")
	if memory == "" {
		log.Println("Memory is required!")
		os.Exit(1)
	}
	memoryInt, _ := util.ParseInt(memory)
	threads := os.Getenv("THREADS")
	if threads == "" {
		log.Println("Threads is required!")
		os.Exit(1)
	}
	threadsInt, _ := util.ParseInt(threads)

	time := os.Getenv("TIME")
	if time == "" {
		log.Println("Time is required!")
		os.Exit(1)
	}
	timeInt, _ := util.ParseInt(time)
	keyLen := os.Getenv("KEY_LEN")
	if keyLen == "" {
		log.Println("KeyLen is required!")
		os.Exit(1)
	}
	keyLenInt, _ := util.ParseInt(keyLen)

	return &PasswordConfig{
		Memory:  uint32(memoryInt),
		Threads: uint8(threadsInt),
		Time:    uint32(timeInt),
		KeyLen:  uint32(keyLenInt),
	}

}
