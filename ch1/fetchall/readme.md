To obtain `exec1.11.input`, I produce `list.html` (not in repo for copy right concern) 
from website: https://www.similarweb.com/top-websites/
By using browser's 'developer tool' (e.g. `option+cmd+i` for chrome on Mac), and copy
following element:
```html
#app > div > main > div > section.data-section.tw-table > div > div > div.tw-table__main-content > div.tw-table__body.tw-table__body--main
```

I got above content by `copy selector` on the target element. Then I paste it to list.html
. I'm sure there are tools to do this (and it's also a nice coding quiz), but let's just 
use simple copy and paste here for now.

Then `./extract.sh > exec1.11.input` to produce website list that can be consumed by:
`make run-exec1.11`

Observe how `fetchall` print error message if website does not response, as exercise 1.11
 asked.
