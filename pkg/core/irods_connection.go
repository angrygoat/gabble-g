// package represents the base connection to an iRODS agent
package irods_connection

import "net"

// a connection to an iRODS agent
type IrodsConnection stuct {
	IrodsAccount IrodsAccount
	AuthResult AuthResult
	IPConn IrodsConnection 
}

// represents the result of an authentication
type AuthResult struct {
	AuthorizingAccount IrodsAccount
	AuthorizedAccount IrodsAccount
	authorized bool = false
}

// auth type for iRODS login
type AuthType string

const (
    STANDARD   AuthType = "Standard"
    PAM AuthType = "PAM"
)

type IrodsAccount struct {
	Host string
	Port int = 1247
	Zone string
	UserName string 
	Password string 
	AdminUserNameForProxy string 

}

