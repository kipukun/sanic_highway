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
	"github.com/kipukun/sanic_highway/db"
	"github.com/kipukun/sanic_highway/http"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Println("[*] initializing db...")
	d, err := db.Init("postgres://postgres:cock@127.0.0.1/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer d.Conn.Close()

	user := flag.Bool("user", false, "whether to insert a new user or not")
	flag.Parse()
	if *user {
		err = insert(d)
		if err != nil {
			fmt.Printf("error inserting user: %v\n", err)
			return
		}
		fmt.Println("\n[*] user successfully created!")
		return
	}

	fmt.Println("[*] starting up http...")
	srv := http.Init(d)
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
