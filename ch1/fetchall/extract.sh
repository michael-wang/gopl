grep 'domain\">[^<]' ./list.html | sed 's/[^>]*>\([^<]*\).*/\1/'
