/* Written by Dave Richards.
 *
 * This is the top-level plugin interface, where you produce a new instance of an user lookup service.
 */
package discodove_interface_userlookup

import (
	"github.com/mediumdaver/discodove-data-store-plugin"
)
 
/* Send this down the userlookup channel to request userlookup via discodove, and it will do the hard
 * work to figure out what plugins to use. 
 * username 	: username to authenticate
 * responseChan : the channel, of type DiscoDoveUserLookupResponse, down which the userlookup will send
 * 				  the response - the userlookup will not close this channel.
 */
type DiscoDoveUserLookupRequest struct {
	Username string  
	ResponseChan chan DiscoDoveUserLookupResponse
}

// For use with authResult in DiscoDoveAuthResponse
const (
	UserFound = iota
	UserNotFound
	UserLookupError
)

/* The userlookup response to your request
 * lookupResult 	: the results of a PerformAuthentication request
 * datastore 		: a datastore instance for this user.
 */
type DiscoDoveUserLookupResponse struct {
	LookupResult int
	Datastore discodove_interface_datastore.DiscoDoveDataStore
}