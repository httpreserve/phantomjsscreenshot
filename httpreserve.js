//httpreserve screenshot functionality
//install in /usr/js for httpreserve's needs
var system = require('system');

if (system.args.length != 3) {
	console.log('URL and output args required');
	phantom.exit(1);
}

address = system.args[1];
output = system.args[2];

var page = require('webpage').create();

//viewportSize being the actual size of the headless browser
page.viewportSize = { width: 1024, height: 768 };
page.clipRect = { left: 0, top: 0, width: 1024, height: 768 };

//set a sensible user-agent
page.settings.userAgent = 'httpreserve-websnapshot-0.0.1';

//the rest of the code is the same as the previous example
page.open(address, function() {
  page.render(output, {quality: '100'});
  phantom.exit();
});
