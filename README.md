  #Commandline options
  -e	Generates and prints an example message (Optional)
  -m int
    	Used if you would like to use a message attribute for routing simulation. (Optional) (default 1)
  -n string
    	Create a prefix to the message attribute.  Default is msg-(random). (Optional) (default "msg-%v")
  -p string
    	The topic's projectid. (Required)
  -r float
    	The number of message you would like to generate per second. (Optional) (default 1000)
  -s int
    	The size of the data body for the message in bytes. (Optional) (default 1000)
  -t string
    	Name of the topic to connect to.  If the topic is not found it will be created.(Required)
