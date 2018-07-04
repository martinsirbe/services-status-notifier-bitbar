# Services Status Notifier BitBar plugin
A simple BitBar plugin for MacOS user which uses `say` command and
to inform about unhealthy services based on provided [**statuspage.io**][2] ID's.

setup
-----
```bash
go get -u github.com/martinsirbe/services-status-notifier-bitbar
make install
```  

Move the `services-status-notifier.1m.sh` script to your BitBar plugins directory.  

service checks
---------------------

To update services checks, add [**statuspage.io**][2] service ID to main.go `services` and run `make install`.

dependencies
------------
* [**BitBar**][1]
* [**statuspage.io**][2]
* [**logrus**][3]


[1]: https://getbitbar.com/
[2]: https://www.statuspage.io/
[3]: https://github.com/sirupsen/logrus
