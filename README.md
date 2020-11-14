# SGMB
### Simple Golang Message Broker
It's a simple message broker under development for one of my client that he need for IOT usage.
At the first release it has ability to listen on both TCP and UDP protocol and routing between the sender and receiver of messages. 
#### How it works
Based on a config file it will activate and listen on configured port with TCP and UDP protocol.\
Messages have pre-defined format.

Format messages:\
`$mobile-number*device-id*message-body-or-command#`\
`@device-id*mobile-number*message-body-or-command#`\
message will split by `*`\
`@` and `$` are symbol identifier which `@` is for devices and `$` for mobile numbers\
firs part is sender id and second part receiver id, and the last part that will delimit by `#` is message body or command name\
Example commands:\
`RDY` : It's for checking connectivity from devices and will receive every minute\
`@device-id*device-id*RDY#`\
the response is as same as request message

`QUITE` : It's for closing connection\
`@device-id*device-id*QUITE#`

Any other message body parts will act as send message command\
`@device-id*mobile-number*message body any format#`\
The response should be `JB_OK` in case the receiver message id is active and listening and message has been sent to it 
or `JB_NOK` when the receiver is not connect. 

#### Service config file
`/etc/systemd/system/sgmb.service`
```
[Unit]
Description=Simple Golang Message Brocker
ConditionPathExists=/home/pi/App
After=network.target
 
[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024

Restart=on-failure
RestartSec=10
startLimitIntervalSec=60

WorkingDirectory=/home/pi/App
ExecStart=/home/pi/App/main

# make sure log directory exists and owned by syslog
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/sgmb-service
ExecStartPre=/bin/chown root:root /var/log/sgmb-service
ExecStartPre=/bin/chmod 755 /var/log/sgmb-service
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=sgmb-service
 
[Install]
WantedBy=multi-user.target
```
#### Logging with unix rotation daily config
```
/PATH_TO_PROJECT/storage/logs/sgmb.log {
  	su root root
	daily
	missingok
        rotate 60
	#compress
	create
	copytruncate
	dateext
	dateformat -%Y-%m-%d
	dateyesterday
}
```
