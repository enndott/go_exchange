go_exchange
==============

by: nicholas ward
--------------

Command line access to currency exchange rates from the openexchange.org API. Consider this a v0.01 - just playing with Go to see how it deals with JSON. 

Two options for running - 

* _go run go_exchange_ - produces full list of currencies and exchange rates
* _go run go_exchange EUR,GBP_ - produces a exchange rates by currency code

If you have any feedback please share it. I would love to learn more about go error handling, making sure that the user input is clean, and if there is a better way to deal with the JSON file. 