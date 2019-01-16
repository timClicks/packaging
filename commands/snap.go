// Copyright 2019 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the LGPLv3, see LICENCE file for details.

package commands

var getProxy = `
{
	grep -s ^http_proxy= /etc/environment  | sed 's/^http_proxy=//'  | sed 's/"//g';
	grep -s ^https_proxy= /etc/environment | sed 's/^https_proxy=//' | sed 's/"//g';

}
`[1:]

var setProxy = `
if ! grep -qs ^http_proxy= /etc/environment; then
	echo 'http_proxy="%s"' >> /etc/environment
else
	sed -i 's/^http_proxy=.*/http_proxy=%s/' /etc/environment
fi

if ! grep -qs ^https_proxy= /etc/environment; then
	echo 'https_proxy="%s"' >> /etc/environment
else
	sed -i 's/^https_proxy=.*/https_proxy=%s/' /etc/environment
fi
`[1:]

var unsetProxy = `
if grep -qs ^http_proxy= /etc/environment; then
	sed -i 's/^http_proxy=.*//' /etc/environment
fi

if grep -qs ^https_proxy= /etc/environment; then
	sed -i 's/^https_proxy=.*//' /etc/environment
fi
`[1:]

// snapCmder provides commands that are relevant for snap-based systems.
var snapCmder = packageCommander{
	update:        "snap refresh",
	upgrade:       `snap refresh "%s"`,
	listInstalled: "snap list",
	install:       "snap install",
	cleanup:       "snap refresh",
	listAvailable: "snap list",
	remove:        "snap remove",
	purge:         "snap remove",
	search:        "snap find %s",
	isInstalled:   `snap list | grep "^%s"`,
	getProxy:      getProxy,
	setProxy:      setProxy,
	setNoProxy:    unsetProxy,
}
