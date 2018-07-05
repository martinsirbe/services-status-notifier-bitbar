# Services Status Notifier BitBar plugin
A simple [**BitBar**][1] plugin which utilises `say` command to inform about
unhealthy services based on specified [**statuspage.io**][2] ID's.

setup
-----
```bash
go get -u github.com/martinsirbe/services-status-notifier-bitbar
make install
```  

Move the `services-status-notifier.1m.sh` script to your BitBar plugins directory.  
Update `GOPATH` in the `services-status-notifier.1m.sh`.

service checks
---------------------

To update services checks, add [**statuspage.io**][2] service ID to main.go `services` slice and run `make install`.

dependencies
------------
* [**BitBar**][1]
* [**statuspage.io**][2]
* [**golang**][3]

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[1]: https://getbitbar.com/
[2]: https://www.statuspage.io/
[3]: https://golang.org/