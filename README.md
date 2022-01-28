# username-finder
First app written on Golang: just to touch Goroutines and Mocks
Currently only one endpoint is impemented. WIP. List of endpoints described below
## API methods:
	POST: /username
		input : JSON body ["url1", "url2", ..."urln"]
		outout : JSON  ["valid-url1", "valid-url2", ..."valid-urln"]

	POST: /qr
		input : JSON body ["url1", "url2", ..."urln"]
		return : qr code in text format for valid urls