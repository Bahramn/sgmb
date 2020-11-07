# SGMB

## Simple Golang Message Broker

## Service config file
`/etc/systemd/system/sgmb.service`
```
[Unit]
Description=Simple Golang Message Brocker
ConditionPathExists=/home/pi/App
After=network.target
 
[Service]
Type=simple
User=pi
Group=pi
LimitNOFILE=1024

Restart=on-failure
RestartSec=10
startLimitIntervalSec=60

WorkingDirectory=/home/pi/App
ExecStart=sudo /home/pi/App/main

# make sure log directory exists and owned by syslog
PermissionsStartOnly=true
ExecStartPre=/bin/mkdir -p /var/log/sgmb-service
ExecStartPre=/bin/chown pi:pi /var/log/sgmb-service
ExecStartPre=/bin/chmod 755 /var/log/sgmb-service
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=sgmb-service
 
[Install]
WantedBy=multi-user.target
```
## Logging with unix rotation daily
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
