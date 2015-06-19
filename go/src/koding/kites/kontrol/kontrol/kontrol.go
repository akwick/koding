package kontrol

import (
	"errors"
	"fmt"
	"io/ioutil"
	"koding/db/mongodb/modelhelper"
	"log"
	"time"

	"github.com/koding/kite"
	"github.com/koding/kite/config"
	"github.com/koding/kite/kontrol"
)

const Version = "0.0.6"

func New(c *Config) *kontrol.Kontrol {
	modelhelper.Initialize(c.MongoURL)

	publicKey, err := ioutil.ReadFile(FindPath(c.PublicKey))
	if err != nil {
		log.Fatalln(err.Error())
	}

	privateKey, err := ioutil.ReadFile(FindPath(c.PrivateKey))
	if err != nil {
		log.Fatalln(err.Error())
	}

	kiteConf := config.MustGet()

	if c.Environment != "" {
		kiteConf.Environment = c.Environment
	}

	if c.Region != "" {
		kiteConf.Region = c.Region
	}

	if c.Port != 0 {
		kiteConf.Port = c.Port
	}

	kon := kontrol.New(kiteConf, Version, string(publicKey), string(privateKey))
	kon.AddAuthenticator("sessionID", authenticateFromSessionID)
	kon.MachineAuthenticate = authenticateMachine

	switch c.Storage {
	case "etcd":
		kon.SetStorage(kontrol.NewEtcd(c.Machines, kon.Kite.Log))
	case "postgres":
		postgresConf := &kontrol.PostgresConfig{
			Host:     c.Postgres.Host,
			Port:     c.Postgres.Port,
			Username: c.Postgres.Username,
			Password: c.Postgres.Password,
			DBName:   c.Postgres.DBName,
		}

		kon.SetStorage(kontrol.NewPostgres(postgresConf, kon.Kite.Log))
	default:
		panic(fmt.Sprintf("storage is not found: '%'", c.Storage))
	}

	if c.TLSKeyFile != "" && c.TLSCertFile != "" {
		kon.Kite.UseTLSFile(c.TLSCertFile, c.TLSKeyFile)
	}

	return kon
}

func authenticateFromSessionID(r *kite.Request) error {
	username, err := findUsernameFromSessionID(r.Auth.Key)
	if err != nil {
		return err
	}

	r.Username = username

	return nil
}

func findUsernameFromSessionID(sessionID string) (string, error) {
	session, err := modelhelper.GetSession(sessionID)
	if err != nil {
		return "", err
	}

	return session.Username, nil
}

func authenticateMachine(authType string, r *kite.Request) error {
	switch authType {
	case "password":
		password, err := r.Client.TellWithTimeout(
			"kite.getPass",
			10*time.Minute,
			"Enter password: ",
		)

		if err != nil {
			return err
		}

		_, err = modelhelper.CheckAndGetUser(r.Client.Kite.Username, password.MustString())
		if err != nil {
			return err
		}
	case "token":
		var args struct {
			Token string
		}

		if err := r.Args.One().Unmarshal(&args); err != nil {
			return err
		}

		if args.Token == "" {
			return errors.New("token is empty")
		}

		// Try to fetch the token and remove it. If it doesn't exist it'll will
		// return an error. If it's exist it'll be deleted and a nil error
		// (means success) will be returned. The underlying implementation uses
		// findAndModify so it's consistent across each kontrol.
		session, err := modelhelper.GetSessionFromToken(args.Token)
		if err != nil {
			return err
		}

		if err := modelhelper.RemoveToken(session.ClientId); err != nil {
			return err
		}

		// prevent using a wrong username
		r.Client.Kite.Username = session.Username
		r.Client.Username = session.Username
		return nil
	default:
		return errors.New("authentication type for machine registration is not defined")
	}

	// everything is ok, succefully validated
	return nil
}
