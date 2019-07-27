package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/kipukun/sanic_highway/config"
	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/http"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var c config.Config
	fmt.Println("[*] reading config from config.toml...")
	f, err := os.Open("config.toml")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[!] config.toml does not exist. using default config")
			c = config.DefaultConfig
		} else {
			log.Fatal(err)
		}
	}
	if c != config.DefaultConfig {
		c, err = config.Load(f)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
	fmt.Println("[*] initializing db...")
	d, err := db.Init(c)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	user := flag.Bool("user", false, "whether to insert a new user or not")
	flag.Parse()
	if *user {
		err = insert(d)
		if err != nil {
			fmt.Printf("\nerror inserting user: %v\n", err)
			return
		}
		fmt.Println("\n[*] user successfully created!")
		return
	}

	fmt.Printf("[*] starting up http at addr %s...\n", c.Web.Addr)
	srv := http.Init(d, c)
	srv.Start()

}

// insert inserts a new user
func insert(d *db.Database) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("[?] username: ")
	user, _ := reader.ReadString('\n')

	fmt.Print("[?] password: ")
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword(pass, 8)
	if err != nil {
		return err
	}
	_, err = d.InsertUser.Exec(id, strings.TrimSpace(user), string(hash))
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code.Name() {
			case "unique_violation":
				return errors.New("username already exists")
			}
		}
		return err
	}

	return nil
}
