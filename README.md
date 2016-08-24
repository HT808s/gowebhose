webhose.io client for Golang
============================

A simple way to access the `webhose.io <https://webhose.io>`_ API from your Golang code

.. code-block:: golang

    import gwb "github.com/ht808s/gowebhose"

    func main() {
      webhose := gwb.Webhose{
        Token: YOUR_API_KEY,
        Parameters: map[string]string{
          "language": "english",
          "country": "US",
        }
      }
      posts := webhose.Search("github")
      for _, post := range posts {
          fmt.Println(post.Title)
      }
    }

API Key
-------

To make use of the webhose.io API, you need to obtain a token that would be
used on every request. To obtain an API key, create an account at
https://webhose.io/auth/signup, and then go into
https://webhose.io/dashboard to see your token.

Installing
----------
You can install from source:

.. code-block:: bash

    $ go get -u "github.com/ht808s/gowebhose"



Query objects
^^^^^^^^^^^^^

Query object correspond to the advanced search options that appear on https://webhose.io/use

Query object have the following members:

* ``all_terms`` - a list of strings, all of which must appear in the results
* ``some_terms`` - a list of strings, some of which must appear in the results
* ``phrase`` - a phrase that must appear verbatim in the results
* ``exclude`` - terms that should not appear in the results
* ``site_type`` - one or more of ``discussions``, ``news``, ``blogs``
* ``language`` - one or more of language names, in lowercase english
* ``site`` - one or more of site names, top level only (i.e., yahoo.com and not news.yahoo.com)
* ``title`` - terms that must appear in the title
* ``body_text`` - term that must appear in the body text

Query objects implement the ``__str__()`` method, which shows the resulting search string.
