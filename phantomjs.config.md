#File: /usr/bin/phantomjs                                                       

#!/bin/sh
LD_LIBRARY_PATH="/usr/lib/phantomjs:$LD_LIBRARY_PATH"

#export these values for it to work headless
export QT_QPA_PLATFORM=offscreen
export QT_QPA_FONTDIR=/usr/share/fonts

export LD_LIBRARY_PATH
exec "/usr/lib/phantomjs/phantomjs" "$@"

