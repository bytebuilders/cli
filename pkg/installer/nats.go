package installer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nkeys"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

func confDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(errors.Wrap(err, "failed to detect current working dir"))
	}
	return filepath.Join(dir, "ace-installer")
}

func natsDir() string {
	return filepath.Join(confDir(), "nats")
}

func genNatsCredentials() (map[string]string, error) {
	fmt.Println("Configuration directory: ", natsDir())

	nc := map[string]string{}

	oKp, oPub, oSeed, oJwt, err := createOperator("Operator")
	if err != nil {
		return nil, err
	}
	if err := storeOperator(natsDir(), "Operator", oPub, oSeed, oJwt, nc); err != nil {
		return nil, err
	}

	sKp, sPub, sSeed, sJwt, err := createAccount("SYS", oKp)
	if err != nil {
		return nil, err
	}
	if err := storeAccount(natsDir(), "Operator", "SYS", sPub, sSeed, sJwt, nc); err != nil {
		return nil, err
	}

	_, suPub, suSeed, suJwt, err := createUser("sys", sKp)
	if err != nil {
		return nil, err
	}
	if err := storeUser(natsDir(), "Operator", "SYS", "sys", suPub, suSeed, suJwt, nc); err != nil {
		return nil, err
	}

	aKp, aPub, aSeed, aJwt, err := createAccount("Admin", oKp)
	if err != nil {
		return nil, err
	}
	if err := storeAccount(natsDir(), "Operator", "Admin", aPub, aSeed, aJwt, nc); err != nil {
		return nil, err
	}

	_, auPub, auSeed, auJwt, err := createUser("admin", aKp)
	if err != nil {
		return nil, err
	}
	if err := storeUser(natsDir(), "Operator", "Admin", "admin", auPub, auSeed, auJwt, nc); err != nil {
		return nil, err
	}

	{
		data, err := yaml.Marshal(nc)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(filepath.Join(natsDir(), "nats-credentials.yaml"), data, 0o644)
		if err != nil {
			return nil, err
		}

		ncEnc := map[string]string{}
		for k, v := range nc {
			ncEnc[k] = base64.StdEncoding.EncodeToString([]byte(v))
		}
		data, err = yaml.Marshal(ncEnc)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(filepath.Join(natsDir(), "nats-credentials.enc.yaml"), data, 0o644)
		if err != nil {
			return nil, err
		}
	}

	return nc, nil
}

func createOperator(name string) (nkeys.KeyPair, string, []byte, string, error) {
	oKp, err := nkeys.CreateOperator()
	if err != nil {
		return nil, "", nil, "", err
	}
	oPub, err := oKp.PublicKey()
	if err != nil {
		return nil, "", nil, "", err
	}

	oSeed, err := oKp.Seed()
	if err != nil {
		return nil, "", nil, "", err
	}
	claim := jwt.OperatorClaims{
		ClaimsData: jwt.ClaimsData{
			Audience:  oPub,
			Expires:   time.Now().AddDate(100, 0, 0).Unix(), // never expire
			ID:        oPub,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "AppsCode Inc.",
			Name:      oPub,
			NotBefore: time.Now().Unix(),
			Subject:   oPub,
		},
		Operator: jwt.Operator{
			SigningKeys: jwt.StringList{oPub},
		},
	}
	// claim := jwt.NewOperatorClaims(oPub)
	claim.Name = name
	oJwt, err := claim.Encode(oKp)
	if err != nil {
		return nil, "", nil, "", err
	}

	return oKp, oPub, oSeed, oJwt, nil
}

func createAccount(name string, oKp nkeys.KeyPair) (nkeys.KeyPair, string, []byte, string, error) {
	aKp, err := nkeys.CreateAccount()
	if err != nil {
		return nil, "", nil, "", err
	}
	aPub, err := aKp.PublicKey()
	if err != nil {
		return nil, "", nil, "", err
	}
	aSeed, err := aKp.Seed()
	if err != nil {
		return nil, "", nil, "", err
	}
	claim := jwt.NewAccountClaims(aPub)
	claim.Name = name
	if name != "SYS" {
		claim.Limits.JetStreamLimits = jwt.JetStreamLimits{
			MemoryStorage: -1,
			DiskStorage:   -1,
			Streams:       -1,
			Consumer:      -1,
		}
	}
	aJwt, err := claim.Encode(oKp)
	if err != nil {
		return nil, "", nil, "", err
	}

	return aKp, aPub, aSeed, aJwt, nil
}

func createUser(name string, aKp nkeys.KeyPair) (nkeys.KeyPair, string, []byte, string, error) {
	uKp, err := nkeys.CreateUser()
	if err != nil {
		return nil, "", nil, "", err
	}
	uSeed, err := uKp.Seed()
	if err != nil {
		return nil, "", nil, "", err
	}

	uPub, err := uKp.PublicKey()
	if err != nil {
		return nil, "", nil, "", err
	}

	uClaim := jwt.NewUserClaims(uPub)
	uClaim.Name = name

	uJwt, err := uClaim.Encode(aKp)
	if err != nil {
		return nil, "", nil, "", err
	}

	return uKp, uPub, uSeed, uJwt, nil
}

func storeOperator(dir, op string, pub string, seed []byte, jwt string, nc map[string]string) error {
	return storeInfo(
		filepath.Join(dir, "keys", pub+".nk"),
		filepath.Join(dir, "stores", op, op+".jwt"),
		filepath.Join(dir, "creds", op, op+".creds"),
		pub,
		seed,
		jwt,
		nc,
	)
}

func storeAccount(dir, op, ac string, pub string, seed []byte, jwt string, nc map[string]string) error {
	return storeInfo(
		filepath.Join(dir, "keys", pub+".nk"),
		filepath.Join(dir, "stores", op, "accounts", ac, ac+".jwt"),
		filepath.Join(dir, "creds", op, "accounts", ac, ac+".creds"),
		pub,
		seed,
		jwt,
		nc,
	)
}

func storeUser(dir, op, ac, user string, pub string, seed []byte, jwt string, nc map[string]string) error {
	return storeInfo(
		filepath.Join(dir, "keys", pub+".nk"),
		filepath.Join(dir, "stores", op, "accounts", ac, "users", user+".jwt"),
		filepath.Join(dir, "creds", op, "accounts", ac, "users", user+".creds"),
		pub,
		seed,
		jwt,
		nc,
	)
}

func storeInfo(keyFile, jwtFile, credFile string, pub string, seed []byte, jwt string, nc map[string]string) error {
	// /keys
	if err := os.MkdirAll(filepath.Dir(keyFile), 0o755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(keyFile, seed, 0o600); err != nil {
		return err
	}

	// /jwt
	if err := os.MkdirAll(filepath.Dir(jwtFile), 0o755); err != nil {
		return err
	}
	if err := ioutil.WriteFile(jwtFile, []byte(jwt), 0o600); err != nil {
		return err
	}
	nc[filepath.Base(jwtFile)] = jwt

	// /creds
	creds, err := getNatsCredsData(jwt, seed)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(credFile), 0o755); err != nil {
		return err
	}
	if err = ioutil.WriteFile(credFile, creds, 0o666); err != nil {
		return err
	}
	nc[filepath.Base(credFile)] = string(creds)

	{
		filename := strings.TrimSuffix(filepath.Base(jwtFile), ".jwt")
		nc[filename+".pub"] = pub
		nc[filename+".nk"] = string(seed)
	}

	return nil
}

// getNatsCredsData returns a decorated file with a decorated JWT and decorated seed
func getNatsCredsData(jwtString string, seed []byte) ([]byte, error) {
	w := bytes.NewBuffer(nil)
	jd, err := jwt.DecorateJWT(jwtString)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(jd)
	if err != nil {
		return nil, err
	}

	d, err := jwt.DecorateSeed(seed)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(d)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}
