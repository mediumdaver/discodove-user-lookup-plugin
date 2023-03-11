/* Written by Dave Richards.
 *
 * This is the top-level plugin interface, where you produce a new instance of an user lookup service.
 */
package discodove_interface_userlookup

import (
	"log/syslog"
	"github.com/spf13/viper"
)
 

type DiscoDoveUserLookupPlugin interface { 

	/* This will be called once when we load this plugin, if you feel compelled to set something up, perhaps a 
	 * control/query/admin thread or something, then do it here in a controlled manner - similarly if 
	 * you want to pool connections, etc....  We assume that each plugin can scale itself, we do no magic
	 * to allow for scalability, so you might want some worker threads.
	 *
	 * Each plugin is responsible for creating it's own syslog connection as *syslog.Writer has a mutex, and 
	 * I don't want the user lookup threads to be blocking on writing to syslog - so you need to scale logging yourself.
	 * 
	 * We use Viper for config, and discodove will pass in the config directives for your module, but as it's viper you
	 * can access the entire discodove config too.  Feel free to specify your own config directives.
	 *
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove" - please
	 *            prefix your log messages with this and perhaps your own identifier e.g. "ldap lookup"
	 * syslogFacility : which facility to use in syslog.
	 * conf: a Viper subtree configuration for this service as specified in the discodove config.
	 */
	Initialize(name string, syslogFacility syslog.Priority, conf *viper.Viper) error

	/* This function is called to lookup a user and return the user structure that can be passed over to 
	 * the mail store. You should expect many concurrent calls (as go routines) of this function.
	 */
	LookupUser(user string) (store string, e error)
}
