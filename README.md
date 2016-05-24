# GA Beacon for Cloud Foundry [![.](http://gaproxy.gstack.io/UA-74118635-2/github.com/gstackio/ga-beacon/readme?pixel)](https://github.com/gstackio/ga-beacon)

This project is an adaptation of the
[original GA Beacon](https://github.com/igrigorik/ga-beacon) code for use in
any [Cloud Foundry](https://www.cloudfoundry.org/) deployment, and possibly
Heroku with few work. (Instead of GAE-compatible clouds like
[Google App Engine](https://cloud.google.com/appengine/)
or [AppScale](https://github.com/AppScale/appscale).)

The point is to trigger Google Analytics hits in contexts where JavaScript
code cannot be embedded, like in emails, in Github `README.md` files or in
Google Docs, Slides or Spreadsheets. More information about this can be found
in the original documentation at [DOCUMENTATION.md](DOCUMENTATION.md).


# How to build, run and test locally

The code doesn't rely the GAE API anymore, so after checkouting this repo, you
simply build with `go build` and run it locally with `./ga-beacon`.

Using your real `UA-XXXXX-X` “Tracking ID” and watching the
_Real-Time › Overview_ console of you Google Analytics account, you can test
the result with:

    open http://ga-beacon.local.gstack.me:8080/UA-XXXXX-X/pif/paf/pouf

… and immediately see a hit in the console!

(If your local 8080 TCP port is already in use, then the `PORT` environment
variable is here for you. Running `PORT=8888 ./ga-beacon` will be the way to
tweak the default setting.)


# How to push to Cloud Foundry

After properly targeting your Org and Space, just push with `cf push` or
anything similar.

You can tweak the default settings in the [manifest.yml](manifest.yml).


# Possible improvements

Logging statements have been converted to plain `fmt.Printf()` calls. So,
basically the logging is always at DEBUG level. Improving this is left as an
exercise to the reader. Submissions for such improvements are welcome.

Support for deploying to [Heroku](https://www.heroku.com/) is also left as an
exercise to the reader. Submission for this are also welcome. We already
favored using a `Procfile`, over the `command:` attribute of the
`manifest.yml` to help in such support.


# Authors and license

Copyright © 2014 Ilya Grigorik (original author)

Copyright © 2016 Benjamin Gandon (for Cloud Foundry adaptations)

`ga-beacon` is released under the terms of the [MIT License](LICENSE).
