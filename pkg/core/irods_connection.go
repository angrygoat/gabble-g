// package represents the base connection to an iRODS agent
package irods_connection

import {
	"net"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"irods_context"
}

// a connection to an iRODS agent
type IrodsConnection stuct {
	IrodsAccount IrodsAccount
	AuthResult AuthResult
	IPConn IrodsConnection 
	ConnStatus ConnStatus = CONN_INIT
}

// represents the result of an authentication
type AuthResult struct {
	AuthorizingAccount IrodsAccount
	AuthorizedAccount IrodsAccount
	authorized bool = false
}

// auth type for iRODS login
type AuthType string
// connection status
type ConnStatus string

const (
    STANDARD   AuthType = "Standard"
    PAM AuthType = "PAM"
)

// enumerates status of an agent connection
const (
	CONN_INIT   ConnStatus = "init"
	CONN_PRE_AUTH   ConnStatus = "pre-auth"
	CONN_CONNECTED  ConnStatus = "connected"
	CONN_CLOSED   ConnStatus = "closed"
)

type IrodsAccount struct {
	Host string
	Port int = 1247
	Zone string
	UserName string 
	Password string 
	AdminUserNameForProxy string 

}

func ConnectionFactory(irodsAccount IrodsAccount, irodsContext IrodsContext)



